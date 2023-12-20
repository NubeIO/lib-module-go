package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetPlugin(pluginUUID string, args nargs.Args, opts ...*Opts) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/%s", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var pluginConf *model.Plugin
	err = json.Unmarshal(res, &pluginConf)
	if err != nil {
		return nil, err
	}
	return pluginConf, nil
}

func (g *GRPCMarshaller) GetPluginByName(name string, args nargs.Args, opts ...*Opts) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/name/%s", name)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil, opts...)
	if err != nil {
		return nil, err
	}
	var pluginConf *model.Plugin
	err = json.Unmarshal(res, &pluginConf)
	if err != nil {
		return nil, err
	}
	return pluginConf, nil
}

func (g *GRPCMarshaller) CreateModuleDir(name string, opts ...*Opts) (*string, error) {
	api := fmt.Sprintf("/api/modules/name/%s/data-dir", name)
	res, err := g.DbHelper.CallDBHelper(nhttp.PUT, api, nargs.Args{}, nil, opts...)
	if err != nil {
		return nil, err
	}
	r := string(res)
	return &r, nil
}

func (g *GRPCMarshaller) UpdatePluginMessage(name string, body *model.Plugin, opts ...*Opts) error {
	api := fmt.Sprintf("/api/plugins/name/%s/message", name)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body, opts...)
	if err != nil {
		return err
	}
	return nil
}
