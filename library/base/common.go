package base

import(
	"fmt"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}