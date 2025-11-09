// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)

// Log prints time (if TimeFormat != ""), [Level], optional [tid], then the message.
// It respects Cfg.LogLevel and colorizes:
//   • [Level] and [tid] with the provided colorize
//   • string-ish args inside the body
//   • timestamp with Cfg.LogTimeColor (if non-nil)
func Log(level LogLevel, colorize Colorizer, format string, args ...any) {
	if Cfg.LogLevel < level {
		return
	}

	// colorize string-ish args
	if colorize != nil {
		for i, a := range args {
			switch v := a.(type) {
			case string:
				args[i] = colorize(v)
			case error:
				args[i] = colorize(v.Error())
			case fmt.Stringer:
				args[i] = colorize(v.String())
			}
		}
	}

	body := fmt.Sprintf(format, args...)
	if !strings.HasSuffix(body, "\n") {
		body += "\n"
	}

	// timestamp (only if TimeFormat is non-empty)
	ts := ""
	if strings.TrimSpace(Cfg.TimeFormat) != "" {
		raw := time.Now().Format(Cfg.TimeFormat)
		if Cfg.LogTimeColor != nil {
			raw = Cfg.LogTimeColor(raw)
		}
		ts = raw + " "
	}

	// [Level] (and [tid] if enabled)
	levelStr := level.String()
	if colorize != nil {
		levelStr = colorize(levelStr)
	}
	prefix := "[" + levelStr + "] "

	if Cfg.UseTid != nil && *Cfg.UseTid {
		tid := getTid()
		tidStr := strconv.Itoa(tid)
		if colorize != nil {
			tidStr = colorize(tidStr)
		}
		prefix = "[" + levelStr + "][" + tidStr + "] "
	}

	LoggerOutputMutex.Lock()
	_, _ = io.WriteString(LoggerOutput, ts+prefix+body)
	LoggerOutputMutex.Unlock()
}
