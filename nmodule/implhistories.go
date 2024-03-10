package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateHistories(histories []*model.History, opts ...*Opts) (bool, error) {
	api := "/api/histories"
	_, err := g.CallDBHelperWithParser(nhttp.POST, api, histories, opts...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *GRPCMarshaller) GetHistories(historyRequest *dto.HistoryRequest, opts ...*Opts) (*dto.HistoryResponse, error) {
	api := "/api/histories"
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, historyRequest, opts...)
	if err != nil {
		return nil, err
	}
	var history *dto.HistoryResponse
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (g *GRPCMarshaller) GetHistoriesFromSqlite(historyRequest *dto.HistoryRequest, opts ...*Opts) (*dto.HistoryResponse, error) {
	api := "/api/histories-sqlite"
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, historyRequest, opts...)
	if err != nil {
		return nil, err
	}
	var history *dto.HistoryResponse
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (g *GRPCMarshaller) GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string, opts ...*Opts) (*model.History, error) {
	api := fmt.Sprintf("/api/histories/point-uuid/%s/host-uuid/%s/latest", pointUUID, hostUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetHistoriesForSync(opts ...*Opts) (*dto.HistorySync, error) {
	api := fmt.Sprintf("/api/histories/sync")
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var historySync *dto.HistorySync
	err = json.Unmarshal(res, &historySync)
	if err != nil {
		return nil, err
	}
	return historySync, nil
}

// DeleteHistories required: opts[0].Args.TimestampLt
func (g *GRPCMarshaller) DeleteHistories(opts ...*Opts) error {
	api := "/api/histories"
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
