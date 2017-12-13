package common

import (
	"github.com/dimorinny/adbaster/adb"
	"github.com/kuzmich-toolkit/kuzmich-runner/runner"
	"github.com/kuzmich-toolkit/kuzmich-runner/sharding"
)

type Runner struct {
	shardNumber,
	countOfShards int

	client adb.Client
}

func NewCommonRunner(client adb.Client, shardNumber, countOfShards int) runner.Runner {
	return &Runner{
		shardNumber:   shardNumber,
		countOfShards: countOfShards,

		client: client,
	}
}

func (r *Runner) run(shard sharding.Shard) runner.Result {

}
