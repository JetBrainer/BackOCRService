package model

import (
	"regexp"
)

//func invoiceNum(text []string) string{
//	Schet := "СЧЕТ"
//	result := ""
//
//	for i:=0; i<len(text);i++{
//		if strings.Contains(text[i],Schet){
//			result = text[i+2]
//		}
//	}
//	return result
//}
func invoiceNumAndData(text string) string{
	pattern := "(....|....(\\s|\\s\\s)..(\\s|\\s\\s).{4,5}.)(\\s|\\s\\s)N.(\\s|\\s\\s)[0-9]{2,4}(\\s|\\s\\s)..(\\s|\\s\\s)[1-3][1-9]((\\s|\\s\\s)\\D{3,8}|\\.(0|1)[1-2]\\.)(\\d{4}|\\d{2})"
	Str := match(pattern,text)
	return Str
}
func match(pattern string, text string) string {
	matched := regexp.MustCompile(pattern).FindString(text)
	return matched
}


func payer(text string)string{
	pattern := "(\\W{3}(\\s|\\s\\s)(«|\\\")\\W{1,}\\d+\\W+\\d+|(^\\W+:|^.{1,}:)(\\s|\\s\\s)\\W{1,}\\d{1,}\\s\\W{1,}\\d{1,}(\\s|\\s\\s)\\W+)"
	Str := match(pattern,text)
	return Str
}
func producer(text string) string{
	pattern := "(П.{1,10}|^.........:)\\s.*(”|“)"
	//\W{3}(\s|\s\s)(\.|\")\W{1,}|(^\W+:|^.{1,}:)(\s|\s\s)\W{1,}“
	Str := match(pattern,text)
	return Str
}

func requisites(text string) string{
	//pattern1 := "(?m)(^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\sБ|^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\s^[Сс][Чч](,\\s|.\\s|\\s\\s).*\\s.*\\s.*)"
	pattern2 := "[4]\\d{19}"
	Str := match(pattern2, text)
	return Str
}

func sumWithTax(text string) string{
	pattern := "....у(\\s|\\s\\s)(\\d\\,\\s|\\d\\'\\s|\\d\\s|\\d)\\d{3}\\,(\\d{2}|\\d{3}\\,|}|\\d{3}\\s)"
	Str := match(pattern, text)
	return Str
}

func amount(text string) string{
	pattern := "[^\\d]{12}(\\s|\\s\\s)\\d(\\.|\\,)"
	Str := match(pattern,text)
	return Str
}

func followed(text string) string{
	pattern := "Ру.{1,}\\s\\W{1,}"
	Str := match(pattern,text)
	return Str
}

func sumTax(text string) string{
	pattern := "Су\\W{1,}.*\\s.*\\s.*(\\s\\d.*)"
	Str := match(pattern, text)
	return Str
}

func prodName(text string) string{
	pattern := "(?m)(^[Тт]о...(а|))\\s.*\\s.*\\s.*"
	Str := match(pattern, text)
	return Str
}