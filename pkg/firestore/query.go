package firestore

import (
	"context"
	"errors"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/rhyth-me/backend/domain/model"
)

// GetUserByGoogleID -
func GetUserByGoogleID(c *firestore.Client, googleID string) (*model.User, error) {

	iter := c.Collection(os.Getenv("USERS_COLLECTION")).
		Where("google.id", "==", googleID).
		Documents(context.Background())

	docs, err := iter.GetAll()
	if err != nil {
		return nil, errors.New("500")
	}

	if len(docs) < 1 {
		return nil, nil
	}

	var User *model.User
	docs[0].DataTo(&User)

	return User, nil

}

// GetUserByScreenName -
func GetUserByScreenName(c *firestore.Client, sn string) (*model.User, error) {

	iter := c.Collection(os.Getenv("USERS_COLLECTION")).
		Where("profile.screenName", "==", sn).
		Documents(context.Background())

	docs, err := iter.GetAll()
	if err != nil {
		return nil, errors.New("500")
	}

	if len(docs) < 1 {
		return nil, nil
	}

	var User *model.User
	docs[0].DataTo(&User)

	return User, nil

}
