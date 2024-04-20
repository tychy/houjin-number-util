package main

import (
	"errors"
	"slices"
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

func TestGenerateGovernmentHoujinNumber(t *testing.T) {
	t.Run("generate government houjin number", func(t *testing.T) {
		houjinNumber := GenerateGovernmentHoujinNumber()
		if Validate(houjinNumber) != nil {
			t.Errorf("expected nil, got %s", houjinNumber)
		}

		code := houjinNumber[1:7]
		if !slices.Contains[[]string](GovermentCodes, code) {
			t.Errorf("expected not in %v, got %s", ToukijoCodes, code)
		}
	})

}

func TestGenerateRegisteredHoujinNumber(t *testing.T) {
	t.Run("generate registered houjin number", func(t *testing.T) {
		houjinNumber := GenerateRegisteredHoujinNumber()
		if Validate(houjinNumber) != nil {
			t.Errorf("expected nil, got %s", houjinNumber)
		}

		code := houjinNumber[1:5]
		if !slices.Contains[[]string](ToukijoCodes, code) {
			t.Errorf("expected not in %v, got %s", ToukijoCodes, code)
		}
		org := houjinNumber[5:7]
		if !slices.Contains[[]string](OrganizationCodes, org) {
			t.Errorf("expected not in %v, got %s", OrganizationCodes, org)
		}
	})
}

func TestGenerateNonRegisteredHoujinNumber(t *testing.T) {
	t.Run("generate non-registered houjin number", func(t *testing.T) {
		houjinNumber := GenerateNonRegisteredHoujinNumber()
		if Validate(houjinNumber) != nil {
			t.Errorf("expected nil, got %s", houjinNumber)
		}

		if houjinNumber[1] != '7' {
			t.Errorf("expected 7, got %c", houjinNumber[1])
		}
	})
}

func TestGenerate(t *testing.T) {
	t.Run("generate check digit", func(t *testing.T) {
		houjinNumber := Generate()
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
