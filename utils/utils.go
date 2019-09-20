package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/letanthang/my_framework/config"
	"github.com/letanthang/my_framework/db/types"
)

func GetMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5(text string) string {
	ba := md5.Sum([]byte(text))
	bs := ba[:]
	return string(bs)
}

func GetToken(id int, userType, name, phone, email string) (string, error) {
	claims := types.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		ClientID:    id,
		ClientType:  userType,
		ClientName:  name,
		ClientPhone: phone,
		ClientEmail: email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Config.Encryption.JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
