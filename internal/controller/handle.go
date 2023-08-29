package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"user-service/internal/core/entity"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}

func SendVerifyEmail(u UserController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Println(err.Error())
			}

			var req request.RequestSendVerificationEmail
			json.Unmarshal(reqBody, &req)

			var res *response.Response = u.service.SendVerificationEmail(req)

			if res.Status == true {
				w.WriteHeader(http.StatusCreated)
			} else {
				if res.Err_code == entity.InternalErrorCode {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			}

			resJson, _ := json.Marshal(res)
			w.Header().Add("Content-Type", "application/json")
			w.Write(resJson)
		}
	}
}

func VerifyEmail(u UserController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

		}
	}
}

func SignUp(u UserController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "view/signup.html")
		} else if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Println(err.Error())
			}

			var req request.RequestSignUp
			json.Unmarshal(reqBody, &req)

			res := u.service.SignUp(req)
			resJson, _ := json.Marshal(res)
			w.Header().Add("Content-Type", "application/json")

			if res.Status {
				w.WriteHeader(http.StatusCreated)
			} else {
				if res.Err_code == entity.InternalErrorCode {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}

			}

			w.Write(resJson)
		}
	}
}

func SignIn(u UserController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "view/signin.html")
		} else if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Println(err.Error())
			}

			var req request.RequestSignin
			json.Unmarshal(reqBody, &req)

			res := u.service.Signin(req)
			resJson, _ := json.Marshal(res)
			w.Header().Add("Content-Type", "application/json")

			if res.Status {
				w.WriteHeader(http.StatusOK)
			} else {
				if res.Err_code == entity.InternalErrorCode {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}

			}
			w.Write(resJson)
		}
	}
}
