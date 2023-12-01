/* I heard someone saying that "I will probably earn money from bet if I bet by followed the last result". So I write a program to bust his hope. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	pocket := 10000
	test_amount := 10000
	counter := 0
	capital := 0

	for test_amount > counter {

		single_result := gambling_round()
		capital += single_result
		counter += 1
	}

	expected_value := capital / counter

	fmt.Println(expected_value - pocket)

}

func gambling_round() int {

	pocket := 10000
	bet_amount := 100
	counter := 48

	bet_result, _ := gambling_single()

	for counter > 0 && pocket > 0 {
		if pocket-bet_amount > 0 {
			pocket -= bet_amount
		} else {
			bet_amount = pocket
			pocket = 0
		}

		table_result, table_notAlls := gambling_single()

		switch {
		case (bet_result == table_result) && (table_notAlls == true):
			pocket += bet_amount * 2
			bet_amount = 100
		default:
			bet_amount *= 2
		}

		bet_result = table_result
		counter -= 1

	}

	return pocket
}

func gambling_single() (bool, bool) {

	rand.Seed(time.Now().UTC().UnixNano())

	a1 := rand.Intn(1000)
	b1 := rand.Intn(1000)
	c1 := rand.Intn(1000)

	a := a1%6 + 1
	b := b1%6 + 1
	c := c1%6 + 1

	var size, isNotAlls bool

	if (a + b + c) >= 11 {
		size = true
	} else {
		size = false
	}

	if a == b && b == c && c == a {
		isNotAlls = false
	} else {
		isNotAlls = true
	}

	return size, isNotAlls

}
