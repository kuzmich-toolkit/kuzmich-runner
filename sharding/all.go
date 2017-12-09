package sharding

import "github.com/dimorinny/adbaster/model"

type AllDevicesManager struct{}

func NewAllDevicesManager() Manager {
	return &AllDevicesManager{}
}

func (m *AllDevicesManager) shard(devices []model.Device) []Shard {
	result := []Shard{}

	if len(devices) > 0 {
		result = append(
			result,
			Shard{
				Title:   "All devices",
				Devices: devices,
			},
		)
	}

	return result
}
