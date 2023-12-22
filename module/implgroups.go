package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateGroup(body *model.Group, opts ...*Opts) (*model.Group, error) {
	api := "/api/groups"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var group *model.Group
	err = json.Unmarshal(res, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GRPCMarshaller) GetGroups(opts ...*Opts) ([]*model.Group, error) {
	api := "/api/groups"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var groups []*model.Group
	err = json.Unmarshal(res, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *GRPCMarshaller) GetGroup(uuid string, opts ...*Opts) (*model.Group, error) {
	api := fmt.Sprintf("/api/groups/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var group *model.Group
	err = json.Unmarshal(res, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GRPCMarshaller) UpdateHostsStatus(uuid string, opts ...*Opts) (*model.Group, error) {
	api := fmt.Sprintf("/api/groups/%s/update-hosts-status", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var group *model.Group
	err = json.Unmarshal(res, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GRPCMarshaller) UpdateGroup(uuid string, body *model.Group, opts ...*Opts) (*model.Group, error) {
	api := fmt.Sprintf("/api/groups/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var group *model.Group
	err = json.Unmarshal(res, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GRPCMarshaller) DeleteGroup(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/groups/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
