package validator

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Custom validation function for Uzbek phone numbers starting with +998
func uzbPhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	var uzbPhonePattern = `^\+998\d{9}$`
	re := regexp.MustCompile(uzbPhonePattern)
	return re.MatchString(phone)
}

// Custom validation function for checking if a date string follows the "2006-01-02" format
func dateValidator(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	var datePattern = `^\d{4}-\d{2}-\d{2}$`
	re := regexp.MustCompile(datePattern)

	if !re.MatchString(date) {
		return false
	}

	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
