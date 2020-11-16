package app

import "regexp"

// Shows Invoice Number
type InvNumber string

func (i *InvNumber) Match(text string) string{
	matched := regexp.MustCompile(string(*i)).FindString(text)
	return matched
}

// Shows sender of Money
type Sender string

func (s *Sender) Match(text string) string{
	matched := regexp.MustCompile(string(*s)).FindString(text)
	return matched
}

// Shows sender bin
type SenderBIN string

func (sbi *SenderBIN) Match(text string) string{
	matched := regexp.MustCompile(string(*sbi)).FindString(text)
	return matched
}

// Shows sender IIK
type SenderIIK string

func (si *SenderIIK) Match(text string) string{
	matched := regexp.MustCompile(string(*si)).FindString(text)
	return matched
}

// Shows Sender Bank
type SenderBank string

func (sb *SenderBank) Match(text string) string{
	matched := regexp.MustCompile(string(*sb)).FindString(text)
	return matched
}

// Shows Sender BIK
type SenderBankBIK string

func (sbb *SenderBankBIK) Match(text string) string{
	matched := regexp.MustCompile(string(*sbb)).FindString(text)
	return matched
}

// Shows Sender Code
type SenderCode string

func (sc *SenderCode) Match(text string) string{
	matched := regexp.MustCompile(string(*sc)).FindString(text)
	return matched
}

// Shows payer Code
type PayerCode string

func (pc *PayerCode) Match(text string) string{
	matched := regexp.MustCompile(string(*pc)).FindString(text)
	return matched
}

// Shows Producer Company
type ProduceCompany string

func (pcy *ProduceCompany) Match(text string) string{
	matched := regexp.MustCompile(string(*pcy)).FindString(text)
	return matched
}

// Shows Producer Company KBE
type ProduceCompanyKBE string

func (pck *ProduceCompanyKBE) Match(text string) string{
	matched := regexp.MustCompile(string(*pck)).FindString(text)
	return matched
}

// Shows Producer Company BIN
type ProduceCompanyBIN string

func (pcb *ProduceCompanyBIN) Match(text string) string{
	matched := regexp.MustCompile(string(*pcb)).FindString(text)
	return matched
}

// Shows Producer Company IIK
type ProduceCompanyIIK string

func (pci *ProduceCompanyIIK) Match(text string) string{
	matched := regexp.MustCompile(string(*pci)).FindString(text)
	return matched
}

// Shows Sun
type Sum string

func (sm *Sum) Match(text string) string{
	matched := regexp.MustCompile(string(*sm)).FindString(text)
	return matched
}

// Shows NDS
type SumNDS string

func (sn *SumNDS) Match(text string) string{
	matched := regexp.MustCompile(string(*sn)).FindString(text)
	return matched
}

// Shows Date
type DateProd string

func (dp *DateProd) Match(text string) string{
	matched := regexp.MustCompile(string(*dp)).FindString(text)
	return matched
}



