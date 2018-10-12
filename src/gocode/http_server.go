package main

import (
	"fmt"
	"net/http"

	"./c10"
	"./c11"
	"./c12"
	"./c13"
	"./c14"
	"./c15"
	"./c17"
	"./c4"
	"./c5"
	"./c6"
	"./c7"
	"./c8"
	"./c9"
	"./v2c2"
	"github.com/gorilla/context"
)

func main() {
	mux := http.NewServeMux()
	//	files := http.FileServer(http.Dir("../../WebContent"))
	//	mux.Handle("/csj/", http.StripPrefix("/csj/", files))
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
	mux.Handle("/csj/c8/repeatvisitor", &c8.RepeatVisitor{})
	mux.Handle("/csj/c8/cookietest", &c8.CookieTest{})
	mux.Handle("/csj/c8/clientaccesscount", &c8.ClientAccessCount{})
	mux.Handle("/csj/c8/registrationform", &c8.RegistrationForm{})
	mux.Handle("/csj/c8/registrationservlet", &c8.RegistrationServlet{})
	mux.Handle("/csj/c9/showsession", &c9.ShowSession{})
	mux.Handle("/csj/c9/showitems", &c9.ShowItems{})
	mux.Handle("/csj/c9/kidsbookspage", c9.NewKidsBooksPage())
	mux.Handle("/csj/c9/techbookspage", c9.NewTechBooksPage())
	mux.Handle("/csj/c9/orderpage", &c9.OrderPage{})
	mux.Handle("/csj/c10/OrderConfirmation.tpl", &c10.OrderConfirmation{})
	mux.Handle("/csj/c11/Expressions.jsp", &c11.Expressions{})
	mux.Handle("/csj/c11/ThreeParams.jsp", &c11.ThreeParams{})
	mux.Handle("/csj/c11/BGColor.jsp", &c11.BGColor{})
	mux.Handle("/csj/c11/AccessCounts.jsp", &c11.AccessCounts{})
	mux.Handle("/csj/c12/ImportAttribute.jsp", &c12.ImportAttribute{})
	mux.Handle("/csj/c13/WhatsNew.jsp", &c13.WhatsNew{})
	mux.Handle("/csj/c14/StringBean.jsp", &c14.StringBean{})
	mux.Handle("/csj/c14/SaleEntry1-Form.jsp", &c14.SaleEntry1Form{})
	mux.Handle("/csj/c14/SaleEntry1.jsp", &c14.SaleEntry1{})
	mux.Handle("/csj/c14/BakedBeanDisplay-page.jsp", &c14.BakedBeanDisplayPage{})
	mux.Handle("/csj/c14/BakedBeanDisplay-request.jsp", &c14.BakedBeanDisplayRequest{})
	mux.Handle("/csj/c14/BakedBeanDisplay-session.jsp", &c14.BakedBeanDisplaySession{})
	mux.Handle("/csj/c14/BakedBeanDisplay-application.jsp", &c14.BakedBeanDisplayApplication{})
	mux.Handle("/csj/c15/showbalance", &c15.ShowBalance{})
	mux.Handle("/csj/c17/northwindservlet", &c17.Northwind{})

	// volume 2
	tServlet1 := v2c2.NewTServlet1()
	mux.Handle("/csj/v2c2/TestServlet1", tServlet1)
	//	mux.Handle("/csj/v2c2/", tServlet1)
	mux.Handle("/csj/v2c2/", v2c2.Reporter(tServlet1))
	mux.Handle("/", &v2c2.NotFound{})

	server := http.Server{
		Addr:    "127.0.0.1:6725",
		Handler: context.ClearHandler(mux),
	}

	//	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	err := server.ListenAndServe()
	fmt.Println(err)
}
