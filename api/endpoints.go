package api

import (
	"fmt"
	"time"
	"github.com/gorilla/mux"
	"sort"
	"net/http"
	"encoding/json"
	"github.com/pborman/uuid"
	"AgileEngineCodingChallengeApiBank/templates"
)



func handleError(err error) {

	if err != nil {
		panic(err)
	}
}

func HandleTransactionPost(w http.ResponseWriter, r *http.Request) {

	var transaction Transaction

	handleError((json.NewDecoder(r.Body).Decode(&transaction)))

	handleError(transaction.Type.validate())
	handleError(transaction.Amount.validate())

	tr := TransactionRecord{
		ID: uuid.New(),
		Date: time.Now(),
		Transaction: Transaction{
			Type: transaction.Type,
			Amount: transaction.Amount,
		},
	}

	func () {
		lockObj.Lock()
		defer lockObj.Unlock()

		TransactionsDB[tr.ID] = tr

		switch tr.Type {
		case "debit":
			userAccount.CurrentBalance -= tr.Amount
		case "credit":
			userAccount.CurrentBalance += tr.Amount
		}
	}()

	ret := TransactionResponse{ID: tr.ID}
	json.NewEncoder(w).Encode(&ret)
}

func HandleTransactionGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tr TransactionRecord
	var exists bool


	func () {
		lockObj.RLock()
		defer lockObj.RUnlock()
		tr, exists = TransactionsDB[id]
	}()

	if (!exists){

		ret := ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		}

		handleError(json.NewEncoder(w).Encode(&ret))
		return
	}
	handleError(json.NewEncoder(w).Encode(&tr))
}

func HandleTransactionList(w http.ResponseWriter, r *http.Request) {
	handleError(json.NewEncoder(w).Encode(getTransactionList()))
}

func getTransactionList () *[]TransactionRecord{
	var records []TransactionRecord
	func() {
		lockObj.RLock()
		defer lockObj.RUnlock()

		records = make([]TransactionRecord, 0, len(TransactionsDB))
		for _, tr := range TransactionsDB {
			records = append(records, tr)
		}
	}()

	sort.Slice(records, func(i, j int) bool {
		return records[i].Date.After(records[j].Date)
	})

	return &records
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	lockObj.RLock()
	defer lockObj.RUnlock()
	handleError(json.NewEncoder(w).Encode(&userAccount))
}

func HandleUI(w http.ResponseWriter, r *http.Request) {

	ctx := make(map[string]interface{})
	transactions := *getTransactionList()
	ctx["transactions"] = transactions

	if len(transactions) > 0 {
		ctx["title"] = fmt.Sprintf("%v transactions found", len(transactions))
	}else{
		ctx["title"] = "no transactions found"
	}

	handleError(templates.Stream("account.history", ctx, w))
}

type endpointFunc func(w http.ResponseWriter, r *http.Request)

type ServiceAPIHandler struct {
	serve        endpointFunc
	responseType apiResponse
}

func (sah ServiceAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h := w.Header()
	switch sah.responseType {

	case ResponseJSON:
		h.Set("Content-Type",  "application/json; charset=UTF-8")
		h.Set("Cache-Control",  "no-store")
		h.Set("Pragma",  "no-cache")

	case ResponseHTML:
		h.Set("Content-Type",  "text/html")
		h.Set("Cache-Control",  "no-store")
		h.Set("Pragma",  "no-cache")
	}

	//panic handler
	defer func() {

		if r := recover(); r != nil {

			w.WriteHeader(500)

			type servicePanicError struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			json.NewEncoder(w).Encode(&servicePanicError{
				Status: 500,
				Error:  fmt.Sprintf("fatal error - %v", r),
			})
		}

	}()
	sah.serve(w, r)
}

func NewEndpoint(f endpointFunc, r apiResponse) http.Handler {
	return ServiceAPIHandler{
		serve: f,
		responseType: r,
	}
}

/*
type notFoundHandler struct{}

func (h notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
*/