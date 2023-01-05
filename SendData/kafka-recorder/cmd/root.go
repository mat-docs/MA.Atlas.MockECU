// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/atrico/container"

	"gitlab.com/atrico/kafka-recorder/pkg"
	"gitlab.com/atrico/kafka-recorder/settings"
)

type RootCmd commandInfo

func RegisterCmdRoot(c container.Container) {
	c.Singleton(func(config settings.RootSettings) RootCmd { return RootCmd(createRootCommand(config)) })
}

func createRootCommand(config settings.RootSettings) commandInfo {
	cmd := &cobra.Command{
		Use:          pkg.Name,
		Short:        pkg.Summary,
		Long:         fmt.Sprintf("%s\n%s", pkg.Description, pkg.Version),
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return setupLogger(config)
		},
	}
	settings.AddFileFlag(cmd.PersistentFlags())
	settings.AddBrokersFlag(cmd.PersistentFlags())
	settings.AddTopicFlag(cmd.PersistentFlags())
	settings.AddTimeoutFlag(cmd.PersistentFlags())
	settings.AddConfigFileFlag(cmd.PersistentFlags())
	settings.AddVerboseFlag(cmd.PersistentFlags())
	settings.AddLogLevelFlag(cmd.PersistentFlags())
	settings.AddLogWriterFlag(cmd.PersistentFlags())
	settings.AddLogTimeFormatFlag(cmd.PersistentFlags())
	settings.AddLogStackTraceFlag(cmd.PersistentFlags())
	return commandInfo{cmd, "."}
}

func setupLogger(config settings.RootSettings) (err error) {
	// Writer
	parameters := strings.Split(config.LogWriter(), "=")
	writerSpec := strings.Split(strings.ToLower(parameters[0]), ":")
	var writer io.Writer
	noColor := false
	switch writerSpec[0] {
	case "stdout":
		writer = os.Stdout
	case "stderr":
		writer = os.Stderr
	case "file":
		if len(parameters) > 1 {
			writer, err = os.OpenFile(parameters[1], os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			noColor = true
		} else {
			err = errors.New("no filename given")
		}
	}
	if err == nil {
		jsonFormat, badFormat := false, false
		if len(writerSpec) > 1 {
			if writerSpec[1] == "json" {
				jsonFormat = true
			} else {
				badFormat = true
			}
		}
		if !jsonFormat {
			writer = zerolog.ConsoleWriter{Out: writer, TimeFormat: config.LogTimeFormat(), NoColor: noColor}
		}
		log.Logger = log.Output(writer)
		if badFormat {
			log.Warn().Msgf("invalid logger format: '%s'", writerSpec[1])
		}
		if err == nil {
			level := zerolog.InfoLevel
			switch strings.ToLower(config.LogLevel()) {
			case "5", "panic":
				level = zerolog.PanicLevel
			case "4", "fatal":
				level = zerolog.FatalLevel
			case "3", "error", "err":
				level = zerolog.ErrorLevel
			case "2", "warning", "warn":
				level = zerolog.WarnLevel
			case "1", "information", "info":
				level = zerolog.InfoLevel
			case "0", "debug":
				level = zerolog.DebugLevel
			case "-1", "trace", "all":
				level = zerolog.TraceLevel
			default:
				log.Warn().Msgf("invalid logger level: '%s'", config.LogLevel())
			}
			zerolog.SetGlobalLevel(level)
			if config.LogStackTrace() {
				zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
			}
		}
	}
	return err
}

const keyConfigFile = "ConfigFile"

func initConfig() {
	// Config file
	cfgFile := viper.GetString("ConfigFile")
	if cfgFile != "" {
		// Use config file from the flag.
		if config, err := tryReadConfigFilepath(cfgFile); err == nil {
			viper.MergeConfigMap(config)
		} else {
			// Fail if specified config cannot be read
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		var err error
		var notFoundError viper.ConfigFileNotFoundError
		// Standard name for config
		cfgName := fmt.Sprintf(".%s", pkg.Name)
		// Try executable directory
		if config, err2 := tryReadConfigNameDir(cfgName, func() (string, error) { _, f, _, _ := runtime.Caller(0); return filepath.Dir(f), nil }); err2 == nil {
			viper.MergeConfigMap(config)
		} else if err == nil && !errors.As(err2, &notFoundError) {
			err = err2
		}
		// Try home directory
		if config, err2 := tryReadConfigNameDir(cfgName, func() (string, error) { return homedir.Dir() }); err2 == nil {
			viper.MergeConfigMap(config)
		} else if err == nil && !errors.As(err2, &notFoundError) {
			err = err2
		}
		// Try current working directory
		if config, err2 := tryReadConfigNameDir(cfgName, func() (string, error) { return os.Getwd() }); err2 == nil {
			viper.MergeConfigMap(config)
		} else if err == nil {
			err = err2
		}
		if err != nil && !errors.As(err, &notFoundError) {
			// Fail if config failed to be parsed
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func tryReadConfigFilepath(filePath string) (config map[string]interface{}, err error) {
	v := viper.New()
	v.SetConfigFile(filePath)
	// log.Debug().Msgf("trying config: %s", filePath)
	return tryReadConfigImpl(v)
}

func tryReadConfigNameDir(name string, getDir func() (dir string, err error)) (config map[string]interface{}, err error) {
	var dir string
	if dir, err = getDir(); err == nil {
		v := viper.New()
		v.SetConfigName(name)
		v.AddConfigPath(dir)
		// log.Debug().Msgf("trying config: '%s' from '%s'", name, dir)
		config, err = tryReadConfigImpl(v)
	}
	return config, err
}

func tryReadConfigImpl(v *viper.Viper) (config map[string]interface{}, err error) {
	if err = v.ReadInConfig(); err == nil {
		settings.GetVerboseService().VerbosePrintln("Using config file:", v.ConfigFileUsed())
		config = v.AllSettings()
		viper.Set(keyConfigFile, v.ConfigFileUsed())
	}
	return config, err
}
