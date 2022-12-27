package Mail

import (
	"fmt"
	"io"
	"mock-grpc/environment"
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
	req.Header.Add("X-RapidAPI-Key", environment.XRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", environment.XRapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
