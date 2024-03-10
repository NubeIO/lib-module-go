package nmodule

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) SendEmail(body *model.Email, opts ...*Opts) (*model.Email, error) {
	api := "/api/emails/send"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var email *model.Email
	err = json.Unmarshal(res, &email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (g *GRPCMarshaller) GetAttachmentDir(opts ...*Opts) (*string, error) {
	api := "/api/emails/attachment-dir"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var dir *string
	err = json.Unmarshal(res, &dir)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
