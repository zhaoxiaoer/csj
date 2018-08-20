package main

import (
	"fmt"
	"net/http"

	"./c4"
	"./c5"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("../../WebContent"))
	mux.Handle("/csj/", http.StripPrefix("/csj/", files))
	mux.Handle("/csj/c4/submitresume", &c4.SubmitResume{})
	mux.Handle("/csj/c4/badcodeservlet", &c4.BadCode{})
	mux.Handle("/csj/c4/submitinsuranceinfo", &c4.SubmitInsuranceInfo{})
	mux.Handle("/csj/c4/bidservlet", &c4.BidServ{})
	mux.Handle("/csj/c5/showrequestheaders", &c5.ShowRequestHeaders{})
	mux.Handle("/csj/c5/longservlet", &c5.LongServ{})

	server := http.Server{
		Addr:    "127.0.0.1:6725",
		Handler: mux,
	}

	//	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	err := server.ListenAndServe()
	fmt.Println(err)
}