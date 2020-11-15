package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(cond bool) validation.RuleFunc{
	return func(v interface{}) error {
		if cond{
			return validation.Validate(v, validation.Required)
		}
		return nil
	}
}