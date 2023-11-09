package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

const (
	EXPECTED_MESSAGE = "ready for queries and writes"
	SMTP_FROM        = "no-reply@regionostergotland.se"
	SMTP_SERVER      = "smtp.lio.se:25"
	SMTP_SUBJECT     = "InfluxDB ran into a problem:"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: influxmonitor <healthendpoint> <emailaddress>")
		fmt.Println("Example: influxmonitor http://localhost:8086/health my@email.com")
		os.Exit(1)
	}
	endpoint := os.Args[1]
	email := os.Args[2]
	var message string
	respons, err := http.Get(endpoint)
	if err != nil || respons == nil {
		message = fmt.Sprintf("%v is unreachable\n", endpoint)
	}
	if respons != nil {
		if respons.StatusCode != http.StatusOK {
			message = fmt.Sprintf("Status: Expected %v but got %v\n", http.StatusOK, respons.StatusCode)
		}
		var result map[string]string
		body, err := io.ReadAll(respons.Body)
		if err != nil {
			message += fmt.Sprintln("failed to read response body")
		}
		json.Unmarshal([]byte(body), &result)

		if result["message"] != EXPECTED_MESSAGE {
			message += fmt.Sprintf("Message: Expected \"%v\" but got \"%v\"\n", EXPECTED_MESSAGE, result["message"])
		}
	}
	if message != "" {
		hostname, _ := os.Hostname()
		err = smtp.SendMail(SMTP_SERVER, nil, SMTP_FROM, []string{email}, []byte(fmt.Sprintf("Subject: %s %v\n", SMTP_SUBJECT, hostname)+fmt.Sprintf("\n%s\n", message)))
		if err != nil {
			log.Fatal("Sending email failed with: ", err)
		}
	}
}
