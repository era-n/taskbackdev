package token

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/era-n/taskbackdev/config"
	"github.com/era-n/taskbackdev/models"
	"github.com/era-n/taskbackdev/models/response"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type JWTData struct {
	jwt.StandardClaims
	Guid string `json:"guid"`
}

var (
	client = config.InitMongoConn()
	cfg, _ = config.LoadConfig()
)

func NewPairOfToken(guid string) (r response.AuthResponse, err error) {
	exp := time.Now().Add(time.Duration(time.Minute * 30)).Unix()

	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
		Guid: guid,
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS512, claims) // SHA512

	token, err := tokenString.SignedString([]byte(cfg.Secret))

	refreshToken := uuid.New()

	refreshTokenBase64 := base64.StdEncoding.EncodeToString(refreshToken[:])

	err = saveToDB([]byte(refreshTokenBase64), guid)

	return response.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshTokenBase64,
	}, err
}

func ValidateRefreshToken(guid string, token string) error {
	user := models.User{}

	coll := client.Database("test").Collection("users")
	err := coll.FindOneAndDelete(context.TODO(), bson.D{{"guid", guid}}).Decode(&user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Token.Value), []byte(token))

	return err
}

func saveToDB(token []byte, guid string) error {
	refreshToken := models.RefreshToken{}
	user := models.User{
		Guid: guid,
	}

	hashed, _ := bcrypt.GenerateFromPassword(token, 10)
	refreshToken.Value = string(hashed)
	refreshToken.CreatedAt = time.Now()
	refreshToken.ExpiresAt = time.Now().Add(time.Duration(time.Hour * 24 * 7)) // неделя
	user.Token = refreshToken

	coll := client.Database("test").Collection("users")
	_, err := coll.InsertOne(context.TODO(), user)

	return err
}
