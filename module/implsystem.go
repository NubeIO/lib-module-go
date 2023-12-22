package module

import (
	"encoding/json"
	"github.com/NubeIO/lib-date/datelib"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-networking/scanner"
	systats "github.com/NubeIO/lib-system"
	"github.com/NubeIO/lib-system/ostats"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
)

func (g *GRPCMarshaller) RunScanner(body *dto.Scanner, opts ...*Opts) (*scanner.Hosts, error) {
	api := "/api/system/scanner"
	res, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	if err != nil {
		return nil, err
	}
	var hosts *scanner.Hosts
	err = json.Unmarshal(res, &hosts)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (g *GRPCMarshaller) RebootHost(opts ...*Opts) error {
	api := "/api/system/reboot"
	_, err := g.DbHelper.CallDBHelper(nhttp.POST, api, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (g *GRPCMarshaller) GetDeviceInfo(opts ...*Opts) (*dto.DeviceInfo, error) {
	api := "/api/system/device"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var deviceInfo *dto.DeviceInfo
	err = json.Unmarshal(res, &deviceInfo)
	if err != nil {
		return nil, err
	}
	return deviceInfo, nil
}

func (g *GRPCMarshaller) GetNetworkInterfaces(opts ...*Opts) ([]*networking.NetworkInterfaces, error) {
	api := "/api/system/network-interfaces"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var interfaces []*networking.NetworkInterfaces
	err = json.Unmarshal(res, &interfaces)
	if err != nil {
		return nil, err
	}
	return interfaces, nil
}

func (g *GRPCMarshaller) HostTime(opts ...*Opts) (*datelib.Time, error) {
	api := "/api/system/time"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var time *datelib.Time
	err = json.Unmarshal(res, &time)
	if err != nil {
		return nil, err
	}
	return time, nil
}

func (g *GRPCMarshaller) GetSystem(opts ...*Opts) (*systats.System, error) {
	api := "/api/system/info"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var system *systats.System
	err = json.Unmarshal(res, &system)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func (g *GRPCMarshaller) GetMemoryUsage(opts ...*Opts) (*dto.MemoryUsage, error) {
	api := "/api/system/usage"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var memoryUsage *dto.MemoryUsage
	err = json.Unmarshal(res, &memoryUsage)
	if err != nil {
		return nil, err
	}
	return memoryUsage, nil
}

func (g *GRPCMarshaller) GetMemory(opts ...*Opts) (*systats.Memory, error) {
	api := "/api/system/memory"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var memory *systats.Memory
	err = json.Unmarshal(res, &memory)
	if err != nil {
		return nil, err
	}
	return memory, nil
}

// GetTopProcesses required: opts[0].Args.Count, opts[0].Args.Sort
func (g *GRPCMarshaller) GetTopProcesses(opts ...*Opts) ([]*systats.Process, error) {
	api := "/api/system/processes"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var processes []*systats.Process
	err = json.Unmarshal(res, &processes)
	if err != nil {
		return nil, err
	}
	return processes, nil
}

func (g *GRPCMarshaller) GetSwap(opts ...*Opts) (*systats.Swap, error) {
	api := "/api/system/swap"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var swap *systats.Swap
	err = json.Unmarshal(res, &swap)
	if err != nil {
		return nil, err
	}
	return swap, nil
}

func (g *GRPCMarshaller) DiscUsage(opts ...*Opts) ([]*ostats.MountingPoint, error) {
	api := "/api/system/disc"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var mp []*ostats.MountingPoint
	err = json.Unmarshal(res, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func (g *GRPCMarshaller) DiscUsagePretty(opts ...*Opts) ([]*dto.Disk, error) {
	api := "/api/system/disc"
	res, err := g.DbHelper.CallDBHelper(nhttp.GET, api, nil, opts...)
	if err != nil {
		return nil, err
	}
	var discs []*dto.Disk
	err = json.Unmarshal(res, &discs)
	if err != nil {
		return nil, err
	}
	return discs, nil
}
