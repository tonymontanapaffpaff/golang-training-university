package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type IAuth interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type Service struct {
	client *redis.Client
}

func NewAuth(client *redis.Client) *Service {
	return &Service{client: client}
}

type AccessDetails struct {
	TokenUuid string
	UserId    string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

//Save token metadata to Redis
func (tk *Service) CreateAuth(userId string, td *TokenDetails) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(td.TokenUuid, userId, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(td.RefreshUuid, userId, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

//Check the metadata saved
func (tk *Service) FetchAuth(tokenUuid string) (string, error) {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	userid, err := tk.client.Get(tokenUuid).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

//Once a user row in the token table
func (tk *Service) DeleteTokens(authD *AccessDetails) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UserId)
	//delete access token
	deletedAt, err := tk.client.Del(authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *Service) DeleteRefresh(refreshUuid string) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//delete refresh token
	deleted, err := tk.client.Del(refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
