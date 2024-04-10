package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrInvalidCharacter          = errors.New("invalid character")
	ErrInvalidHoujinNumberLength = errors.New("invalid houjin number length")
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

func Validate(houjinNumber string) error {
	if err := validateHoujinNumber(13, houjinNumber); err != nil {
		return err
	}

	checkDigit := calculateCheckDigit(houjinNumber[1:])
	if checkDigit != string(houjinNumber[0]) {
		return ErrInvalidCheckDigit
	}
	return nil
}

func Generate() string {
	n := rand.Intn(900000000000) + 100000000000 // 12 digits
	str := fmt.Sprintf("%d", n)
	return calculateCheckDigit(str) + str
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
