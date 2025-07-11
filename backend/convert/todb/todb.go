package todb

import (
	"github.com/torfstack/kayvault/backend/models"
	sqlc "github.com/torfstack/kayvault/sql/gen"
)

func Secret(in models.Secret) sqlc.Secret {
	return sqlc.Secret{
		ID:    *in.ID,
		Value: []byte(in.Value),
		Key:   in.Key,
		Url:   in.Url,
		Tags:  tagsString(in.Tags),
	}
}

func InsertSecretParams(in models.Secret, userID int64) sqlc.InsertSecretParams {
	return sqlc.InsertSecretParams{
		Value:  []byte(in.Value),
		Key:    in.Key,
		Url:    in.Url,
		Tags:   tagsString(in.Tags),
		UserID: userID,
	}
}

func UpdateSecretParams(in models.Secret, userID int64) sqlc.UpdateSecretParams {
	return sqlc.UpdateSecretParams{
		ID:     *in.ID,
		Value:  []byte(in.Value),
		Key:    in.Key,
		Url:    in.Url,
		Tags:   tagsString(in.Tags),
		UserID: userID,
	}
}

func InsertUserParams(in models.User) sqlc.InsertUserParams {
	return sqlc.InsertUserParams{
		Subject:  in.Subject,
		Email:    in.Email,
		FullName: in.FullName,
	}
}

func tagsString(tags []string) string {
	if len(tags) == 0 {
		return ""
	}
	t := ""
	for _, tag := range tags {
		t += tag + ","
	}
	return t[:len(t)-1]
}
