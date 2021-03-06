package log

import (
        "bytes"
)

var (
        logInvokePoint string
)


func init() {
        b := &bytes.Buffer{}        
        b.WriteByte('\n')
        b.Write(bytes.Repeat([]byte{' '}, 64))
        b.WriteString("logInvokePoint")
        logInvokePoint = b.String()
}

// Error error log
func Error(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                logger.Error(sDescribe, append(args, logInvokePoint, getCodePosition(2))...)
        } else {
	        logger.Error(sDescribe, args...)
        }
}

// Info info log
func Info(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                logger.Info(sDescribe, append(args, logInvokePoint, getCodePosition(2))...)
        } else {
	        logger.Info(sDescribe, args...)
        }
}

// Debug debug log
func Debug(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                logger.Debug(sDescribe, append(args, logInvokePoint, getCodePosition(2))...)
        } else {
	        logger.Debug(sDescribe, args...)
        }
}

// Warn warning log
func Warn(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                logger.Warn(sDescribe, append(args, logInvokePoint, getCodePosition(2))...)
        } else {
	        logger.Warn(sDescribe, args...)
        }
}

// Crit critical log
func Crit(sDescribe string, args ...interface{}) {
        if pkgLvl == StrLvlDebug {
                logger.Crit(sDescribe, append(args, logInvokePoint, getCodePosition(2))...)
        } else {
                logger.Crit(sDescribe, args...)
        }
}


