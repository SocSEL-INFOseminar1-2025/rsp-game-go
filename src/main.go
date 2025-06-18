package main

import (
	"fmt"
	"math/rand"
	"time" // timeパッケージを追加
)

func main() {
	// 乱数のシードを設定
	rand.Seed(time.Now().UnixNano())

	var my_hand int
	var enemyHand int
	var rspResult int

	for { // 勝敗が決まるまで繰り返す
		fmt.Println("出す手を決めてください")
		fmt.Println("0:グー 1:チョキ 2:パー")
		fmt.Scan(&my_hand)

		// 入力値のバリデーション
		if my_hand < 0 || my_hand > 2 {
			fmt.Println("無効な入力です。0, 1, 2のいずれかを入力してください。")
			continue // 再度入力を促す
		}

		enemyHand = getRandom()

		fmt.Println("じゃんけんぽん！")
		fmt.Printf("あなたは %s、相手は %s\n", handToString(my_hand), handToString(enemyHand)) // どちらが出したか表示
		rspResult = rspBattle(my_hand, enemyHand)

		printResult(rspResult)

		if rspResult != 0 { // 引き分けでなければループを終了
			break
		} else {
			fmt.Println("あいこでしょ！") // 引き分けの場合はもう一度
		}
	}

	fmt.Println("また遊んでね！")
}

// 0~2の値を乱数から生成する
func getRandom() int {
	return rand.Intn(3) // 0, 1, 2 のいずれかを返す
}

// じゃんけんの結果を判定
func rspBattle(myHand, enemyHand int) int {
	// 引き分け:0 勝利:1 敗北:2
	switch {
	case myHand == enemyHand:
		return 0 // 引き分け
	case (myHand == 0 && enemyHand == 1) || // グー vs チョキ
		(myHand == 1 && enemyHand == 2) || // チョキ vs パー
		(myHand == 2 && enemyHand == 0): // パー vs グー
		return 1 // 勝利
	default:
		return 2 // 敗北
	}
}

// 結果を表示
func printResult(result int) {
	if result == 0 {
		fmt.Println("Draw")
	} else if result == 1 {
		fmt.Println("You WIN!")
	} else if result == 2 {
		fmt.Println("You LOSE")
	}
}

// 手の値を文字列に変換するヘルパー関数
func handToString(hand int) string {
	switch hand {
	case 0:
		return "グー"
	case 1:
		return "チョキ"
	case 2:
		return "パー"
	default:
		return "不明"
	}
}