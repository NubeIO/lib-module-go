package module

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

func (g *GRPCMarshaller) CreateDevice(body *model.Device, opts ...*Opts) (*model.Device, error) {
	api := "/api/devices"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
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

func (g *GRPCMarshaller) GetDevices(opts ...*Opts) ([]*model.Device, error) {
	api := "/api/devices"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetDevice(uuid string, opts ...*Opts) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetDeviceByName(networkName, deviceName string, opts ...*Opts) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/name/%s/%s", networkName, deviceName)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetOneDeviceByArgs(opts ...*Opts) (*model.Device, error) {
	api := "/api/devices/one/args"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) UpdateDevice(uuid string, body *model.Device, opts ...*Opts) (*model.Device, error) {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
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

func (g *GRPCMarshaller) UpdateDeviceErrors(uuid string, body *model.Device, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/error", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdateDeviceDescendantsErrors(deviceUUID, message, messageLevel, messageCode string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/error/descendants", deviceUUID)
	device := &model.Device{
		CommonFault: model.CommonFault{
			Message:      message,
			MessageLevel: messageLevel,
			MessageCode:  messageCode,
		},
	}
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, device, opts...)
	return err
}

func (g *GRPCMarshaller) ClearDeviceDescendantsErrors(deviceUUID string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/error/descendants", deviceUUID)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteDevice(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
