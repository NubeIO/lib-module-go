package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) GetHistoryLogByHostUUID(hostUUID string, opts ...*Opts) (*model.HistoryLog, error) {
	api := fmt.Sprintf("/api/history-logs/host-uuid/%s", hostUUID)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) UpdateHistoryLog(body *model.HistoryLog, opts ...*Opts) (bool, error) {
	api := "/api/history-logs"
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	return err == nil, err
}
