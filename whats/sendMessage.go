package whats

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// SendWhatsapp send message to whatsapp
func SendWhatsapp(number string, message string) {

	env := func(key string) string {
		return os.Getenv(key)
	}

	apiurl:= env("API_URL")
    data := url.Values{}
    			 data.Set("token", env("API_TOKEN"))
			 data.Set("to", number)
			 data.Set("body", message)

   
	payload := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", apiurl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

 
	fmt.Println(string(body))
}

