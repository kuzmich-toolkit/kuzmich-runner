package sharding

import (
	"github.com/dimorinny/adbaster/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingShardsWithSdkManagerAndDifferentSdk(t *testing.T) {
	manager := NewSdkManager()
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
			Sdk:          21,
			BatteryLevel: 100,
		},
	}

	expected := []Shard{
		{
			Title: "Sdk level based shard for: 22 api level",
			Devices: []model.Device{
				devices[0],
			},
		},
		{
			Title: "Sdk level based shard for: 21 api level",
			Devices: []model.Device{
				devices[1],
			},
		},
	}
	result := manager.shard(devices)

	assert.Equal(
		t,
		len(result),
		len(expected),
	)
	assert.Subset(
		t,
		manager.shard(devices),
		expected,
	)
}

func TestCreatingShardsWithSdkManagerAndSameSdk(t *testing.T) {
	manager := NewSdkManager()
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
			Title:   "Sdk level based shard for: 22 api level",
			Devices: devices,
		},
	}
	result := manager.shard(devices)

	assert.Equal(
		t,
		expected,
		result,
	)
}

func TestCreatingShardsWithSdkManagerReturnsNoShardsForNoDevices(t *testing.T) {
	manager := NewSdkManager()
	devices := []model.Device{}

	assert.Empty(
		t,
		manager.shard(devices),
	)
}
