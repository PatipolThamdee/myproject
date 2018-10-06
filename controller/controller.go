package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PatipolThamdee/myproject/common"
	"github.com/PatipolThamdee/myproject/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

func TestCtrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthenCtrl")

	// json := simplejson.New()
	// json.Set("foo", "bar")

	// payload, err := json.MarshalJSON()
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(payload)
	// param, err := common.GetParam(c)
	// if err != nil {
	// 	common.BadRequest(c, err)
	// 	return err
	// }
	// fmt.Println(`Here`)
	// res, err, code := api.TestApi()
	// if err != nil {
	// 	common.Response(c, map[string]interface{}{"Message": err.Error()}, code)
	// 	return err
	// }

	// common.Success(c, res)
	// return nil

	common.ReadJSON()
}

func TestCtrlPost(w http.ResponseWriter, r *http.Request) {
	// r.Header.Set("content-type", "application/json")
	// fmt.Println("AuthenCtrl")
	// param := r.FormValue("test")
	// fmt.Println(param)
	decoder := json.NewDecoder(r.Body)
	// b, err := ioutil.ReadAll(r.Body)
	var msg map[string]interface{}
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

	// fmt.Println(msg["test"])

	// var msg map[string]interface{}
	// err = json.Unmarshal(b, &msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// fmt.Println(msg["test"])
	// fmt.Println(msg["kuy"])

}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`checkToken`)
	decoded := context.Get(r, "decoded")
	var user common.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`getToken`)

	var user common.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	response, err := model.GetToken(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": user.Username,
	// 	"password": user.Password,
	// })
	// tokenString, error := token.SignedString([]byte("secret"))
	// if error != nil {
	// 	fmt.Println(error)
	// }
	json.NewEncoder(w).Encode(response)
}
