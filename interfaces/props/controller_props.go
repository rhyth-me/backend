// Package props is a scaffold file for props of controllers
package props

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/stripe/stripe-go/v72/client"
)

// ControllerProps is passed from Bootstrap() to all controllers
type ControllerProps struct {
	Auth      *auth.Client
	Firestore *firestore.Client
	Stripe    *client.API
}
