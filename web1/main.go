package main
import (
   "fmt"
   "net/http"
   "io/ioutil"
)
const url= "https://lco.dev"

func main() {
   fmt.Println("LCO WEB REQUEST")

   response, err := http.Get(url)

   if err != nil {
      panic(err)
   }

   fmt.Printf("Response of type: %T\n", response)

   defer response.Body.Close() 
   
   data, err := ioutil.ReadAll(response.Body)

   if err != nil {
      panic(err)
   }
   content := string(data)
   fmt.Println(content)


}