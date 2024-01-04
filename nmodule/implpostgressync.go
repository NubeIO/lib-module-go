package nmodule

import (
	"encoding/json"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strconv"
)

func (g *GRPCMarshaller) GetHistoriesForPostgresSync(lastSyncId int, opts ...*Opts) ([]*model.History, error) {
	api := "/api/postgres-sync/histories"
	lastSyncIdStr := strconv.Itoa(lastSyncId)
	if len(opts) > 0 {
		opts[0].Args = &nargs.Args{IdGt: &lastSyncIdStr}
	} else {
		opts = append(opts, &Opts{Args: &nargs.Args{IdGt: &lastSyncIdStr}})
	}
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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

func (g *GRPCMarshaller) GetPointsForPostgresSync(opts ...*Opts) ([]*dto.PointForPostgresSync, error) {
	api := "/api/postgres-sync/points"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.PointForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetNetworksTagsForPostgresSync(opts ...*Opts) ([]*dto.NetworkTagForPostgresSync, error) {
	api := "/api/postgres-sync/networks-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.NetworkTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetDevicesTagsForPostgresSync(opts ...*Opts) ([]*dto.DeviceTagForPostgresSync, error) {
	api := "/api/postgres-sync/devices-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.DeviceTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetPointsTagsForPostgresSync(opts ...*Opts) ([]*dto.PointTagForPostgresSync, error) {
	api := "/api/postgres-sync/points-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.PointTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetNetworksMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.NetworkMetaTagForPostgresSync, error) {
	api := "/api/postgres-sync/networks-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.NetworkMetaTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetDevicesMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.DeviceMetaTagForPostgresSync, error) {
	api := "/api/postgres-sync/devices-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.DeviceMetaTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetPointsMetaTagsForPostgresSync(opts ...*Opts) ([]*dto.PointMetaTagForPostgresSync, error) {
	api := "/api/postgres-sync/points-meta-tags"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var r []*dto.PointMetaTagForPostgresSync
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *GRPCMarshaller) GetLastSyncHistoryIdForPostgresSync(opts ...*Opts) (int, error) {
	api := "/api/postgres-sync/last-sync-history-id"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
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
	res, err := g.CallDBHelperWithParser(nhttp.PATCH, api, log, opts...)
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
