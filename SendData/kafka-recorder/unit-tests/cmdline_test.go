// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package unit_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"gitlab.com/atrico/container"
	"gitlab.com/atrico/testing/v2/assert"
	"gitlab.com/atrico/testing/v2/is"
	"gitlab.com/atrico/testing/v2/random"
	"gitlab.com/atrico/viperEx/v2"

	"gitlab.com/atrico/kafka-recorder/cmd"
	"gitlab.com/atrico/kafka-recorder/pkg"
	"gitlab.com/atrico/kafka-recorder/settings"
	"gitlab.com/atrico/kafka-recorder/unit-tests/common"
)

// SECTION-START: Options

var rg = random.NewValueGenerator()

var OptionFile = OptionSet{
	"Default": NewSimpleOption("--file", func() (value string) { rg.Value(&value); return value }, func(s *common.MockSettings, value string) { s.FileVar = value }),
	"Short":   NewSimpleOption("-f", func() (value string) { rg.Value(&value); return value }, func(s *common.MockSettings, value string) { s.FileVar = value }),
}
var OptionBrokers = OptionSet{
	"Default": NewSliceOption("--broker", func() (value []string) { rg.Value(&value); return value }, func(s *common.MockSettings, value []string) { s.BrokersVar = value }),
	"Short":   NewSliceOption("-b", func() (value []string) { rg.Value(&value); return value }, func(s *common.MockSettings, value []string) { s.BrokersVar = value }),
}
var OptionTopic = OptionSet{
	"Default": NewSimpleOption("--topic", func() (value string) { rg.Value(&value); return value }, func(s *common.MockSettings, value string) { s.TopicVar = value }),
}
var OptionTimeout = OptionSet{
	"Default": NewSimpleOption("--timeout", func() (value int) { rg.Value(&value); return value }, func(s *common.MockSettings, value int) { s.TimeoutVar = value }),
	"Short":   NewSimpleOption("-t", func() (value int) { rg.Value(&value); return value }, func(s *common.MockSettings, value int) { s.TimeoutVar = value }),
}
var OptionVerbose = OptionSet{
	"Default":     NewBooleanOption("--verbose", func(s *common.MockSettings) { s.VerboseVar = true }),
	"=True":       NewBooleanOption("--verbose=true", func(s *common.MockSettings) { s.VerboseVar = true }),
	"=False":      NewBooleanOption("--verbose=false", func(s *common.MockSettings) { s.VerboseVar = false }),
	"Short":       NewBooleanOption("-v", func(s *common.MockSettings) { s.VerboseVar = true }),
	"Short=True":  NewBooleanOption("-v=true", func(s *common.MockSettings) { s.VerboseVar = true }),
	"Short=False": NewBooleanOption("-v=false", func(s *common.MockSettings) { s.VerboseVar = false }),
}
var OptionLogLevel = OptionSet{
	"Default": NewSimpleOption("--log-level", func() interface{} { return "debug" }, func(s *common.MockSettings, value interface{}) { s.LogLevelVar = value.(string) }),
}
var OptionLogWriter = OptionSet{
	"Default": NewSimpleOption("--log-writer", func() (value string) { rg.Value(&value); return value }, func(s *common.MockSettings, value string) { s.LogWriterVar = value }),
}
var OptionLogTimeFormat = OptionSet{
	"Default": NewSimpleOption("--log-timeformat", func() (value string) { rg.Value(&value); return value }, func(s *common.MockSettings, value string) { s.LogTimeFormatVar = value }),
}
var OptionLogStackTrace = OptionSet{
	"Default": NewBooleanOption("--log-stacktrace", func(s *common.MockSettings) { s.LogStackTraceVar = true }),
	"=True":   NewBooleanOption("--log-stacktrace=true", func(s *common.MockSettings) { s.LogStackTraceVar = true }),
	"=False":  NewBooleanOption("--log-stacktrace=false", func(s *common.MockSettings) { s.LogStackTraceVar = false }),
}

// SECTION-END

// SECTION-START: TestCases
// ----------------------------------------------------------------------------------------------------------------------------
// Test cases
// ----------------------------------------------------------------------------------------------------------------------------
type CmdlineTestCase struct {
	Command []string
	Args    []CommandArgument
	Options []Option
}

