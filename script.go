package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"bufio"
	"strings"
	"net/smtp"
)

func watchLogFile(filename string, keyword string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), keyword) {
			sendAlert(fmt.Sprintf("Found keyword '%s' in log file", keyword))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sendAlert(message string) {
	// Sending email alert (SMTP server configuration)
	auth := smtp.PlainAuth("", "your-email@example.com", "your-email-password", "smtp.example.com")
	err := smtp.SendMail("smtp.example.com:587", auth, "your-email@example.com", []string{"alert@example.com"}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Alert sent!")
}

func main() {
	// Watch the log file for a keyword
	watchLogFile("/var/log/syslog", "ERROR")
}
