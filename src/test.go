package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	request, _ := http.NewRequest("GET", "http://zodream.localhost/api.php/goods", nil)
	response, _ := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		body := string(str)
		fmt.Println(body)
	}
}
