package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateTicketComment(body *model.TicketComment, opts ...*Opts) (*model.TicketComment, error) {
	api := "/api/tickets/comments"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var ticketComment *model.TicketComment
	err = json.Unmarshal(res, &ticketComment)
	if err != nil {
		return nil, err
	}
	return ticketComment, nil
}

func (g *GRPCMarshaller) UpdateTicketComment(uuid string, body *model.TicketComment, opts ...*Opts) (*model.TicketComment, error) {
	api := fmt.Sprintf("/api/tickets/comments/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var ticketComment *model.TicketComment
	err = json.Unmarshal(res, &ticketComment)
	if err != nil {
		return nil, err
	}
	return ticketComment, nil
}

func (g *GRPCMarshaller) DeleteTicketComment(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/tickets/comments/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
