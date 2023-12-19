package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) GetSchedules() ([]*model.Schedule, error) {
	api := "/api/schedules"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var schedules []*model.Schedule
	if err = json.Unmarshal(res, &schedules); err != nil {
		return nil, err
	}
	return schedules, nil
}

func (g *GRPCMarshaller) UpdateScheduleAllProps(uuid string, body *model.Schedule) (*model.Schedule, error) {
	api := fmt.Sprintf("/api/schedules/%s/all-props", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return nil, err
	}
	var schedule *model.Schedule
	err = json.Unmarshal(res, &schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
