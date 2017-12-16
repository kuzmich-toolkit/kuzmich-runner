package arguments

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCommand() *cobra.Command {
	return &cobra.Command{
		Run: func(command *cobra.Command, args []string) {},
	}
}

func executeCommand(command *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(command, args...)
	return output, err
}

func executeCommandC(command *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	command.SetOutput(buf)
	command.SetArgs(args)

	c, err = command.ExecuteC()

	return c, buf.String(), err
}

func TestParsingBooleanFlag(t *testing.T) {
	command := createCommand()
	command.Flags().String("test", "", "")

	executeCommand(command, "--test", "false")

	value, err := parseBooleanFlag(command, "test")

	assert.Equal(
		t,
		false,
		value,
	)
	assert.Nil(
		t,
		err,
	)
}

func TestParsingBooleanFlagWhenInvalidStringIsPassed(t *testing.T) {
	command := createCommand()
	command.Flags().String("test", "", "")

	executeCommand(command, "--test", "true1")

	_, err := parseBooleanFlag(command, "test")

	assert.NotNil(
		t,
		err,
	)
}

func TestParsingMapFlag(t *testing.T) {
	command := createCommand()
	command.Flags().String("test", "", "")

	executeCommand(command, "--test", "hello world hello1 world1")

	result, err := parseMapFlag(command, "test")

	assert.Equal(
		t,
		map[string]string{
			"hello":  "world",
			"hello1": "world1",
		},
		result,
	)
	assert.Nil(
		t,
		err,
	)
}

func TestParsingMapFlagReturnsEmptyMapForEmptyString(t *testing.T) {
	command := createCommand()
	command.Flags().String("test", "", "")

	executeCommand(command, "--test", "")

	result, err := parseMapFlag(command, "test")

	assert.Equal(
		t,
		map[string]string{},
		result,
	)
	assert.Nil(
		t,
		err,
	)
}

func TestParsingMapFlagReturnsErrorForNonEvenNumberOfWords(t *testing.T) {
	command := createCommand()
	command.Flags().String("test", "", "")

	executeCommand(command, "--test", "test1 test2 test3")

	result, err := parseMapFlag(command, "test")

	assert.Nil(
		t,
		result,
	)
	assert.NotNil(
		t,
		err,
	)
}
