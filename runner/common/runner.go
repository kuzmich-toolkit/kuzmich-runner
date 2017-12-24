package common

import (
	"errors"
	"fmt"
	"github.com/dimorinny/adbaster/adb"
	"github.com/dimorinny/adbaster/model"
	"github.com/kuzmich-toolkit/kuzmich-runner/arguments"
	"github.com/kuzmich-toolkit/kuzmich-runner/runner"
	"github.com/kuzmich-toolkit/kuzmich-runner/sharding"
	"log"
	"strconv"
)

type Runner struct {
	arguments arguments.CommonRunnerArguments
	client    adb.Client
}

//func NewCommonRunner(
//	client adb.Client,
//	arguments arguments.CommonRunnerArguments,
//) runner.Runner {
//	return &Runner{
//		client:    client,
//		arguments: arguments,
//	}
//}

func (r *Runner) run(shard sharding.Shard) runner.ShardResult {
	devicesCount := len(shard.Devices)

	for index, device := range shard.Devices {
		// TODO: think about handling results
		r.runInstrumentationTests(device.Identifier, devicesCount, index)
	}

	// TODO: parse results

	return runner.ShardResult{}
}

func (r *Runner) runInstrumentationTests(
	device model.DeviceIdentifier,
	numShards,
	shardIndex int,
) {
	_, _, err := r.client.RunInstrumentationTests(
		device,
		model.InstrumentationParams{
			TestPackage: r.arguments.TestPackage,
			Arguments: r.makeInstrumentationArguments(
				map[string]string{
					"numShards":  strconv.Itoa(numShards),
					"shardIndex": strconv.Itoa(shardIndex),
				},
			),
		},
	)
	if err != nil {
		log.Fatal(
			errors.New(
				fmt.Sprintf(
					"Error during running instrumentation tests in device: %s. Info: %s",
					device,
					err.Error(),
				),
			),
		)
	}

	// TODO: think about handling results
}

// create instrumentation arguments from shard arguments
// and arguments passed as command line argument
func (r *Runner) makeInstrumentationArguments(arguments map[string]string) model.InstrumentationArguments {
	result := model.InstrumentationArguments{}

	for key, value := range r.arguments.InstrumentationArguments {
		result[key] = value
	}

	for key, value := range arguments {
		result[key] = value
	}

	return result
}
