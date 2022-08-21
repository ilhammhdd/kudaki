package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type AddCartItem struct{}

func (aci *AddCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := aci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddCartItem{Producer: kafka.NewProduction(), Consumer: kafka.NewConsumption()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (aci *AddCartItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"item_uuid":     RegexUUIDV4,
			"item_amount":   RegexNumber,
			"duration_from": RegexNumber,
			"duration":      RegexNumber},
		request: r}

	return restValidation.Validate()
}

type RetrieveCartItems struct{}

func (rci *RetrieveCartItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := adapters.RetrieveCartItems{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &edha).WriteResponse(&w)
}

func (rci *RetrieveCartItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

type DeleteCartItem struct{}

func (dci *DeleteCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.DeleteCartItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

func (dci *DeleteCartItem) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules:  map[string]string{"cart_item_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

type UpdateCartItem struct{}

func (uci *UpdateCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := uci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.UpdateCartItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

func (uci *UpdateCartItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"cart_item_uuid": RegexUUIDV4,
			"total_item":     RegexNumber},
		request: r}

	return restValidation.Validate()
}
