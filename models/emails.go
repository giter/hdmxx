package models;

import (
    "fmt"
	  "net/smtp"
    "crypto/tls"
)

const HOST        = "smtp.exmail.qq.com"
const SERVER_NAME = "smtp.exmail.qq.com:465"

const LOGIN_NAME  = "xxx"
const PASSWORD    = "yyy"
const FROM        = "zzz"

var TLS_CONFIG  = &tls.Config {
		InsecureSkipVerify: true,
		ServerName: HOST,
}

func NewEmail(To string, Subject string , Body string) (err error) {

		auth := smtp.PlainAuth("",LOGIN_NAME, PASSWORD, HOST)

	  headers := make(map[string]string)

    headers["From"] = FROM
    headers["To"] = To

    headers["Subject"] = Subject

    message := ""

    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }

    message += "\r\n" + Body

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