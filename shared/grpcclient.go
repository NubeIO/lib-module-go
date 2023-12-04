package shared

import (
	"context"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/proto"
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

func (m *GRPCClient) Call(method, api, args string, body []byte) ([]byte, error) {
	log.Debug("gRPC Call client has been called...")
	resp, err := m.client.Call(context.Background(), &proto.Request{
		Method: method,
		Api:    api,
		Args:   args,
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
	r, err := m.Impl.Call(method, req.Api, req.Args, req.Body)
	if err != nil {
		return &proto.Response{R: nil, E: []byte(err.Error())}, nil
	}
	return &proto.Response{R: r, E: nil}, nil
}

func (m *GRPCDBHelperServer) PatchWithOpts(ctx context.Context, req *proto.PatchWithOptsRequest) (*proto.Response, error) {
	r, err := m.Impl.PatchWithOpts(req.Path, req.Uuid, req.Body, req.Opts)
	if err != nil {
		return &proto.Response{R: nil, E: []byte(err.Error())}, nil
	}
	return &proto.Response{R: r, E: nil}, nil
}

func (m *GRPCDBHelperServer) SetErrorsForAll(ctx context.Context, request *proto.SetErrorsForAllRequest) (*proto.ErrorResponse, error) {
	err := m.Impl.SetErrorsForAll(
		request.Path,
		request.Uuid,
		request.Message,
		request.MessageLevel,
		request.MessageCode,
		request.DoPoints,
	)
	if err != nil {
		return &proto.ErrorResponse{E: []byte(err.Error())}, nil
	}
	return &proto.ErrorResponse{E: nil}, nil
}

func (m *GRPCDBHelperServer) ClearErrorsForAll(ctx context.Context, request *proto.ClearErrorsForAllRequest) (*proto.ErrorResponse, error) {
	err := m.Impl.ClearErrorsForAll(request.Path, request.Uuid, request.DoPoints)
	if err != nil {
		return &proto.ErrorResponse{E: []byte(err.Error())}, nil
	}
	return &proto.ErrorResponse{E: nil}, nil
}

func (m *GRPCDBHelperServer) WizardNewNetworkDevicePoint(ctx context.Context, request *proto.WizardNewNetworkDevicePointRequest) (*proto.BoolResponse, error) {
	_, err := m.Impl.WizardNewNetworkDevicePoint(request.Plugin, request.Network, request.Device, request.Point)
	if err != nil {
		return &proto.BoolResponse{R: false, E: []byte(err.Error())}, nil
	}
	return &proto.BoolResponse{R: true, E: nil}, nil
}

func (m *GRPCDBHelperServer) CreateModuleDataDir(ctx context.Context, request *proto.DataDirRequest) (*proto.DataDirResponse, error) {
	r, err := m.Impl.CreateModuleDataDir(request.Name)
	if err != nil {
		return &proto.DataDirResponse{Dir: "", E: []byte(err.Error())}, nil
	}
	return &proto.DataDirResponse{Dir: r, E: nil}, nil
}

func (m *GRPCDBHelperServer) MQTTPublish(ctx context.Context, request *proto.MQTTPublishRequest) (*proto.ErrorResponse, error) {
	err := m.Impl.MQTTPublish(request.Topic, request.Qos, request.Retain, request.Body)
	if err != nil {
		return &proto.ErrorResponse{E: []byte(err.Error())}, nil
	}
	return &proto.ErrorResponse{E: nil}, nil
}

func (m *GRPCDBHelperServer) MQTTPublishNonBuffer(ctx context.Context, request *proto.MQTTPublishNonBufferRequest) (*proto.ErrorResponse, error) {
	err := m.Impl.MQTTPublishNonBuffer(request.Topic, request.Qos, request.Retain, request.Body)
	if err != nil {
		return &proto.ErrorResponse{E: []byte(err.Error())}, nil
	}
	return &proto.ErrorResponse{E: nil}, nil
}
