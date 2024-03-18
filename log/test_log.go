package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type TestLoggerOps func(t *TestLogger)

// WithTestLoggerLevel use for set custom level for TestLogger
// as level use zerolog.Level
func WithTestLoggerLevel(lvl zerolog.Level) TestLoggerOps {
	return func(t *TestLogger) {
		t.lvl = lvl
	}
}

// WithTestLoggerWriter use for set custom writer for TestLogger
func WithTestLoggerWriter(w io.Writer) TestLoggerOps {
	return func(t *TestLogger) {
		t.w = w
	}
}

// WithTestLoggerName use for set custom name for TestLogger
// as name use string
func WithTestLoggerName(name string) TestLoggerOps {
	return func(t *TestLogger) {
		t.testName = name
	}
}

type TestLogger struct {
	z zerolog.Logger

	// --- custom fields ---
	testName string
	w        io.Writer
	lvl      zerolog.Level
}

const testNameKey = "test_name"

func (t *TestLogger) Printf(format string, v ...any) {
	t.z.Info().Msgf(format, v...)
}

// NewTestLogger returns new instance of TestLogger
// you need provide stdout for your log
// this function will init default logger with default formatter
//
// as default TestLogger use Info level
//
// as default TestLogger use os.Stdout
//
// as default TestLogger use "unknown" as test name
func NewTestLogger(opts ...TestLoggerOps) *TestLogger {
	tl := &TestLogger{
		testName: "unknown",
		lvl:      zerolog.InfoLevel,
		w:        os.Stdout,
	}

	for _, opt := range opts {
		opt(tl)
	}

	tl.z = zerolog.New(tl.w).With().Str(testNameKey, tl.testName).Logger()
	return tl
}
