package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"

)

func main(){
	fmt.Println("Welcome babes")
	//PerformGetRequest()
	PerformForm()
}


func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
    defer res.Body.Close()


	fmt.Println("status code:",res.StatusCode)
	fmt.Println("Content length", res.ContentLength)



	var response strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteC, _ := response.Write(content)

    fmt.Println("Byte count is",byteC)
	fmt.Println(response.String())
    
	//content, err := ioutil.ReadAll(res.Body)
    //fmt.Println(string(content))

}




func PerformPost(){
	const myurl= "http://localhost:8000/post"


	//fake json payload
	requestBody := strings.NewReader(`
	     {
			"name": "Tanish",
			"Toni" : "Kroos"
		 }
	`)
	res, err := http.Post(myurl,"application/json",requestBody)
	if err!= nil{
		panic(err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err!= nil{
		panic(err)
	}

	fmt.Println(string(content))

}

func PerformForm(){
	const myurl = "http://localhost:8000/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname","Tanish")
	data.Add("lastname","Saini")

	res, err := http.PostForm(myurl,data)
	if err!= nil{
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))


}











