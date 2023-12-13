package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetPlugin(pluginUUID string, args nargs.Args) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/%s", pluginUUID)
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
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

func (g *GRPCMarshaller) GetPluginByName(name string, args nargs.Args) (*model.Plugin, error) {
	api := fmt.Sprintf("/api/plugins/name/%s", name)
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
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

func (g *GRPCMarshaller) CreateModuleDir(name string) (*string, error) {
	api := fmt.Sprintf("/api/modules/name/%s/data-dir", name)
	res, err := g.DbHelper.CallDBHelper(http.PUT, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	r := string(res)
	return &r, nil
}

func (g *GRPCMarshaller) UpdatePluginMessage(name string, body *model.Plugin) error {
	api := fmt.Sprintf("/api/plugins/name/%s/message", name)
	_, err := g.CallDBHelperWithParser(http.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return err
	}
	return nil
}
