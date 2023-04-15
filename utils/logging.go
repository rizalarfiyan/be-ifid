package utils

import (
	"fmt"
	"log"
)

const (
	typeWarning = 34
	typeInfo    = 33
	typeSuccess = 32
	typeError   = 31
)

func baseLogging(color int, state string, messages ...any) []any {
	return append([]any{fmt.Sprintf("\x1b[%dm[%s]\x1b[0m", color, state)}, messages...)
}

func Warning(messages ...any) {
	log.Println(baseLogging(typeWarning, "WARNING", messages...)...)
}

func Info(messages ...any) {
	log.Println(baseLogging(typeInfo, " INFO  ", messages...)...)
}

func Success(messages ...any) {
	log.Println(baseLogging(typeSuccess, "SUCCESS", messages...)...)
}

func SafeError(messages ...any) {
	log.Println(baseLogging(typeError, " ERROR ", messages...)...)
}

func Error(messages ...any) {
	log.Fatalln(baseLogging(typeError, " ERROR ", messages...)...)
}
