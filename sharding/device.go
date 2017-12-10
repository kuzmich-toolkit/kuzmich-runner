package sharding

import "github.com/dimorinny/adbaster/model"

type PerDeviceManager struct{}

func NewPerDeviceManager() Manager {
	return &PerDeviceManager{}
}

func (m *PerDeviceManager) shard(devices []model.Device) []Shard {
	shards := []Shard{}

	for _, device := range devices {
		shards = append(
			shards,
			Shard{
				Title: string(device.Identifier),
				Devices: []model.Device{
					device,
				},
			},
		)
	}

	return shards
}
