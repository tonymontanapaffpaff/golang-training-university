package data

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data/auth"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserData struct {
	Collection *mongo.Collection
	IAuth      auth.IAuth
	IToken     auth.IToken
}

func NewUserData(collection *mongo.Collection, IAuth auth.IAuth, IToken auth.IToken) *UserData {
	return &UserData{Collection: collection, IAuth: IAuth, IToken: IToken}
}

func (d *UserData) Login(user User) (map[string]string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"username", user.Username}}
	var result User

	err := d.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("can't get user with given ID, error: %w", err)
	}

	// compare the user from the request, with the one we defined:
	if user.Password != result.Password {
		return nil, fmt.Errorf("please provide valid login details")
	}

	ts, err := d.IToken.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	saveErr := d.IAuth.CreateAuth(user.ID, ts)
	if saveErr != nil {
		return nil, saveErr
	}
	return map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}, nil
}

func (d *UserData) Logout(req *http.Request) error {
	// if metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := d.IToken.ExtractTokenMetadata(req)
	if metadata != nil {
		return d.IAuth.DeleteTokens(metadata)
	}
	return nil
}

func (d *UserData) Refresh(refreshToken string) (map[string]string, HttpErr) {
	// verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})
	// if there is an error, the token must have expired
	if err != nil {
		return nil, HttpErr{
			Err:        fmt.Errorf("refresh token expired"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	// check is token valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, HttpErr{
			StatusCode: http.StatusUnauthorized,
		}
	}
	// since token is valid, get the uuid
	claims, ok := token.Claims.(jwt.MapClaims) // the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) // convert the interface to string
		if !ok {
			return nil, HttpErr{
				Err:        fmt.Errorf("unauthorized"),
				StatusCode: http.StatusUnprocessableEntity,
			}
		}
		userId, roleOk := claims["user_id"].(string)
		if roleOk == false {
			return nil, HttpErr{
				Err:        fmt.Errorf("unauthorized"),
				StatusCode: http.StatusUnprocessableEntity,
			}
		}
		// delete the previous Refresh Token
		delErr := d.IAuth.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			return nil, HttpErr{
				Err:        fmt.Errorf("unauthorized"),
				StatusCode: http.StatusUnauthorized,
			}
		}
		// create new pairs of refresh and access tokens
		ts, createErr := d.IToken.CreateToken(userId)
		if createErr != nil {
			return nil, HttpErr{
				Err:        createErr,
				StatusCode: http.StatusForbidden,
			}
		}
		// save the tokens metadata to redis
		saveErr := d.IAuth.CreateAuth(userId, ts)
		if saveErr != nil {
			return nil, HttpErr{
				Err:        saveErr,
				StatusCode: http.StatusForbidden,
			}
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		return tokens, HttpErr{}
	} else {
		return nil, HttpErr{
			Err:        fmt.Errorf("refresh expired"),
			StatusCode: 0,
		}
	}
}
