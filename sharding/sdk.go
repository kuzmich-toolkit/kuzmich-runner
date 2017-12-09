package sharding

import (
	"fmt"
	"github.com/dimorinny/adbaster/model"
)

type SdkManager struct{}

func NewSdkManager() Manager {
	return &SdkManager{}
}

func (a *SdkManager) shard(devices []model.Device) []Shard {
	groupedDevices := map[int][]model.Device{}
	result := []Shard{}

	for _, device := range devices {
		currentDevices, contains := groupedDevices[device.Sdk]

		if contains {
			groupedDevices[device.Sdk] = append(currentDevices, device)
		} else {
			groupedDevices[device.Sdk] = []model.Device{
				device,
			}
		}
	}

	for sdk, sdkDevices := range groupedDevices {
		result = append(
			result,
			Shard{
				Title:   fmt.Sprintf("Sdk level based shard for: %d api leve", sdk),
				Devices: sdkDevices,
			},
		)
	}

	return result
}
