package main
import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var limits int = 0
var loan int = 0
type response struct{
	Type  string		`json:"type"`
	Amount int		`json:"amount"`
}


func writeHead(w http.ResponseWriter) http.ResponseWriter{
	w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("content-type", "application/json")
	return w
}


func queryLimits(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	r := response{
		Type: "limits",
		Amount : limits,
	}
	b, err := json.Marshal(&r)
	if err == nil {
		fmt.Fprint(w, string(b))
	}else{
		fmt.Fprint(w, err)
	}
}

func queryLoan(w http.ResponseWriter, req *http.Request) {
	writeHead(w)
	r := response{
		Type: "loan",
		Amount : loan,
	}
	b, err := json.Marshal(&r)
	if err == nil {
		fmt.Fprint(w, string(b))
	}else{
		fmt.Fprint(w, err)
	}
}

func clearLoan(w http.ResponseWriter, req *http.Request) {
	w = writeHead(w)
	req.ParseForm()


	amount, found := req.Form["Amount"]
	if !found {
		fmt.Fprintf(w,"wrong arguments")
		return
	}

	a, _ := strconv.Atoi(amount[0])
	loan -= a/123
	fmt.Fprintf(w, []byte{1})

}

func transferLimits(w http.ResponseWriter, req *http.Request) {
	w = writeHead(w)
	req.ParseForm()

	amount, found := req.Form["Amount"]
	if !found {
		fmt.Fprintf(w,"wrong arguments")
		return
	}
	//if len(amount) == 0{
	//	fmt.Fprintf(w,"no values")
	//}

	a, _ := strconv.Atoi(amount[0])
	limits += a
	fmt.Fprint(w, "ok")
}

func withdrawLimits(w http.ResponseWriter, req *http.Request) {
	w = writeHead(w)
	req.ParseForm()
	amount, found := req.Form["Amount"]
	if !found {
		fmt.Fprintf(w,"wrong arguments")
		return
	}
	a, _ := strconv.Atoi(amount[0])

	if a > limits {
		fmt.Fprint(w, "you don't have enough limits")
	}else{
		limits -= a
		loan += a
	}
	fmt.Fprint(w, "ok")
}


var router = mux.NewRouter()

func main() {
	router.HandleFunc("/limits", queryLimits).Methods("GET")
	router.HandleFunc("/loan", queryLoan).Methods("GET")
	router.HandleFunc("/tLimits", transferLimits).Methods("POST")
	router.HandleFunc("/wLimits", withdrawLimits).Methods("POST")
	router.HandleFunc("/clear", clearLoan).Methods("POST")
	http.Handle("/", router)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}

