//Demo for Marshalling and Unmarshalling of JSON
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "Zama",
		Last:  "Md",
		Age:   33,
	}
	p2 := person{
		First: "John",
		Last:  "St",
		Age:   40,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	fmt.Println("--- After marshalling")
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	jsonString := string(bs)
	fmt.Println("String format: ", jsonString)

	fmt.Println("--- After unmarshalling")

	jsonString2 := `[{"First":"Zama","Last":"Md","Age":33},{"First":"John","Last":"St","Age":40}]`
	bs2 := []byte(jsonString2)
	var people2 []person2

	err = json.Unmarshal(bs2, &people2)

	if err != nil {
		fmt.Println("Error while unmarshalling: ", err)
	}

	for i, v := range people2 {
		fmt.Println(i, "\t - ", v)
	}
}
