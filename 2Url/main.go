 package main
 import (
	"fmt"
	"net/url"
 )

 const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
 func main() {
	fmt.Println("Welcome to handling urls")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	if err!= nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)

	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query param are %T \n",qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

    for _, val := range qparams{
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",

	}
    anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl)




 }