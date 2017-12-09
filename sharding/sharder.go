package sharding

import "github.com/dimorinny/adbaster/model"

type (
	Manager interface {
		shard(devices []model.Device) []Shard
	}

	Shard struct {
		Title   string
		Devices []model.Device
	}
)
