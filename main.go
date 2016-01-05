// JSON and go – http://blog.golang.org/json-and-go
package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	String string
	//	Int    int
}

func main() {
	var data Data
	JSONstring := "{\"String\":\"hello world!\", \"Int\": 42}" // If there is more info in JSON than receiving struct can accept, info is discard

	err := json.Unmarshal([]byte(JSONstring), &data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
