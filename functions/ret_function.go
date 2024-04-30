package functions

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, int) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		fmt.Println(v)
		initials = append(initials, v[:2])
	}

	var ret string

	for _, v := range initials {
		ret = ret + v + " "
	}

	return ret, len(names)

}

func Return_muntiple_function() {
	fn1, len1 := getInitials("saber yin gusion")
	fmt.Println(fn1, "total initials", len1)

	fn2, len2 := getInitials("cloude wanwan")
	fmt.Println(fn2, "total initials", len2)

	fn3, len3 := getInitials("triezla")
	fmt.Println(fn3, "total initials", len3)
}
