package Mail

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Mail(email string) {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	MailURL := os.Getenv("MailURL")
	url := MailURL
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

	payload := strings.NewReader(code)
	req, _ := http.NewRequest("POST", url, payload)

	XRapidAPIKey := os.Getenv("XRapidAPIKey")
	XRapidAPIHost := os.Getenv("XRapidAPIHost")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", XRapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", XRapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
