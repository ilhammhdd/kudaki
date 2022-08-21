package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type RetrieveOwnerOrderHistories struct{}

func (rooh *RetrieveOwnerOrderHistories) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rooh.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveOwnerOrderHistories{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rooh *RetrieveOwnerOrderHistories) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"order_status": RegexOrderStatus,
			"offset":       RegexNumber,
			"limit":        RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type RetrieveTenantOrderHistories struct{}

func (rtoh *RetrieveTenantOrderHistories) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rtoh.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveTenantOrderHistories{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rtoh *RetrieveTenantOrderHistories) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"order_status": RegexOrderStatus,
			"offset":       RegexNumber,
			"limit":        RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type TenantReviewOwnerOrder struct{}

func (tro *TenantReviewOwnerOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := tro.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.TenantReviewOwnerOrder{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (tro *TenantReviewOwnerOrder) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"rating":           RegexNotEmpty,
			"owner_order_uuid": RegexUUIDV4,
			"review":           RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

type ApproveOwnerOrder struct{}

func (ao *ApproveOwnerOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ao.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.ApproveOwnerOrder{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ao *ApproveOwnerOrder) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"owner_order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type DisapproveOwnerOrder struct{}

func (do *DisapproveOwnerOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := do.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DisapproveOwnerOrder{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (do *DisapproveOwnerOrder) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"owner_order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type CheckOut struct{}

func (co *CheckOut) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := co.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.CheckOut{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (co *CheckOut) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"cart_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type OwnerConfirmReturnment struct{}

func (ocr *OwnerConfirmReturnment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ocr.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.OwnerConfirmReturnment{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ocr *OwnerConfirmReturnment) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"owner_order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

// ----------------------------------------------------------------

type OwnerOrderRented struct{}

func (oor *OwnerOrderRented) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := oor.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.OwnerOrderRented{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (oor *OwnerOrderRented) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"owner_order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}
