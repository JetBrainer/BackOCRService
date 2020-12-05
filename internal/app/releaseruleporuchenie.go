package app

type DocPlatPoruchenie struct {
	DocType					string	`json:"doc_type"`
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
	d.DocType = "Платежное поручение"

	invNum := InvNumber("[П|п].*")
	d.InvoiceNumber = invNum.Match(text)

	sender := Sender("[О].*(\\s|\\s\\s)д.*(\\s|\\s\\s).*(\\s|\\s\\s).*")
	d.SenderCompany = sender.Match(text)

	senderBIN := SenderBIN("БИН.*")
	d.SenderCompanyBIN = senderBIN.Match(text)

	senderIIK := SenderIIK("k.*")
	d.SenderCompanyIIK = senderIIK.Match(text)

	senderBank := SenderBank("[Б|б].{3}(\\s|\\s\\s)о.*(\\s|\\s\\s)Б.*(\\s|\\s\\s).*")
	d.SenderCompanyBank = senderBank.Match(text)

	senderBankBIK := SenderBankBIK("БИК.*(\\s|\\s\\s).*\\s.*")
	d.SenderCompanyBankBIK = senderBankBIK.Match(text)

	senderCode := SenderCode("код(\\s|\\s\\s)\\d\\d")
	d.SenderCompanyCode = senderCode.Match(text)

	payerCode := PayerCode("Код.*(\\s|\\s\\s).*(\\s|\\s\\s).*(\\s|\\s\\s).*(\\s|\\s\\s).*(\\s|\\s\\s).*")
	d.PayerCompanyCode = payerCode.Match(text)

	produce := ProduceCompany("Бе.*(\\s|\\s\\s).*")
	d.ProducerCompany = produce.Match(text)

	produceKBE := ProduceCompanyKBE("КБ.*(\\s|\\s\\s).*")
	d.ProducerCompanyKBE = produceKBE.Match(text)

	produceBIN := ProduceCompanyBIN("БИН.*")
	d.ProducerCompanyBIN = produceBIN.Match(text)

	produceIIK := ProduceCompanyIIK("kZ\\w{18}")
	d.ProducerCompanyIIK = produceIIK.Match(text)

	sum := Sum("С.{4}(\\s|\\s\\s)\\d{1,10}(\\s|\\s\\s).*")
	d.ProducerSenderSum = sum.Match(text)

	sumNDS := SumNDS("НДС.*(\\s|\\s\\s).*")
	d.ProducerSenderSumNDS = sumNDS.Match(text)

	dateRelease := DateProd("\\d\\d\\.\\d\\d.*")
	d.DateRelease = dateRelease.Match(text)
}