package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateSchedule(body *model.Schedule, opts ...*Opts) (*model.Schedule, error) {
	api := "/api/schedules"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
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

func (g *GRPCMarshaller) GetSchedules(opts ...*Opts) ([]*model.Schedule, error) {
	api := "/api/schedules"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var schedules []*model.Schedule
	if err = json.Unmarshal(res, &schedules); err != nil {
		return nil, err
	}
	return schedules, nil
}

func (g *GRPCMarshaller) GetSchedule(uuid string, opts ...*Opts) (*model.Schedule, error) {
	api := fmt.Sprintf("/api/schedules/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var schedule *model.Schedule
	if err = json.Unmarshal(res, &schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (g *GRPCMarshaller) GetOneScheduleByArgs(opts ...*Opts) (*model.Schedule, error) {
	api := "/api/schedules/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var schedule *model.Schedule
	if err = json.Unmarshal(res, &schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (g *GRPCMarshaller) UpdateSchedule(uuid string, body *model.Schedule, opts ...*Opts) (*model.Schedule, error) {
	api := fmt.Sprintf("/api/schedules/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var schedule *model.Schedule
	if err = json.Unmarshal(res, &schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (g *GRPCMarshaller) ScheduleWrite(uuid string, body *dto.ScheduleData, opts ...*Opts) (*dto.ScheduleData, error) {
	api := fmt.Sprintf("/api/schedules/%s/write", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var scheduleData *dto.ScheduleData
	if err = json.Unmarshal(res, &scheduleData); err != nil {
		return nil, err
	}
	return scheduleData, nil
}

func (g *GRPCMarshaller) UpdateScheduleAllProps(uuid string, body *model.Schedule, opts ...*Opts) (*model.Schedule, error) {
	api := fmt.Sprintf("/api/schedules/%s/all-props", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
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

func (g *GRPCMarshaller) DeleteSchedule(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/schedules/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
