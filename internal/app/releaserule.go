package app

import (
	"github.com/JetBrainer/BackOCRService/internal/app/model"
)

func RuleDocUsage(text string) (string,string,string,string,string,string,string,string,string){
	//invAndDate := &model.DocStruct{InvNumAndData:"fbsdofbsdoifoid"}
	//invAndDate.Match(jValue.JustText())

	invAndData := model.InvNumAndData("(....|....(\\s|\\s\\s)..(\\s|\\s\\s).{4,5}.)(\\s|\\s\\s)N.(\\s|\\s\\s)[0-9]{2,4}(\\s|\\s\\s)..(\\s|\\s\\s)[1-3][1-9]((\\s|\\s\\s)\\D{3,8}|\\.(0|1)[1-2]\\.)(\\d{4}|\\d{2})")
	dateNumAndInv := invAndData.Match(text)

	payer := model.Payer("(\\W{3}(\\s|\\s\\s)(«|\\\")\\W{1,}\\d+\\W+\\d+|(^\\W+:|^.{1,}:)(\\s|\\s\\s)\\W{1,}\\d{1,}\\s\\W{1,}\\d{1,}(\\s|\\s\\s)\\W+)")
	pay := payer.Match(text)

	produce := model.Producer("(П.{1,10}|^.........:)\\\\s.*(”|“)")
	producer := produce.Match(text)

	requis := model.Requisites("(?m)(^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\sБ|^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s.*\\\\s.*)")
	requisite := requis.Match(text)

	sumWTax := model.SumWithTax("....у(\\s|\\s\\s)(\\d\\,\\s|\\d\\'\\s|\\d\\s|\\d)\\d{3}\\,(\\d{2}|\\d{3}\\,|}|\\d{3}\\s)")
	sumNTax := sumWTax.Match(text)

	amount := model.Amount("[^\\\\d]{12}(\\\\s|\\\\s\\\\s)\\\\d(\\\\.|\\\\,)")
	amountOf := amount.Match(text)

	follow := model.Followed("Ру.{1,}\\s\\W{1,}")
	signed := follow.Match(text)

	fullSum := model.SumTax("Су\\W{1,}.*\\s.*\\s.*(\\s\\d.*)")
	fullPrice := fullSum.Match(text)

	prodN := model.ProdName("(?m)(^[Тт]о...(а|))\\s.*\\s.*\\s.*")
	prod := prodN.Match(text)

	return dateNumAndInv,pay,producer,requisite,sumNTax,amountOf,signed,fullPrice,prod
}