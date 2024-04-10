package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("invalid character", func(t *testing.T) {
		houjinNumber := "123456789012a"
		err := Validate(houjinNumber)
		if !errors.Is(err, ErrInvalidCharacter) {
			t.Errorf("expected %v, got %v", ErrInvalidCharacter, err)
		}
	})
	t.Run("invalid houjin number length", func(t *testing.T) {
		houjinNumber := "123456789012"
		err := Validate(houjinNumber)
		if !errors.Is(err, ErrInvalidHoujinNumberLength) {
			t.Errorf("expected %v, got %v", ErrInvalidHoujinNumberLength, err)
		}
	})
	t.Run("invalid check digit", func(t *testing.T) {
		houjinNumber := "1234567890123"
		err := Validate(houjinNumber)
		if !errors.Is(err, ErrInvalidCheckDigit) {
			t.Errorf("expected %v, got %v", ErrInvalidCheckDigit, err)
		}
	})

	t.Run("valid houjin number", func(t *testing.T) {
		houjinNumber := "8700110005901"
		err := Validate(houjinNumber)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}

func TestGenerate(t *testing.T) {
	t.Run("generate check digit", func(t *testing.T) {
		houjinNumber := Generate()
		fmt.Println(houjinNumber)
		if Validate(houjinNumber) != nil {
			t.Errorf("expected nil, got %s", houjinNumber)
		}
	})
}

func TestCalculateCheckDigit(t *testing.T) {
	t.Run("invalid character", func(t *testing.T) {
		houjinNumber := "12345678901a"
		_, err := CalculateCheckDigit(houjinNumber)
		if !errors.Is(err, ErrInvalidCharacter) {
			t.Errorf("expected %v, got %v", ErrInvalidCharacter, err)
		}
	})
	t.Run("invalid houjin number length", func(t *testing.T) {
		houjinNumber := "12345678901"
		_, err := CalculateCheckDigit(houjinNumber)
		if !errors.Is(err, ErrInvalidHoujinNumberLength) {
			t.Errorf("expected %v, got %v", ErrInvalidHoujinNumberLength, err)
		}
	})

	t.Run("calculate check digit", func(t *testing.T) {
		houjinNumber := "700110005901"
		checkDigit, err := CalculateCheckDigit(houjinNumber)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if checkDigit != "8" {
			t.Errorf("expected 8, got %s", checkDigit)
		}
	})
}
