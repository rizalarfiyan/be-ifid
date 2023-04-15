package utils

import (
	"encoding/json"
	"log"
)

func PrettyPrint(i interface{}) {
	byteStr, _ := json.MarshalIndent(i, "", "\t")
	log.Println(baseLogging(typeWarning, "DEBUG", "\n"+string(byteStr)+"\n")...)
}
