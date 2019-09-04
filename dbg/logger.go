package dbg

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var debugEnabled = false
var debugLog *log.Logger

var mu = sync.Mutex{}

func Enable(enable bool, logFile ...io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	if enable {
		lf := []io.Writer{os.Stdout}
		if len(logFile) > 0 {
			lf = logFile
		}
		debugLog = log.New(io.MultiWriter(lf...), "DEBUG:", log.Lmicroseconds | log.Lshortfile)
	}
	debugEnabled = enable
}

func Enabled() bool {
	return debugEnabled
}

// Debug uses the default logger and logs with the Info severity.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprint(v...))
	}
}

// DebugDepth acts as Info but uses depth to determine which call frame to log.
// DebugDepth(0, "msg") is the same as Info("msg").
func DebugDepth(depth int, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2 + depth, fmt.Sprint(v...))
	}
}

// Debugln uses the default logger and logs with the Info severity.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprintln(v...))
	}
}

// Debugf uses the default logger and logs with the Info severity.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprintf(format, v...))
	}
}

