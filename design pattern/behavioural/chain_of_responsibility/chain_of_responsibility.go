package cor

import "fmt"

type ILogger interface {
	SetNext(ILogger)
	Log(int, string)
}

type Logger struct {
	NextLogger ILogger
}

func (l *Logger) SetNext(logger ILogger) {
	l.NextLogger = logger
}

func (l Logger) Log(level int, msg string) {
	if level > 1 {
		if l.NextLogger == nil {
			fmt.Println("This level doesn't support")
			return
		}
		l.NextLogger.Log(level, msg)
		return
	}
	fmt.Println("Logger: ", msg)
}

type DebugLogger struct {
	NextLogger ILogger
}

func (d *DebugLogger) SetNext(logger ILogger) {
	d.NextLogger = logger
}

func (d *DebugLogger) Log(level int, msg string) {
	if level > 2 {
		if d.NextLogger == nil {
			fmt.Println("This level doesn't support")
			return
		}
		d.NextLogger.Log(level, msg)
		return
	}
	fmt.Println("Debug logger: ", msg)
}

// func main() {
// 	logger := &cor.Logger{}
// 	debug := &cor.DebugLogger{}
// 	logger.SetNext(debug)
// 	logger.Log(1, "hello")
// 	logger.Log(2, "hello")
// }
