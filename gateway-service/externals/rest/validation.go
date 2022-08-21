package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/user"

	"github.com/google/uuid"

	grpc_exteral "github.com/ilhammhdd/kudaki-gateway-service/externals/grpc"

	"google.golang.org/grpc"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

const (
	RegexEmail                       = `^[A-Za-z0-9](([_\.\-]?[a-zA-Z0-9]+)*)@([A-Za-z0-9]+)(([\.\-]?[a-zA-Z0-9]+)*)\.([A-Za-z]{2,})$`
	RegexEmailErrMessage             = "not a valid email address"
	RegexPassword                    = `^[\w_!@#$%*]{6,30}$`
	RegexPasswordErrMessage          = "not a valid, allowed alphanumeric with _!@#$%* symbols, minimal 6 and maximal 30 in length"
	RegexNotEmpty                    = `.*\S.*`
	RegexNotEmptyErrMessage          = "can't be empty"
	RegexURL                         = `^((((h)(t)|(f))(t)(p)((s)?))\://)?(www.|[a-zA-Z0-9].)[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,6}(\:[0-9]{1,5})*(/($|[a-zA-Z0-9\.\,\;\?\'\\\+&amp;%\$#\=~_\-]+))*$`
	RegexURLErrMessage               = "not a valid url"
	RegexJWT                         = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`
	RegexJWTErrMessage               = "not a valid jwt"
	RegexNumber                      = `^[0-9]+`
	RegexNumberErrMessage            = "not a number"
	RegexLatitude                    = `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`
	RegexLatitudeErrMessage          = "not a latitude value"
	RegexLongitude                   = `^[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`
	RegexLongitudeErrMessage         = "not a longitude value"
	RegexRole                        = `^(USER|ORGANIZER)$`
	RegexRoleErrMessage              = "not a valid one"
	RegexUUIDV4                      = `^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`
	RegexUUIDV4ErrMessage            = "not a valid UUID v4 form"
	RegexUnitofMeasurement           = `^(MM|CM|DM|M|DAM|HM|KM)$`
	RegexUnitofMeasurementErrMessage = "not a valid unit of measurement"
	RegexPriceDuration               = `^(DAY|WEEK|MONTH|YEAR)$`
	RegexPriceDurationErrMessage     = "not a valid price duration"
	RegexOrderStatus                 = `^(PENDING|APPROVED|DISAPPROVED|PROCESSED|RENTED|DONE)$`
	RegexOrderStatusErrMessage       = "not a valid order status"
)

type RestValidator interface {
	Validate() (errs *[]string, ok bool)
}

type RestValidation struct {
	Rules   map[string]string
	request *http.Request
}

func (rv RestValidation) ValidateIfExists() (existedParam *[]string, errs *[]string, valid bool) {

	var tempExistedParam []string
	var tempErrs []string
	rv.request.ParseMultipartForm(50000000)
	valid = true

	for param, rule := range rv.Rules {
		if rv.request.MultipartForm == nil {
			valid = false
			tempErrs = append(tempErrs, "multipart form required")
		} else if val, ok := rv.request.MultipartForm.Value[param]; ok {
			tempExistedParam = append(tempExistedParam, param)
			regexOk, regexErr := regexp.MatchString(rule, val[0])
			errorkit.ErrorHandled(regexErr)
			if !regexOk {
				tempErrs = append(tempErrs, fmt.Sprintf("%v %s", param, GetRegexErrorMessage(rule)))
				valid = false
			}
		}
	}

	return &tempExistedParam, &tempErrs, valid
}

func (rv RestValidation) Validate() (*[]string, bool) {
	rv.request.ParseMultipartForm(50000000)
	valid := true
	var errs []string

	for param, rule := range rv.Rules {
		if rv.request.MultipartForm == nil {
			valid = false
			errs = append(errs, "multipart form required")
		} else if val, ok := rv.request.MultipartForm.Value[param]; !ok {
			valid = false
			errs = append(errs, fmt.Sprintf("%s multipart/form-data parameter not exists", param))
		} else {
			regexOk, regexErr := regexp.MatchString(rule, val[0])
			errorkit.ErrorHandled(regexErr)
			if !regexOk {
				valid = false
				errs = append(errs, fmt.Sprintf("%s %s", param, GetRegexErrorMessage(rule)))
			}
		}
	}

	return &errs, valid
}

func (rv RestValidation) ValidateJSON() (*[]string, bool) {

	return nil, false
}

func GetRegexErrorMessage(rule string) string {
	switch rule {
	case RegexEmail:
		return RegexEmailErrMessage
	case RegexPassword:
		return RegexPasswordErrMessage
	case RegexNotEmpty:
		return RegexNotEmptyErrMessage
	case RegexURL:
		return RegexURLErrMessage
	case RegexJWT:
		return RegexJWTErrMessage
	case RegexNumber:
		return RegexNumberErrMessage
	case RegexLatitude:
		return RegexLatitudeErrMessage
	case RegexLongitude:
		return RegexLongitudeErrMessage
	case RegexRole:
		return RegexRoleErrMessage
	case RegexUUIDV4:
		return RegexUUIDV4ErrMessage
	case RegexPriceDuration:
		return RegexPriceDurationErrMessage
	case RegexUnitofMeasurement:
		return RegexUnitofMeasurementErrMessage
	case RegexOrderStatus:
		return RegexOrderStatusErrMessage
	default:
		return "regex error message not defined"
	}
}

func MethodValidator(m string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != m {
			resBody := adapters.ResponseBody{Errs: &[]string{fmt.Sprintf("method not allowed, need %s method", m)}}
			adapters.NewResponse(http.StatusMethodNotAllowed, &resBody).WriteResponse(&w)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func HeaderParamValidator(rules map[string]string, h http.Header) (*[]string, bool) {

	valid := true
	var errs []string

	for param, rule := range rules {
		if vals, ok := h[param]; !ok {
			valid = false
			errs = append(errs, fmt.Sprintf("%s header parameter not exists", param))
		} else {
			for _, val := range vals {
				regexOk, regexErr := regexp.MatchString(rule, val)
				errorkit.ErrorHandled(regexErr)
				if !regexOk {
					valid = false
					errs = append(errs, fmt.Sprintf("%s %s", param, GetRegexErrorMessage(rule)))
				}
			}
		}
	}

	return &errs, valid
}

type URLParamValidation struct {
	Rules  map[string]string
	Values url.Values
}

func (upv *URLParamValidation) Validate() (*[]string, bool) {

	valid := true
	var errs []string

	if len(upv.Values) == 0 {
		valid = false
		errs = append(errs, "url parameters needed")

		return &errs, valid
	}

	for param, rule := range upv.Rules {
		vals, ok := upv.Values[param]
		if !ok && len(upv.Values) > 0 {
			valid = false
			errs = append(errs, fmt.Sprintf("%s url parameter not exists", param))
		} else {
			for _, val := range vals {
				regexOk, regexErr := regexp.MatchString(rule, val)
				errorkit.ErrorHandled(regexErr)
				if !regexOk {
					valid = false
					errs = append(errs, fmt.Sprintf("%s %s", param, GetRegexErrorMessage(rule)))
				}
			}
		}
	}

	return &errs, valid
}

func (upv *URLParamValidation) ValidateIfExists() (*[]string, bool) {

	valid := true
	var errs []string

	if len(upv.Values) == 0 {
		valid = false
		errs = append(errs, "url parameters needed")

		return &errs, valid
	}

	for param, rule := range upv.Rules {
		vals, _ := upv.Values[param]
		for _, val := range vals {
			regexOk, regexErr := regexp.MatchString(rule, val)
			errorkit.ErrorHandled(regexErr)
			if !regexOk {
				valid = false
				errs = append(errs, fmt.Sprintf("%s %s", param, GetRegexErrorMessage(rule)))
			}
		}
	}

	return &errs, valid
}

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rules := map[string]string{
			"Kudaki-Token": RegexJWT}

		errs, valid := HeaderParamValidator(rules, r.Header)
		if !valid {
			resBody := adapters.ResponseBody{Errs: errs}
			adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

			return
		}

		uar := grpc_exteral.AuthenticateUser{
			Uid: uuid.New().String(),
			Jwt: r.Header.Get("Kudaki-Token")}

		conn, err := grpc.Dial(os.Getenv("USER_AUTH_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
		errorkit.ErrorHandled(err)

		defer conn.Close()

		client := grpc_exteral.NewUserClient(conn)

		ua, err := client.UserAuthentication(r.Context(), &uar)
		if err != nil {
			resBody := adapters.ResponseBody{Errs: &[]string{err.Error()}}
			adapters.NewResponse(http.StatusUnauthorized, &resBody).WriteResponse(&w)

			return
		}

		if ua.EventStatus.HttpCode != http.StatusOK {
			resBody := adapters.ResponseBody{Errs: &ua.EventStatus.Errors}
			adapters.NewResponse(int(ua.EventStatus.HttpCode), &resBody).WriteResponse(&w)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func Authorize(role []user.UserRole, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headerErr, ok := HeaderParamValidator(
			map[string]string{"Kudaki-Token": RegexJWT},
			r.Header)

		if !ok {
			resBody := adapters.ResponseBody{Errs: headerErr}
			adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

			return
		}

		conn, err := grpc.Dial(os.Getenv("USER_AUTH_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
		errorkit.ErrorHandled(err)

		uc := grpc_exteral.NewUserClient(conn)

		uar := &grpc_exteral.AuthorizeUser{
			Jwt:       r.Header.Get("Kudaki-Token"),
			UserRoles: role,
			Uid:       uuid.New().String()}

		uad, err := uc.UserAuthorization(r.Context(), uar)
		errorkit.ErrorHandled(err)

		if uad.EventStatus.HttpCode != http.StatusOK {
			resBody := adapters.ResponseBody{Errs: &uad.EventStatus.Errors}
			adapters.NewResponse(http.StatusUnauthorized, &resBody).WriteResponse(&w)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func TestAuthenticateJWT(w http.ResponseWriter, r *http.Request) {
	resBody := adapters.ResponseBody{}
	adapters.NewResponse(http.StatusOK, &resBody).WriteResponse(&w)
}

func TestAuthorizeUser(w http.ResponseWriter, r *http.Request) {
	resBody := adapters.ResponseBody{}
	adapters.NewResponse(http.StatusOK, &resBody).WriteResponse(&w)
}
