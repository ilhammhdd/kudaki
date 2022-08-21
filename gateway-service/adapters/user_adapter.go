package adapters

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/google/uuid"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type Signup struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (s *Signup) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := events.Signup{
		Email:       r.MultipartForm.Value["email"][0],
		FullName:    r.MultipartForm.Value["full_name"][0],
		Password:    r.MultipartForm.Value["password"][0],
		PhoneNumber: r.MultipartForm.Value["phone_number"][0],
		Role:        r.MultipartForm.Value["role"][0],
		Uid:         uuid.New().String(),
	}
	out, err := proto.Marshal(&outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
	/* outEvent := events.SignupRequested{
		Email:       r.MultipartForm.Value["email"][0],
		FullName:    r.MultipartForm.Value["full_name"][0],
		Password:    r.MultipartForm.Value["password"][0],
		PhoneNumber: r.MultipartForm.Value["phone_number"][0],
		Photo:       r.MultipartForm.Value["photo"][0],
		Role:        r.MultipartForm.Value["role"][0],
		Uid:         uuid.New().String(),
	}
	out, err := proto.Marshal(&outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out */
}

func (s *Signup) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.Signedup)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (s *Signup) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.Signedup

	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (s *Signup) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: s.Consumer,
		InTopic:/* events.UserTopic_SIGNED_UP.String() */ events.UserAuthServiceEventTopic_SIGNED_UP.String(),
		OutTopic:/* events.UserTopic_SIGN_UP_REQUESTED.String() */ events.UserAuthServiceCommandTopic_SIGN_UP.String(),
		Producer:       s.Producer,
		InEventChecker: s}
}

type Login struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (l *Login) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.LoginRequested) */ new(events.Login)
	outEvent.Email = r.MultipartForm.Value["email"][0]
	outEvent.Password = r.MultipartForm.Value["password"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (l *Login) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.Loggedin)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = map[string]string{"token": inEvent.User.Token}
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (l *Login) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.Loggedin
	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}

	return nil, false
}

func (l *Login) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: l.Consumer,
		InTopic:/* events.UserTopic_LOGGED_IN.String() */ events.UserAuthServiceEventTopic_LOGGEDIN.String(),
		InEventChecker: l,
		OutTopic:/* events.UserTopic_LOGIN_REQUESTED.String() */ events.UserAuthServiceCommandTopic_LOGIN.String(),
		Producer: l.Producer}
}

type VerifyUser struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (vu *VerifyUser) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.VerifyUserRequested) */ new(events.VerifyUser)
	outEvent.Uid = uuid.New().String()
	outEvent.VerifyToken = r.URL.Query().Get("verify_token")

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (vu *VerifyUser) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.UserVerified)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (vu *VerifyUser) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.UserVerified
	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (vu *VerifyUser) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: vu.Consumer,
		InTopic:/* events.UserTopic_USER_VERIFIED.String() */ events.UserAuthServiceEventTopic_USER_VERIFIED.String(),
		InEventChecker: vu,
		OutTopic:/* events.UserTopic_VERIFY_USER_REQUESTED.String() */ events.UserAuthServiceCommandTopic_VERIFY_USER.String(),
		Producer: vu.Producer,
	}
}

type ChangePassword struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (cp *ChangePassword) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.ChangePasswordRequested) */ new(events.ChangePassword)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.NewPassword = r.MultipartForm.Value["new_password"][0]
	outEvent.OldPassword = r.MultipartForm.Value["old_password"][0]
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (cp *ChangePassword) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PasswordChanged)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (cp *ChangePassword) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.PasswordChanged

	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (cp *ChangePassword) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: cp.Consumer,
		InTopic:/* events.UserTopic_PASSWORD_CHANGED.String() */ events.UserAuthServiceEventTopic_PASSWORD_CHANGED.String(),
		InEventChecker: cp,
		OutTopic:/* events.UserTopic_CHANGE_PASSWORD_REQUESTED.String() */ events.UserAuthServiceCommandTopic_CHANGE_PASSWORD.String(),
		Producer: cp.Producer}
}

type ResetPasswordSendEmail struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rpse *ResetPasswordSendEmail) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.SendResetPasswordEmailRequested) */ new(events.SendResetPasswordEmail)
	outEvent.Email = r.MultipartForm.Value["email"][0]
	outEvent.Uid = uuid.New().String()

	log.Println("outEvent : ", outEvent)

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rpse *ResetPasswordSendEmail) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ResetPasswordEmailSent)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rpse *ResetPasswordSendEmail) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ResetPasswordEmailSent

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}

	return nil, false
}

func (rpse *ResetPasswordSendEmail) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: rpse.Consumer,
		InTopic:/* events.UserTopic_RESET_PASSWORD_EMAIL_SENT.String() */ events.UserAuthServiceEventTopic_RESET_PASSWORD_EMAIL_SENT.String(),
		InEventChecker: rpse,
		OutTopic:/* events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED.String() */ events.UserAuthServiceCommandTopic_SEND_RESET_PASSWORD_EMAIL.String(),
		Producer: rpse.Producer}
}

