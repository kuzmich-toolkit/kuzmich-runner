package arguments

type (
	sharedArguments struct {
		Apk,
		TestApk,
		Package,
		TestPackage,
		Output string

		InstrumentationArguments map[string]string

		Verbose bool
	}

	CommonRunnerArguments struct {
		sharedArguments
	}
)

func NewCommonRunnerArguments() *CommonRunnerArguments {
	return &CommonRunnerArguments{
		sharedArguments: sharedArguments{
			InstrumentationArguments: map[string]string{},
		},
	}
}
