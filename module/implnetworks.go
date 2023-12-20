package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) CreateNetwork(body *model.Network, opts ...*Opts) (*model.Network, error) {
	api := "/api/networks"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, nargs.Args{}, body, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworks(args nargs.Args, opts ...*Opts) ([]*model.Network, error) {
	api := "/api/networks"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) GetNetwork(uuid string, args nargs.Args, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworkByName(networkName string, args nargs.Args, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/name/%s", networkName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworkByPlugin(pluginUUID string, args nargs.Args, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetOneNetworkByArgs(args nargs.Args, opts ...*Opts) (*model.Network, error) {
	api := "/api/networks/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworksByPlugin(pluginUUID string, args nargs.Args, opts ...*Opts) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s/all", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) GetNetworksByPluginName(pluginName string, args nargs.Args, opts ...*Opts) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-name/%s/all", pluginName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) UpdateNetwork(uuid string, body *model.Network, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body, opts...)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) UpdateNetworkErrors(uuid string, body *model.Network, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/error", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdateNetworkDescendantsErrors(networkUUID, message, messageLevel, messageCode string, withPoints bool, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/error/descendants", networkUUID)
	network := &model.Network{
		CommonFault: model.CommonFault{
			Message:      message,
			MessageLevel: messageLevel,
			MessageCode:  messageCode,
		},
	}
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{WithPoints: withPoints}, network, opts...)
	return err
}

func (g *GRPCMarshaller) ClearNetworkDescendantsErrors(networkUUID string, withPoints bool, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/error/descendants", networkUUID)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nargs.Args{WithPoints: withPoints}, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteNetwork(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nargs.Args{}, nil, opts...)
	return err
}