type ResetPassword struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rp *ResetPassword) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.PasswordReseted

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (rp *ResetPassword) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: rp.Consumer,
		InTopic:/* events.UserTopic_PASSWORD_RESETED.String() */ events.UserAuthServiceEventTopic_PASSWORD_RESETED.String(),
		InEventChecker: rp,
		OutTopic:/* events.UserTopic_RESET_PASSWORD_REQUESTED.String() */ events.UserAuthServiceCommandTopic_RESET_PASSWORD.String(),
		Producer: rp.Producer}
}

func (rp *ResetPassword) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PasswordReseted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rp *ResetPassword) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* &events.ResetPasswordRequested{ */ &events.ResetPassword{
		NewPassword: r.MultipartForm.Value["new_password"][0],
		ResetToken:  r.URL.Query().Get("reset_token"),
		Uid:         uuid.New().String()}

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

type AddAddress struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (aa *AddAddress) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddAddress)

	lat, err := strconv.ParseFloat(r.MultipartForm.Value["latitude"][0], 32)
	errorkit.ErrorHandled(err)
	long, err := strconv.ParseFloat(r.MultipartForm.Value["longitude"][0], 32)
	errorkit.ErrorHandled(err)

	outEvent.FullAddress = r.MultipartForm.Value["full_address"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Latitude = float32(lat)
	outEvent.Longitude = float32(long)
	outEvent.ReceiverName = r.MultipartForm.Value["receiver_name"][0]
	outEvent.ReceiverPhoneNumber = r.MultipartForm.Value["receiver_phone_number"][0]
	outEvent.Uid = uuid.New().String()
	outEvent.ZipCode = r.MultipartForm.Value["zip_code"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (aa *AddAddress) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.AddressAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (aa *AddAddress) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.AddressAdded

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (aa *AddAddress) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       aa.Consumer,
		InEventChecker: aa,
		InTopic:        events.UserInfoServiceEventTopic_ADDRESS_ADDED.String(),
		OutTopic:       events.UserInfoServiceCommandTopic_ADD_ADDRESS.String(),
		Producer:       aa.Producer}
}

type UpdateAddress struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ua *UpdateAddress) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	lat, err := strconv.ParseFloat(r.MultipartForm.Value["latitude"][0], 32)
	errorkit.ErrorHandled(err)
	long, err := strconv.ParseFloat(r.MultipartForm.Value["longitude"][0], 32)
	errorkit.ErrorHandled(err)

	outEvent := &events.UpdateAddress{
		AddressUuid:         r.MultipartForm.Value["address_uuid"][0],
		FullAddress:         r.MultipartForm.Value["full_address"][0],
		KudakiToken:         r.Header.Get("Kudaki-Token"),
		Latitude:            float32(lat),
		Longitude:           float32(long),
		ReceiverName:        r.MultipartForm.Value["receiver_name"][0],
		ReceiverPhoneNumber: r.MultipartForm.Value["receiver_phone_number"][0],
		Uid:                 uuid.New().String(),
		ZipCode:             r.MultipartForm.Value["zip_code"][0]}

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)
	return outEvent.Uid, out
}

func (ua *UpdateAddress) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.AddressUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (ua *UpdateAddress) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.AddressUpdated

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (ua *UpdateAddress) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ua.Consumer,
		InEventChecker: ua,
		InTopic:        events.UserInfoServiceEventTopic_ADDRESS_UPDATED.String(),
		OutTopic:       events.UserInfoServiceCommandTopic_UPDATE_ADDRESS.String(),
		Producer:       ua.Producer}
}

type UpdateProfile struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (up *UpdateProfile) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.UpdateProfile)

	outEvent.FullName = r.MultipartForm.Value["full_name"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.PhoneNumber = r.MultipartForm.Value["phone_number"][0]
	outEvent.Photo = r.MultipartForm.Value["photo"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (up *UpdateProfile) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ProfileUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (up *UpdateProfile) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ProfileUpdated

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (up *UpdateProfile) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       up.Consumer,
		InEventChecker: up,
		InTopic:        events.UserInfoServiceEventTopic_PROFILE_UPDATED.String(),
		OutTopic:       events.UserInfoServiceCommandTopic_UPDATE_PROFILE.String(),
		Producer:       up.Producer}
}

type RetrieveAddresses struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ra *RetrieveAddresses) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveAddresses)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ra *RetrieveAddresses) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.AddressesRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	type resDataSchema struct {
		Addresses json.RawMessage `json:"addresses"`
	}

	resData := resDataSchema{
		Addresses: inEvent.Result}

	resBody.Data = resData

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (ra *RetrieveAddresses) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.AddressesRetrieved

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (ra *RetrieveAddresses) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ra.Consumer,
		InEventChecker: ra,
		InTopic:        events.UserInfoServiceEventTopic_ADDRESSES_RETRIEVED.String(),
		OutTopic:       events.UserInfoServiceCommandTopic_RETRIEVE_ADDRESSES.String(),
		Producer:       ra.Producer}
}

type RetrieveProfile struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ra *RetrieveProfile) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveProfile)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ra *RetrieveProfile) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ProfileRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (ra *RetrieveProfile) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ProfileRetrieved

	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (ra *RetrieveProfile) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ra.Consumer,
		InEventChecker: ra,
		InTopic:        events.UserInfoServiceEventTopic_PROFILE_RETRIEVED.String(),
		OutTopic:       events.UserInfoServiceCommandTopic_RETRIEVE_PROFILE.String(),
		Producer:       ra.Producer}
}
