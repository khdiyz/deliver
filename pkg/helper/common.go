package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func IsValidBirthYear(birthYear string) (bool, error) {
	if len(birthYear) != 4 {
		return false, fmt.Errorf("invalid birth year: %s", birthYear)
	}

	birthYearInt, err := strconv.Atoi(birthYear)
	if err != nil {
		return false, fmt.Errorf("invalid birth year: %s", birthYear)
	}

	currentYear := time.Now().Year()

	if birthYearInt < 1900 && birthYearInt > currentYear {
		return false, fmt.Errorf("birth year must be between 1900 and %d", currentYear)
	}

	return true, nil
}

func IsValidPhoneNumber(phoneNumber string) (bool, error) {
	var uzbPhonePattern = `^\+998\d{9}$`

	re := regexp.MustCompile(uzbPhonePattern)

	isValid := re.MatchString(phoneNumber)

	if !isValid {
		return isValid, fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	return true, nil
}

func TruncateTime(input time.Time) time.Time {
	return input.Truncate(24 * time.Hour)
}
