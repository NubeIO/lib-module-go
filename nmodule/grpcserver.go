package nmodule

import (
	"context"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/proto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Module

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) Init(ctx context.Context, req *proto.InitRequest) (*proto.Empty, error) {
	log.Debug("gRPC Init server has been called...")
	conn, err := m.broker.Dial(req.AddServer)
	if err != nil {
		return nil, err
	}
	// defer conn.Close() // TODO: we haven't closed this
	dbHelper := &GRPCDBHelperClient{proto.NewDBHelperClient(conn)}
	err = m.Impl.Init(dbHelper, req.ModuleName)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (m *GRPCServer) Enable(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	log.Debug("gRPC Enable server has been called...")
	err := m.Impl.Enable()
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (m *GRPCServer) Disable(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	log.Debug("gRPC Disable server has been called...")
	err := m.Impl.Disable()
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (m *GRPCServer) ValidateAndSetConfig(ctx context.Context, req *proto.ConfigBody) (*proto.Response, error) {
	log.Debug("gRPC Disable server has been called...")
	bytes, err := m.Impl.ValidateAndSetConfig(req.Config)
	if err != nil {
		return nil, err
	}
	return &proto.Response{R: bytes}, nil
}

func (m *GRPCServer) GetInfo(ctx context.Context, req *proto.Empty) (*proto.InfoResponse, error) {
	log.Debug("gRPC GetInfo server has been called...")
	r, err := m.Impl.GetInfo()
	if err != nil {
		return nil, err
	}
	return &proto.InfoResponse{
		Name:       r.Name,
		Author:     r.Author,
		Website:    r.Website,
		License:    r.License,
		HasNetwork: r.HasNetwork,
	}, nil
}

func ConvertHeadersToHTTP(protoHeaders []*proto.Header) http.Header {
	headers := make(http.Header)
	for _, protoHeader := range protoHeaders {
		for _, value := range protoHeader.Values {
			headers.Add(protoHeader.Key, value)
		}
	}
	return headers
}

func (m *GRPCServer) CallModule(ctx context.Context, req *proto.RequestModule) (*proto.Response, error) {
	log.Debug("gRPC CallModule server has been called...") // when server calls it, it lands second (it is in module)
	method, err := nhttp.StringToMethod(req.Method)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	r, err := m.Impl.CallModule(method, req.UrlString, ConvertHeadersToHTTP(req.Headers), req.Body)
	if err != nil {
		return nil, err
	}
	return &proto.Response{R: r}, nil
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
	var apiArgs *nargs.Args
	if req.Args != nil {
		apiArgs, err = nargs.DeserializeArgs(*req.Args)
		if err != nil {
			return nil, err
		}
	}

	var r []byte
	if req.HostUUID != nil {
		r, err = m.Impl.CallDBHelper(method, req.Api, req.Body, &Opts{Args: apiArgs, HostUUID: req.HostUUID})
	} else {
		r, err = m.Impl.CallDBHelper(method, req.Api, req.Body)
	}
	if err != nil {
		return &proto.Response{R: nil, E: []byte(err.Error())}, nil
	}
	return &proto.Response{R: r, E: nil}, nil
}
