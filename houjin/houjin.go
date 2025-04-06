package houjin

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

// エラー定義
var (
	// ErrInvalidCharacter は法人番号に無効な文字が含まれている場合に返されます
	ErrInvalidCharacter = errors.New("invalid character")
	// ErrInvalidHoujinNumberLength は法人番号の長さが無効な場合に返されます
	ErrInvalidHoujinNumberLength = errors.New("invalid houjin number length")
	// ErrInvalidHoujinNumber は法人番号が無効な場合に返されます
	ErrInvalidHoujinNumber = errors.New("invalid houjin number")
	// ErrInvalidCheckDigit はチェックデジットが無効な場合に返されます
	ErrInvalidCheckDigit = errors.New("invalid check digit")
)

// validateHoujinNumber は法人番号の文字列長と文字の妥当性を検証します
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

// ValidateCheckSum は法人番号のチェックサムが正しいかを検証します
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

// validateToukijoCodeOrgCode は登記所コードと組織区分コードの妥当性を検証します
func validateToukijoCodeOrgCode(code, org string) error {
	if !slices.Contains(ToukijoCodes, code) {
		return ErrInvalidHoujinNumber
	}
	if !slices.Contains(OrganizationCodes, org) {
		return ErrInvalidHoujinNumber
	}
	return nil
}

// ValidateHoujinNumber は法人番号の妥当性を検証します
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

// selectRandomPattern はパターンのスライスからランダムに1つを選択します
func selectRandomPattern(patterns []string) string {
	n := rand.Intn(len(patterns))
	return patterns[n]
}

// pow はx^yを計算します（yは1~12の整数）
func pow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

// generateHoujinNumber は指定された桁数のランダムな数字を生成します
func generateHoujinNumber(len int) string {
	n := rand.Intn(9*pow(10, len-1)) + pow(10, (len-1)) // len digits
	return fmt.Sprintf("%d", n)
}

// GovermentCodes は政府機関のコードリストです
var GovermentCodes = []string{
	"000011",
	"000012",
	"000013",
	"000020",
	"000030",
}

// GenerateGovernmentHoujinNumber は政府機関の法人番号を生成します
func GenerateGovernmentHoujinNumber() string {
	str := selectRandomPattern(GovermentCodes) + generateHoujinNumber(6)
	return calculateCheckDigit(str) + str
}

// OrganizationCodes は組織区分コードリストです
var OrganizationCodes = []string{
	"01",
	"02",
	"03",
	"04",
	"05",
}

// GenerateRegisteredHoujinNumber は設立登記法人の法人番号を生成します
func GenerateRegisteredHoujinNumber() string {
	str := selectRandomPattern(ToukijoCodes) + selectRandomPattern(OrganizationCodes) + generateHoujinNumber(6)
	return calculateCheckDigit(str) + str
}

// GenerateNonRegisteredHoujinNumber は設立登記のない法人の法人番号を生成します
func GenerateNonRegisteredHoujinNumber() string {
	str := "7" + generateHoujinNumber(11)
	return calculateCheckDigit(str) + str
}

// Generate はランダムな法人番号を生成します
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

// calculateCheckDigit は12桁の基礎番号からチェックデジットを計算します
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

// CalculateCheckDigit は12桁の基礎番号からチェックデジットを計算します
func CalculateCheckDigit(houjinNumber string) (string, error) {
	if err := validateHoujinNumber(12, houjinNumber); err != nil {
		return "", err
	}
	return calculateCheckDigit(houjinNumber), nil
}
