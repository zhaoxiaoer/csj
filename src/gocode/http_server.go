package main

import (
	"fmt"
	"net/http"

	"./c4"
	"./c5"
	"./c6"
	"./c7"
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
	mux.Handle("/csj/c5/browserinsult", &c5.BrowserInsult{})
	mux.Handle("/csj/c6/wrongdestination", &c6.WrongDestination{})
	mux.Handle("/csj/c6/searchengineform", &c6.SearchEngineForm{})
	mux.Handle("/csj/c6/searchengines", &c6.SearchEngines{})
	mux.Handle("/csj/c7/applesandoranges", &c7.ApplesAndOranges{})

	server := http.Server{
		Addr:    "127.0.0.1:6725",
		Handler: mux,
	}

	//	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	err := server.ListenAndServe()
	fmt.Println(err)
}
