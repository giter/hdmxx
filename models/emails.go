package models;

import (
		"strings"
    "fmt"
	  "net/smtp"
		"net/mail"
    "crypto/tls"
		"encoding/base64"
)

const HOST        = "smtp.exmail.qq.com"
const SERVER_NAME = "smtp.exmail.qq.com:465"

const LOGIN_NAME  = "XXX"
const PASSWORD    = "YYY"
const FROM        = "ZZZ"

var TLS_CONFIG  = &tls.Config {
		InsecureSkipVerify: true,
		ServerName: HOST,
}

func encodeRFC2047(String string) string {

	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func NewEmail(To string, Subject string , Body string) (err error) {

		if To == "" {
			return
		}

		auth := smtp.PlainAuth("",LOGIN_NAME, PASSWORD, HOST)

	  headers := make(map[string]string)

    headers["From"] = FROM
    headers["To"] = To

		headers["Content-Type"] = "text/plain; charset=\"utf-8\""
		headers["Content-Transfer-Encoding"] = "base64"

    headers["Subject"] = encodeRFC2047(Subject)

    message := ""

    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }

    message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(Body))

		conn, err := tls.Dial("tcp", SERVER_NAME, TLS_CONFIG)

    if err != nil {
        return
    }

    c, err := smtp.NewClient(conn, HOST)

    if err != nil {
        return
    }

		defer c.Quit()

    // Auth
    if err = c.Auth(auth); err != nil {
        return
    }

    // To && From
    if err = c.Mail(FROM); err != nil {
        return
    }

    if err = c.Rcpt(To); err != nil {
        return
    }

    // Data
    w, err := c.Data()

    if err != nil {
        return
    }

    _, err = w.Write([]byte(message))

    if err != nil {
        return
    }

    err = w.Close()

    return
}
