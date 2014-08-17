package models;

import (

	"io/ioutil"
	"net"
	"net/http"
)

func CheckHttp(Method string, Url string)(Status int,err error) {

			var resp *http.Response

			Status = 1

			if Method == "GET" {
				resp,err = http.Get(Url)
			}else if Method == "POST" {
				resp,err = http.Post(Url,"application/form-data-url", nil)
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

func CheckNet(Protocol string, Address string, Port int, Input string, Result string)(Status int,err error){

	var conn net.Conn
	if conn, err = net.Dial(Protocol, Address + ":" + string(Port)); err != nil {

		Status = 0
		return
	}

	defer conn.Close()

	if len(Input) > 0 {
		panic("Not Implemented!")
	}else{
		Status = 1
	}


	return
}
