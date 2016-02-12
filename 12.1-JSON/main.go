// example of using encoding/json package
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Person structure with json tags associated
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// buffer used to indent json below
	var out bytes.Buffer

	// create a person structure with data
	person := &Person{
		Name:    "John",
		Age:     23,
		Hobbies: []string{"e-sports", "football"},
	}
	// convert the structure into json format
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	// indent the json produced
	json.Indent(&out, jsonPerson, "", "\t")
	out.WriteTo(os.Stdout)
	out.Reset() // flush the buffer

	// read the row bytes from json file
	rowByte, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	// populate the structure from json file
	json.Unmarshal(rowByte, &person)
	fmt.Println(person)
}
