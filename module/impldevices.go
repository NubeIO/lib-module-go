package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) CreateDevice(body *model.Device) (*model.Device, error) {
	api := "/api/devices"
	res, err := g.CallDBHelperWithParser(http.POST, api, nargs.Args{}, body)
	if err != nil {
		return nil, err
	}
	var device *model.Device
	err = json.Unmarshal(res, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GRPCMarshaller) GetDevices(args nargs.Args) ([]*model.Device, error) {
	api := "/api/devices"
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var devices []*model.Device
	err = json.Unmarshal(res, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (g *GRPCMarshaller) GetDevice(uuid string, args nargs.Args) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var device *model.Device
	err = json.Unmarshal(res, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GRPCMarshaller) GetDeviceByName(networkName, deviceName string, args nargs.Args) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/name/%s/%s", networkName, deviceName)
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var device *model.Device
	err = json.Unmarshal(res, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GRPCMarshaller) GetOneDeviceByArgs(args nargs.Args) (*model.Device, error) {
	api := "/api/devices/one/args"
	res, err := g.DbHelper.CallDBHelper(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var device *model.Device
	err = json.Unmarshal(res, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GRPCMarshaller) UpdateDevice(uuid string, body *model.Device) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	res, err := g.CallDBHelperWithParser(http.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return nil, err
	}
	var device *model.Device
	err = json.Unmarshal(res, &device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GRPCMarshaller) UpdateDeviceErrors(uuid string, body *model.Device) error {
	api := fmt.Sprintf("/api/devices/%s/error", uuid)
	_, err := g.CallDBHelperWithParser(http.PATCH, api, nargs.Args{}, body)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdateDeviceDescendantsErrors(deviceUUID, message, messageLevel, messageCode string) error {
	api := fmt.Sprintf("/api/devices/%s/error/descendants", deviceUUID)
	device := &model.Device{
		CommonFault: model.CommonFault{
			Message:      message,
			MessageLevel: messageLevel,
			MessageCode:  messageCode,
		},
	}
	_, err := g.CallDBHelperWithParser(http.PATCH, api, nargs.Args{}, device)
	return err
}

func (g *GRPCMarshaller) ClearDeviceDescendantsErrors(deviceUUID string) error {
	api := fmt.Sprintf("/api/devices/%s/error/descendants", deviceUUID)
	_, err := g.DbHelper.CallDBHelper(http.DELETE, api, nargs.Args{}, nil)
	return err
}

func (g *GRPCMarshaller) DeleteDevice(uuid string) error {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(http.DELETE, api, nargs.Args{}, nil)
	return err
}
