package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateLocation(body *model.Location, opts ...*Opts) (*model.Location, error) {
	api := "/api/locations"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var location *model.Location
	err = json.Unmarshal(res, &location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (g *GRPCMarshaller) GetLocations(opts ...*Opts) ([]*model.Location, error) {
	api := "/api/locations"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var locations []*model.Location
	err = json.Unmarshal(res, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (g *GRPCMarshaller) GetLocation(uuid string, opts ...*Opts) (*model.Location, error) {
	api := fmt.Sprintf("/api/locations/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var location *model.Location
	err = json.Unmarshal(res, &location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (g *GRPCMarshaller) UpdateLocation(uuid string, body *model.Location, opts ...*Opts) (*model.Location, error) {
	api := fmt.Sprintf("/api/locations/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var location *model.Location
	err = json.Unmarshal(res, &location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (g *GRPCMarshaller) DeleteLocation(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/locations/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
