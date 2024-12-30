package logger

import "log"

func Init() {
	log.SetFlags(log.Llongfile | log.Lmsgprefix)
}
