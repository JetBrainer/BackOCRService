package model

import (
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"log"
)

func RuleUsageLocal(config *apiserver.Config){
	//res, err := config.ParseFromLocal("images/photo_2020-11-09_23-08-13.jpg")
	res2, err := config.ParseFromURL("https://delovoymir.biz/res/images/uploaded/articles/img/488934743833454.png")
	log.Println(res2.JustText())
	if err != nil{
		log.Println("Parse Local Error", err)
	}
	//newText := strings.Split(res.JustText()," ")
	numWord := invoiceNumAndData(res2.JustText())
	numWord2 := payer(res2.JustText())
	numWord3 := producer(res2.JustText())
	numWord4 := requisites(res2.JustText())
	numWord5 := sumWithTax(res2.JustText())
	numWord6 := amount(res2.JustText())
	numWord7 := followed(res2.JustText())
	numWord8 := sumTax(res2.JustText())
	numWord9 := prodName(res2.JustText())
	log.Println(numWord)
	log.Println(numWord2)
	log.Println(numWord3)
	log.Println(numWord4)
	log.Println(numWord5)
	log.Println(numWord6)
	log.Println(numWord7)
	log.Println(numWord8)
	log.Println(numWord9)
}