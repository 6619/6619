package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		猜数字游戏
		生成随机整数，提示用户输入数字，比较，当用户输入较大，提示太大类了；当用户输入较小，提示太小了，当用户输入正确，提示经过N次对了，太聪明了。
		用户最多猜5次，如果5次内没有猜对，提示太笨了，游戏结束。

		扩展：当成功或失败，提示用户是否继续，输入：yes、y、Y则继续，从新生成随机数，让用户猜
	*/

	var (
		randNum  int
		inputNum int
		count    int
		inputStr string
	)
	for {
		fmt.Println("确认进行游戏（y、Y、yes）")
		fmt.Scan(&inputStr)
		if inputStr == "yes" || inputStr == "y" || inputStr == "Y" {
			fmt.Printf("现在正式开始游戏，生成一个数字")

			// 生成随机数
			rand.Seed(time.Now().Unix())
			randNum = rand.Intn(100)
			fmt.Println(randNum)
			for i := 1; i <= 7; i++ {
				fmt.Println("请输入您猜测的数字：")
				fmt.Scan(&inputNum)
				fmt.Printf("您猜测的数字是: %d\n", inputNum)
				if inputNum == randNum {
					fmt.Printf("恭喜您只用了 %d 次就猜对了，您真是太聪明了！！！\n", i)
					break
				} else if inputNum > randNum {
					fmt.Println("太大了！！！")
				} else if inputNum < randNum {
					fmt.Println("太小了！！！")
				}
				count++
			}
			if inputNum != randNum {
				fmt.Printf("您的运气不太好，猜了 %d 次都没有猜中，不如进行下一句游戏。\n", count)
			}
		} else {
			fmt.Println("游戏结束")
			break
		}
	}

}
