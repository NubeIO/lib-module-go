package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) CreateBulkHistory(histories []*model.History, opts ...*Opts) (bool, error) {
	api := "/api/histories"
	_, err := g.CallDBHelperWithParser(nhttp.POST, api, nargs.Args{}, histories, opts...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *GRPCMarshaller) CreateBulkPointHistory(histories []*model.PointHistory, opts ...*Opts) (bool, error) {
	api := "/api/histories/points"
	_, err := g.CallDBHelperWithParser(nhttp.POST, api, nargs.Args{}, histories, opts...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *GRPCMarshaller) GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string, opts ...*Opts) (*model.History, error) {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s/host-uuid/%s/latest", pointUUID, hostUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
	if err != nil {
		return nil, err
	}
	var history *model.History
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (g *GRPCMarshaller) GetPointHistoriesMissingTimestamps(pointUUID string, opts ...*Opts) ([]string, error) {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s/missing-timestamps", pointUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
	if err != nil {
		return nil, err
	}
	var missingTimestamps []string
	err = json.Unmarshal(res, &missingTimestamps)
	if err != nil {
		return nil, err
	}
	return missingTimestamps, nil
}
