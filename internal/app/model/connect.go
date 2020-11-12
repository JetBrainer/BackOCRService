package model

import (
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"log"
)

func RuleUsageLocal(res apiserver.OCRText) (string,string,string,string,string,string,string,string,string){
	log.Println(res.JustText())

	// Rules
	numAndData := invoiceNumAndData(res.JustText())
	pay1 := payer(res.JustText())
	produce := producer(res.JustText())
	requis := requisites(res.JustText())
	sumWithTaxes := sumWithTax(res.JustText())
	amountOfProd := amount(res.JustText())
	signer := followed(res.JustText())
	fullSum := sumTax(res.JustText())
	product := prodName(res.JustText())

	return numAndData,pay1,produce,requis,sumWithTaxes, amountOfProd,signer,fullSum,product
}