package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) CreateHost(body *model.Host, opts ...*Opts) (*model.Host, error) {
	api := "/api/hosts"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var host *model.Host
	err = json.Unmarshal(res, &host)
	if err != nil {
		return nil, err
	}
	return host, nil
}

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

func (g *GRPCMarshaller) GetHost(uuid string, opts ...*Opts) (*model.Host, error) {
	api := fmt.Sprintf("/api/hosts/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var host *model.Host
	err = json.Unmarshal(res, &host)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func (g *GRPCMarshaller) UpdateHost(uuid string, body *model.Host, opts ...*Opts) (*model.Host, error) {
	api := fmt.Sprintf("/api/hosts/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var host *model.Host
	err = json.Unmarshal(res, &host)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func (g *GRPCMarshaller) UpsertHostTags(uuid string, body []*model.Tag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/hosts/%s/tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteHost(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/hosts/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
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
