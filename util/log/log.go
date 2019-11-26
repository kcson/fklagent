package log

import (
	"fasoo.com/fklagent/constant"
	"fasoo.com/fklagent/util/config"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var eLog *log.Logger
var iLog *log.Logger
var dLog *log.Logger

var lLevel constant.LOG_LEVEL
var logConsole bool = false

func init() {
	lLevel = constant.LOG_LEVEL(config.P.GetInt("log.level", 1))
	logPath := config.P.GetString("log.path", "./log")
	logPrefix := config.P.GetString("log.fileprefix", "happygolf")
	logConsole = config.P.GetBool("log.console", false)

	maxSize := config.P.GetInt("log.maxsize", 20)
	maxBackups := config.P.GetInt("log.maxbackups", 10)
	maxAge := config.P.GetInt("log.maxage", 30)

	eLog = log.New(os.Stdout, "[ERROR]", log.Ldate|log.Ltime)
	iLog = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	dLog = log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime)

	fmt.Println(logPath + string(os.PathSeparator) + logPrefix + "_error.log")
	eLog.SetOutput(&lumberjack.Logger{
		Filename:   logPath + string(os.PathSeparator) + logPrefix + "_error.log",
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})
	iLog.SetOutput(&lumberjack.Logger{
		Filename:   logPath + string(os.PathSeparator) + logPrefix + "_info.log",
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})
	dLog.SetOutput(&lumberjack.Logger{
		Filename:   logPath + string(os.PathSeparator) + logPrefix + "_debug.log",
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})

	INFO("init log config!!")
}

func ERROR(msg string) {
	if lLevel < constant.LOG_LEVEL_ERROR {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	msg = parseFilePath(file) + ":" + strconv.Itoa(line) + " " + msg
	eLog.Println(msg)
	if logConsole {
		fmt.Println("[ERROR]" + msg)
	}

}

func INFO(msg string) {
	if lLevel < constant.LOG_LEVEL_INFO {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	msg = parseFilePath(file) + ":" + strconv.Itoa(line) + " " + msg
	iLog.Println(msg)
	if logConsole {
		fmt.Println("[INFO]" + msg)
	}
}

func DEBUG(msg string) {
	if lLevel < constant.LOG_LEVEL_DEBUG {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	msg = parseFilePath(file) + ":" + strconv.Itoa(line) + " " + msg
	dLog.Println(msg)
	if logConsole {
		fmt.Println("[DEBUG]" + msg)
	}
}

func parseFilePath(fPath string) (fileName string) {
	_, fileName = filepath.Split(fPath)

	return
}
