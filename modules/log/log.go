package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"morningo/config"
	"os"
	"runtime/debug"
	"time"
)

// ---------------------------------------------------------------
// Logger
// ---------------------------------------------------------------
//
// there are three kinds of logger：
//
// 1） access logger
//     log every request visited which used to counts the ip
//     and other indicators
//
// 2） error logger
//	   record the panic error
//
// 3） info logger
//     log something the developer wants to output
//
// ---------------------------------------------------------------

var (
	ErrorWriter io.Writer
	InfoWriter  io.Writer
)

const (
	LeveL_WARNING = "warning"
	LeveL_INFO    = "info"
	LeveL_DEBUG   = "debug"
	LeveL_ERROR   = "error"
	LeveL_SERIOUS = "serious"
)

type E struct {
	Function string
	Error    error
	Title    string
	Info     M
	Level    string
	Context  *gin.Context
}

type M map[string]interface{}

func init() {
	InitAllLogger()
}

func InitAllLogger() {
	// init access.log
	if config.GetEnv().AccessLog {
		gin.DefaultWriter = InitLogger(config.GetEnv().AccessLogPath)
	}

	// init error.log
	if config.GetEnv().ErrorLog {
		ErrorWriter = InitLogger(config.GetEnv().ErrorLogPath)
	}

	// init info.log
	if config.GetEnv().InfoLog {
		InfoWriter = InitLogger(config.GetEnv().InfoLogPath)
	}
}

func InitLogger(path string) io.Writer {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	if config.GetEnv().Debug {
		return io.MultiWriter(file, os.Stdout)
	} else {
		return io.MultiWriter(file)
	}
}

func Error(err interface{}) {
	if config.GetEnv().ErrorLog {
		fmt.Fprintf(ErrorWriter, "%s", "\n")
		fmt.Fprintf(ErrorWriter, "%s", "["+time.Now().Format("2006-01-02 15:04:05")+"] app.ERROR: ")
		fmt.Fprintf(ErrorWriter, "%s", err)
		fmt.Fprintf(ErrorWriter, "%s", "\nStack trace:\n")
		fmt.Fprintf(ErrorWriter, "%s", debug.Stack())
		fmt.Fprintf(ErrorWriter, "%s", "\n")
	}
}

func Info(info E) {
	if config.GetEnv().InfoLog {
		fmt.Fprintf(InfoWriter, "%s", "time="+time.Now().Format("2006-01-02 15:04:05")+" ")

		if info.Level == "" {
			info.Level = "info"
		}

		fmt.Fprintf(InfoWriter, "level=%s ", info.Level)

		if info.Context != nil {
			fmt.Fprintf(InfoWriter, "method=%s path=%s ", info.Context.Request.Method, info.Context.Request.URL.Path)
		}

		if info.Function != "" {
			fmt.Fprintf(InfoWriter, "function=%s ", info.Function)
		}

		if info.Title != "" {
			fmt.Fprintf(InfoWriter, "title=%s ", info.Title)
		}

		if info.Error != nil {
			fmt.Fprintf(InfoWriter, "error=%s ", info.Error)
		}

		for k, v := range info.Info {
			fmt.Fprintf(InfoWriter, "%s=%v ", k, v)
		}
		fmt.Fprintf(InfoWriter, "\n")
	}
}

func Println(a ...interface{}) {
	if config.GetEnv().Debug {
		fmt.Println(a...)
	}
}

func Printf(format string, a ...interface{}) {
	if config.GetEnv().Debug {
		fmt.Printf(format, a...)
	}
}
