package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"strconv"
)

func (g *GRPCMarshaller) CreatePoint(body *model.Point, opts ...*Opts) (*model.Point, error) {
	api := "/api/points"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
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

func (g *GRPCMarshaller) GetPoints(body *dto.Filter, opts ...*Opts) ([]*model.Point, error) {
	api := "/api/points"
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, body, opts...)
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

func (g *GRPCMarshaller) GetPoint(uuid string, opts ...*Opts) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetPointByName(networkName, deviceName, pointName string, opts ...*Opts) (*model.Point,
	error) {
	api := fmt.Sprintf("/api/points/name/%s/%s/%s", networkName, deviceName, pointName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetOnePointByArgs(opts ...*Opts) (*model.Point, error) {
	api := "/api/points/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetPointWithParent(uuid string, opts ...*Opts) (*dto.PointWithParent, error) {
	api := fmt.Sprintf("/api/points/%s/with-parents", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var pointWithParent *dto.PointWithParent
	err = json.Unmarshal(res, &pointWithParent)
	if err != nil {
		return nil, err
	}
	return pointWithParent, nil
}

func (g *GRPCMarshaller) GetPointWithParentByName(networkName, deviceName, pointName string, opts ...*Opts) (*dto.PointWithParent, error) {
	api := fmt.Sprintf("/api/points/name/%s/%s/%s/with-parents", networkName, deviceName, pointName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var pointWithParent *dto.PointWithParent
	err = json.Unmarshal(res, &pointWithParent)
	if err != nil {
		return nil, err
	}
	return pointWithParent, nil
}

func (g *GRPCMarshaller) CountPoints(body *dto.Filter, opts ...*Opts) (int, error) {
	api := fmt.Sprintf("/api/points/count")
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, body, opts...)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(res))
}

func (g *GRPCMarshaller) UpdatePoint(uuid string, body *model.Point, opts ...*Opts) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
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

func (g *GRPCMarshaller) PointWrite(uuid string, body *dto.PointWriter, opts ...*Opts) (*dto.PointWriteResponse, error) {
	api := fmt.Sprintf("/api/points/%s/write", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	var pwr *dto.PointWriteResponse
	err = json.Unmarshal(res, &pwr)
	if err != nil {
		return nil, err
	}
	return pwr, nil
}

func (g *GRPCMarshaller) PointWriteByName(networkName, deviceName, pointName string, body *dto.PointWriter, opts ...*Opts) (*dto.PointWriteResponse, error) {
	api := fmt.Sprintf("/api/points/name/%s/%s/%s/write", networkName, deviceName, pointName)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	var pwr *dto.PointWriteResponse
	err = json.Unmarshal(res, &pwr)
	if err != nil {
		return nil, err
	}
	return pwr, nil
}

func (g *GRPCMarshaller) UpdatePointFault(uuid string, body *model.CommonFault, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/%s/fault", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdatePointState(uuid string, body datatype.PointState, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/%s/state", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpsertPoint(uuid string, body *model.Point, opts ...*Opts) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
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

func (g *GRPCMarshaller) UpsertPointMetaTags(uuid string, body []*model.PointMetaTag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/%s/meta-tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertPointTags(uuid string, body []*model.Tag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/%s/tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeletePoint(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteOnePointByArgs(opts ...*Opts) error {
	api := "/api/points/one/args"
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeletePointByName(name string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/points/name/%s", name)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
