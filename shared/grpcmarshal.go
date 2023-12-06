package shared

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

type Marshaller interface {
	GetNetworks(args nargs.Args) ([]*model.Network, error)
	GetNetwork(uuid string, args nargs.Args) (*model.Network, error)
	GetNetworkByName(networkName string, args nargs.Args) (*model.Network, error)
	GetNetworkByPlugin(pluginUUID string, args nargs.Args) (*model.Network, error)
	GetOneNetworkByArgs(args nargs.Args) (*model.Network, error)
	GetNetworksByPluginName(pluginName string, args nargs.Args) ([]*model.Network, error)
	GetNetworksByPlugin(pluginUUID string, args nargs.Args) ([]*model.Network, error)

	GetDevices(args nargs.Args) ([]*model.Device, error)
	GetDevice(uuid string, args nargs.Args) (*model.Device, error)
	GetDeviceByName(networkName, deviceName string, args nargs.Args) (*model.Device, error)
	GetOneDeviceByArgs(args nargs.Args) (*model.Device, error)

	GetPoints(args nargs.Args) ([]*model.Point, error)
	GetPoint(uuid string, args nargs.Args) (*model.Point, error)
	GetPointByName(networkName, deviceName, pointName string, args nargs.Args) (*model.Point, error)
	GetOnePointByArgs(args nargs.Args) (*model.Point, error)

	// CreateBulkPointsHistories(histories []*model.PointHistory) (bool, error)
	//
	// CreateNetwork(body *model.Network) (*model.Network, error)
	// CreateDevice(body *model.Device) (*model.Device, error)
	// CreatePoint(body *model.Point) (*model.Point, error)
	//
	// UpdateNetwork(uuid string, body *model.Network) (*model.Network, error)
	// UpdateDevice(uuid string, body *model.Device) (*model.Device, error)
	// UpdatePoint(uuid string, body *model.Point, opts ...interfaces.UpdatePointOpts) (*model.Point, error)
	//
	// UpdateNetworkErrors(uuid string, body *model.Network) error
	// UpdateDeviceErrors(uuid string, body *model.Device) error
	// UpdatePointErrors(uuid string, body *model.Point) error
	// UpdatePointSuccess(uuid string, body *model.Point) error
	//
	// DeleteNetwork(uuid string) error
	// DeleteDevice(uuid string) error
	// DeletePoint(uuid string) error
	//
	// PointWrite(uuid string, pointWriter *model.PointWriter) (*common.PointWriteResponse, error)
	//
	// GetSchedules() ([]*model.Schedule, error)
	// UpdateScheduleAllProps(uuid string, body *model.Schedule) (*model.Schedule, error)
	//
	// GetPlugin(pluginUUID string, args nargs.Args) (*model.Plugin, error)
	// GetPluginByName(name string, args nargs.Args) (*model.Plugin, error)
	// SetErrorsForAllDevicesOnNetwork(networkUUID, message, messageLevel, messageCode string, doPoints bool) error
	// ClearErrorsForAllDevicesOnNetwork(networkUUID string, doPoints bool) error
	// SetErrorsForAllPointsOnDevice(deviceUUID, message, messageLevel, messageCode string) error
	// ClearErrorsForAllPointsOnDevice(deviceUUID string) error
	// WizardNewNetworkDevicePoint(plugin string, network *model.Network, device *model.Device, point *model.Point) (bool, error)
	// DeviceNameExistsInNetwork(deviceName, networkUUID string) (*model.Device, bool)
	//
	// GetHosts(args nargs.Args) ([]*model.Host, error)
	// GetHostsWithHistoryEnabled(args nargs.Args) ([]*model.Host, error)
	// GetHistoryLogByHostUUID(hostUUID string) (*model.HistoryLog, error)
	// CloneHostThingsToCloud(host *model.Host) error
	// CreateBulkHistory(histories []*model.History) (bool, error)
	// UpdateBulkHistoryLogs(logs []*model.HistoryLog) (bool, error)
	//
	// GetHistoryPostgresLogLastSyncHistoryId() (int, error)
	// GetHistoriesForPostgresSync(lastSyncId int) ([]*model.History, error)
	// UpdateHistoryPostgresLog(log *model.HistoryPostgresLog) (*model.HistoryPostgresLog, error)
	// GetPointsForPostgresSync() ([]*interfaces.PointForPostgresSync, error)
	// GetNetworksTagsForPostgresSync() ([]*interfaces.NetworkTagForPostgresSync, error)
	// GetDevicesTagsForPostgresSync() ([]*interfaces.DeviceTagForPostgresSync, error)
	// GetPointsTagsForPostgresSync() ([]*interfaces.PointTagForPostgresSync, error)
	// GetNetworksMetaTagsForPostgresSync() ([]*model.NetworkMetaTag, error)
	// GetDevicesMetaTagsForPostgresSync() ([]*model.DeviceMetaTag, error)
	// GetPointsMetaTagsForPostgresSync() ([]*model.PointMetaTag, error)
	// GetPointHistoriesMissingTimestamps(pointUUID string) ([]string, error)
	// GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string) (*model.History, error)
	//
	// Publish(topic string, qos mqttclient.QOS, retain bool, payload string) error
	// PublishNonBuffer(topic string, qos mqttclient.QOS, retain bool, payload interface{}) error
}

