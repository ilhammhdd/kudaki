package usecases

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/kudaki_event"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type KudakiEventAdded struct {
	DBO DBOperator
}

func (ae *KudakiEventAdded) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ae.initInOutEvent(in)

	outEvent.DokuInvoice = ae.initDokuInvoice(inEvent)
	// respStatCode := ae.reqPayment(outEvent.DokuInvoice)
	if outEvent.DokuInvoice == nil {
		outEvent.EventStatus.Errors = []string{"payment request to doku failed"}
		outEvent.EventStatus.HttpCode = http.StatusInternalServerError
	}

	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ae *KudakiEventAdded) initInOutEvent(in proto.Message) (inEvent *events.KudakiEventAdded, outEvent *events.KudakiEventDokuInvoiceIssued) {
	inEvent = in.(*events.KudakiEventAdded)

	outEvent = new(events.KudakiEventDokuInvoiceIssued)
	outEvent.EventStatus = new(events.Status)
	outEvent.Organizer = inEvent.Organizer
	outEvent.Uid = inEvent.Uid

	return
}

func (ae *KudakiEventAdded) initDokuInvoice(inEvent *events.KudakiEventAdded) *kudaki_event.DokuInvoice {
	durationFromTime := time.Unix(inEvent.KudakiEvent.DurationFrom.Seconds, 0)
	durationToTime := time.Unix(inEvent.KudakiEvent.DurationFrom.Seconds, 0)

	adDurationFromTime := time.Unix(inEvent.KudakiEvent.AdDurationFrom.Seconds, 0)
	adDurationToTime := time.Unix(inEvent.KudakiEvent.AdDurationTo.Seconds, 0)

	var amount float32
	amount = (float32(inEvent.KudakiEvent.AdDurationTo.Seconds) - float32(inEvent.KudakiEvent.AdDurationFrom.Seconds)) / 86400
	log.Println("amount : ", amount)
	basket := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v",
		adDurationFromTime.Format(time.RFC3339),
		adDurationToTime.Format(time.RFC3339),
		inEvent.KudakiEvent.Description,
		durationFromTime.Format(time.RFC3339),
		durationToTime.Format(time.RFC3339),
		inEvent.KudakiEvent.Name,
		inEvent.KudakiEvent.OrganizerUserUuid,
		inEvent.KudakiEvent.Uuid,
		inEvent.KudakiEvent.Venue)

	currency, err := strconv.ParseInt(os.Getenv("CURRENCY"), 10, 32)
	errorkit.ErrorHandled(err)

	t := time.Now()
	tFormatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	tParsed, _ := strconv.ParseInt(tFormatted, 10, 64)
	transId := t.Unix()
	transIdStr := fmt.Sprintf("%v", transId)

	errorkit.ErrorHandled(err)

	words := fmt.Sprintf("%v%v%v%v", amount, os.Getenv("MALLID"), os.Getenv("SHARED_KEY"), transId)
	h := sha1.New()
	io.WriteString(h, words)

	di := &kudaki_event.DokuInvoice{
		Amount:                amount,
		Basket:                basket,
		Currency:              int32(currency),
		Email:                 inEvent.Organizer.Email,
		KudakiEventUuid:       inEvent.KudakiEvent.Uuid,
		Name:                  ae.retrieveProfile(inEvent.Organizer.Uuid).FullName,
		PurchaseAmount:        float32(amount * 5000),
		PurchaseCurrency:      int32(currency),
		RequestDateTime:       tParsed,
		SessionId:             inEvent.Uid,
		TransactionIdMerchant: transIdStr,
		Uuid:                  uuid.New().String(),
		Words:                 base64.StdEncoding.EncodeToString(h.Sum(nil)),
	}

	log.Println("initialized doku invoice : ", di)

	return di
}

func (ae *KudakiEventAdded) reqPayment(dokuInvoice *kudaki_event.DokuInvoice) (stat int) {
	var body bytes.Buffer
	mpWriter := multipart.NewWriter(&body)

	mpWriter.WriteField("MALLID", os.Getenv("MALLID"))
	mpWriter.WriteField("CHAINMERCHANT", os.Getenv("CHAINMERCHANT"))
	mpWriter.WriteField("AMOUNT", fmt.Sprintf("%v", dokuInvoice.Amount))
	mpWriter.WriteField("PURCHASEAMOUNT", fmt.Sprintf("%v", dokuInvoice.PurchaseAmount))
	mpWriter.WriteField("TRANSIDMERCHANT", fmt.Sprintf("%v", dokuInvoice.TransactionIdMerchant))
	mpWriter.WriteField("WORDS", string(dokuInvoice.Words))
	mpWriter.WriteField("REQUESTDATETIME", fmt.Sprintf("%v", dokuInvoice.RequestDateTime))
	mpWriter.WriteField("CURRENCY", os.Getenv("CURRENCY"))
	mpWriter.WriteField("PURCHASECURRENCY", os.Getenv("CURRENCY"))
	mpWriter.WriteField("SESSIONID", dokuInvoice.SessionId)
	mpWriter.WriteField("NAME", dokuInvoice.Name)
	mpWriter.WriteField("EMAIL", dokuInvoice.Email)
	mpWriter.WriteField("BASKET", dokuInvoice.Basket)

	mpWriter.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodPost, "https://staging.doku.com/Suite/Receive", &body)
	errorkit.ErrorHandled(err)
	resp, err := client.Do(req)
	errorkit.ErrorHandled(err)

	log.Println("doku payment response status : ", resp.StatusCode)

	return resp.StatusCode
}

func (ae *KudakiEventAdded) retrieveProfile(userUUID string) *user.Profile {
	row, err := ae.DBO.QueryRow("SELECT id,uuid,full_name,photo,created_at FROM kudaki_user.profiles WHERE user_uuid = ?;", userUUID)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	var createdAt int64
	row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &createdAt)
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	profile.CreatedAt = createdAtProto

	return &profile
}
