// Generated 2023-01-05 17:28:57 by go-framework v2.1.4
package unit_tests

import (
	"fmt"
	"strings"

	"gitlab.com/atrico/kafka-recorder/unit-tests/common"
)

// ----------------------------------------------------------------------------------------------------------------------------
// Options
// ----------------------------------------------------------------------------------------------------------------------------

type Option interface {
	Set()
	Cmdline() string
	ModifySettings(settings *common.MockSettings)
}
type OptionSet map[string]Option
type OptionSetList []OptionSet

// ----------------------------------------------------------------------------------------------------------------------------
// Simple option (--opt <val>)
// ----------------------------------------------------------------------------------------------------------------------------

func NewSimpleOption[T any](option string, generator func() T, modifier func(s *common.MockSettings, value T)) Option {
	opt := simpleOption[T]{
		option:    option,
		generator: generator,
		modifier:  modifier,
	}
	return &opt
}

type simpleOption[T any] struct {
	value     T
	option    string
	generator func() T
	modifier  func(s *common.MockSettings, value T)
}

func (o *simpleOption[T]) Set() {
	o.value = o.generator()
}

func (o *simpleOption[T]) Cmdline() string {
	return fmt.Sprintf("%s %v", o.option, o.value)
}

func (o *simpleOption[T]) ModifySettings(settings *common.MockSettings) {
	o.modifier(settings, o.value)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Boolean option (--opt/--opt=true/--opt=false)
// ----------------------------------------------------------------------------------------------------------------------------

func NewBooleanOption(option string, modifier func(s *common.MockSettings)) Option {
	opt := boolOption{
		option:   option,
		modifier: modifier,
	}
	return &opt
}

type boolOption struct {
	value    bool
	option   string
	modifier func(s *common.MockSettings)
}

func (o *boolOption) Set() {
	// Nothing to do
}

func (o *boolOption) Cmdline() string {
	return fmt.Sprintf("%s", o.option)
}

func (o *boolOption) ModifySettings(settings *common.MockSettings) {
	o.modifier(settings)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Slice option (--opt 1 --opt 2 --opt 3)
// ----------------------------------------------------------------------------------------------------------------------------

func NewSliceOption[T any](option string, generator func() []T, modifier func(s *common.MockSettings, value []T)) Option {
	opt := sliceOption[T]{
		option:    option,
		generator: generator,
		modifier:  modifier,
	}
	return &opt
}

type sliceOption[T any] struct {
	value     []T
	option    string
	generator func() []T
	modifier  func(s *common.MockSettings, value []T)
}

func (o *sliceOption[T]) Set() {
	o.value = o.generator()
}

func (o *sliceOption[T]) Cmdline() string {
	items := make([]string, 0)
	for _, item := range o.value {
		items = append(items, fmt.Sprintf("%s %v", o.option, item))
	}
	return strings.Join(items, " ")
}

func (o *sliceOption[T]) ModifySettings(settings *common.MockSettings) {
	o.modifier(settings, o.value)
}
