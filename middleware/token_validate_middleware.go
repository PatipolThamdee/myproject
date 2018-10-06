package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type Exception struct {
	Message string `json:"message"`
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println(`middleware`)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		// decoder := json.NewDecoder(req.Body)
		// var msg map[string]interface{}
		// _ = decoder.Decode(&msg)
		// authorizationHeader := msg["authorization"].(string)
		fmt.Println(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			/* Bearer or any string and a space needed before actual token
			Ex. "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6ImdnNjkiLCJ1c2VybmFtZSI6Imp1bmUifQ.s8-IxLFFxLfgEzXmk1_6AJxydcrtmP4C8gQ_c5MS-2Q"
			if you wanna use only token just change to if bearerToen == 1 and bearerToken[0]
			*/
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					fmt.Println(`ooo`)

					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					fmt.Println(`token valid`)
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
		// fmt.Println(`pppppp`)

	})
}
