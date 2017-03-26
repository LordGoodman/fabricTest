package main
import(
//	"encodeing/json"
//	"io"
	"fmt"
	"net/http"
)

var chaincodeName string


func login(w http.ResponseWriter, req *http.Request) {
	//fmt.Println()
	//req.ParseForm()
	//param_user, found1 := req.Form["user"]
	//param_pass, found2 := req.Form["pass"]

	//if !(found1 && found2) {
	//	fmt.Fprintf(w, "Don't do this")
	//
	req.ParseForm()
	userName, found1 := req.Form["userName"]
	pass, found2     := req.Form["pass"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Please offer the user and pass")
	}

	if userName == "Tom" && pass == "123" {
		fmt.Fprintf(w, "ok")
	}else userName == "alice" && pass == "123" {
		fmt.Fprintf(w, "ok")
	}else userName == "bob"   &&  pass == "123" {
		fmt.Fprintf(w, "ok")
	}else {
		fmt.Fprintf(w, "user or pass error!")
	}

}

func topupHandle(w http.ResponseWriter, req *http.Request) {
	req.ParesForm()
	amount, found1 := req.Form["amount"]
	user, found2   := req.Form["user"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
	}

	err := topup(chaincodeName, amount, user)
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprinrf(w, "ok")
}


func investHandle(w http.ResponseWriter, req *http.Request) {
	req.ParesForm()
	amount, found1 := req.Form["amount"]
	user, found2   := req.Form["user"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
	}

	err := invest(chaincodeName, amount, user)
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprinrf(w, "ok")
}


func cashoutHandle(w http.ResponseWriter, req *http.Request) {
	req.ParesForm()
	amount, found1 := req.Form["amount"]
	user, found2   := req.Form["user"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
	}

	err := cashout(chaincodeName, amount, user)
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprinrf(w, "ok")
}

func transferHandle(w http.ResponseWriter, req *http.Request) {
	req.ParesForm()
	amount, found1 := req.Form["amount"]
	from, found2   := req.Form["from"]
	to, found3     := req.Form]["to"]


	if !(found1 && found2) {
		fmt.Fprintf(w, "Wrong arguments")
	}

	err := topup(chaincodeName, amount, from, to)
	if err != nil {
		fmt.Fprintf(w, "faild")
	}

	fmt.Fprinrf(w, "ok")
}

func queryHandle(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user, found := req.Form["user"]
	if !found {
		fmt.Fprintf(w, "wrong arguments")
	}

	err := 
	}

func main() {
	
	http.HandleFunc("/login", login)
	http.HandleFunc("/topup", topupHandle)
	http.HandleFunc("/invest", investHandle)
	http.HandleFunc("/cashout", cashoutHandle)
	http.HandleFunc("/transfer", transferHandle)
	http.ListenAndServe(":8080", nil)
}




