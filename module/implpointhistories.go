package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreatePointHistories(histories []*model.PointHistory, opts ...*Opts) (bool, error) {
	api := "/api/histories/points"
	_, err := g.CallDBHelperWithParser(nhttp.POST, api, histories, opts...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *GRPCMarshaller) GetPointHistories(opts ...*Opts) ([]*model.PointHistory, error) {
	api := "/api/histories/points"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var histories []*model.PointHistory
	err = json.Unmarshal(res, &histories)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (g *GRPCMarshaller) GetPointHistoriesByPointUUID(pointUUID string, opts ...*Opts) ([]*model.PointHistory, error) {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s", pointUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var histories []*model.PointHistory
	err = json.Unmarshal(res, &histories)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (g *GRPCMarshaller) GetLatestPointHistoryByPointUUID(pointUUID string, opts ...*Opts) (*model.PointHistory, error) {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s/one", pointUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var history *model.PointHistory
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (g *GRPCMarshaller) GetPointHistoriesByPointUUIDs(pointUUIDs []*string, opts ...*Opts) ([]*model.PointHistory, error) {
	api := "/api/histories/points/point-uuid"
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, pointUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var histories []*model.PointHistory
	err = json.Unmarshal(res, &histories)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

// GetPointHistoriesForSync required: opts[0].Args.Id, opts[0].Args.TimestampLt
func (g *GRPCMarshaller) GetPointHistoriesForSync(opts ...*Opts) ([]*model.PointHistory, error) {
	api := "/api/histories/points/sync"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var histories []*model.PointHistory
	err = json.Unmarshal(res, &histories)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (g *GRPCMarshaller) GetPointHistoriesMissingTimestamps(pointUUID string, opts ...*Opts) ([]string, error) {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s/missing-timestamps", pointUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) DeletePointHistoriesByPointUUID(pointUUID string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/histories/points/point-uuid/%s", pointUUID)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
