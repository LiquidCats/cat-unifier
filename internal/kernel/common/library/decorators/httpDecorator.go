package decorators

import (
	"cat-unifier/internal/kernel/common/contracts"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type IHttpDecorator interface {
	HandleGetBalance(writer http.ResponseWriter, request *http.Request)
	HandleGetBlock(writer http.ResponseWriter, request *http.Request)
	HandleGetTransaction(writer http.ResponseWriter, request *http.Request)
}

type httpDecorator struct {
	reader contracts.IReader
}

func (h *httpDecorator) HandleGetBalance(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := fmt.Sprintf("%s", vars["account"])
	result := h.reader.GetBalance(id)

	h.responseWithJson(writer, 200, result)
}

func (h *httpDecorator) HandleGetBlock(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := fmt.Sprintf("%s", vars["block"])
	withTransactions, _ := strconv.ParseBool(request.URL.Query().Get("withTransactions"))
	result := h.reader.GetBlock(id, withTransactions)

	h.responseWithJson(writer, 200, result)
}

func (h *httpDecorator) HandleGetTransaction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := fmt.Sprintf("%s", vars["transaction"])
	result := h.reader.GetTransaction(id)

	h.responseWithJson(writer, 200, result)
}

func (h *httpDecorator) responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *httpDecorator) respondWithError(w http.ResponseWriter, code int, message string) {
	h.responseWithJson(w, code, map[string]string{"error": message})
}

func NewHttpDecorator(reader contracts.IReader) IHttpDecorator {
	return &httpDecorator{reader: reader}
}
