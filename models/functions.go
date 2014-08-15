package models;

import (

	"io/ioutil"
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

				if body,err := ioutil.ReadAll(resp.Body) ; err != nil || len(body) == 0 {
					Status = 0
				}
			}

			return
}
