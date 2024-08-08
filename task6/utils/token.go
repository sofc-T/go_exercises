package Utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


type SignedDetails struct{
	Email 	string
	Name 	string
	Uid 	string
	Role	string
	jwt.StandardClaims 
}

var jwt_secret_key string

func ImportJWTSecretKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil{
		return "", errors.New("secret key not defined in env")
	}
	jwt_secret_key = os.Getenv("jwt_secret_key")
	if jwt_secret_key == "" {
		return "",  errors.New("secret key not defined in env")
	}

	return jwt_secret_key, nil
}

func GenerateTokens(email string, name string, UID string, role string ) (*string , *string, error){

	if jwt_secret_key == ""{
		ImportJWTSecretKey()
	}

	claims := &SignedDetails{
		email, name, UID, role, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		email, name, UID, role, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwt_secret_key))
	if err != nil{
		return nil, nil, errors.New("couldn't Generate Token")
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(jwt_secret_key))

	if err != nil{
		return nil, nil, errors.New("couldn't Generate Refresh Token")
	}


	return &token, &refreshToken, nil 
}

func ValidateToken(signedToken string) (*SignedDetails , error){
	if jwt_secret_key == ""{
		ImportJWTSecretKey()
	}

	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func (token *jwt.Token) (interface{}, error ){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
			return []byte(jwt_secret_key), nil
		},
	)

	if err != nil{
		return nil, errors.New("wrong Credentails")
	}

	claims, ok := token.Claims.(*SignedDetails)
	
	if !ok {
		return nil, errors.New("wrong Credentails")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("expired Token")
	}


	return claims, nil

} 


