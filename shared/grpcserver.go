package shared

import (
	"context"
	"errors"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/proto"
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

func (m *GRPCServer) Call(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Debug("gRPC Call server has been called...")
	r, err := m.Impl.Call(req.Method, req.Api, req.Args, req.Body)
	if err != nil {
		return nil, err
	}
	return &proto.Response{R: r}, nil
}

// GRPCDBHelperClient is an implementation of DBHelper that talks over RPC.
type GRPCDBHelperClient struct{ client proto.DBHelperClient }

func (m *GRPCDBHelperClient) Call(method http.Method, api, args string, body []byte) ([]byte, error) {
	resp, err := m.client.Call(context.Background(), &proto.Request{
		Method: string(method),
		Api:    api,
		Args:   args,
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

func (m *GRPCDBHelperClient) SetErrorsForAll(path, uuid, message, messageLevel, messageCode string, doPoints bool) error {
	resp, err := m.client.SetErrorsForAll(context.Background(), &proto.SetErrorsForAllRequest{
		Path:         path,
		Uuid:         uuid,
		Message:      message,
		MessageLevel: messageLevel,
		MessageCode:  messageCode,
		DoPoints:     doPoints,
	})
	if err != nil {
		return err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return errors.New(errStr)
	}
	return nil
}

func (m *GRPCDBHelperClient) PatchWithOpts(path, uuid string, body []byte, opts []byte) ([]byte, error) {
	resp, err := m.client.PatchWithOpts(context.Background(), &proto.PatchWithOptsRequest{
		Path: path,
		Uuid: uuid,
		Body: body,
		Opts: opts,
	})
	if err != nil {
		log.Error("PatchWithOpts: ", err)
		return nil, err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		log.Error("PatchWithOpts: ", errStr)
		return nil, errors.New(errStr)
	}
	return resp.R, nil
}

func (m *GRPCDBHelperClient) ClearErrorsForAll(path, uuid string, doPoints bool) error {
	resp, err := m.client.ClearErrorsForAll(context.Background(), &proto.ClearErrorsForAllRequest{
		Path:     path,
		Uuid:     uuid,
		DoPoints: doPoints,
	})
	if err != nil {
		return err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return errors.New(errStr)
	}
	return nil
}

func (m *GRPCDBHelperClient) WizardNewNetworkDevicePoint(plugin string, network, device, point []byte) (bool, error) {
	resp, err := m.client.WizardNewNetworkDevicePoint(context.Background(), &proto.WizardNewNetworkDevicePointRequest{
		Plugin:  plugin,
		Network: network,
		Device:  device,
		Point:   point,
	})
	if err != nil {
		return false, err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return false, errors.New(errStr)
	}
	return true, nil
}

func (m *GRPCDBHelperClient) CreateModuleDataDir(name string) (string, error) {
	resp, err := m.client.CreateModuleDataDir(context.Background(), &proto.DataDirRequest{Name: name})
	if err != nil {
		return "", err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return "", errors.New(errStr)
	}
	return resp.Dir, nil
}

func (m *GRPCDBHelperClient) MQTTPublish(topic string, qos []byte, retain bool, body string) error {
	resp, err := m.client.MQTTPublish(context.Background(), &proto.MQTTPublishRequest{
		Topic:  topic,
		Qos:    qos,
		Retain: retain,
		Body:   body,
	})
	if err != nil {
		return err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return errors.New(errStr)
	}
	return nil
}

func (m *GRPCDBHelperClient) MQTTPublishNonBuffer(topic string, qos []byte, retain bool, body []byte) error {
	resp, err := m.client.MQTTPublishNonBuffer(context.Background(), &proto.MQTTPublishNonBufferRequest{
		Topic:  topic,
		Qos:    qos,
		Retain: retain,
		Body:   body,
	})
	if err != nil {
		return err
	}
	if resp.E != nil {
		errStr := string(resp.E)
		return errors.New(errStr)
	}
	return nil
}
