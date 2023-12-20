package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetHistoryLogByHostUUID(hostUUID string, opts ...*Opts) (*model.HistoryLog, error) {
	api := fmt.Sprintf("/api/history-logs/host-uuid/%s", hostUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
	if err != nil {
		return nil, err
	}
	var historyLog *model.HistoryLog
	err = json.Unmarshal(res, &historyLog)
	if err != nil {
		return nil, err
	}
	return historyLog, nil
}

func (g *GRPCMarshaller) UpdateBulkHistoryLogs(logs []*model.HistoryLog, opts ...*Opts) (bool, error) {
	api := "/api/history-logs"
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, logs, opts...)
	if err != nil {
		return false, err
	}
	return true, nil
}
