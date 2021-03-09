package firestore

import (
	"context"
	"errors"

	"github.com/rhyth-me/backend/domain/model"
)

// GetUserByScreenName - Fetch user by screenName.
func GetUserByScreenName(sn string) (*model.User, error) {
	ctx := context.Background()

	iter := Client.Collection(Users).
		Where("profile.screenName", "==", sn).
		Documents(ctx)

	docs, err := iter.GetAll()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if len(docs) < 1 {
		return nil, errors.New("User Not Found")
	}

	var User *model.User
	docs[0].DataTo(&User)

	return User, nil

}

// GetUserByGoogleID - Fetch user by google ID.
func GetUserByGoogleID(googleID string) (*model.User, error) {
	ctx := context.Background()

	iter := Client.Collection(Users).
		Where("google.id", "==", googleID).
		Documents(ctx)

	docs, err := iter.GetAll()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if len(docs) < 1 {
		return nil, errors.New("User Not Found")
	}

	var User *model.User
	docs[0].DataTo(&User)

	return User, nil

}

// StoreUser - Save the user to the DB.
func StoreUser(user *model.User) (*model.User, error) {
	ctx := context.Background()

	_, err := Client.Collection(Users).Doc(user.Google.ID).Set(ctx, user)

	return user, err

}
