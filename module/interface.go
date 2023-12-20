package module

import (
	"context"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"net/http"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

type DBHelper interface {
	CallDBHelper(method nhttp.Method, api string, body []byte, opts ...*Opts) ([]byte, error)
}

type Info struct {
	Name       string
	Author     string
	Website    string
	License    string
	HasNetwork bool
}

// Module is the interface that we're exposing as a plugin.
type Module interface {
	ValidateAndSetConfig(config []byte) ([]byte, error)
	Init(dbHelper DBHelper, moduleName string) error
	Enable() error
	Disable() error
	GetInfo() (*Info, error)
	CallModule(method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error)
}

// NubeModule is the implementation of plugin.Plugin so we can serve/consume this.
type NubeModule struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Module
}

func (p *NubeModule) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterModuleServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *NubeModule) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewModuleClient(c),
		broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &NubeModule{}
