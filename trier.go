package etrier

import "runtime"

var (
	handlers = make(map[uintptr]func(err error))
	defaultHandler = func(err error) {
		println("error ignored: handler not defined")
	}
)

// SetHandler - sets handler for current and all functions called by it;
// It is guaranteed that the error will be not nil when the handler is called.
func SetHandler(handler func(err error)) bool {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return false
	}
	handlers[runtime.FuncForPC(pc).Entry()] = handler
	return true
}

func GetHandler() func(err error) {
	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			return defaultHandler
		}

		handler, ok := handlers[runtime.FuncForPC(pc).Entry()]
		if !ok {
			continue
		}
		return handler
	}
}

// Try - if err is not nil, calls handler.
func Try(err error) {
	if err != nil {
		GetHandler()(err)
	}
}

// Try1 - if err is not nil, calls handler.
func Try1[A any](a A, err error) (A) {
	Try(err)
	return a
}

// Try2 - if err is not nil, calls handler.
func Try2[A, B any](a A, b B, err error) (A, B) {
	Try(err)
	return a, b
}

// Try3 - if err is not nil, calls handler.
func Try3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	Try(err)
	return a, b, c
}

// Try4 - if err is not nil, calls handler.
func Try4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	Try(err)
	return a, b, c, d
}
