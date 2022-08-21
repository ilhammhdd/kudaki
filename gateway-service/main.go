package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/user"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/rest"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			env := strings.Split(val, "=")
			env[1] = strings.ReplaceAll(env[1], ".exe", "")
			log.Printf("env[0], env[1] is %s, %s", env[0], env[1])
			os.Setenv(env[0], env[1])
		}
	}
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- restListener

	wp.PoolWG.Wait()
}

func restListener() {
	http.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OKE"))
		w.WriteHeader(http.StatusOK)
	}))
	/*
		kudaki file service
	*/
	http.Handle("/file", rest.MethodRouting{
		PostHandler: new(rest.StoreFile),
		GetHandler:  new(rest.RetrieveFile),
	})
	/*
		kudaki event aggregate
	*/
	http.Handle("/event/publish", rest.Authenticate(rest.Authorize([]user.UserRole{user.UserRole_ADMIN}, new(rest.PublishKudakiEvent))))
	http.Handle("/event", rest.MethodRouting{
		PostHandler:   rest.Authenticate(rest.Authorize([]user.UserRole{user.UserRole_ORGANIZER}, new(rest.AddKudakiEvent))),
		DeleteHandler: rest.Authenticate(rest.Authorize([]user.UserRole{user.UserRole_ORGANIZER}, new(rest.DeleteKudakiEvent))),
		GetHandler:    new(rest.RetrieveKudakiEvent),
	})
	http.Handle("/event-payment/doku/redirect", new(rest.RedirectDoku))
	http.Handle("/event-payment/doku/notify", new(rest.NotifyDoku))
	http.Handle("/event-payment/doku/identify", new(rest.IdentifyDoku))
	http.Handle("/event-payment/doku/review", new(rest.ReviewDoku))
	http.Handle("/event-payment/transactions",
		rest.MethodValidator(
			http.MethodGet,
			rest.Authenticate(
				rest.Authorize(
					[]user.UserRole{
						user.UserRole_ORGANIZER,
						user.UserRole_KUDAKI_TEAM,
						user.UserRole_ADMIN},
					new(rest.RetrieveOrganizerInvoices)))))
	http.Handle("/event-payment/doku/payment-request", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.PaymentRequest))))
	/*
		mountain aggregate
	*/
	http.Handle("/recommendation/item", rest.MethodRouting{
		DeleteHandler: rest.Authenticate(new(rest.DeleteRecommendedGearItem)),
		PostHandler:   rest.Authenticate(new(rest.AddRecommendedGearItem)),
	})
	http.Handle("/recommendation", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddRecommendedGear)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteRecommendedGear)),
	})
	http.Handle("/recommendations", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveRecommendedGears))))
	http.Handle("/recommendation/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveRecommendedGearItems))))
	http.Handle("/recommendation/upvote", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.UpVoteRecommendedGear))))
	http.Handle("/recommendation/downvote", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.DownVoteRecommendedGear))))
	http.Handle("/mountain", rest.MethodValidator(http.MethodGet, new(rest.RetrieveMountains)))
	/*
		order aggregate
	*/
	http.Handle("/order/owner", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveOwnerOrderHistories))))
	http.Handle("/order/tenant", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveTenantOrderHistories))))
	http.Handle("/order/owner-order-review", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.TenantReviewOwnerOrder))))
	http.Handle("/order/owner/approve", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.ApproveOwnerOrder))))
	http.Handle("/order/checkout", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.CheckOut))))
	http.Handle("/order/owner/disapprove", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.DisapproveOwnerOrder))))
	http.Handle("/order/confirm-returnment/owner", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.OwnerConfirmReturnment))))
	http.Handle("/order/owner/rented", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.OwnerOrderRented))))
	/*
		rental aggregate
	*/
	http.Handle("/rental/cart/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddCartItem)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteCartItem)),
		PatchHandler:  rest.Authenticate(new(rest.UpdateCartItem)),
	})
	http.Handle("/rental/cart/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveCartItems))))
	/*
		store aggregate
	*/
	http.Handle("/storefront/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddStorefrontItem)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteStorefrontItem)),
		PutHandler:    rest.Authenticate(new(rest.UpdateStorefrontItem)),
	})
	http.Handle("/storefront/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveStorefrontItems))))
	http.Handle("/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItems))))
	http.Handle("/items/search", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.SearchItems))))
	// http.Handle("/item-review/review", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.ReviewItem))))
	http.Handle("/item-review/reviews", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItemReviews))))
	http.Handle("/item-review/review/comment", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.CommentItemReview))))
	http.Handle("/item-review/review/comments", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItemReviewComments))))
	/*
		user aggregate
	*/
	http.Handle("/login", rest.MethodValidator(http.MethodPost, new(rest.Login)))
	http.Handle("/user/password/reset", rest.MethodRouting{
		GetHandler:   nil,
		PatchHandler: new(rest.ResetPassword),
		PostHandler:  new(rest.ResetPasswordSendEmail),
	})
	http.Handle("/signup", rest.MethodValidator(http.MethodPost, new(rest.Signup)))
	http.Handle("/user-info/address", rest.MethodRouting{
		PostHandler: rest.Authenticate(new(rest.AddAddress)),
		PutHandler:  rest.Authenticate(new(rest.UpdateAddress)),
	})
	http.Handle("/user-info/profile", rest.MethodRouting{
		PatchHandler: rest.Authenticate(new(rest.UpdateProfile)),
		GetHandler:   rest.Authenticate(new(rest.RetrieveProfile)),
	})
	http.Handle("/user-info/addresses", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveAddresses))))
	http.Handle("/user/password/change", rest.MethodValidator(http.MethodPatch, rest.Authenticate(new(rest.ChangePassword))))
	http.Handle("/user/verify", rest.MethodValidator(http.MethodGet, new(rest.VerifyUser)))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT"))}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
