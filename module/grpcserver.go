package module

import (
	"context"
	"errors"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/parser"
	"github.com/NubeIO/lib-module-go/proto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
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

func (m *GRPCServer) CallModule(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Debug("gRPC CallModule server has been called...") // when server calls it, it lands second (it is in module)
	method, err := http.StringToMethod(req.Method)
	if err != nil {
		return nil, err
	}
	apiArgs, err := parser.DeserializeArgs(req.Args)
	if err != nil {
		return nil, err
	}
	r, err := m.Impl.CallModule(method, req.Api, *apiArgs, req.Body)
	if err != nil {
		return nil, err
	}
	return &proto.Response{R: r}, nil
}

// GRPCDBHelperClient is an implementation of DBHelper that talks over RPC.
type GRPCDBHelperClient struct{ client proto.DBHelperClient }

func (m *GRPCDBHelperClient) CallDBHelper(method http.Method, api string, args nargs.Args, body []byte) ([]byte, error) {
	// This should call at first from module
	apiArgs, err := parser.SerializeArgs(args)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.CallDBHelper(context.Background(), &proto.Request{
		Method: string(method),
		Api:    api,
		Args:   *apiArgs,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return nil, errors.New(errStr)
	}
	return resp.R, nil
}
