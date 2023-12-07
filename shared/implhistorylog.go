package shared

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetHistoryLogByHostUUID(hostUUID string) (*model.HistoryLog, error) {
	api := fmt.Sprintf("/api/history-logs/host-uuid/%s", hostUUID)
	res, err := g.DbHelper.CallDBHelper(http.GET, api, nargs.Args{}, nil)
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

func (g *GRPCMarshaller) UpdateBulkHistoryLogs(logs []*model.HistoryLog) (bool, error) {
	api := "/api/history-logs"
	_, err := g.CallDBHelperWithParser(http.PATCH, api, nargs.Args{}, logs)
	if err != nil {
		return false, err
	}
	return true, nil
}
