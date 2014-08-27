package models;

import (

	"time"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"encoding/base64"
	"bytes"
	"log"
)

func CheckHttp(s Site)(Status int,err error) {

			Method := s.Method
			Url := s.Url
			CTimeout := s.CTimeout
			RTimeout := s.RTimeout

			var transport =	&http.Transport{
					Dial: (&net.Dialer{
						Timeout: time.Duration(CTimeout) * time.Millisecond,
					}).Dial,
				}

			var client = &http.Client{
				Transport: transport,
				Timeout: time.Duration(RTimeout) * time.Millisecond,
			}

			var resp *http.Response

			Status = 1

			if Method == "GET" {
				resp,err = client.Get(Url)
			}else if Method == "POST" {
				resp,err = client.Post(Url,"application/form-data-url", nil)
			}

			if err != nil {

				Status = 0

			}else{

				defer resp.Body.Close()

				var body []byte

				if body,err = ioutil.ReadAll(resp.Body) ; err != nil || len(body) == 0 {
					Status = 0
				}
			}

			return
}

func CheckNet(Protocol string, s Site)(Status int, err error) {

	Address := s.Address
	Port := s.Port

	CTimeout := s.CTimeout
	RTimeout := s.RTimeout

	dialer := &net.Dialer{
		Timeout: time.Duration(CTimeout) * time.Millisecond,
	}


	Input := s.Input
	Result := s.Result

	var tinput []byte
	var toutput []byte

	if len(Input) > 0 {

		if tinput,err = base64.StdEncoding.DecodeString(Input) ; err != nil {
			Status = 0
			return
		}
	}

	if len(Result) > 0 {
		if toutput,err = base64.StdEncoding.DecodeString(Result) ; err != nil {
			Status = 0
			return
		}
	}

	var conn net.Conn
	if conn, err = dialer.Dial(Protocol, Address + ":" + strconv.Itoa(Port)); err != nil {

		Status = 0
		return
	}

	log.Println("Connected")

	var n int

	defer conn.Close()

	conn.SetDeadline(time.Now().Add(time.Duration(RTimeout) * time.Millisecond))

	if tinput != nil {

		log.Println("Inputing...")

		if n,err = conn.Write(tinput);err != nil || n<len(tinput) {
			Status = 0
			return
		}

		log.Println(Input)
	}


	if toutput != nil {

		log.Println("Receiving...")
		soutput := make([]byte,len(toutput) + 10)
		if n,err = conn.Read(soutput) ; err != nil || n != len(toutput) || !bytes.Equal(soutput[0:n],toutput) {
			log.Println(soutput)
			s.Status = 0
			return
		}
		log.Println(Result)
	}

	Status = 1


	return
}
