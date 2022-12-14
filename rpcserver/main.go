package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

type SmsArgs struct {
	Request, Content string
}

type EmailArgs struct {
	Receiver, Body string
}

type SmsService struct{}
type EmailService struct{}

type Response struct {
	Result string
}

func (s *SmsService) SendSMS(r *http.Request, args *SmsArgs, result *Response) error {
	*result = Response{Result: `All is ok`}
	log.Println(`All is ok`)
	return nil
}

func (s *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *Response) error {
	*result = Response{Result: `All is ok`}
	return nil
}

func main() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), `application/json`)

	sms := new(SmsService)
	email := new(EmailService)

	_ = rpcServer.RegisterService(sms, `sms`)
	_ = rpcServer.RegisterService(email, `email`)

	router := mux.NewRouter()
	router.Handle(`/delivery`, rpcServer)
	log.Fatal(http.ListenAndServe(`:8081`, router))
}
