package app

import (
	"log"
	"reflect"
)

func MultiplexDocument(text string) string{
	invSchet := DocStr{}
	invSchet.RuleDocUsage(text)

	vSchet := reflect.ValueOf(invSchet)


	var counter int
	for i:=0;i< vSchet.NumField();i++{
		values := vSchet.Field(i).String()
		if values != ""{
			counter++
		}
	}
	if counter == vSchet.NumField(){
		return "Счет Фактура"
	}

	counter = 0
	docPlat := DocPlatPoruchenie{}
	docPlat.RuleDocUsage(text)

	docPlatSchet := reflect.ValueOf(docPlat)
	for i:=0;i< docPlatSchet.NumField();i++{
		values := docPlatSchet.Field(i).String()
		if values != ""{
			counter++
		}
	}
	if counter == docPlatSchet.NumField(){
		return "Платежное Поручение"
	}

	return ""
}


func GetMultiplexer(text string) Document{
	valType := MultiplexDocument(text)

	switch valType {
	case "Счет Фактура":
		log.Println("Счет фактура")
		invSchet := DocStr{}
		return &invSchet
	case "Платежное Поручение":
		log.Println("Платежное Поручение")
		docPlat := DocPlatPoruchenie{}
		return &docPlat
	case "":
		log.Println("Не распознано")
		return nil
	}

	return nil
}