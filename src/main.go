package main

import (
	"fmt"
	"math/rand"
)


func main() {
	var my_hand int
	var enemyHand int
	var rspResult int

	fmt.Println("出す手を決めてください")
	fmt.Println("0:グー 1:チョキ 2:パー")
	fmt.Scan(&my_hand)

	enemyHand = getRandom()

	fmt.Println("じゃんけんぽん！")
	fmt.Println(enemyHand)
	rspResult = rspBattle(my_hand, enemyHand)

	printResult(rspResult)
	fmt.Println("また遊んでね！")
}

func getRandom() int {
	// 0~2の値を乱数から生成する
	randomNum := rand.Intn(100)
	return randomNum % 2
}

func rspBattle(myHand, enemyHand int) int {
	// 引き分け:0 勝利:1 敗北:2
	switch myHand {
	case 0:
		switch enemyHand {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		default:
			return 999
		}
	case 1:
		switch enemyHand {
		case 0:
			return 2
		case 1:
			return 0
		case 2:
			return 1
		default:
			return 999
		}
	case 2:
		switch enemyHand {
		case 0:
			return 1
		case 1:
			return 2
		case 2:
			return 0
		default:
			return 999
		}
	}

	return 999
}

func printResult(result int) {
	if result == 0 {
		fmt.Println("Draw")
	} else if result == 1 {
		fmt.Println("You WIN!")
	} else if result == 2 {
		fmt.Println("You LOSE")
	}
}