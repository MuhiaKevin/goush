package validator

import (
	"fmt"
	"net/url"
	"strings"
)

type Validator struct {
    FieldErrors map[string]string
}


func (v *Validator) valid() bool  {
    return len(v.FieldErrors) == 0
}

func (v *Validator) Valid() bool {
    return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
     if v.FieldErrors == nil {
        v.FieldErrors = make(map[string]string)
    }

     if _, exists := v.FieldErrors[key]; !exists {
        v.FieldErrors[key] = message
    }
}

func (v *Validator) CheckField(ok bool, key, message string) {
    if !ok {
        v.AddFieldError(key, message)
    }
}


func NotBlank(value string) bool {
    return strings.TrimSpace(value) != ""
}



func IsURL(value string) bool {
    _, err := url.ParseRequestURI(value)
    fmt.Println(value)

    if err != nil {
        return false
    }

    return true
}
