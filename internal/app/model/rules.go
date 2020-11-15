package model

import (
	"regexp"
)

type Docs interface{
	Match(string) string
}

// Shows us invoiceNum and Data
type InvNumAndData string

// Take regexp pattern and shows
// result due to pattern
func (in *InvNumAndData)Match(text string) string {
	matched := regexp.MustCompile(string(*in)).FindString(text)

	return matched
}

// Shows us payer in Invoice like
// 'Плательщик'
type Payer string

func (p *Payer)Match(text string) string {
	matched := regexp.MustCompile(string(*p)).FindString(text)
	return matched
}

// Shows us producer in Invoice 'Поставщик'
type Producer string

func (p *Producer)Match(text string) string {
	matched := regexp.MustCompile(string(*p)).FindString(text)
	return matched
}

// Shows Bank number data
type Requisites string

func (p *Requisites)Match(text string) string {
	matched := regexp.MustCompile(string(*p)).FindString(text)
	return matched
}

// Shows our full result of sum with tax 'НДС'
type SumWithTax string

func (s *SumWithTax)Match(text string) string {
	matched := regexp.MustCompile(string(*s)).FindString(text)
	return matched
}

// Shows amount of stuff
type Amount string

func (a *Amount)Match(text string) string {
	matched := regexp.MustCompile(string(*a)).FindString(text)
	return matched
}

// Shows who signed the invoice
type Followed string

func (f *Followed)Match(text string) string {
	matched := regexp.MustCompile(string(*f)).FindString(text)
	return matched
}

// Shows full sum with product
type SumTax string

func (st *SumTax)Match(text string) string {
	matched := regexp.MustCompile(string(*st)).FindString(text)
	return matched
}

// Shows product name
type ProdName string

func (pn *ProdName)Match(text string) string {
	matched := regexp.MustCompile(string(*pn)).FindString(text)
	return matched
}
