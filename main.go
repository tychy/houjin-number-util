package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

var (
	ErrInvalidCharacter          = errors.New("invalid character")
	ErrInvalidHoujinNumberLength = errors.New("invalid houjin number length")
	ErrInvalidHoujinNumber       = errors.New("invalid houjin number")
	ErrInvalidCheckDigit         = errors.New("invalid check digit")
)

func validateHoujinNumber(length int, houjinNumber string) error {
	if len(houjinNumber) != length {
		return ErrInvalidHoujinNumberLength
	}
	for i := 0; i < length; i++ {
		if houjinNumber[i] < '0' || houjinNumber[i] > '9' {
			return ErrInvalidCharacter
		}
	}
	return nil
}

func ValidateCheckSum(houjinNumber string) error {
	if err := validateHoujinNumber(13, houjinNumber); err != nil {
		return err
	}

	checkDigit := calculateCheckDigit(houjinNumber[1:])
	if checkDigit != string(houjinNumber[0]) {
		return ErrInvalidCheckDigit
	}
	return nil
}

func validateToukijoCodeOrgCode(code, org string) error {
	if !slices.Contains(ToukijoCodes, code) {
		return ErrInvalidHoujinNumber
	}
	if !slices.Contains(OrganizationCodes, org) {
		return ErrInvalidHoujinNumber
	}
	return nil
}

func ValidateHoujinNumber(houjinNumber string) error {
	if err := ValidateCheckSum(houjinNumber); err != nil {
		return err
	}

	top := houjinNumber[1]

	switch top {
	case '0':
		govCode := houjinNumber[1:7]
		if slices.Contains(GovermentCodes, govCode) {
			return nil
		}
		return validateToukijoCodeOrgCode(houjinNumber[1:5], houjinNumber[5:7])

	case '1', '2', '3', '4', '5':
		code := houjinNumber[1:5]
		org := houjinNumber[5:7]
		return validateToukijoCodeOrgCode(code, org)

	case '7':
		return nil
	default:
		return ErrInvalidHoujinNumber
	}
}

func selectRandomPattern(patterns []string) string {
	n := rand.Intn(len(patterns))
	return patterns[n]
}

// x^yを計算
// yは1~12の整数
// yは最大12なので効率の悪いが簡単な実装をする
func pow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

func generateHoujinNumber(len int) string {
	n := rand.Intn(9*pow(10, len-1)) + pow(10, (len-1)) // len digits
	return fmt.Sprintf("%d", n)

}

var GovermentCodes = []string{
	"000011",
	"000012",
	"000013",
	"000020",
	"000030",
}

func GenerateGovernmentHoujinNumber() string {
	str := selectRandomPattern(GovermentCodes) + generateHoujinNumber(6)
	return calculateCheckDigit(str) + str
}

var OrganizationCodes = []string{
	"01",
	"02",
	"03",
	"04",
	"05",
}

func GenerateRegisteredHoujinNumber() string {
	str := selectRandomPattern(ToukijoCodes) + selectRandomPattern(OrganizationCodes) + generateHoujinNumber(6)
	return calculateCheckDigit(str) + str
}

func GenerateNonRegisteredHoujinNumber() string {
	str := "7" + generateHoujinNumber(11)
	return calculateCheckDigit(str) + str
}

func Generate() string {
	// ramdomly call one of the three functions
	n := rand.Intn(10)
	switch n {
	case 0: // 10%
		return GenerateGovernmentHoujinNumber()
	case 1: // 10%
		return GenerateNonRegisteredHoujinNumber()
	default: // 80%
		return GenerateRegisteredHoujinNumber()
	}
}

func calculateCheckDigit(houjinNumber string) string {
	var sumOne, sumTwo int
	for i := 0; i < 12; i++ {
		if i%2 == 0 {
			sumOne += int(houjinNumber[i] - '0')
		} else {
			sumTwo += int(houjinNumber[i] - '0')
		}
	}

	checkDigit := 9 - (sumOne*2+sumTwo)%9
	return fmt.Sprintf("%d", checkDigit)
}

func CalculateCheckDigit(houjinNumber string) (string, error) {
	if err := validateHoujinNumber(12, houjinNumber); err != nil {
		return "", err
	}
	return calculateCheckDigit(houjinNumber), nil
}
