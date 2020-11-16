package app

import (
	"github.com/JetBrainer/BackOCRService/internal/app/model"
)

type DocStr struct {
	DataNum 	string `json:"data_num"`
	Payer 		string `json:"payer"`
	Producer 	string `json:"producer"`
	Requis		string `json:"requis"`
	SumNTax		string `json:"sum_n_tax"`
	AmountOf	string `json:"amount_of"`
	Signed		string `json:"signed"`
	FullPrice	string `json:"full_price"`
	Prod		string `json:"prod"`
}

func (d *DocStr)RuleDocUsage(text string) {

	invAndData := model.InvNumAndData("(....|....(\\s|\\s\\s)..(\\s|\\s\\s).{4,5}.)(\\s|\\s\\s)N.(\\s|\\s\\s)[0-9]{2,4}(\\s|\\s\\s)..(\\s|\\s\\s)[1-3][1-9]((\\s|\\s\\s)\\D{3,8}|\\.(0|1)[1-2]\\.)(\\d{4}|\\d{2})")
	d.DataNum = invAndData.Match(text)

	payer := model.Payer("(\\W{3}(\\s|\\s\\s)(«|\\\")\\W{1,}\\d+\\W+\\d+|(^\\W+:|^.{1,}:)(\\s|\\s\\s)\\W{1,}\\d{1,}\\s\\W{1,}\\d{1,}(\\s|\\s\\s)\\W+)")
	d.Payer = payer.Match(text)

	produce := model.Producer("(П.{1,10}|^.........:)\\\\s.*(”|“)")
	d.Producer = produce.Match(text)

	requis := model.Requisites("(?m)(^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\sБ|^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s.*\\\\s.*)")
	d.Requis = requis.Match(text)

	sumWTax := model.SumWithTax("....у(\\s|\\s\\s)(\\d\\,\\s|\\d\\'\\s|\\d\\s|\\d)\\d{3}\\,(\\d{2}|\\d{3}\\,|}|\\d{3}\\s)")
	d.SumNTax = sumWTax.Match(text)

	amount := model.Amount("[^\\\\d]{12}(\\\\s|\\\\s\\\\s)\\\\d(\\\\.|\\\\,)")
	d.AmountOf = amount.Match(text)

	follow := model.Followed("Ру.{1,}\\s\\W{1,}")
	d.Signed = follow.Match(text)

	fullSum := model.SumTax("Су\\W{1,}.*\\s.*\\s.*(\\s\\d.*)")
	d.FullPrice = fullSum.Match(text)

	prodN := model.ProdName("(?m)(^[Тт]о...(а|))\\s.*\\s.*\\s.*")
	d.Prod = prodN.Match(text)

}