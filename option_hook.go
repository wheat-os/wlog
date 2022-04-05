package wlog

type Hook interface {
	// hook 处理
	Handler(level Level, content []byte) error
	// 处理 hook err 中发生的错误
	DealWithErr(err error)
}

type shortHook struct {
	handler func(level Level, content []byte) error
	dealErr func(err error)
}

// hook 处理
func (h *shortHook) Handler(level Level, content []byte) error {
	return h.handler(level, content)
}

// 处理 hook err 中发生的错误
func (h *shortHook) DealWithErr(err error) {
	h.dealErr(err)
}

// 简单的 hook 新建接口
func ShortHook(
	handler func(level Level, content []byte) error,
	dealErr func(err error),
) Hook {
	return &shortHook{handler: handler, dealErr: dealErr}
}

type SimpleHook struct {
	level Level
	hook  Hook
}

// hook 处理
func (s *SimpleHook) Handler(logLevel Level, content []byte) error {
	if s.level&logLevel == 0 || s.hook == nil {
		return nil
	}

	if err := s.hook.Handler(logLevel, content); err != nil {
		s.DealWithErr(err)
	}

	return nil
}

// 处理 hook err 中发生的错误
func (s *SimpleHook) DealWithErr(err error) {
	s.hook.DealWithErr(err)
}

func WithHook(level Level, hook Hook) OptionFunc {
	return func(o *options) {
		o.hook = &SimpleHook{
			level: level,
			hook:  hook,
		}
	}
}

// 更高级的 hook，用来直接处理一组 hook
type HookGroup struct {
	level Level
	hooks []Hook // hook 处理
}

func (s *HookGroup) Handler(logLevel Level, content []byte) error {
	if s.level&logLevel == 0 {
		return nil
	}

	for _, hook := range s.hooks {
		if err := hook.Handler(logLevel, content); err != nil {
			hook.DealWithErr(err)
		}
	}

	return nil
}

// 处理 hook err 中发生的错误
func (s *HookGroup) DealWithErr(err error) {
	panic("restricted call")
}

// hook 组合
func WithHookGroup(logLevel Level, hooks ...Hook) OptionFunc {
	return func(o *options) {
		o.hook = &HookGroup{
			hooks: hooks,
			level: logLevel,
		}
	}
}
