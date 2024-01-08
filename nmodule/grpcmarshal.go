package nmodule

import (
	"encoding/json"
	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-networking/scanner"
	systats "github.com/NubeIO/lib-system"
	"github.com/NubeIO/lib-system/ostats"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
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
	GetOneNetworkByArgs(opts ...*Opts) (*model.Network, error)
	GetNetworkByPlugin(pluginUUID string, opts ...*Opts) (*model.Network, error)
	GetNetworksByPlugin(pluginUUID string, opts ...*Opts) ([]*model.Network, error)
	GetNetworkByPluginName(pluginName string, opts ...*Opts) (*model.Network, error)
	GetNetworksByPluginName(pluginName string, opts ...*Opts) ([]*model.Network, error)
	UpdateNetwork(uuid string, body *model.Network, opts ...*Opts) (*model.Network, error)
	CountNetworks(opts ...*Opts) (int, error)
	UpdateNetworkErrors(uuid string, body *model.Network, opts ...*Opts) error
	UpdateNetworkDescendantsErrors(networkUUID, message, messageLevel, messageCode string, withPoints bool, opts ...*Opts) error
	UpsertNetworkMetaTags(uuid string, body []*model.NetworkMetaTag, opts ...*Opts) error
	UpsertNetworkTags(uuid string, body []*model.Tag, opts ...*Opts) error
	DeleteNetwork(uuid string, opts ...*Opts) error
	DeleteOneNetworkByArgs(opts ...*Opts) error
	DeleteNetworkByName(name string, opts ...*Opts) error
	ClearNetworkDescendantsErrors(networkUUID string, withPoints bool, opts ...*Opts) error

	CreateDevice(body *model.Device, opts ...*Opts) (*model.Device, error)
	GetDevices(opts ...*Opts) ([]*model.Device, error)
	GetDevice(uuid string, opts ...*Opts) (*model.Device, error)
	GetDeviceByName(networkName, deviceName string, opts ...*Opts) (*model.Device, error)
	GetOneDeviceByArgs(opts ...*Opts) (*model.Device, error)
	CountDevices(opts ...*Opts) (int, error)
	UpdateDevice(uuid string, body *model.Device, opts ...*Opts) (*model.Device, error)
	UpdateDeviceErrors(uuid string, body *model.Device, opts ...*Opts) error
	UpdateDeviceDescendantsErrors(deviceUUID, message, messageLevel, messageCode string, opts ...*Opts) error
	UpsertDeviceMetaTags(uuid string, body []*model.DeviceMetaTag, opts ...*Opts) error
	UpsertDeviceTags(uuid string, body []*model.Tag, opts ...*Opts) error
	DeleteDevice(uuid string, opts ...*Opts) error
	DeleteOneDeviceByArgs(opts ...*Opts) error
	DeleteDeviceByName(name string, opts ...*Opts) error
	ClearDeviceDescendantsErrors(deviceUUID string, opts ...*Opts) error

	CreatePoint(body *model.Point, opts ...*Opts) (*model.Point, error)
	GetPoints(opts ...*Opts) ([]*model.Point, error)
	GetPoint(uuid string, opts ...*Opts) (*model.Point, error)
	GetPointByName(networkName, deviceName, pointName string, opts ...*Opts) (*model.Point, error)
	GetOnePointByArgs(opts ...*Opts) (*model.Point, error)
	GetPointWithParent(uuid string, opts ...*Opts) (*dto.PointWithParent, error)
	GetPointWithParentByName(networkName, deviceName, pointName string, opts ...*Opts) (*dto.PointWithParent, error)
	CountPoints(opts ...*Opts) (int, error)
	UpdatePoint(uuid string, body *model.Point, opts ...*Opts) (*model.Point, error)
	PointWrite(uuid string, body *dto.PointWriter, opts ...*Opts) (*dto.PointWriteResponse, error)
	PointWriteByName(networkName, deviceName, pointName string, body *dto.PointWriter, opts ...*Opts) (*dto.PointWriteResponse, error)
	UpdatePointErrors(uuid string, body *model.Point, opts ...*Opts) error
	UpdatePointSuccess(uuid string, body *model.Point, opts ...*Opts) error
	UpsertPoint(uuid string, body *model.Point, opts ...*Opts) (*model.Point, error)
	UpsertPointMetaTags(uuid string, body []*model.PointMetaTag, opts ...*Opts) error
	UpsertPointTags(uuid string, body []*model.Tag, opts ...*Opts) error
	DeletePoint(uuid string, opts ...*Opts) error
	DeleteOnePointByArgs(opts ...*Opts) error
	DeletePointByName(name string, opts ...*Opts) error

	CreateSchedule(body *model.Schedule, opts ...*Opts) (*model.Schedule, error)
	GetSchedules(opts ...*Opts) ([]*model.Schedule, error)
	GetSchedule(uuid string, opts ...*Opts) (*model.Schedule, error)
	GetOneScheduleByArgs(opts ...*Opts) (*model.Schedule, error)
	UpdateSchedule(uuid string, body *model.Schedule, opts ...*Opts) (*model.Schedule, error)
	ScheduleWrite(uuid string, body *dto.ScheduleData, opts ...*Opts) (*dto.ScheduleData, error)
	UpdateScheduleAllProps(uuid string, body *model.Schedule, opts ...*Opts) (*model.Schedule, error)
	DeleteSchedule(uuid string, opts ...*Opts) error

	CreateLocation(body *model.Location, opts ...*Opts) (*model.Location, error)
	GetLocations(opts ...*Opts) ([]*model.Location, error)
	GetLocation(uuid string, opts ...*Opts) (*model.Location, error)
	UpdateLocation(uuid string, body *model.Location, opts ...*Opts) (*model.Location, error)
	DeleteLocation(uuid string, opts ...*Opts) error

	CreateGroup(body *model.Group, opts ...*Opts) (*model.Group, error)
	GetGroups(opts ...*Opts) ([]*model.Group, error)
	GetGroup(uuid string, opts ...*Opts) (*model.Group, error)
	UpdateHostsStatus(uuid string, opts ...*Opts) (*model.Group, error)
	UpdateGroup(uuid string, body *model.Group, opts ...*Opts) (*model.Group, error)
	DeleteGroup(uuid string, opts ...*Opts) error

	CreateHost(body *model.Host, opts ...*Opts) (*model.Host, error)
	GetHosts(opts ...*Opts) ([]*model.Host, error)
	GetHost(uuid string, opts ...*Opts) (*model.Host, error)
	UpdateHost(uuid string, body *model.Host, opts ...*Opts) (*model.Host, error)
	UpsertHostTags(uuid string, body []*model.Tag, opts ...*Opts) error
	DeleteHost(uuid string, opts ...*Opts) error
	CloneHostThingsToCloud(hostUUID string, opts ...*Opts) error

	GetPlugins(opts ...*Opts) ([]*model.Plugin, error)
	GetPlugin(uuid string, opts ...*Opts) (*model.Plugin, error)
	GetPluginByName(name string, opts ...*Opts) (*model.Plugin, error)
	UpdatePluginMessage(name string, body *model.Plugin, opts ...*Opts) error
	CreateModuleDir(name string, opts ...*Opts) (*string, error)

	CreateHistories(histories []*model.History, opts ...*Opts) (bool, error)
	GetHistories(historyRequest *dto.HistoryRequest, opts ...*Opts) (*dto.HistoryResponse, error)
	GetLatestHistoryByHostAndPointUUID(hostUUID, pointUUID string, opts ...*Opts) (*model.History, error)
	DeleteHistories(opts ...*Opts) error

	CreatePointHistories(histories []*model.PointHistory, opts ...*Opts) (bool, error)
	GetPointHistories(opts ...*Opts) ([]*model.PointHistory, error)
	GetPointHistoriesByPointUUID(pointUUID string, opts ...*Opts) ([]*model.PointHistory, error)
	GetLatestPointHistoryByPointUUID(pointUUID string, opts ...*Opts) (*model.PointHistory, error)
	GetPointHistoriesByPointUUIDs(pointUUIDs []*string, opts ...*Opts) ([]*model.PointHistory, error)
	GetPointHistoriesForSync(opts ...*Opts) ([]*model.PointHistory, error)
	GetPointHistoriesMissingTimestamps(pointUUID string, opts ...*Opts) ([]string, error)
	DeletePointHistoriesByPointUUID(pointUUID string, opts ...*Opts) error

	CreateAlert(body *model.Alert, opts ...*Opts) (*model.Alert, error)
	GetAlerts(opts ...*Opts) ([]*model.Alert, error)
	GetAlert(uuid string, opts ...*Opts) (*model.Alert, error)
	UpdateAlertStatus(uuid string, body *dto.AlertStatus, opts ...*Opts) (*model.Alert, error)
	UpdateAlertTeams(uuid string, teamUUIDs []*string, opts ...*Opts) ([]*model.AlertTeam, error)
	UpsertAlertMetaTags(uuid string, body []*model.AlertMetaTag, opts ...*Opts) error
	UpsertAlertTags(uuid string, body []*model.Tag, opts ...*Opts) error
	DeleteAlert(uuid string, opts ...*Opts) error
	DeleteAlertTransaction(transactionUUID string, opts ...*Opts) error

	CreateTeam(body *model.Team, opts ...*Opts) (*model.Team, error)
	GetTeams(opts ...*Opts) ([]*model.Team, error)
	GetTeam(uuid string, opts ...*Opts) (*model.Team, error)
	UpdateTeam(uuid string, body *model.Team, opts ...*Opts) (*model.Team, error)
	UpsertTeamMembers(uuid string, memberUUIDs []*string, opts ...*Opts) ([]*model.Member, error)
	UpsertTeamViews(uuid string, viewUUIDs []*string, opts ...*Opts) ([]*model.TeamView, error)
	UpsertTeamMetaTags(uuid string, body []*model.TeamMetaTag, opts ...*Opts) error
	DeleteTeam(uuid string, opts ...*Opts) error

	CreateTicket(body *model.Ticket, opts ...*Opts) (*model.Ticket, error)
	GetTickets(opts ...*Opts) ([]*model.Ticket, error)
	GetTicket(uuid string, opts ...*Opts) (*model.Ticket, error)
	UpdateTicket(uuid string, body *model.Ticket, opts ...*Opts) (*model.Ticket, error)
	UpsertTicketPriority(uuid string, body *dto.TicketPriority, opts ...*Opts) error
	UpsertTicketStatus(uuid string, body *dto.TicketStatus, opts ...*Opts) error
	UpsertTicketTeams(uuid string, teamUUIDs []*string, opts ...*Opts) ([]*model.TicketTeam, error)
	UpsertTicketMembers(uuid string, memberUUIDs []*string, opts ...*Opts) ([]*model.TicketMember, error)
	DeleteTicket(uuid string, opts ...*Opts) error

	CreateTicketComment(body *model.TicketComment, opts ...*Opts) (*model.TicketComment, error)
	UpdateTicketComment(uuid string, body *model.TicketComment, opts ...*Opts) (*model.TicketComment, error)
	DeleteTicketComment(uuid string, opts ...*Opts) error

	RunScanner(body *dto.Scanner, opts ...*Opts) (*scanner.Hosts, error)
	RebootHost(opts ...*Opts) error
	GetDeviceInfo(opts ...*Opts) (*dto.DeviceInfo, error)
	GetNetworkInterfaces(opts ...*Opts) ([]*networking.NetworkInterfaces, error)
	HostTime(opts ...*Opts) (*datelib.Time, error)
	GetSystem(opts ...*Opts) (*systats.System, error)
	GetMemoryUsage(opts ...*Opts) (*dto.MemoryUsage, error)
	GetMemory(opts ...*Opts) (*systats.Memory, error)
	GetTopProcesses(opts ...*Opts) ([]*systats.Process, error)
	GetSwap(opts ...*Opts) (*systats.Swap, error)
	DiscUsage(opts ...*Opts) ([]*ostats.MountingPoint, error)
	DiscUsagePretty(opts ...*Opts) ([]*dto.Disk, error)

	GetHistoriesForPostgresSync(opts ...*Opts) ([]*model.History, error)
	GetPointsForPostgresSync(opts ...*Opts) ([]*dto.PointForPostgresSync, error)
	GetNetworksTagsForPostgresSync(opts ...*Opts) ([]*dto.NetworkTagForPostgresSync, error)
	GetDevicesTagsForPostgresSync(opts ...*Opts) ([]*dto.DeviceTagForPostgresSync, error)
	GetPointsTagsForPostgresSync(opts ...*Opts) ([]*dto.PointTagForPostgresSync, error)
	GetNetworksMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.NetworkMetaTagForPostgresSync, error)
	GetDevicesMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.DeviceMetaTagForPostgresSync, error)
	GetPointsMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.PointMetaTagForPostgresSync, error)
	GetLastSyncHistoryIdForPostgresSync(opts ...*Opts) (int, error)
	UpdateLastSyncHistoryRowForPostgresSync(log *model.HistoryPostgresLog, opts ...*Opts) (*model.HistoryPostgresLog, error)

	GetHistoryLogByHostUUID(hostUUID string, opts ...*Opts) (*model.HistoryLog, error)
	UpdateBulkHistoryLogs(logs []*model.HistoryLog, opts ...*Opts) (bool, error)

	Publish(topic string, qos datatype.QOS, retain bool, payload string, opts ...*Opts) error
	PublishNonBuffer(topic string, qos datatype.QOS, retain bool, payload string, opts ...*Opts) error
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