type GRPCMarshaller struct {
	DbHelper DBHelper
}

func (g *GRPCMarshaller) GetNetworks(args nargs.Args) ([]*model.Network, error) {
	api := "/api/networks"
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) GetNetwork(uuid string, args nargs.Args) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/%s", uuid)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworkByName(networkName string, args nargs.Args) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/name/%s", networkName)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworkByPlugin(pluginUUID string, args nargs.Args) (*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s", pluginUUID)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetOneNetworkByArgs(args nargs.Args) (*model.Network, error) {
	api := "/api/networks/one/args"
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var network *model.Network
	err = json.Unmarshal(res, &network)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GRPCMarshaller) GetNetworksByPlugin(pluginUUID string, args nargs.Args) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-uuid/%s/all", pluginUUID)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) GetNetworksByPluginName(pluginName string, args nargs.Args) ([]*model.Network, error) {
	api := fmt.Sprintf("/api/networks/plugin-name/%s/all", pluginName)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var networks []*model.Network
	err = json.Unmarshal(res, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GRPCMarshaller) GetDevices(args nargs.Args) ([]*model.Device, error) {
	api := "/api/devices"
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
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
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
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
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
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
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
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

func (g *GRPCMarshaller) GetPoints(args nargs.Args) ([]*model.Point, error) {
	api := "/api/points"
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var points []*model.Point
	err = json.Unmarshal(res, &points)
	if err != nil {
		return nil, err
	}
	return points, nil
}

func (g *GRPCMarshaller) GetPoint(uuid string, args nargs.Args) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/%s", uuid)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) GetPointByName(networkName, deviceName, pointName string, args nargs.Args) (*model.Point, error) {
	api := fmt.Sprintf("/api/points/name/%s/%s/%s", networkName, deviceName, pointName)
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GRPCMarshaller) GetOnePointByArgs(args nargs.Args) (*model.Point, error) {
	api := "/api/points/one/args"
	res, err := g.DbHelper.Call(http.GET, api, args, nil)
	if err != nil {
		return nil, err
	}
	var point *model.Point
	err = json.Unmarshal(res, &point)
	if err != nil {
		return nil, err
	}
	return point, nil
}

