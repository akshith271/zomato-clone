package Mail

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// func main() {
func Mail(email string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	code := fmt.Sprintf(`{
	    "personalizations": [
	        {
	            "to": [
	                {
	                    "email": `+`"%s"`+`
	                }
	            ],
	            "subject": "Account confirmation"
	        }
	    ],
	    "from": {
	        "email": "akshithbharadwaj@beautifulcode.in"
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your Account creation is successful !!!\nThank You for joining Us"`+`
			}
	    ]
	}`, email)
	// fmt.Println(code)

	payload := strings.NewReader(code)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "04e9f7a9f9msh6a6e005a692ad20p163e6fjsna668112a5f00")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
