package instrumentation

import (
	"fmt"
	"github.com/dimorinny/adbaster/adb/instrumentation"
	"github.com/dimorinny/adbaster/model"
)

func LogInstrumentationEvent(device model.Device, event instrumentation.Event) {
	// TODO: more logging

	switch event {
	case instrumentation.TestsRunStartedEvent{}:
		logForDevice(
			device,
			fmt.Sprintf(
				"Tests run started. Count of tests: %d",
				event.(instrumentation.TestsRunStartedEvent).NumberOfTests,
			),
		)

	case instrumentation.TestsRunFailedEvent{}:
		logForDevice(
			device,
			fmt.Sprintf(
				"Tests run failed with message: %s",
				event.(instrumentation.TestsRunFailedEvent).Message,
			),
		)

	case instrumentation.TestsRunFinishedEvent{}:
		logForDevice(
			device,
			"Tests run finished.",
		)
	}
}

func logForDevice(device model.Device, message string) {
	fmt.Println(
		fmt.Sprintf(
			"%s: %s",
			device.Identifier,
			message,
		),
	)
}
