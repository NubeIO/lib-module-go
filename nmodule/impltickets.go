package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateTicket(body *model.Ticket, opts ...*Opts) (*model.Ticket, error) {
	api := "/api/tickets"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var ticket *model.Ticket
	err = json.Unmarshal(res, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (g *GRPCMarshaller) GetTickets(opts ...*Opts) ([]*model.Ticket, error) {
	api := "/api/tickets"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var tickets []*model.Ticket
	err = json.Unmarshal(res, &tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (g *GRPCMarshaller) GetTicket(uuid string, opts ...*Opts) (*model.Ticket, error) {
	api := fmt.Sprintf("/api/tickets/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var ticket *model.Ticket
	err = json.Unmarshal(res, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (g *GRPCMarshaller) UpdateTicket(uuid string, body *model.Ticket, opts ...*Opts) (*model.Ticket, error) {
	api := fmt.Sprintf("/api/tickets/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var ticket *model.Ticket
	err = json.Unmarshal(res, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (g *GRPCMarshaller) UpsertTicketPriority(uuid string, body *dto.TicketPriority, opts ...*Opts) error {
	api := fmt.Sprintf("/api/tickets/%s/priority", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpsertTicketStatus(uuid string, body *dto.TicketStatus, opts ...*Opts) error {
	api := fmt.Sprintf("/api/tickets/%s/status", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpsertTicketTeams(uuid string, teamUUIDs []*string, opts ...*Opts) ([]*model.TicketTeam, error) {
	api := fmt.Sprintf("/api/tickets/%s/teams", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, teamUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var ticketTeam []*model.TicketTeam
	err = json.Unmarshal(res, &ticketTeam)
	if err != nil {
		return nil, err
	}
	return ticketTeam, nil
}

func (g *GRPCMarshaller) UpsertTicketMembers(uuid string, memberUUIDs []*string, opts ...*Opts) ([]*model.TicketMember, error) {
	api := fmt.Sprintf("/api/tickets/%s/members", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PUT, api, memberUUIDs, opts...)
	if err != nil {
		return nil, err
	}
	var teamMembers []*model.TicketMember
	err = json.Unmarshal(res, &teamMembers)
	if err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func (g *GRPCMarshaller) DeleteTicket(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/tickets/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
