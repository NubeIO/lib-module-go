package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateAlert(body *model.Alert, opts ...*Opts) (*model.Alert, error) {
	api := "/api/alerts"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var alert *model.Alert
	err = json.Unmarshal(res, &alert)
	if err != nil {
		return nil, err
	}
	return alert, nil
}

func (g *GRPCMarshaller) GetAlerts(opts ...*Opts) ([]*model.Alert, error) {
	api := "/api/alerts"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var alerts []*model.Alert
	err = json.Unmarshal(res, &alerts)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (g *GRPCMarshaller) GetAlert(uuid string, opts ...*Opts) (*model.Alert, error) {
	api := fmt.Sprintf("/api/alerts/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var alert *model.Alert
	err = json.Unmarshal(res, &alert)
	if err != nil {
		return nil, err
	}
	return alert, nil
}

func (g *GRPCMarshaller) UpdateAlertStatus(uuid string, body *dto.AlertStatus, opts ...*Opts) (*model.Alert, error) {
	api := fmt.Sprintf("/api/alerts/%s/status", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var alert *model.Alert
	err = json.Unmarshal(res, &alert)
	if err != nil {
		return nil, err
	}
	return alert, nil
}

func (g *GRPCMarshaller) UpdateAlertTeams(uuid string, teamUUIDs []*string, opts ...*Opts) ([]*model.AlertTeam, error) {
	api := fmt.Sprintf("/api/alerts/%s/teams", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, teamUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var alertTeams []*model.AlertTeam
	err = json.Unmarshal(res, &alertTeams)
	if err != nil {
		return nil, err
	}
	return alertTeams, nil
}

func (g *GRPCMarshaller) UpsertAlertMetaTags(uuid string, body []*model.AlertMetaTag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/alerts/%s/meta-tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertAlertTags(uuid string, body []*model.Tag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/alerts/%s/tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteAlert(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/alerts/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteAlertTransaction(transactionUUID string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/alerts/transactions/%s", transactionUUID)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
