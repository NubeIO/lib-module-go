package nmodule

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
)

func (g *GRPCMarshaller) PostgresRawQuery(body *dto.QueryBody, opts ...*Opts) (*dto.QueryResponse, error) {
	api := "/api/postgres/query-data"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var history *dto.QueryResponse
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}
