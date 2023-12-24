package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateTeam(body *model.Team, opts ...*Opts) (*model.Team, error) {
	api := "/api/teams"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var team *model.Team
	err = json.Unmarshal(res, &team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (g *GRPCMarshaller) GetTeams(opts ...*Opts) ([]*model.Team, error) {
	api := "/api/teams"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	err = json.Unmarshal(res, &teams)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (g *GRPCMarshaller) GetTeam(uuid string, opts ...*Opts) (*model.Team, error) {
	api := fmt.Sprintf("/api/teams/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var team *model.Team
	err = json.Unmarshal(res, &team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (g *GRPCMarshaller) UpdateTeam(uuid string, body *model.Team, opts ...*Opts) (*model.Team, error) {
	api := fmt.Sprintf("/api/teams/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var team *model.Team
	err = json.Unmarshal(res, &team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (g *GRPCMarshaller) UpsertTeamMembers(uuid string, memberUUIDs []*string, opts ...*Opts) ([]*model.Member, error) {
	api := fmt.Sprintf("/api/teams/%s/members", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, memberUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var members []*model.Member
	err = json.Unmarshal(res, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (g *GRPCMarshaller) UpsertTeamViews(uuid string, viewUUIDs []*string, opts ...*Opts) ([]*model.TeamView, error) {
	api := fmt.Sprintf("/api/teams/%s/views", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, viewUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var views []*model.TeamView
	err = json.Unmarshal(res, &views)
	if err != nil {
		return nil, err
	}
	return views, nil
}

func (g *GRPCMarshaller) UpsertTeamMetaTags(uuid string, body []*model.TeamMetaTag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/teams/%s/meta-tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteTeam(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/teams/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
