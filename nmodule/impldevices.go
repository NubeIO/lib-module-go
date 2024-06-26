package nmodule

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"strconv"
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

func (g *GRPCMarshaller) GetDevices(body *dto.Filter, opts ...*Opts) ([]*model.Device, error) {
	api := "/api/devices"
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, body, opts...)
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

func (g *GRPCMarshaller) CountDevices(body *dto.Filter, opts ...*Opts) (int, error) {
	api := fmt.Sprintf("/api/devices/count")
	res, err := g.CallDBHelperWithParser(nhttp.GET, api, body, opts...)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(res))
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

func (g *GRPCMarshaller) UpdateDeviceFault(uuid string, body *model.CommonFault, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/fault", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) UpdateDeviceDescendantsFault(deviceUUID string, body *model.CommonFault, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/fault/descendants", deviceUUID)
	_, err := g.CallDBHelperWithParser(nhttp.PATCH, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertDeviceMetaTags(uuid string, body []*model.DeviceMetaTag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/meta-tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) UpsertDeviceTags(uuid string, body []*model.Tag, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s/tags", uuid)
	_, err := g.CallDBHelperWithParser(nhttp.PUT, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteDevice(uuid string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/%s", uuid)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteOneDeviceByArgs(opts ...*Opts) error {
	api := "/api/devices/one/args"
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}

func (g *GRPCMarshaller) DeleteDeviceByName(name string, opts ...*Opts) error {
	api := fmt.Sprintf("/api/devices/name/%s", name)
	_, err := g.DbHelper.CallDBHelper(nhttp.DELETE, api, nil, opts...)
	return err
}
