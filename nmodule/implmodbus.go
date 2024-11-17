package nmodule

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
)

func (g *GRPCMarshaller) ModbusWriteValueManually(body *dto.ManualPointWriteValue, opts ...*Opts) (*dto.ManualPointWriteValueResponse, error) {
	api := "/api/modules/module-core-modbus/api/manual-write-value"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var manualPointWriteValueResponse *dto.ManualPointWriteValueResponse
	err = json.Unmarshal(res, &manualPointWriteValueResponse)
	if err != nil {
		return nil, err
	}
	return manualPointWriteValueResponse, nil
}

func (g *GRPCMarshaller) ModbusSetCommissioning(body *dto.CommissioningToolRequest, opts ...*Opts) (*dto.CommissioningToolResponse, error) {
	api := "/api/modules/module-core-modbus/api/commissioning"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var commissioningToolResp *dto.CommissioningToolResponse
	err = json.Unmarshal(res, &commissioningToolResp)
	if err != nil {
		return nil, err
	}
	return commissioningToolResp, nil
}
