package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJSON() {
	jsonValue := map[string]interface{}{}
	// var jsonValue map[string]interface{}
	byteValue, _ := ioutil.ReadFile("././config/cfg.json")

	json.Unmarshal(byteValue, &jsonValue)
	jsonString := jsonValue[`JWT.TOKEN.SECRET`].(string)
	fmt.Println(jsonString)
	decoded, _ := base64.StdEncoding.DecodeString(jsonString)
	fmt.Println(string(decoded))
}
