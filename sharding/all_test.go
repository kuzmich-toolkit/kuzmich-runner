package sharding

import (
	"github.com/dimorinny/adbaster/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingShardsWithAllManager(t *testing.T) {
	manager := NewAllDevicesManager()
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
			Title:   "All devices",
			Devices: devices,
		},
	}

	assert.Equal(
		t,
		expected,
		manager.shard(devices),
	)
}

func TestCreatingShardsWithAllManagerReturnsNoShardsForNoDevices(t *testing.T) {
	manager := NewAllDevicesManager()
	devices := []model.Device{}

	assert.Empty(
		t,
		manager.shard(devices),
	)
}
