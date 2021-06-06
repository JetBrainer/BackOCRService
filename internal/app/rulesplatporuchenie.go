package app

import (
	"regexp"
	"strings"
)

// Shows Invoice Number
type InvNumber string

func (i *InvNumber) Match(text string) string {
	matched := regexp.MustCompile(string(*i)).FindString(text)
	val := strings.Split(matched, " ")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows sender of Money
type Sender string

func (s *Sender) Match(text string) string {
	matched := regexp.MustCompile(string(*s)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows sender bin
type SenderBIN string

func (sbi *SenderBIN) Match(text string) string {
	matched := regexp.MustCompile(string(*sbi)).FindAllString(text, 2)
	if matched == nil {
		return ""
	}
	val := strings.Split(matched[0], " ")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows sender IIK
type SenderIIK string

func (si *SenderIIK) Match(text string) string {
	matched := regexp.MustCompile(string(*si)).FindString(text)
	val2 := strings.TrimSpace(matched)
	return val2
}

// Shows Sender Bank
type SenderBank string

func (sb *SenderBank) Match(text string) string {
	matched := regexp.MustCompile(string(*sb)).FindString(text)
	val := strings.Split(matched, " ")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Sender BIK
type SenderBankBIK string

func (sbb *SenderBankBIK) Match(text string) string {
	matched := regexp.MustCompile(string(*sbb)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Sender Code
type SenderCode string

func (sc *SenderCode) Match(text string) string {
	matched := regexp.MustCompile(string(*sc)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows payer Code
type PayerCode string

func (pc *PayerCode) Match(text string) string {
	matched := regexp.MustCompile(string(*pc)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Producer Company
type ProduceCompany string

func (pcy *ProduceCompany) Match(text string) string {
	matched := regexp.MustCompile(string(*pcy)).FindString(text)
	if matched == "" {
		return ""
	}
	val := strings.Split(matched, " ")
	val2 := strings.Join(val[1:len(val)-1], " ")
	return val2
}

// Shows Producer Company KBE
type ProduceCompanyKBE string

func (pck *ProduceCompanyKBE) Match(text string) string {
	matched := regexp.MustCompile(string(*pck)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Producer Company BIN
type ProduceCompanyBIN string

func (pcb *ProduceCompanyBIN) Match(text string) string {
	matched := regexp.MustCompile(string(*pcb)).FindAllString(text, 2)
	if matched == nil || len(matched) != 1{
		return ""
	}
	val := strings.Split(matched[0], " ")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Producer Company IIK
type ProduceCompanyIIK string

func (pci *ProduceCompanyIIK) Match(text string) string {
	matched := regexp.MustCompile(string(*pci)).FindString(text)
	return matched
}

// Shows Sun
type Sum string

func (sm *Sum) Match(text string) string {
	matched := regexp.MustCompile(string(*sm)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows NDS
type SumNDS string

func (sn *SumNDS) Match(text string) string {
	matched := regexp.MustCompile(string(*sn)).FindString(text)
	val := strings.Split(matched, "\r\n")
	val2 := strings.TrimSpace(val[len(val)-1])
	return val2
}

// Shows Date
type DateProd string

func (dp *DateProd) Match(text string) string {
	matched := regexp.MustCompile(string(*dp)).FindString(text)
	val2 := strings.TrimSpace(matched)
	return val2
}
