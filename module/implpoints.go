package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) CreatePoint(body *model.Point) (*model.Point, error) {
	api := "/api/points"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, nargs.Args{}, body)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) GetPoints(args nargs.Args) ([]*model.Point, error) {
	api := "/api/points"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var points []*model.Point
	err = json.Unmarshal(res, &points)
	if err != nil {
		return nil, err
	}
	return points, nil
}

func (g *GRPCMarshaller) GetPoint(uuid string, args nargs.Args) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) GetPointByName(networkName, deviceName, pointName string, args nargs.Args) (*model.Point,
	error) {
	api := fmt.Sprintf("/api/points/name/%s/%s/%s", networkName, deviceName, pointName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) GetOnePointByArgs(args nargs.Args) (*model.Point, error) {
	api := "/api/points/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) UpdatePoint(uuid string, body *model.Point, args nargs.Args) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, args, body)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) UpdatePointErrors(uuid string, body *model.Point) error {
	api := fmt.Sprintf("/api/points/%s/error", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdatePointSuccess(uuid string, body *model.Point) error {
	api := fmt.Sprintf("/api/points/%s/success", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) PointWrite(uuid string, body *model.PointWriter) (*model.PointWriteResponse, error) {
	api := fmt.Sprintf("/api/points/%s/write", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body)

	var pwr *model.PointWriteResponse
	err = json.Unmarshal(res, &pwr)
	if err != nil {
		return nil, err
	}
	return pwr, nil
}

func (g *GRPCMarshaller) DeletePoint(uuid string) error {
	api := fmt.Sprintf("/api/points/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nargs.Args{}, nil)
	return err
}
