package shared

import (
	"context"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/parser"
	"github.com/NubeIO/lib-module-go/proto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

func (m *GRPCClient) Call(method http.Method, api string, args nargs.Args, body []byte) ([]byte, error) {
	log.Debug("gRPC Call client has been called...")
	apiArgs, _ := parser.SerializeArgs(args)
	resp, err := m.client.Call(context.Background(), &proto.Request{
		Method: string(method),
		Api:    api,
		Args:   *apiArgs,
		Body:   body,
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

func (m *GRPCDBHelperServer) Call(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	method, _ := http.StringToMethod(req.Method)
	apiArgs, _ := parser.DeserializeArgs(req.Args)
	r, err := m.Impl.Call(method, req.Api, *apiArgs, req.Body)
	if err != nil {
		return &proto.Response{R: nil, E: []byte(err.Error())}, nil
	}
	return &proto.Response{R: r, E: nil}, nil
}
