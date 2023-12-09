package shared

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

type Marshaller interface {
	CreateNetwork(body *model.Network) (*model.Network, error)
	GetNetworks(args nargs.Args) ([]*model.Network, error)
	GetNetwork(uuid string, args nargs.Args) (*model.Network, error)
	GetNetworkByName(networkName string, args nargs.Args) (*model.Network, error)
	GetNetworkByPlugin(pluginUUID string, args nargs.Args) (*model.Network, error)
	GetOneNetworkByArgs(args nargs.Args) (*model.Network, error)
	GetNetworksByPlugin(pluginUUID string, args nargs.Args) ([]*model.Network, error)
	GetNetworksByPluginName(pluginName string, args nargs.Args) ([]*model.Network, error)
	UpdateNetwork(uuid string, body *model.Network) (*model.Network, error)
	UpdateNetworkErrors(uuid string, body *model.Network) error
	UpdateNetworkDescendantsErrors(networkUUID, message, messageLevel, messageCode string, withPoints bool) error
	ClearNetworkDescendantsErrors(networkUUID string, withPoints bool) error
	DeleteNetwork(uuid string) error

	CreateDevice(body *model.Device) (*model.Device, error)
	GetDevices(args nargs.Args) ([]*model.Device, error)
	GetDevice(uuid string, args nargs.Args) (*model.Device, error)
	GetDeviceByName(networkName, deviceName string, args nargs.Args) (*model.Device, error)
	GetOneDeviceByArgs(args nargs.Args) (*model.Device, error)
	UpdateDevice(uuid string, body *model.Device) (*model.Device, error)
	UpdateDeviceErrors(uuid string, body *model.Device) error
	UpdateDeviceDescendantsErrors(deviceUUID, message, messageLevel, messageCode string) error
	ClearDeviceDescendantsErrors(deviceUUID string) error
	DeleteDevice(uuid string) error

	CreatePoint(body *model.Point) (*model.Point, error)
	GetPoints(args nargs.Args) ([]*model.Point, error)
	GetPoint(uuid string, args nargs.Args) (*model.Point, error)
	GetPointByName(networkName, deviceName, pointName string, args nargs.Args) (*model.Point, error)
	GetOnePointByArgs(args nargs.Args) (*model.Point, error)
	UpdatePoint(uuid string, body *model.Point, args nargs.Args) (*model.Point, error)
	UpdatePointErrors(uuid string, body *model.Point) error
	UpdatePointSuccess(uuid string, body *model.Point) error
	PointWrite(uuid string, body *model.PointWriter) (*model.PointWriteResponse, error)
	DeletePoint(uuid string) error

	GetSchedules() ([]*model.Schedule, error)
	UpdateScheduleAllProps(uuid string, body *model.Schedule) (*model.Schedule, error)

	GetHosts(args nargs.Args) ([]*model.Host, error)
	CloneHostThingsToCloud(hostUUID string) error

	GetPlugin(pluginUUID string, args nargs.Args) (*model.Plugin, error)
	GetPluginByName(name string, args nargs.Args) (*model.Plugin, error)
	CreateModuleDir(name string) (*string, error)

	CreateBulkHistory(histories []*model.History) (bool, error)
	CreateBulkPointHistory(histories []*model.PointHistory) (bool, error)
	GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string) (*model.History, error)
	GetPointHistoriesMissingTimestamps(pointUUID string) ([]string, error)

	GetHistoriesForPostgresSync(lastSyncId int) ([]*model.History, error)
	GetPointsForPostgresSync() ([]*model.PointForPostgresSync, error)
	GetNetworksTagsForPostgresSync() ([]*model.NetworkTagForPostgresSync, error)
	GetDevicesTagsForPostgresSync() ([]*model.DeviceTagForPostgresSync, error)
	GetPointsTagsForPostgresSync() ([]*model.PointTagForPostgresSync, error)
	GetNetworksMetaTagsForPostgresSync() ([]*model.NetworkMetaTag, error)
	GetDevicesMetaTagsForPostgresSync() ([]*model.DeviceMetaTag, error)
	GetPointsMetaTagsForPostgresSync() ([]*model.PointMetaTag, error)
	GetLastSyncHistoryIdForPostgresSync() (int, error)
	UpdateLastSyncHistoryRowForPostgresSync(log *model.HistoryPostgresLog) (*model.HistoryPostgresLog, error)

	GetHistoryLogByHostUUID(hostUUID string) (*model.HistoryLog, error)
	UpdateBulkHistoryLogs(logs []*model.HistoryLog) (bool, error)

	Publish(topic string, qos model.QOS, retain bool, payload string) error
	PublishNonBuffer(topic string, qos model.QOS, retain bool, payload string) error
}

func New(dbHelper DBHelper) *GRPCMarshaller {
	return &GRPCMarshaller{DbHelper: dbHelper}
}

type GRPCMarshaller struct {
	DbHelper DBHelper
}

func (g *GRPCMarshaller) CallDBHelperWithParser(method http.Method, api string, args nargs.Args, body interface{}) (
	[]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return g.DbHelper.CallDBHelper(method, api, args, b)
}
