package utils

import (
	"fmt"
	"log"
)

const (
	warning = 34
	info    = 33
	success = 32
	error   = 31
)

func baseLogging(color int, state string, messages ...any) []any {
	return append([]any{fmt.Sprintf("\x1b[%dm[%s]\x1b[0m", color, state)}, messages...)
}

func Warning(messages ...any) {
	log.Println(baseLogging(warning, "WARNING", messages...)...)
}

func Info(messages ...any) {
	log.Println(baseLogging(info, " INFO  ", messages...)...)
}

func Success(messages ...any) {
	log.Println(baseLogging(success, "SUCCESS", messages...)...)
}

func SafeError(messages ...any) {
	log.Println(baseLogging(error, " ERROR ", messages...)...)
}

func Error(messages ...any) {
	log.Fatalln(baseLogging(error, " ERROR ", messages...)...)
}
