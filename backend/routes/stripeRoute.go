
package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	//"github.com/stripe/stripe-go/v72/customer"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"os"
)

type userPayment struct {
	Amount int64 `json:"amount"`
	Email string `json:"email"`
}
//func getPayment(s *stripe.PaymentIntent) (*stripe.PaymentIntent){

//}
func paymentRequest(c *gin.Context){
	var up userPayment;
	var upPtr *userPayment = &up
	var err error = c.BindJSON(upPtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"error receiving paymeny"});
		return;
	}
	// -------getting sripe key----------------//
	envErr := godotenv.Load(".env");
	if envErr != nil {
		panic(envErr);
	}      
	var STRIPE_SECRET string = os.Getenv("STRIPE_SECRET_KEY");
	stripe.Key = STRIPE_SECRET;
	//--------------customer----------------//
	//--------------payment----------------//
	fmt.Print(up.Amount);
   params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(up.Amount*100),   /////up.Amount
        Currency: stripe.String(string(stripe.CurrencyUSD)),
        PaymentMethodTypes: stripe.StringSlice([]string{
            "card",
        }),
		//Source:&stripe.SourceParams{Token: }, // stripe.String("tok_visa")
        ReceiptEmail: stripe.String("matsuneterence@gmail.com"),
    }
    pi, err := paymentintent.New(params)
	// checking for errors and returning json----------------//
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not process patmeny"});
		return;
	}
	c.IndentedJSON(200, gin.H{"message":pi})

}

func StripeRoutes(g *gin.RouterGroup){
	g.POST("/stp", paymentRequest)
}

