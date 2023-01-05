// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package unit_tests

import (
	"fmt"
	"strings"

	"gitlab.com/atrico/kafka-recorder/unit-tests/common"
)

type CommandArgument interface {
	Set()
	Cmdline() string
	ModifySettings(settings *common.MockSettings)
}

func NewCommandArgument[T any](generator func() T, modifier func(s *common.MockSettings, value T)) CommandArgument {
	arg := commandArgument[T]{
		generator: generator,
		modifier:  modifier,
	}
	return &arg
}

func NewCommandSliceArgument[T any](generator func() []T, modifier func(s *common.MockSettings, value []T)) CommandArgument {
	arg := commandSliceArgument[T]{
		generator: generator,
		modifier:  modifier,
	}
	return &arg
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type commandArgument[T any] struct {
	value     T
	generator func() T
	modifier  func(s *common.MockSettings, value T)
}

func (a *commandArgument[T]) Set() {
	a.value = a.generator()
}

func (a *commandArgument[T]) Cmdline() string {
	return fmt.Sprintf("%v", a.value)
}

func (a *commandArgument[T]) ModifySettings(settings *common.MockSettings) {
	a.modifier(settings, a.value)
}

type commandSliceArgument[T any] struct {
	value     []T
	generator func() []T
	modifier  func(s *common.MockSettings, value []T)
}

func (a *commandSliceArgument[T]) Set() {
	a.value = a.generator()
}

func (a *commandSliceArgument[T]) Cmdline() string {
	cmdline := strings.Builder{}
	for i, item := range a.value {
		if i > 0 {
			cmdline.WriteString(" ")
		}
		cmdline.WriteString(fmt.Sprintf("%v", item))
	}
	return cmdline.String()
}

func (a *commandSliceArgument[T]) ModifySettings(settings *common.MockSettings) {
	a.modifier(settings, a.value)
}
