package module

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetHosts(opts ...*Opts) ([]*model.Host, error) {
	api := "/api/hosts"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) CloneHostThingsToCloud(hostUUID string, opts ...*Opts) error {
	api := "/api/host/clone-things-to-cloud"
	if len(opts) > 0 {
		opts[0].Args = &nargs.Args{HostUUID: &hostUUID}
	} else {
		opts = append(opts, &Opts{Args: &nargs.Args{HostUUID: &hostUUID}})
	}
	_, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	return err
}
