package api

import (
	"encoding/json"
	"net/http"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type PaymentAPI struct {
	data *data.PaymentData
}

func ServePaymentResource(r *mux.Router, data data.PaymentData) {
	api := &PaymentAPI{data: &data}
	r.HandleFunc("/payments", api.GetAllPayments).Methods("GET")
}

func (a PaymentAPI) GetAllPayments(writer http.ResponseWriter, request *http.Request) {
	payments, err := a.data.ReadAll()
	if err != nil {
		log.Error(err)
		_, err = writer.Write([]byte("got an error when tried to get list of payments"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Debugf("GetAllPayments method result: %v", payments)
	err = json.NewEncoder(writer).Encode(payments)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
