package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Helllo JSON")
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React", 299, "LearnCodeOnline.in", "abc123", []string{"webdev", "js"}},
		{"MERN", 299, "LearnCodeOnline.in", "abc153", []string{"webdev", "js"}},
		{"Terrform", 299, "LearnCodeOnline.in", "abc12546", []string{"DevOps", "js"}},
		{"DSA", 299, "LearnCodeOnline.in", "abc13", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
				"webdev",
				"js"
		]
	}
	`)

	var lcoCourse course
	check := json.Valid(jsonDataFromWeb)
	if check {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("Json not valid")
	}
	// want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
