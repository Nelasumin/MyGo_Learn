package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	maxNum := 2
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("请输入你猜的数字")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("一个错误 请努力尝试", err)
			continue
		}
		input = strings.Trim(input, "\r\n")

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("随机数错1误的指")
			continue
		}

		fmt.Println("你猜的数字", guess)
		fmt.Println("随机的数字", secretNumber)
		if guess > secretNumber {
			fmt.Println("猜的比较大")
		} else if guess < secretNumber {
			fmt.Println("猜的比较小")
		} else if guess == secretNumber {
			fmt.Println("Bye~~~~~~~~")
			break
		}
	}
}