// func (g *GRPCMarshaller) CreateBulkPointsHistories(histories []*model.PointHistory) (bool, error) {
// 	hist, err := json.Marshal(histories)
// 	if err != nil {
// 		return false, err
// 	}
// 	_, err = g.DbHelper.Post("bulk_points_histories", hist)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (g *GRPCMarshaller) GetPointByName(networkName, deviceName, pointName string, args nargs.Args) (*model.Point,
// 	error) {
// 	serializedArgs, err := args.SerializeArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	uuid := strings.Join([]string{networkName, deviceName, pointName}, common.Separator)
// 	res, err := g.DbHelper.Get("point_by_name", uuid, serializedArgs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var point *model.Point
// 	err = json.Unmarshal(res, &point)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return point, nil
// }
//
// func (g *GRPCMarshaller) CreateNetwork(body *model.Network) (*model.Network, error) {
// 	net, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Post("networks", net)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var network *model.Network
// 	err = json.Unmarshal(res, &network)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return network, nil
// }
//
// func (g *GRPCMarshaller) CreateDevice(body *model.Device) (*model.Device, error) {
// 	dev, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Post("devices", dev)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var device *model.Device
// 	err = json.Unmarshal(res, &device)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return device, nil
// }
//
// func (g *GRPCMarshaller) CreatePoint(body *model.Point) (*model.Point, error) {
// 	pnt, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Post("points", pnt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var point *model.Point
// 	err = json.Unmarshal(res, &point)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return point, nil
// }
//
// func (g *GRPCMarshaller) UpdateNetwork(uuid string, body *model.Network) (*model.Network, error) {
// 	net, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Patch("networks", uuid, net)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var network *model.Network
// 	err = json.Unmarshal(res, &network)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return network, nil
// }
//
// func (g *GRPCMarshaller) UpdateDevice(uuid string, body *model.Device) (*model.Device, error) {
// 	dev, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Patch("devices", uuid, dev)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var device *model.Device
// 	err = json.Unmarshal(res, &device)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return device, nil
// }
//
// func (g *GRPCMarshaller) UpdatePoint(uuid string, body *model.Point, opts ...interfaces.UpdatePointOpts) (*model.Point, error) {
// 	var res []byte
// 	var err error
// 	pnt, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(opts) > 0 {
// 		optsByte, err := json.Marshal(opts)
// 		if err != nil {
// 			return nil, err
// 		}
// 		res, err = g.DbHelper.PatchWithOpts("points", uuid, pnt, optsByte)
// 	} else {
// 		res, err = g.DbHelper.Patch("points", uuid, pnt)
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	var point *model.Point
// 	err = json.Unmarshal(res, &point)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return point, nil
// }
//
// func (g *GRPCMarshaller) UpdateNetworkErrors(uuid string, body *model.Network) error {
// 	dev, err := json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = g.DbHelper.Patch("network_errors", uuid, dev)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) UpdateDeviceErrors(uuid string, body *model.Device) error {
// 	dev, err := json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = g.DbHelper.Patch("device_errors", uuid, dev)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) UpdatePointErrors(uuid string, body *model.Point) error {
// 	point, err := json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = g.DbHelper.Patch("point_errors", uuid, point)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) UpdatePointSuccess(uuid string, body *model.Point) error {
// 	point, err := json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = g.DbHelper.Patch("point_success", uuid, point)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) DeleteNetwork(uuid string) error {
// 	_, err := g.DbHelper.Delete("networks", uuid)
// 	return err
// }
//
// func (g *GRPCMarshaller) DeleteDevice(uuid string) error {
// 	_, err := g.DbHelper.Delete("devices", uuid)
// 	return err
// }
//
// func (g *GRPCMarshaller) DeletePoint(uuid string) error {
// 	_, err := g.DbHelper.Delete("points", uuid)
// 	return err
// }
//
// func (g *GRPCMarshaller) PointWrite(uuid string, pointWriter *model.PointWriter) (*common.PointWriteResponse, error) {
// 	pw, err := json.Marshal(pointWriter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Patch("point_write", uuid, pw)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pwr *common.PointWriteResponse
// 	err = json.Unmarshal(res, &pwr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pwr, nil
// }
//
// func (g *GRPCMarshaller) GetSchedules() ([]*model.Schedule, error) {
// 	res, err := g.DbHelper.GetWithoutParam("schedules", "")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var schedules []*model.Schedule
// 	if err = json.Unmarshal(res, &schedules); err != nil {
// 		return nil, err
// 	}
//
// 	return schedules, nil
// }
//
// func (g *GRPCMarshaller) UpdateScheduleAllProps(uuid string, body *model.Schedule) (*model.Schedule, error) {
// 	sch, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Patch("schedules", uuid, sch)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var schedule *model.Schedule
// 	err = json.Unmarshal(res, &schedule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return schedule, nil
// }
//
// func (g *GRPCMarshaller) GetPlugin(pluginUUID string, args nargs.Args) (*model.Plugin, error) {
// 	serializedArgs, err := args.SerializeArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Get("plugin_by_uuid", pluginUUID, serializedArgs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pluginConf *model.Plugin
// 	err = json.Unmarshal(res, &pluginConf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pluginConf, nil
// }
//
// func (g *GRPCMarshaller) GetPluginByName(name string, args nargs.Args) (*model.Plugin, error) {
// 	serializedArgs, err := args.SerializeArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.Get("plugin_by_name", name, serializedArgs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pluginConf *model.Plugin
// 	err = json.Unmarshal(res, &pluginConf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pluginConf, nil
// }
//
// func (g *GRPCMarshaller) SetErrorsForAllDevicesOnNetwork(networkUUID, message, messageLevel, messageCode string, doPoints bool) error {
// 	err := g.DbHelper.SetErrorsForAll("devices_on_network", networkUUID, message, messageLevel, messageCode, doPoints)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) ClearErrorsForAllDevicesOnNetwork(networkUUID string, doPoints bool) error {
// 	err := g.DbHelper.ClearErrorsForAll("devices_on_network", networkUUID, doPoints)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) SetErrorsForAllPointsOnDevice(deviceUUID, message, messageLevel, messageCode string) error {
// 	err := g.DbHelper.SetErrorsForAll("points_on_device", deviceUUID, message, messageLevel, messageCode, false)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) ClearErrorsForAllPointsOnDevice(deviceUUID string) error {
// 	err := g.DbHelper.ClearErrorsForAll("points_on_device", deviceUUID, false)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) WizardNewNetworkDevicePoint(plugin string, network *model.Network, device *model.Device, point *model.Point) (bool, error) {
// 	net, err := json.Marshal(network)
// 	if err != nil {
// 		return false, err
// 	}
// 	dev, err := json.Marshal(device)
// 	if err != nil {
// 		return false, err
// 	}
// 	pnt, err := json.Marshal(point)
// 	if err != nil {
// 		return false, err
// 	}
// 	chk, err := g.DbHelper.WizardNewNetworkDevicePoint(plugin, net, dev, pnt)
// 	return chk, err
// }
//
// func (g *GRPCMarshaller) DeviceNameExistsInNetwork(deviceName, networkUUID string) (*model.Device, bool) {
// 	network, err := g.GetNetwork(networkUUID, nargs.Args{})
// 	if err != nil {
// 		return nil, false
// 	}
// 	for _, dev := range network.Devices {
// 		if dev.Name == deviceName {
// 			return dev, true
// 		}
// 	}
// 	return nil, false
// }
//
// func (g *GRPCMarshaller) GetHosts(args nargs.Args) ([]*model.Host, error) {
// 	serializedArgs, err := args.SerializeArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.GetWithoutParam("hosts", serializedArgs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var hosts []*model.Host
// 	err = json.Unmarshal(res, &hosts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return hosts, nil
// }
//
// func (g *GRPCMarshaller) GetHostsWithHistoryEnabled(args nargs.Args) ([]*model.Host, error) {
// 	serializedArgs, err := args.SerializeArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := g.DbHelper.GetWithoutParam("hosts_with_history", serializedArgs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var hosts []*model.Host
// 	err = json.Unmarshal(res, &hosts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return hosts, nil
// }
//
// func (g *GRPCMarshaller) GetHistoryLogByHostUUID(hostUUID string) (*model.HistoryLog, error) {
// 	res, err := g.DbHelper.Get("history_log_by_id", hostUUID, "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var historyLog *model.HistoryLog
// 	err = json.Unmarshal(res, &historyLog)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return historyLog, nil
// }
//
// func (g *GRPCMarshaller) CloneHostThingsToCloud(host *model.Host) error {
// 	hst, err := json.Marshal(host)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = g.DbHelper.Post("clone_host_things_to_cloud", hst)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) CreateBulkHistory(histories []*model.History) (bool, error) {
// 	hist, err := json.Marshal(histories)
// 	if err != nil {
// 		return false, err
// 	}
// 	_, err = g.DbHelper.Post("bulk_history", hist)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (g *GRPCMarshaller) UpdateBulkHistoryLogs(logs []*model.HistoryLog) (bool, error) {
// 	histLog, err := json.Marshal(logs)
// 	if err != nil {
// 		return false, err
// 	}
// 	_, err = g.DbHelper.PatchWithoutParam("bulk_history_logs", histLog)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (g *GRPCMarshaller) GetHistoryPostgresLogLastSyncHistoryId() (int, error) {
// 	res, err := g.DbHelper.GetWithoutParam("postgres_history_last_sync_id", "")
// 	if err != nil {
// 		return 0, err
// 	}
// 	var q int
// 	err = json.Unmarshal(res, &q)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return q, nil
// }
//
// func (g *GRPCMarshaller) GetHistoriesForPostgresSync(lastSyncId int) ([]*model.History, error) {
// 	lastSyncIdStr := strconv.Itoa(lastSyncId)
// 	res, err := g.DbHelper.Get("histories_for_postgres_sync", lastSyncIdStr, "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var history []*model.History
// 	err = json.Unmarshal(res, &history)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return history, nil
// }
//
// func (g *GRPCMarshaller) UpdateHistoryPostgresLog(log *model.HistoryPostgresLog) (*model.HistoryPostgresLog, error) {
// 	histLog, err := json.Marshal(log)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := g.DbHelper.PatchWithoutParam("history_postgres_log", histLog)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var updatedLog *model.HistoryPostgresLog
// 	err = json.Unmarshal(resp, &updatedLog)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return updatedLog, nil
// }
//
// func (g *GRPCMarshaller) GetPointsForPostgresSync() ([]*interfaces.PointForPostgresSync, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("points_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*interfaces.PointForPostgresSync
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetNetworksTagsForPostgresSync() ([]*interfaces.NetworkTagForPostgresSync, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("networks_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*interfaces.NetworkTagForPostgresSync
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetDevicesTagsForPostgresSync() ([]*interfaces.DeviceTagForPostgresSync, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("devices_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*interfaces.DeviceTagForPostgresSync
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetPointsTagsForPostgresSync() ([]*interfaces.PointTagForPostgresSync, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("points_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*interfaces.PointTagForPostgresSync
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetNetworksMetaTagsForPostgresSync() ([]*model.NetworkMetaTag, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("networks_meta_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*model.NetworkMetaTag
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetDevicesMetaTagsForPostgresSync() ([]*model.DeviceMetaTag, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("devices_meta_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*model.DeviceMetaTag
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetPointsMetaTagsForPostgresSync() ([]*model.PointMetaTag, error) {
// 	resp, err := g.DbHelper.GetWithoutParam("points_meta_tags_for_postgres_sync", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var r []*model.PointMetaTag
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) GetPointHistoriesMissingTimestamps(pointUUID string) ([]string, error) {
// 	res, err := g.DbHelper.Get("point_histories_missing_timestamps", pointUUID, "")
// 	var missingTimestamps []string
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(res, &missingTimestamps)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return missingTimestamps, nil
// }
//
// func (g *GRPCMarshaller) GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string) (*model.History, error) {
// 	uuid := strings.Join([]string{hostUUID, pointUUID}, common.Separator)
// 	resp, err := g.DbHelper.Get("latest_history_by_host_and_point_uuid", uuid, "")
// 	var r *model.History
// 	err = json.Unmarshal(resp, &r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r, nil
// }
//
// func (g *GRPCMarshaller) Publish(topic string, qos mqttclient.QOS, retain bool, payload string) error {
// 	qosBytes, err := json.Marshal(qos)
// 	if err != nil {
// 		return err
// 	}
// 	err = g.DbHelper.MQTTPublish(topic, qosBytes, retain, payload)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (g *GRPCMarshaller) PublishNonBuffer(topic string, qos mqttclient.QOS, retain bool, payload interface{}) error {
// 	var qosBytes, body []byte
// 	var err error
// 	qosBytes, err = json.Marshal(qos)
// 	if err != nil {
// 		return err
// 	}
// 	body, err = json.Marshal(body)
// 	if err != nil {
// 		return err
// 	}
// 	err = g.DbHelper.MQTTPublishNonBuffer(topic, qosBytes, retain, body)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
