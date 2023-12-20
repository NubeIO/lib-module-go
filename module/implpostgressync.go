package module

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strconv"
)

func (g *GRPCMarshaller) GetHistoriesForPostgresSync(lastSyncId int, opts ...*Opts) ([]*model.History, error) {
	api := "/api/postgres-sync/histories"
	lastSyncIdStr := strconv.Itoa(lastSyncId)
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{IdGt: &lastSyncIdStr}, nil, opts...)
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

func (g *GRPCMarshaller) GetPointsForPostgresSync(opts ...*Opts) ([]*model.PointForPostgresSync, error) {
	api := "/api/postgres-sync/points"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworksTagsForPostgresSync(opts ...*Opts) ([]*model.NetworkTagForPostgresSync, error) {
	api := "/api/postgres-sync/networks-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetDevicesTagsForPostgresSync(opts ...*Opts) ([]*model.DeviceTagForPostgresSync, error) {
	api := "/api/postgres-sync/devices-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetPointsTagsForPostgresSync(opts ...*Opts) ([]*model.PointTagForPostgresSync, error) {
	api := "/api/postgres_sync/points-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetNetworksMetaTagsForPostgresSync(opts ...*Opts) ([]*model.NetworkMetaTag, error) {
	api := "/api/postgres-sync/networks-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetDevicesMetaTagsForPostgresSync(opts ...*Opts) ([]*model.DeviceMetaTag, error) {
	api := "/api/postgres-sync/devices-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetPointsMetaTagsForPostgresSync(opts ...*Opts) ([]*model.PointMetaTag, error) {
	api := "/api/postgres-sync/points-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) GetLastSyncHistoryIdForPostgresSync(opts ...*Opts) (int, error) {
	api := "/api/postgres-sync/last-sync-history-id"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nargs.Args{}, nil, opts...)
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

func (g *GRPCMarshaller) UpdateLastSyncHistoryRowForPostgresSync(log *model.HistoryPostgresLog, opts ...*Opts) (
	*model.HistoryPostgresLog, error) {
	api := "/api/postgres-sync/last-sync-history-row"
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, nargs.Args{}, log, opts...)
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
