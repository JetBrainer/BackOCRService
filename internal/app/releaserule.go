package app


type DocStr struct {
	DocType		string `json:"doc_type"`
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
	d.DocType = "Счет Фактуры"

	invAndData := InvNumAndData("[Сс][Чч][Ее][Тт].*(\\s|\\s\\s)N.*")
	d.DataNum = invAndData.Match(text)

	payer := Payer("(\\W{3}(\\s|\\s\\s)(«|\\\")\\W{1,}\\d+\\W+\\d+|(^\\W+:|^.{1,}:)(\\s|\\s\\s)\\W{1,}\\d{1,}\\s\\W{1,}\\d{1,}(\\s|\\s\\s)\\W+)")
	d.Payer = payer.Match(text)

	produce := Producer("(П.{1,10}|^.........:)\\s.*(”|“)")
	d.Producer = produce.Match(text)

	requis := Requisites("(?m)(^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\sБ|^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\s^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\s.*\\s.*)")
	d.Requis = requis.Match(text)

	sumWTax := SumWithTax("....у(\\s|\\s\\s)(\\d\\,\\s|\\d\\'\\s|\\d\\s|\\d)\\d{3}\\,(\\d{2}|\\d{3}\\,|}|\\d{3}\\s)")
	d.SumNTax = sumWTax.Match(text)

	amount := Amount("[^\\d]{12}(\\s|\\s\\s)\\d(\\.|\\,)")
	d.AmountOf = amount.Match(text)

	follow := Followed("Ру.{1,}\\s\\W{1,}")
	d.Signed = follow.Match(text)

	fullSum := SumTax("Су\\W{1,}.*\\s.*\\s.*(\\s\\d.*)")
	d.FullPrice = fullSum.Match(text)

	prodN := ProdName("(?m)^[Н].*(\\s|\\s\\s).*(\\s|\\s\\s).*")
	d.Prod = prodN.Match(text)

}