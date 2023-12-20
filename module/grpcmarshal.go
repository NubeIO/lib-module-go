package module

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

type Opts struct {
	Args     *nargs.Args
	HostUUID *string
}

type Marshaller interface {
	CreateNetwork(body *model.Network, opts ...*Opts) (*model.Network, error)
	GetNetworks(opts ...*Opts) ([]*model.Network, error)
	GetNetwork(uuid string, opts ...*Opts) (*model.Network, error)
	GetNetworkByName(networkName string, opts ...*Opts) (*model.Network, error)
	GetNetworkByPlugin(pluginUUID string, opts ...*Opts) (*model.Network, error)
	GetOneNetworkByArgs(opts ...*Opts) (*model.Network, error)
	GetNetworksByPlugin(pluginUUID string, opts ...*Opts) ([]*model.Network, error)
	GetNetworksByPluginName(pluginName string, opts ...*Opts) ([]*model.Network, error)
	UpdateNetwork(uuid string, body *model.Network, opts ...*Opts) (*model.Network, error)
	UpdateNetworkErrors(uuid string, body *model.Network, opts ...*Opts) error
	UpdateNetworkDescendantsErrors(networkUUID, message, messageLevel, messageCode string, withPoints bool, opts ...*Opts) error
	ClearNetworkDescendantsErrors(networkUUID string, withPoints bool, opts ...*Opts) error
	DeleteNetwork(uuid string, opts ...*Opts) error

	CreateDevice(body *model.Device, opts ...*Opts) (*model.Device, error)
	GetDevices(opts ...*Opts) ([]*model.Device, error)
	GetDevice(uuid string, opts ...*Opts) (*model.Device, error)
	GetDeviceByName(networkName, deviceName string, opts ...*Opts) (*model.Device, error)
	GetOneDeviceByArgs(opts ...*Opts) (*model.Device, error)
	UpdateDevice(uuid string, body *model.Device, opts ...*Opts) (*model.Device, error)
	UpdateDeviceErrors(uuid string, body *model.Device, opts ...*Opts) error
	UpdateDeviceDescendantsErrors(deviceUUID, message, messageLevel, messageCode string, opts ...*Opts) error
	ClearDeviceDescendantsErrors(deviceUUID string, opts ...*Opts) error
	DeleteDevice(uuid string, opts ...*Opts) error

	CreatePoint(body *model.Point, opts ...*Opts) (*model.Point, error)
	GetPoints(opts ...*Opts) ([]*model.Point, error)
	GetPoint(uuid string, opts ...*Opts) (*model.Point, error)
	GetPointByName(networkName, deviceName, pointName string, opts ...*Opts) (*model.Point, error)
	GetOnePointByArgs(opts ...*Opts) (*model.Point, error)
	UpdatePoint(uuid string, body *model.Point, opts ...*Opts) (*model.Point, error)
	UpdatePointErrors(uuid string, body *model.Point, opts ...*Opts) error
	UpdatePointSuccess(uuid string, body *model.Point, opts ...*Opts) error
	PointWrite(uuid string, body *model.PointWriter, opts ...*Opts) (*model.PointWriteResponse, error)
	DeletePoint(uuid string, opts ...*Opts) error

	GetSchedules(opts ...*Opts) ([]*model.Schedule, error)
	UpdateScheduleAllProps(uuid string, body *model.Schedule, opts ...*Opts) (*model.Schedule, error)

	GetHosts(opts ...*Opts) ([]*model.Host, error)
	CloneHostThingsToCloud(hostUUID string, opts ...*Opts) error

	GetPlugin(pluginUUID string, opts ...*Opts) (*model.Plugin, error)
	GetPluginByName(name string, opts ...*Opts) (*model.Plugin, error)
	UpdatePluginMessage(name string, body *model.Plugin, opts ...*Opts) error
	CreateModuleDir(name string, opts ...*Opts) (*string, error)

	CreateBulkHistory(histories []*model.History, opts ...*Opts) (bool, error)
	CreateBulkPointHistory(histories []*model.PointHistory, opts ...*Opts) (bool, error)
	GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string, opts ...*Opts) (*model.History, error)
	GetPointHistoriesMissingTimestamps(pointUUID string, opts ...*Opts) ([]string, error)

	GetHistoriesForPostgresSync(lastSyncId int, opts ...*Opts) ([]*model.History, error)
	GetPointsForPostgresSync(opts ...*Opts) ([]*model.PointForPostgresSync, error)
	GetNetworksTagsForPostgresSync(opts ...*Opts) ([]*model.NetworkTagForPostgresSync, error)
	GetDevicesTagsForPostgresSync(opts ...*Opts) ([]*model.DeviceTagForPostgresSync, error)
	GetPointsTagsForPostgresSync(opts ...*Opts) ([]*model.PointTagForPostgresSync, error)
	GetNetworksMetaTagsForPostgresSync(opts ...*Opts) ([]*model.NetworkMetaTag, error)
	GetDevicesMetaTagsForPostgresSync(opts ...*Opts) ([]*model.DeviceMetaTag, error)
	GetPointsMetaTagsForPostgresSync(opts ...*Opts) ([]*model.PointMetaTag, error)
	GetLastSyncHistoryIdForPostgresSync(opts ...*Opts) (int, error)
	UpdateLastSyncHistoryRowForPostgresSync(log *model.HistoryPostgresLog, opts ...*Opts) (*model.HistoryPostgresLog, error)

	GetHistoryLogByHostUUID(hostUUID string, opts ...*Opts) (*model.HistoryLog, error)
	UpdateBulkHistoryLogs(logs []*model.HistoryLog, opts ...*Opts) (bool, error)

	Publish(topic string, qos model.QOS, retain bool, payload string, opts ...*Opts) error
	PublishNonBuffer(topic string, qos model.QOS, retain bool, payload string, opts ...*Opts) error
}

func New(dbHelper DBHelper) *GRPCMarshaller {
	return &GRPCMarshaller{DbHelper: dbHelper}
}

type GRPCMarshaller struct {
	DbHelper DBHelper
}

func (g *GRPCMarshaller) CallDBHelperWithParser(method nhttp.Method, api string, body interface{}, opts ...*Opts) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return g.DbHelper.CallDBHelper(method, api, b, opts...)
}
