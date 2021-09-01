package middleware

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"nutrition/infrastructure"
	"nutrition/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

// Authentication interface have some method interact with token jwt
type Authentication interface {
	GetClaimsData(tokenString string) (*model.User, error)
}

type authenication struct {
	tokenAuthDecode *jwtauth.JWTAuth
	publicKey       interface{}
}

func (ac *authenication) GetClaimsData(tokenString string) (*model.User, error) {
	user := model.User{}
	words := strings.Fields(tokenString)
	if len(words) == 1 {
		token, err := jwt.ParseWithClaims(words[0], &user, func(token *jwt.Token) (interface{}, error) {
			return ac.publicKey, nil
		})
		if err != nil {
			log.Printf("Have problem at get claims data: %v\n", err)
			return nil, err
		}
		if !token.Valid {
			log.Printf("Have problem at get claims data: %v\n", err)
			return nil, err
		}
	} else {
		token, err := jwt.ParseWithClaims(words[1], &user, func(token *jwt.Token) (interface{}, error) {
			return ac.publicKey, nil
		})
		if err != nil {
			log.Printf("Have problem at get claims data: %v\n", err)
			return nil, err
		}
		if !token.Valid {
			log.Printf("Have problem at get claims data: %v\n", err)
			return nil, err
		}
	}

	return &user, nil
}

// NewAuthentication export middleware authentication
func NewAuthentication() (Authentication, error) {
	publicByte, err := ioutil.ReadFile(infrastructure.JwtPublicKeyPath)
	if err != nil {
		log.Printf("Has problem at create authentication, parse private key: %v", err)
		return nil, err
	}
	PublicKeyRS256String := string(publicByte)
	publicKeyBlock, _ := pem.Decode([]byte(PublicKeyRS256String))
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Printf("Has problem at create authentication, parse public key: %v", err)
		return nil, err
	}

	return &authenication{
		tokenAuthDecode: jwtauth.New("RS256", publicKey, nil),
		publicKey:       publicKey,
	}, nil

}

func ParseUserDataFromToken(w http.ResponseWriter, r *http.Request) *model.User {
	tokenString := r.Header.Get("Authorization")
	// not exist token
	if tokenString == "" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return nil
	}
	authentication, _ := NewAuthentication()
	user, err := authentication.GetClaimsData(tokenString)
	//token fail
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return nil
	}
	return user
}

func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			// not exist token
			if tokenString == "" {
				w.WriteHeader(http.StatusForbidden)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			authentication, _ := NewAuthentication()
			_, err := authentication.GetClaimsData(tokenString)
			//token fail
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			//auth success
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
