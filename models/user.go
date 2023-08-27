package models

type User struct {
	Token RefreshToken
	Guid  string `bson:"guid, omitempty"`
}
