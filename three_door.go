package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	threeDoorCommon()
}

func threeDoorCommon() {
	rightCnt := 0
	changeRightCnt := 0
	for i := 0; i < 10000000; i++ {
		randCarLoc, _ := rand.Int(rand.Reader, big.NewInt(3))
		randGuestChoice, _ := rand.Int(rand.Reader, big.NewInt(3))

		// fmt.Println(randCarLoc)
		// fmt.Println(randGuestChoice)
		j := 0
		for ; j < 3; j++ {
			if randCarLoc.Int64() != int64(j) && randGuestChoice.Int64() != int64(j) {
				break
			}
		}

		// fmt.Println(j)
		if randCarLoc.Cmp(randGuestChoice) == 0 {
			rightCnt++
		} else {
			changeRightCnt++

		}

		fmt.Printf("total is: %d, right is %d, changeRight is: %d\n", i+1, rightCnt, changeRightCnt)

	}

}
