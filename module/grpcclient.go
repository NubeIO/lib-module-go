package module

import (
	"context"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/proto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

// GRPCClient is an implementation of Module that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.ModuleClient
}

func (m *GRPCClient) Init(dbHelper DBHelper, moduleName string) error {
	log.Debug("gRPC Init client has been called...")
	dbHelperServer := &GRPCDBHelperServer{Impl: dbHelper}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		proto.RegisterDBHelperServer(s, dbHelperServer)

		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)

	_, err := m.client.Init(context.Background(), &proto.InitRequest{
		AddServer:  brokerID,
		ModuleName: moduleName,
	})

	// s.Stop() // TODO: we haven't closed this
	return err
}

func (m *GRPCClient) Enable() error {
	log.Debug("gRPC Enable client has been called...")
	_, err := m.client.Enable(context.Background(), &proto.Empty{})
	return err
}

func (m *GRPCClient) Disable() error {
	log.Debug("gRPC Disable client has been called...")
	_, err := m.client.Disable(context.Background(), &proto.Empty{})
	return err
}

func (m *GRPCClient) ValidateAndSetConfig(config []byte) ([]byte, error) {
	log.Debug("gRPC ValidateAndSetConfig client has been called...")
	resp, err := m.client.ValidateAndSetConfig(context.Background(), &proto.ConfigBody{Config: config})
	if err != nil {
		return nil, err
	}
	return resp.R, nil
}

func (m *GRPCClient) GetInfo() (*Info, error) {
	log.Debug("gRPC GetInfo client has been called...")
	resp, err := m.client.GetInfo(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}
	return &Info{
		Name:       resp.Name,
		Author:     resp.Author,
		Website:    resp.Website,
		License:    resp.License,
		HasNetwork: resp.HasNetwork,
	}, nil
}

func ConvertHTTPToHeaders(httpHeaders http.Header) []*proto.Header {
	var headers []*proto.Header
	for key, values := range httpHeaders {
		header := &proto.Header{
			Key:    key,
			Values: values,
		}
		headers = append(headers, header)
	}
	return headers
}

func (m *GRPCClient) CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	log.Debug("gRPC Call client has been called...") // when server calls it first it lands here
	resp, err := m.client.CallModule(context.Background(), &proto.RequestModule{
		Method:    string(method),
		UrlString: urlString,
		Headers:   ConvertHTTPToHeaders(headers),
		Body:      body,
	})
	if err != nil {
		return nil, err
	}
	return resp.R, nil
}

// GRPCDBHelperServer is the gRPC server that GRPCDBHelperClient talks to.
type GRPCDBHelperServer struct {
	// This is the real implementation
	Impl DBHelper
}

func (m *GRPCDBHelperServer) CallDBHelper(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	method, err := nhttp.StringToMethod(req.Method)
	if err != nil {
		return nil, err
	}
	apiArgs, err := nargs.DeserializeArgs(req.Args)
	if err != nil {
		return nil, err
	}
	var r []byte
	if req.HostUUID != nil {
		r, err = m.Impl.CallDBHelper(method, req.Api, *apiArgs, req.Body, &Opts{HostUUID: *req.HostUUID})
	} else {
		r, err = m.Impl.CallDBHelper(method, req.Api, *apiArgs, req.Body)
	}
	if err != nil {
		return &proto.Response{R: nil, E: []byte(err.Error())}, nil
	}
	return &proto.Response{R: r, E: nil}, nil
}
