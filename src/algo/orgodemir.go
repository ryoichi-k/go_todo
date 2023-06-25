package algo

import (
	"fmt"
	"math/rand"
)

type Orgodemir struct {
	HP      int
	ATK     int
	SPD     int
	Pattern string //flag which pattern A or B
	//movingNum int    //行動回数
}

// オルゴデミーラ攻撃パターン
// 第一形態
// HP：3000 MP：∞ 攻撃力：182 守備力：127 素早さ：107
// ●1～3回行動（※）
// Aパート：メラゾーマ→打撃→【マグマ】→いてつくはどう→かまいたち→移行
// Bパート：打撃、念じボール、いなずま、イオナズン、あやしいひとみ、移行
// ※Aパート、Bパート共に1～2回行動。移行を挟むと最大3回行動。移行は1行動に入る

func MovementPattern() {
	fmt.Println("in")
	var orgodemir Orgodemir
	orgodemir.HP = 3000
	orgodemir.ATK = 182
	orgodemir.SPD = 107

	patternSlice := []string{"メラゾーマ", "打撃", "マグマ", "いてつくはどう", "かまいたち", "移行",
		"打撃", "念じボール", "いなずま", "イオナズン", "あやしいひとみ", "移行"}

	//あやしいひとみ　→　移行　→ メラゾーマ[11:] [0]
	//イオナズン　→ あやしいひとみ[10:12]

	//各ターンの最初に行動回数をランダムで決定する
	movingNum := rand.Intn(2) //0 or 1
	//現在の配列のインデックス番号
	var currentNum int
	//最初の１ターン
	fmt.Println(patternSlice[0 : movingNum+1])
	currentNum = movingNum + 1

	for {
		//行動回数を決定する
		movingNum := rand.Intn(2) //0 or 1
		lastInd := currentNum + movingNum + 1
		//簡易スライスを溢れた場合12 or 13, 11の時は移行なので[o]にいく
		if lastInd > 12 || currentNum >= 11 {
			//13 -(12-1)
			//overflowingNumberが1
			overflowingNumber := lastInd - (len(patternSlice) - 1)
			lastInd = 11
			if currentNum >= 11 {
				currentNum = 11
			}
			fmt.Println("----------")
			fmt.Println(currentNum)
			fmt.Println(lastInd)

			fmt.Println("patternSlice[currentNum:lastInd]")
			fmt.Println(patternSlice[lastInd])
			currentNum = 0
			lastInd = overflowingNumber
			break
		}
		fmt.Println("----表示----")
		fmt.Println("currentNum")
		fmt.Println(currentNum)
		fmt.Println(lastInd)
		fmt.Println(patternSlice[currentNum:lastInd]) //ここ
		currentNum = currentNum + movingNum + 1
	}
	//fmt.Println("finish")
}
