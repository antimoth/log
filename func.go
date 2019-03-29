package log

import (
        "fmt"
)

// Error error log
func Error(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                fmt.Println(getCodePosition(2))
        }
	logger.Error(sDescribe, args...)
}

// Info info log
func Info(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                fmt.Println(getCodePosition(2))
        }
	logger.Info(sDescribe, args...)
}

// Debug debug log
func Debug(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                fmt.Println(getCodePosition(2))
        }
	logger.Debug(sDescribe, args...)
}

// Warn warning log
func Warn(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                fmt.Println(getCodePosition(2))
        }
	logger.Warn(sDescribe, args...)
}

// Crit critical log
func Crit(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                fmt.Println(getCodePosition(2))
        }
        logger.Crit(sDescribe, args...)
}


