package utils

import (
	"encoding/json"
	"fmt"
)

// Println Println
func Println(v interface{}) {
	o, err := json.Marshal(v)
	fmt.Println("-----------------------------------------")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(o))
	}
	fmt.Println("-----------------------------------------")
}
