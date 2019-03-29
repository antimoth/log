package log

import (
	"os"

	"bytes"
	"fmt"

	log "github.com/inconshreveable/log15"
)

var (
        pkgLvl = StrLvlDebug
        logger = log.New()
        streamHandler = log.StreamHandler(os.Stdout, terminalFormat())
)

func init() {
	logger.SetHandler(getLvlFilterHandler(pkgLvl))
}

func getLvlFilterHandler(lvl string) log.Handler {
        return log.LvlFilterHandler(toLogLevel(lvl), streamHandler)
}

func toLogLevel(lvl string) log.Lvl {
	switch lvl {
	case StrLvlDebug:
		return log.LvlDebug
	case StrLvlInfo:
		return log.LvlInfo
	case StrLvlWarn:
		return log.LvlWarn
	case StrLvlError:
		return log.LvlError
	case StrlvlCrit:
		return log.LvlCrit
	default:
		return log.LvlInfo
	}
}

func toLvlStr(lvl log.Lvl) string {
	switch lvl {
	case log.LvlDebug:
	        return StrLvlDebug
	case log.LvlInfo:
	        return StrLvlInfo
	case log.LvlWarn:
	        return StrLvlWarn
        case log.LvlError:
                return StrLvlError
        case log.LvlCrit:
                return StrlvlCrit
	default:
		return strLvlUnknown
	}
}
func printLogCtx(buf *bytes.Buffer, ctx []interface{}) {
	for i := 0; i < len(ctx); i += 2 {
		k, ok := ctx[i].(string)
		v := fmt.Sprintf("%+v", ctx[i+1])
		if !ok {
			k, v = "ERR", fmt.Sprintf("%+v", k)
		}
		fmt.Fprintf(buf, " %s=%s", k, v)
	}
	buf.WriteByte('\n')
}

func terminalFormat() log.Format {
	return log.FormatFunc(func(r *log.Record) []byte {
		b := &bytes.Buffer{}
		fmt.Fprintf(b, "[%v][%s] %s", toLvlStr(r.Lvl), r.Time.Format(timeFormat), r.Msg)
		// if len(r.Ctx) > 0 && len(r.Msg) < termMsgJust {
		// 	b.Write(bytes.Repeat([]byte{' '}, termMsgJust-len(r.Msg)))
		// }
		printLogCtx(b, r.Ctx)
		return b.Bytes()
	})
}

// SetLvl set new level 
func SetLvl(lvl string)  {
	logger.SetHandler(getLvlFilterHandler(lvl))
	pkgLvl = lvl
}
