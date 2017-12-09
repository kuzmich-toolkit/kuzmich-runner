package sharding

import (
	"github.com/dimorinny/adbaster/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingShardsWithEveryManager(t *testing.T) {
	manager := NewPerDeviceManager()
	devices := []model.Device{
		{
			Identifier:   "identifier1",
			Arch:         "x86",
			Timezone:     "ru_RU",
			HeapSize:     "1000GB",
			Sdk:          22,
			BatteryLevel: 100,
		},
		{
			Identifier:   "identifier2",
			Arch:         "x86",
			Timezone:     "ru_RU",
			HeapSize:     "1000GB",
			Sdk:          22,
			BatteryLevel: 100,
		},
	}

	expected := []Shard{
		{
			Title: string(devices[0].Identifier),
			Devices: []model.Device{
				devices[0],
			},
		},
		{
			Title: string(devices[1].Identifier),
			Devices: []model.Device{
				devices[1],
			},
		},
	}

	assert.Equal(
		t,
		expected,
		manager.shard(devices),
	)
}

func TestCreatingShardsWithEveryManagerReturnsNoShardsForNoDevices(t *testing.T) {
	manager := NewPerDeviceManager()
	devices := []model.Device{}

	assert.Empty(
		t,
		manager.shard(devices),
	)
}
