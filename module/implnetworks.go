package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strconv"
)

func (g *GRPCMarshaller) CreateNetwork(body *model.Network, opts ...*Opts) (*model.Network, error) {
	api := "/api/networks"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
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

func (g *GRPCMarshaller) GetNetworks(opts ...*Opts) ([]*model.Network, error) {
	api := "/api/networks"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetwork(uuid string, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworkByName(networkName string, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/name/%s", networkName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetOneNetworkByArgs(opts ...*Opts) (*model.Network, error) {
	api := "/api/networks/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworkByPlugin(pluginUUID string, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworksByPlugin(pluginUUID string, opts ...*Opts) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s/all", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworkByPluginName(pluginName string, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-name/%s", pluginName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworksByPluginName(pluginName string, opts ...*Opts) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-name/%s/all", pluginName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) CountNetworks(opts ...*Opts) (int, error) {
	api := fmt.Sprintf("/api/networks/count")
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(res))
}

func (g *GRPCMarshaller) UpdateNetwork(uuid string, body *model.Network, opts ...*Opts) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
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
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
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
	if len(opts) > 0 {
		opts[0].Args = &nargs.Args{WithPoints: withPoints}
	} else {
		opts = append(opts, &Opts{Args: &nargs.Args{WithPoints: withPoints}})
	}
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, network, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertNetworkMetaTags(uuid string, body []*model.NetworkMetaTag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/meta-tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertNetworkTags(uuid string, body []*model.Tag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteNetwork(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteOneNetworkByArgs(opts ...*Opts) error {
	api := "/api/networks/one/args"
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteNetworkByName(name string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/name/%s", name)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) ClearNetworkDescendantsErrors(networkUUID string, withPoints bool, opts ...*Opts) error {
	api := fmt.Sprintf("/api/networks/%s/error/descendants", networkUUID)
	if len(opts) > 0 {
		opts[0].Args = &nargs.Args{WithPoints: withPoints}
	} else {
		opts = append(opts, &Opts{Args: &nargs.Args{WithPoints: withPoints}})
	}
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
