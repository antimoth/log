package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// getRuntimeInfo callstack information and code position
func getRuntimeInfo(skip int) string {
	pc, fileP, line, ok := runtime.Caller(skip)
	if !ok {
		return "Get runtime caller error!"
	}

	callerStack := runtime.FuncForPC(pc).Name()
	var index int

	if index = strings.Index(callerStack, ".("); index == -1 {
		index = strings.LastIndex(callerStack, ".")
	}

	// parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	// pl := len(parts)

	// for i := 0; i < pl; i++ {
	// 	if parts[i][0] == '(' {
	// 		packageName = strings.Join(parts[0:i], ".")
	// 		break
	// 	}
	// }
	// if packageName == "" {
	//         packageName = strings.Join(parts[0:pl-1], ".")
	// }

	return fmt.Sprintf("%s@%s:%d", callerStack[0:index], path.Base(fileP), line)
}

// getCodePosition 
func getCodePosition(skip int) string {
	_, fileP, line, ok := runtime.Caller(skip)
	if !ok {
		return "Get runtime caller error!"
	}
        return fmt.Sprintf("%s:%d", fileP, line)
}

