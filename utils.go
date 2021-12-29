package gromer

import (
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/imdario/mergo"
)

var Validator = validator.New()
var ValidatorErrorMap = map[string]string{
	"required": "is required",
}
var upperRegex = regexp.MustCompile("^[^a-z]*$")

type timeTransformer struct {
}

func (t timeTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				srcResult := src.MethodByName("IsZero").Call([]reflect.Value{})
				if !srcResult[0].Bool() {
					dst.Set(src)
				}
			}
			return nil
		}
	}
	return nil
}

func Merge(dst interface{}, src interface{}) error {
	err := mergo.Merge(dst, src, mergo.WithOverride, mergo.WithTransformers(timeTransformer{}))
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
