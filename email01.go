/* RZFeeser | Alta3 Research
   Sending an Email (SMTP) message   */
   
package main

import (
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "ilmka3spam@gmail.com"            // update this to reflect your value
    password := "optum021"     // update this to reflect your value
    to := []string{"wschlichtig@gmail.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    message := []byte("Subject: Testing email from go\r\n\n" + "My super secret message.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
}

