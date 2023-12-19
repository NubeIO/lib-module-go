package module

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strconv"
)

func (g *GRPCMarshaller) GetHistoriesForPostgresSync(lastSyncId int) ([]*model.History, error) {
	api := "/api/postgres-sync/histories"
	lastSyncIdStr := strconv.Itoa(lastSyncId)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{IdGt: &lastSyncIdStr}, nil)
	if err != nil {
		return nil, err
	}
	var history []*model.History
	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (g *GRPCMarshaller) GetPointsForPostgresSync() ([]*model.PointForPostgresSync, error) {
	api := "/api/postgres-sync/points"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.PointForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetNetworksTagsForPostgresSync() ([]*model.NetworkTagForPostgresSync, error) {
	api := "/api/postgres-sync/networks-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.NetworkTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetDevicesTagsForPostgresSync() ([]*model.DeviceTagForPostgresSync, error) {
	api := "/api/postgres-sync/devices-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.DeviceTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetPointsTagsForPostgresSync() ([]*model.PointTagForPostgresSync, error) {
	api := "/api/postgres_sync/points-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.PointTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetNetworksMetaTagsForPostgresSync() ([]*model.NetworkMetaTag, error) {
	api := "/api/postgres-sync/networks-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.NetworkMetaTag
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetDevicesMetaTagsForPostgresSync() ([]*model.DeviceMetaTag, error) {
	api := "/api/postgres-sync/devices-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.DeviceMetaTag
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetPointsMetaTagsForPostgresSync() ([]*model.PointMetaTag, error) {
	api := "/api/postgres-sync/points-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return nil, err
	}
	var r []*model.PointMetaTag
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetLastSyncHistoryIdForPostgresSync() (int, error) {
	api := "/api/postgres-sync/last-sync-history-id"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil)
	if err != nil {
		return 0, err
	}
	var q int
	err = json.Unmarshal(res, &q)
	if err != nil {
		return 0, err
	}
	return q, nil
}

func (g *GRPCMarshaller) UpdateLastSyncHistoryRowForPostgresSync(log *model.HistoryPostgresLog) (
	*model.HistoryPostgresLog, error) {
	api := "/api/postgres-sync/last-sync-history-row"
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, log)
	if err != nil {
		return nil, err
	}
	var updatedLog *model.HistoryPostgresLog
	err = json.Unmarshal(res, &updatedLog)
	if err != nil {
		return nil, err
	}
	return updatedLog, nil
}
