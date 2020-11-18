package app

type DocPlatPoruchenie struct {
	InvoiceNumber 			string	`json:"invoice_number"`
	SenderCompany 			string	`json:"sender_company"`
	SenderCompanyBIN		string	`json:"sender_company_bin"`
	SenderCompanyIIK		string	`json:"sender_company_iik"`
	SenderCompanyBank		string	`json:"sender_company_bank"`
	SenderCompanyBankBIK	string	`json:"sender_company_bank_bik"`
	SenderCompanyCode		string	`json:"sender_company_code"`
	PayerCompanyCode		string	`json:"payer_company_code"`
	ProducerCompany			string	`json:"producer_company"`
	ProducerCompanyKBE		string	`json:"producer_company_kbe"`
	ProducerCompanyBIN		string	`json:"producer_company_bin"`
	ProducerCompanyIIK		string	`json:"producer_company_iik"`
	ProducerSenderSum		string	`json:"producer_sender_sum"`
	ProducerSenderSumNDS	string	`json:"producer_sender_sum_nds"`
	DateRelease				string	`json:"date_release"`
}

// Rule for "Платежное поручение"
// Every value made according to document
func (d *DocPlatPoruchenie) RuleDocUsage(text string) {
	invNum := InvNumber("")
	d.InvoiceNumber = invNum.Match(text)

	sender := Sender("")
	d.SenderCompany = sender.Match(text)

	senderBIN := SenderBIN("")
	d.SenderCompanyBIN = senderBIN.Match(text)

	senderIIK := SenderIIK("")
	d.SenderCompanyIIK = senderIIK.Match(text)

	senderBank := SenderBank("")
	d.SenderCompanyBank = senderBank.Match(text)

	senderBankBIK := SenderBankBIK("")
	d.SenderCompanyBankBIK = senderBankBIK.Match(text)

	senderCode := SenderCode("")
	d.SenderCompanyCode = senderCode.Match(text)

	payerCode := PayerCode("")
	d.PayerCompanyCode = payerCode.Match(text)

	produce := ProduceCompany("")
	d.ProducerCompany = produce.Match(text)

	produceKBE := ProduceCompanyKBE("")
	d.ProducerCompanyKBE = produceKBE.Match(text)

	produceBIN := ProduceCompanyBIN("")
	d.ProducerCompanyBIN = produceBIN.Match(text)

	produceIIK := ProduceCompanyIIK("")
	d.ProducerCompanyIIK = produceIIK.Match(text)

	sum := Sum("")
	d.ProducerSenderSum = sum.Match(text)

	sumNDS := SumNDS("")
	d.ProducerSenderSumNDS = sumNDS.Match(text)

	dateRelease := DateProd("")
	d.DateRelease = dateRelease.Match(text)
}