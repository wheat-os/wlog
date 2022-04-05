package wlog

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// 简单的单 Hook 模型使用
func TestSimpleHook_Handler(t *testing.T) {

	testValue := 0

	logger := New()

	// 定义一个简单的钩子
	hook := ShortHook(
		func(level Level, content []byte) error {
			testValue += 1
			return nil
		},
		func(err error) {},
	)

	// 设置一个支持 DebugLevel，ErrorLevel 的 hook
	logger.SetOptions(WithHook(DebugLevel|ErrorLevel, hook))

	logger.Info("test")
	// info 不触发 hook
	require.Equal(t, testValue, 0)

	logger.Error("test")
	// error 触发 hook +1
	require.Equal(t, testValue, 1)

	testValue = 0

	hook = ShortHook(
		func(level Level, content []byte) error {
			// 模拟一个错误
			if level == InfoLevel {
				return errors.New("def info")
			}
			testValue += 1
			return nil
		},
		func(err error) {
			testValue -= 1
		},
	)

	logger.SetOptions(WithHook(InfoLevel|ErrorLevel, hook))
	logger.Info("")
	logger.Error("")
	logger.Warn("")

	require.Equal(t, testValue, 0)

}

func TestHookGroup_Handler(t *testing.T) {
	testValue := 0
	logger := New()

	add1Hook := ShortHook(
		func(level Level, content []byte) error {
			if level == ErrorLevel {
				return errors.New("error")
			}
			testValue += 1
			return nil
		},
		func(err error) {
			testValue -= 1
		},
	)

	add2Hook := ShortHook(
		func(level Level, content []byte) error {
			if level == ErrorLevel {
				return errors.New("error")
			}
			testValue += 2
			return nil
		},
		func(err error) {
			testValue -= 2
		},
	)

	add3Hook := ShortHook(
		func(level Level, content []byte) error {
			if level == ErrorLevel {
				return errors.New("error")
			}
			testValue += 3
			return nil
		},
		func(err error) {
			testValue -= 3
		},
	)

	logger.SetOptions(WithHookGroup(InfoLevel|ErrorLevel, add1Hook, add2Hook, add3Hook))

	// 同时触发
	logger.Info("")
	require.Equal(t, testValue, 6)

	// 错误处理
	logger.Error("")
	require.Equal(t, testValue, 0)

}
