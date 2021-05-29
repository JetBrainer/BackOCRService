package app

import (
	"log"
	"reflect"
)

func MultiplexDocument(text string) string {
	invSchet := DocStr{}
	invSchet.RuleDocUsage(text)

	var counter int

	docPlat := DocPlatPoruchenie{}
	docPlat.RuleDocUsage(text)

	docPlatSchet := reflect.ValueOf(docPlat)
	for i := 0; i < docPlatSchet.NumField(); i++ {
		values := docPlatSchet.Field(i).String()
		if values != "" {
			counter++
		}
	}
	if counter >= docPlatSchet.NumField()-3 {
		return "Платежное Поручение"
	}

	return "Счет Фактура"
}

func GetMultiplexer(text string) Document {
	switch MultiplexDocument(text) {
	case "Счет Фактура":
		log.Println("Счет фактура")
		return &DocStr{}
	case "Платежное Поручение":
		log.Println("Платежное Поручение")
		return &DocPlatPoruchenie{}
	default:
		log.Println("Не распознано")
		return nil
	}
}
