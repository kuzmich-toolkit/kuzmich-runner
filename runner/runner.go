package runner

import "github.com/kuzmich-toolkit/kuzmich-runner/sharding"

//noinspection GoNameStartsWithPackageName
type (
	Runner interface {
		run(shard sharding.Shard) Result
	}
)
