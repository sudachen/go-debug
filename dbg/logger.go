package dbg

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const tagDebug = "DEBUG: "

var debugEnabled = false
var debugLog *log.Logger

var mu = sync.Mutex{}

func Enable(enable bool, flags int, verbose bool, logFile ...io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	if enable {
		var lf []io.Writer
		if len(logFile) > 0 {
			for _, f := range logFile {
				if f != nil {
					lf = append(lf, f)
				}
			}
		}
		if len(lf) == 0 || verbose {
			lf = append([]io.Writer{os.Stdout}, lf...)
		}
		if flags == 0 {
			flags = log.Ldate | log.Lmicroseconds | log.Lshortfile
		}
		debugLog = log.New(io.MultiWriter(lf...), tagDebug, flags)
	}
	debugEnabled = enable
}

func Enabled() bool {
	return debugEnabled
}

func Debug(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprint(v...))
	}
}

func DebugDepth(depth int, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2+depth, fmt.Sprint(v...))
	}
}

func Debugln(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprintln(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if debugEnabled && debugLog != nil {
		_ = debugLog.Output(2, fmt.Sprintf(format, v...))
	}
}
