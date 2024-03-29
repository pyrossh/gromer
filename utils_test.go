package gromer

import (
	"regexp"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func init() {
	pinCodeRegex := regexp.MustCompile("^[1-9]{1}[0-9]{2}\\s{0,1}[0-9]{3}$")
	RegisterValidation("pincode", "is not in valid format", func(fl validator.FieldLevel) bool {
		return pinCodeRegex.MatchString(fl.Field().String())
	})
}

func TestValidator(t *testing.T) {
	assert.NoError(t, Validator.Var("560001", "pincode"))
	assert.EqualError(t, Validator.Var("ABCD", "pincode"), "Key: '' Error:Field validation for '' failed on the 'pincode' tag")
}

type Todo struct {
	ID        string    `json:"id" validate:"required"`
	OrgID     string    `json:"orgId" validate:"required"`
	Pincode   string    `json:"pincode" validate:"required,pincode"`
	PAN       string    `json:"pan" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func TestValidate(t *testing.T) {
	todo := &Todo{
		ID:      "123",
		OrgID:   "",
		Pincode: "",
		PAN:     "",
	}
	err := Validate(todo)
	validationErrors := err.(validator.ValidationErrors)
	assert.Equal(t, GetValidationError(validationErrors), map[string]string{
		"pincode": "is required",
		"pan":     "is required",
		"orgId":   "is required",
	})

	todo.Pincode = "AWS"
	todo.OrgID = "123"
	err = Validate(todo)
	validationErrors = err.(validator.ValidationErrors)
	assert.EqualValues(t, map[string]string{
		"pincode": "is not in valid format",
		"pan":     "is required",
	}, GetValidationError(validationErrors))
}

type Note struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func TestMerge(t *testing.T) {
	note := &Note{
		ID:        "",
		Text:      "",
		CreatedAt: time.Time{},
	}
	err := Merge(note, &Note{
		ID:        "1",
		Text:      "1",
		CreatedAt: time.Date(2020, 10, 10, 0, 0, 0, 0, time.UTC),
	})
	assert.NoError(t, err)
	assert.EqualValues(t, &Note{
		ID:        "1",
		Text:      "1",
		CreatedAt: time.Date(2020, 10, 10, 0, 0, 0, 0, time.UTC),
	}, note)
	err = Merge(note, &Note{
		ID:        "2",
		Text:      "2",
		CreatedAt: time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC),
	})
	assert.NoError(t, err)
	assert.EqualValues(t, &Note{
		ID:        "2",
		Text:      "2",
		CreatedAt: time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC),
	}, note)
	// TODO: look at this merge
	err = Merge(note, &Note{
		ID:        "2",
		Text:      "",
		CreatedAt: time.Time{},
	})
	assert.NoError(t, err)
	assert.EqualValues(t, &Note{
		ID:        "2",
		Text:      "2",
		CreatedAt: time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC),
	}, note)
}
