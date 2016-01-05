// JSON and go – http://blog.golang.org/json-and-go
package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	String      string
	InnerStruct *InnerStruct
}

type InnerStruct struct {
	Int    int
	String string
	Map    map[string]string
}

func main() {
	inner := &InnerStruct{42, "hello world", map[string]string{"hello": "world"}}
	data := &Data{"JSON IS GREAT", inner}
	fmt.Println(data)

	encoded, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(encoded))
	}
}
