package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type DeleteRecommendedGearItem struct{}

func (drgi *DeleteRecommendedGearItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := drgi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DeleteRecommendedGearItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (drgi *DeleteRecommendedGearItem) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"recommended_gear_item_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type AddRecommendedGear struct{}

func (arg *AddRecommendedGear) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	edha := adapters.AddRecommendedGear{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &edha).WriteResponse(&w)
}

func (arg *AddRecommendedGear) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

type DeleteRecommendedGear struct{}

func (drg *DeleteRecommendedGear) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := drg.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DeleteRecommendedGear{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (drg *DeleteRecommendedGear) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"recommended_gear_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type AddRecommendedGearItem struct{}

func (argi *AddRecommendedGearItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := argi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.AddRecommendedGearItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (argi *AddRecommendedGearItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"recommended_gear_uuid": RegexUUIDV4,
			"item_type":             RegexNotEmpty,
			"total":                 RegexNumber},
		request: r}
	return restValidation.Validate()
}

type RetrieveRecommendedGears struct{}

func (rrg *RetrieveRecommendedGears) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rrg.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveRecommendedGears{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rrg *RetrieveRecommendedGears) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"mountain_uuid": RegexUUIDV4,
			"offset":        RegexNumber,
			"limit":         RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type RetrieveRecommendedGearItems struct{}

func (rrgi *RetrieveRecommendedGearItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rrgi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveRecommendedGearItems{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rrgi *RetrieveRecommendedGearItems) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"recommended_gear_uuid": RegexUUIDV4,
			"offset":                RegexNumber,
			"limit":                 RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type UpVoteRecommendedGear struct{}

func (uvrg *UpVoteRecommendedGear) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := uvrg.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.UpVoteRecommendedGear{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (uvrg *UpVoteRecommendedGear) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"recommended_gear_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type DownVoteRecommendedGear struct{}

func (dvrg *DownVoteRecommendedGear) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dvrg.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DownVoteRecommendedGear{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (dvrg *DownVoteRecommendedGear) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"recommended_gear_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

// ----------------------------------------------------------------------------------------------

type RetrieveMountains struct{}

func (rm *RetrieveMountains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rm.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveMountains{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rm *RetrieveMountains) validate(r *http.Request) (errs *[]string, ok bool) {
	urlParamValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlParamValidation.Validate()
}

// ----------------------------------------------------------------------------------------------

type SearchMountains struct{}

func (sm *SearchMountains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := sm.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.SearchMountains{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (sm *SearchMountains) validate(r *http.Request) (errs *[]string, ok bool) {
	urlParamValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlParamValidation.Validate()
}
