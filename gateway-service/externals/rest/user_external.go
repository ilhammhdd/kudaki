package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type Login struct{}

func (l *Login) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"email":    RegexEmail,
			"password": RegexPassword},
		request: r,
	}

	return restValidation.Validate()
}

func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := l.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.Login{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
	}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type ResetPasswordSendEmail struct{}

func (rpse *ResetPasswordSendEmail) validate(r *http.Request) (*[]string, bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules:   map[string]string{"email": RegexEmail},
		request: r}

	return restValidation.Validate()
}

func (rpse *ResetPasswordSendEmail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := rpse.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.ResetPasswordSendEmail{Consumer: kafka.NewConsumption(), Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type ChangePassword struct{}

func (cp *ChangePassword) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"old_password": RegexPassword,
			"new_password": RegexPassword},
		request: r}

	return restValidation.Validate()
}

func (cp *ChangePassword) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := cp.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.ChangePassword{Consumer: kafka.NewConsumption(), Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type ResetPassword struct{}

func (rp *ResetPassword) validate(r *http.Request) (*[]string, bool) {
	r.ParseMultipartForm(32 << 20)
	var allErrs []string
	var allOk bool = true

	restValidation := RestValidation{
		Rules:   map[string]string{"new_password": RegexPassword},
		request: r}

	if errs, ok := restValidation.Validate(); !ok {
		allErrs = append(allErrs, *errs...)
		allOk = false
	}

	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"reset_token": RegexJWT},
		Values: r.URL.Query()}

	if errs, ok := urlValidation.Validate(); !ok {
		allErrs = append(allErrs, *errs...)
		allOk = false
	}

	return &allErrs, allOk
}

func (rp *ResetPassword) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := rp.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.ResetPassword{Consumer: kafka.NewConsumption(), Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type Signup struct{}

func (s *Signup) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	v := RestValidation{
		Rules: map[string]string{
			"email":        RegexEmail,
			"password":     RegexPassword,
			"role":         RegexRole,
			"full_name":    RegexNotEmpty,
			"phone_number": RegexNotEmpty},
		request: r,
	}

	return v.Validate()
}

func (s *Signup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := s.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.Signup{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type AddAddress struct{}

func (aa *AddAddress) validate(r *http.Request) (*[]string, bool) {
	r.ParseMultipartForm(32 << 20)

	v := RestValidation{
		Rules: map[string]string{
			"receiver_name":         RegexNotEmpty,
			"full_address":          RegexNotEmpty,
			"receiver_phone_number": RegexNotEmpty,
			"latitude":              RegexLatitude,
			"longitude":             RegexLongitude,
			"zip_code":              RegexNotEmpty},
		request: r,
	}

	return v.Validate()
}

func (aa *AddAddress) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := aa.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.AddAddress{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type UpdateAddress struct{}

func (ua *UpdateAddress) validate(r *http.Request) (*[]string, bool) {
	r.ParseMultipartForm(32 << 20)

	v := RestValidation{
		Rules: map[string]string{
			"receiver_name":         RegexNotEmpty,
			"full_address":          RegexNotEmpty,
			"receiver_phone_number": RegexNotEmpty,
			"latitude":              RegexLatitude,
			"longitude":             RegexLongitude,
			"zip_code":              RegexNotEmpty,
			"address_uuid":          RegexUUIDV4},
		request: r,
	}

	return v.Validate()
}

func (ua *UpdateAddress) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := ua.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.UpdateAddress{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type UpdateProfile struct{}

func (up *UpdateProfile) validate(r *http.Request) (*[]string, bool) {
	r.ParseMultipartForm(32 << 20)

	v := RestValidation{
		Rules: map[string]string{
			"photo":        RegexNotEmpty,
			"full_name":    RegexNotEmpty,
			"phone_number": RegexNotEmpty},
		request: r,
	}

	return v.Validate()
}

func (up *UpdateProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := up.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.UpdateProfile{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type VerifyUser struct{}

func (vu *VerifyUser) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules:  map[string]string{"verify_token": RegexJWT},
		Values: r.URL.Query(),
	}
	return paramValidation.Validate()
}

func (vu *VerifyUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := vu.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.VerifyUser{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type RetrieveAddresses struct{}

func (ra *RetrieveAddresses) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	adapter := &adapters.RetrieveAddresses{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type RetrieveProfile struct{}

func (rp *RetrieveProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	adapter := &adapters.RetrieveProfile{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}
