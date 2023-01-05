package cmd

import (
	"strconv"

	"github.com/spf13/viper"
)

type Argument func(args []string, idx int)

func MakeStringArgument(name string) Argument {
	return func(args []string, idx int) {
		viper.Set(name, args[idx])
	}
}
func MakeStringSliceArgument(name string) Argument {
	return func(args []string, idx int) {
		viper.Set(name, args[idx:])
	}
}
func MakeIntArgument(name string) Argument {
	return func(args []string, idx int) {
		if val, err := strconv.ParseInt(args[idx], 10, 64); err == nil {
			viper.Set(name, int(val))
		}
	}
}
func MakeIntSliceArgument(name string) Argument {
	return func(args []string, idx int) {
		var iArr []int
		for _, arg := range args[idx:] {
			if val, err := strconv.ParseInt(arg, 10, 64); err == nil {
				iArr = append(iArr, int(val))
			} else {
				panic(err)
			}
		}
		viper.Set(name, iArr)
	}
}
func ParseArguments(args []string, mandatory, optional []Argument, variable Argument) {
	idx := 0
	for _, setter := range mandatory {
		setter(args, idx)
		idx++
	}
	for _, setter := range optional {
		if idx >= len(args) {
			return
		}
		setter(args, idx)
		idx++
	}
	if variable != nil && idx < len(args) {
		variable(args, idx)
	}
}
