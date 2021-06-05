package validation

import (
    "regexp"
)

type Validator struct {}

func (v *Validator) ValidateString(value string, rule Rule) FieldResult {
    if l := int64(len(value)); l < rule.Min {
        return Short
    } else if l > rule.Max {
        return Long
    }
    
    ok, _ := regexp.MatchString(rule.Regex, value)
    if ok {
        return Valid
    }
    
    return Incorrect
}

func (v *Validator) ValidateInt(value int64, rule Rule) FieldResult {
    if value < rule.Min {
        return Short
    }
    if value > rule.Max {
        return Long
    }
    return Valid
}