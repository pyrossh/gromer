package gromer

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/imdario/mergo"
)

var Validator = validator.New()
var ValidatorErrorMap = map[string]string{
	"required": "is required",
}
var upperRegex = regexp.MustCompile("^[^a-z]*$")

func Merge(dst interface{}, src interface{}) error {
	err := mergo.Merge(dst, src, mergo.WithOverwriteWithEmptyValue)
	if err != nil {
		return err
	}
	return Validate(dst)
}

func Validate(dst interface{}) error {
	return Validator.Struct(dst)
}

func RegisterValidation(k, msg string, validate func(fl validator.FieldLevel) bool) {
	ValidatorErrorMap[k] = msg
	Validator.RegisterValidation(k, validate)
}

func GetValidationError(err validator.ValidationErrors) map[string]string {
	emap := map[string]string{}
	for _, e := range err {
		parts := strings.Split(e.StructNamespace(), ".")
		lowerParts := []string{}
		for _, p := range parts[1:] {
			if upperRegex.MatchString(p) {
				lowerParts = append(lowerParts, strings.ToLower(p))
			} else {
				lowerParts = append(lowerParts, strcase.ToLowerCamel(p))
			}
		}
		k := strings.Join(lowerParts, ".")
		errorMsg, ok := ValidatorErrorMap[e.Tag()]
		if ok {
			emap[k] = errorMsg
		} else {
			emap[k] = "is not valid" // e.Error()
		}
	}
	return emap
}
