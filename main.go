package main

import (
	"fmt"

	"github.com/tychy/houjin-number-util/houjin"
)

func main() {
	// ランダムな法人番号を生成
	randomHoujinNumber := houjin.Generate()
	fmt.Printf("ランダム法人番号: %s\n", randomHoujinNumber)

	// 法人番号が有効かどうかを検証
	err := houjin.ValidateHoujinNumber(randomHoujinNumber)
	if err != nil {
		fmt.Printf("検証エラー: %v\n", err)
	} else {
		fmt.Println("法人番号は有効です")
	}

	// 政府機関の法人番号を生成
	govHoujinNumber := houjin.GenerateGovernmentHoujinNumber()
	fmt.Printf("政府機関法人番号: %s\n", govHoujinNumber)

	// 設立登記法人の法人番号を生成
	regHoujinNumber := houjin.GenerateRegisteredHoujinNumber()
	fmt.Printf("設立登記法人番号: %s\n", regHoujinNumber)

	// 設立登記のない法人の法人番号を生成
	nonRegHoujinNumber := houjin.GenerateNonRegisteredHoujinNumber()
	fmt.Printf("設立登記のない法人番号: %s\n", nonRegHoujinNumber)

	// 無効な法人番号の例
	invalidHoujinNumber := "1234567890123"
	err = houjin.ValidateHoujinNumber(invalidHoujinNumber)
	if err != nil {
		fmt.Printf("無効な法人番号の検証結果: %v\n", err)
	}
}
