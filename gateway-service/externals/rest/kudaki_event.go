package rest

import (
	"log"
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type AddKudakiEvent struct{}

func (ake *AddKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ake.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.AddKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ake *AddKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"venue":            RegexNotEmpty,
			"description":      RegexNotEmpty,
			"duration_from":    RegexNumber,
			"duration_to":      RegexNumber,
			"name":             RegexNotEmpty,
			"ad_duration_from": RegexNotEmpty,
			"ad_duration_to":   RegexNotEmpty,
			"file_path":        RegexNotEmpty},
		request: r}
	return restValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type DeleteKudakiEvent struct{}

func (dke *DeleteKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dke.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DeleteKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (dke *DeleteKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"event_uuid": RegexUUIDV4},
		request: r}
	return restValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type RedirectDoku struct{}

func (ap *RedirectDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *RedirectDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type NotifyDoku struct{}

func (ap *NotifyDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CONTINUE"))
}

func (ap *NotifyDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type IdentifyDoku struct{}

func (ap *IdentifyDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *IdentifyDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type ReviewDoku struct{}

func (ap *ReviewDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *ReviewDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type RetrieveOrganizerInvoices struct{}

func (ap *RetrieveOrganizerInvoices) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.RetrieveOrganizerInvoices{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

func (ap *RetrieveOrganizerInvoices) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type PaymentRequest struct{}

func (ap *PaymentRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.PaymentRequest{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

func (ap *PaymentRequest) validate(r *http.Request) (errs *[]string, ok bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"transaction_id_merchant": RegexNotEmpty,
			"session_id":              RegexUUIDV4,
			"hashed_words":            RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type RetrieveKudakiEvent struct{}

func (dke *RetrieveKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dke.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (dke *RetrieveKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	urlParamValidation := URLParamValidation{
		Rules: map[string]string{
			"kudaki_event_uuid": RegexUUIDV4},
		Values: r.URL.Query(),
	}

	return urlParamValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type PublishKudakiEvent struct{}

func (pke *PublishKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := pke.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.PublishKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (pke *PublishKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"kudaki_event_uuid": RegexUUIDV4},
		request: r}
	return restValidation.Validate()
}
