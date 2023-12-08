package shared

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetHosts(args nargs.Args) ([]*model.Host, error) {
	api := "/api/hosts"
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var hosts []*model.Host
	err = json.Unmarshal(res, &hosts)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (g *GRPCMarshaller) CloneHostThingsToCloud(hostUUID string) error {
	api := "/api/host/clone-things-to-cloud"
	_, err := g.DbHelper.CallDBHelper(http.GET, api, nargs.Args{HostUUID: &hostUUID}, nil)
	return err
}