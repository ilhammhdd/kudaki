package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := asi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (asi *AddStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"name":                RegexNotEmpty,
			"amount":              RegexNumber,
			"unit":                RegexNotEmpty,
			"price":               RegexNumber,
			"description":         RegexNotEmpty,
			"photo":               RegexNotEmpty,
			"price_duration":      RegexPriceDuration,
			"length":              RegexNumber,
			"width":               RegexNumber,
			"height":              RegexNumber,
			"unit_of_measurement": RegexUnitofMeasurement,
			"color":               RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dsi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.DeleteStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (dsi *DeleteStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"item_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := usi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.UpdateStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (usi *UpdateStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"item_uuid":           RegexUUIDV4,
			"name":                RegexNotEmpty,
			"amount":              RegexNumber,
			"unit":                RegexNotEmpty,
			"price":               RegexNumber,
			"description":         RegexNotEmpty,
			"photo":               RegexNotEmpty,
			"price_duration":      RegexPriceDuration,
			"length":              RegexNumber,
			"width":               RegexNumber,
			"height":              RegexNumber,
			"unit_of_measurement": RegexUnitofMeasurement,
			"color":               RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

type RetrieveStorefrontItems struct{}

func (rsi *RetrieveStorefrontItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := rsi.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
	}

	adapter := adapters.RetrieveStorefrontItems{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (rsi *RetrieveStorefrontItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlParamValidation := URLParamValidation{
		Rules: map[string]string{
			"limit":  RegexNumber,
			"offset": RegexNumber},
		Values: r.URL.Query()}

	return urlParamValidation.Validate()
}

// ----------------------------------------------

type RetrieveItems struct{}

func (ri *RetrieveItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ri.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.RetrieveItems{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (ri *RetrieveItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

type SearchItems struct{}

func (si *SearchItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := si.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := adapters.SearchItems{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &edha).WriteResponse(&w)
}

func (si *SearchItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlParamValidation := URLParamValidation{
		Rules: map[string]string{
			"keyword": RegexNotEmpty,
			"offset":  RegexNumber,
			"limit":   RegexNumber},
		Values: r.URL.Query()}

	return urlParamValidation.Validate()
}

type RetrieveItemReviews struct{}

func (rir *RetrieveItemReviews) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rir.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := adapters.RetrieveItemReviews{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &edha).WriteResponse(&w)
}

func (rir *RetrieveItemReviews) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"item_uuid": RegexUUIDV4,
			"offset":    RegexNumber,
			"limit":     RegexNumber},
		Values: r.URL.Query()}
	return paramValidation.Validate()
}

type CommentItemReview struct{}

func (cir *CommentItemReview) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := cir.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.CommentItemReview{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (cir *CommentItemReview) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"item_review_uuid": RegexUUIDV4,
			"comment":          RegexNotEmpty},
		request: r}
	return restValidation.Validate()
}

type RetrieveItemReviewComments struct{}

func (rirc *RetrieveItemReviewComments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rirc.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveItemReviewComments{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rirc *RetrieveItemReviewComments) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"item_review_uuid": RegexUUIDV4,
			"offset":           RegexNumber,
			"limit":            RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}