var CmdlineTestCases = []CmdlineTestCase{
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionFile["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionFile["Short"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionBrokers["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionBrokers["Short"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionTopic["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionTimeout["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionTimeout["Short"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["=True"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["=False"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short=True"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short=False"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogLevel["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogWriter["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogTimeFormat["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["Default"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["=True"]}},
	{Command: []string{"record"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["=False"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionFile["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionFile["Short"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionBrokers["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionBrokers["Short"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionTopic["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionTimeout["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionTimeout["Short"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["=True"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["=False"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short=True"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionVerbose["Short=False"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogLevel["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogWriter["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogTimeFormat["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["Default"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["=True"]}},
	{Command: []string{"replay"}, Args: []CommandArgument{}, Options: []Option{OptionLogStackTrace["=False"]}},
}

// SECTION-END

func addUserTests(tests []CmdlineTestCase) []CmdlineTestCase {
	// Append extra tests here
	return tests
}

// SECTION-START: Test
// ----------------------------------------------------------------------------------------------------------------------------
// Test
// ----------------------------------------------------------------------------------------------------------------------------

func Test_CommandLine(t *testing.T) {
	for _, testCase := range addUserTests(CmdlineTestCases) {
		// Build command line and expectations
		cmdline := strings.Builder{}
		cmdline.WriteString(strings.Join(testCase.Command, " "))
		expected := common.NewMockSettings(testCase.Command)
		resetDefaultArguments(&expected)
		for _, arg := range testCase.Args {
			arg.Set()
			if len(cmdline.String()) > 0 {
				cmdline.WriteString(" ")
			}
			cmdline.WriteString(arg.Cmdline())
			arg.ModifySettings(&expected)
		}
		for _, opt := range testCase.Options {
			opt.Set()
			if len(cmdline.String()) > 0 {
				cmdline.WriteString(" ")
			}
			cmdline.WriteString(opt.Cmdline())
			opt.ModifySettings(&expected)
		}
		t.Run(cmdline.String(), func(t *testing.T) { testCommandLineImpl(t, cmdline.String(), expected) })
	}
}

func testCommandLineImpl(t *testing.T, cmdline string, expected common.MockSettings) {
	// Arrange
	args := strings.Split(cmdline, " ")
	os.Args = make([]string, 1+len(args))
	os.Args[0] = pkg.Name
	copy(os.Args[1:], args)
	cmd := resetCommand()

	// Act
	err := cmd.Execute()

	// Assert
	assert.That[any](t, err, is.Nil, "Error")
	assert.That(t, results["TheCommand"].([]string), is.EqualTo(expected.TheCommand), "Command")
	assert.That(t, results["File"].(string), is.DeepEqualTo(expected.File()), "File")
	assert.That(t, results["Brokers"].([]string), is.DeepEqualTo(expected.Brokers()), "Brokers")
	assert.That(t, results["Topic"].(string), is.DeepEqualTo(expected.Topic()), "Topic")
	assert.That(t, results["Timeout"].(int), is.DeepEqualTo(expected.Timeout()), "Timeout")
	assert.That(t, results["ConfigFile"].(string), is.DeepEqualTo(expected.ConfigFile()), "ConfigFile")
	assert.That(t, results["Verbose"].(bool), is.DeepEqualTo(expected.Verbose()), "Verbose")
	assert.That(t, results["LogLevel"].(string), is.DeepEqualTo(expected.LogLevel()), "LogLevel")
	assert.That(t, results["LogWriter"].(string), is.DeepEqualTo(expected.LogWriter()), "LogWriter")
	assert.That(t, results["LogTimeFormat"].(string), is.DeepEqualTo(expected.LogTimeFormat()), "LogTimeFormat")
	assert.That(t, results["LogStackTrace"].(bool), is.DeepEqualTo(expected.LogStackTrace()), "LogStackTrace")
}

func resetCommand() *cobra.Command {
	// Re-register singletons
	c := container.NewContainer()
	settings.RegisterSettings(c)
	cmd.RegisterCmd(c)
	registerMockApiFactories(c)
	// Reset settings
	viperEx.Reset()
	settings.ResetCaches()
	// Return root cmd
	var cmdFactory cmd.CommandFactory
	c.Make(&cmdFactory)
	return cmdFactory.Create()
}

// Default value not copied across is unit test so blank these
func resetDefaultArguments(expected *common.MockSettings) {
}

// SECTION-END
