package model

import (
	"fmt"
)

func TestApi() (interface{}, error, int) {
	fmt.Println(`LOLOL`)
	return map[string]interface{}{"test": "ok"}, nil, 69
}
