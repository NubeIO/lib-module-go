package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) GetPlugins(opts ...*Opts) ([]*model.Plugin, error) {
	api := "/api/plugins"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var plugins []*model.Plugin
	err = json.Unmarshal(res, &plugins)
	if err != nil {
		return nil, err
	}
	return plugins, nil
}

func (g *GRPCMarshaller) GetPlugin(uuid string, opts ...*Opts) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var plugin *model.Plugin
	err = json.Unmarshal(res, &plugin)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}

func (g *GRPCMarshaller) GetPluginByName(name string, opts ...*Opts) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/name/%s", name)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var plugin *model.Plugin
	err = json.Unmarshal(res, &plugin)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}

func (g *GRPCMarshaller) CreateModuleDir(name string, opts ...*Opts) (*string, error) {
	api := fmt.Sprintf("/api/modules/name/%s/data-dir", name)
	res, err := g.DbHelper.CallDBHelper(nhttp.PUT, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	r := string(res)
	return &r, nil
}

func (g *GRPCMarshaller) UpdatePluginMessage(name string, body *model.Plugin, opts ...*Opts) error {
	api := fmt.Sprintf("/api/plugins/name/%s/message", name)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}
