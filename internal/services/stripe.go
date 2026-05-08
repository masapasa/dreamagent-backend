package services

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/webhook"
)

type StripeService struct {
	secretKey string
}

func NewStripeService(secretKey string) *StripeService {
	stripe.Key = secretKey
	return &StripeService{secretKey: secretKey}
}

// CreateCheckoutSession for Pro subscription (example)
func (s *StripeService) CreateCheckoutSession(userID string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String("https://your-domain.com/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("https://your-domain.com/cancel"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_123"), // replace with real price ID
				Quantity: stripe.Int64(1),
			},
		},
	}
	return session.New(params)
}

// HandleWebhook processes Stripe events and updates subscriptions table
func (s *StripeService) HandleWebhook(payload []byte, signature string) error {
	event, err := webhook.ConstructEvent(payload, signature, s.secretKey) // note: use webhook secret in prod
	if err != nil {
		return err
	}
	// TODO: switch on event.Type, update DB subscriptions for user
	return nil
}
