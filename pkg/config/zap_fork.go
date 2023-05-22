package config

import (
	"go.uber.org/zap/zapcore"
)

type zapFork struct {
	zapcore.LevelEnabler
	cores []zapcore.Core
}

func NewZapFork(cores ...zapcore.Core) zapcore.Core {
	return &zapFork{
		cores: cores,
	}
}

// With adds structured context to the Core.
func (z *zapFork) With(fields []zapcore.Field) zapcore.Core {
	out := &zapFork{
		cores: make([]zapcore.Core, len(z.cores)),
	}
	for i := range z.cores {
		out.cores[i] = z.cores[i].With(fields)
	}
	return out
}

// LevelEnabler decides whether a given logging level is enabled when logging a
// message.
//
// Enablers are intended to be used to implement deterministic filters;
// concerns like sampling are better implemented as a Core.
//
// Each concrete Level value implements a static LevelEnabler which returns
// true for itself and all higher logging levels. For example WarnLevel.Enabled()
// will return true for WarnLevel, ErrorLevel, DPanicLevel, PanicLevel, and
// FatalLevel, but return false for InfoLevel and DebugLevel.
func (z *zapFork) Enabled(level zapcore.Level) bool {
	for _, c := range z.cores {
		if enabled := c.Enabled(level); enabled {
			return true
		}
	}
	return false
}

// Check determines whether the supplied Entry should be logged (using the
// embedded LevelEnabler and possibly some extra logic). If the entry
// should be logged, the Core adds itself to the CheckedEntry and returns
// the result.
//
// Callers must use Check before calling Write.
func (z *zapFork) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return checkedEntry.AddCore(entry, z)
	}
	return nil
}

// Write serializes the Entry and any Fields supplied at the log site and
// writes them to their destination.
//
// If called, Write should always log the Entry and Fields; it should not
// replicate the logic of Check.
func (z *zapFork) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for _, c := range z.cores {
		if !c.Enabled(entry.Level) {
			// Skip disabled cores
			continue
		}
		if err := c.Write(entry, fields); err != nil {
			return err
		}
	}
	return nil
}

// Sync flushes buffered logs (if any).
func (z *zapFork) Sync() error {
	for _, c := range z.cores {
		if err := c.Sync(); err != nil {
			return err
		}
	}
	return nil
}
