/*
 * Created by Visual Studio Code.
 * User: HM
 * Date: 30/07/2018 * Time:14:00 */
package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/PatipolThamdee/myproject/router/routes"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	// "smartbofapp/v1/errorapis"
	// "smartbofapp/v1/router/response"
	// "github.com/valyala/fasthttp"
)

// func main() {
// 	var c config.Server
// 	var a auth.Server

// 	mux := http.NewServeMux()

// 	mux.Handle("/util/config", c)
// 	fmt.Println("Add route......")

// 	mux.Handle("/util/authentication", a)
// 	fmt.Println("Add route......")

// 	fmt.Println("Start Server .......")
// 	http.ListenAndServe(":9090", mux)
// }

func main() {
	r := mux.NewRouter()

	// r.Use(afterMiddlewareHandle)
	// r.Use(addRequestID)

	// app.Bootstarp(r)

	fmt.Println("Start :9090")
	route.InitialRoutes(r)

	//used for Google App Engine
	// http.Handle("/", r)

	//used for local
	log.Fatal(http.ListenAndServe(":9090", r))

	// r.Run(":4002")
	appengine.Main()
}

// display not found message and http code
/*
 @param routing.Context c
 @return error
*/
// func notFound(c *routing.Context) error {
// 	// response.NotFound(c)
// 	js, _ := json.Marshal(&map[string]interface{}{"Code": "404", "Message": "No data found."})
// 	c.Write(js)
// 	c.SetStatusCode(http.StatusNotFound)

// 	c.Next()
// 	return nil
// }

// generate uuid
/*
 @param routing.Context c
 @return error
*/
// func addRequestID(c *routing.Context) error {
// 	u := uuid.New()
// 	rqID := u.String()
// 	c.Set("requestID", rqID)
// 	c.Set("x-go-version", runtime.Version())
// 	c.Response.Header.Add("X-Request-ID", rqID)
// 	c.Next()
// 	return nil
// }

// begin route
/*
 @param routing.Context c
 @return error
*/
// func afterMiddlewareHandle(c *routing.Context) error {
// 	fmt.Println("Before Request")

// 	c.Response.Header.SetContentType("application/json")
// 	defer func(c *routing.Context) error {
// 		if rec := recover(); rec != nil {
// 			if _, ok := rec.(error); ok {
// 				fmt.Println("\n================================ Begin Error-Stack ================================")
// 				debug.PrintStack()
// 				fmt.Println("\n================================ End Error-Stack ================================")
// 			}
// 			errType := reflect.TypeOf(rec).String()
// 			fmt.Println("recrecrecrec", reflect.TypeOf(rec))
// 			if errType == "*stacktrace.stacktrace" {
// 				InternalServerError(c, fmt.Sprint(rec))
// 			} else if errType == "*errorapis.ErrorApi" {
// 				// errAPI := rec.(*errorapis.ErrorApi)
// 				// response.Response(c, &response.ErrorResponse{Code: errAPI.GetCode(), Message: errAPI.Error(), Errors: errAPI.GetTrace()}, errAPI.GetHTTPCode())
// 				InternalServerError(c, fmt.Sprint(rec))
// 			} else if errType == "*response.ErrorResponse" {
// 				// errAPI := rec.(*response.ErrorResponse)
// 				// response.Response(c, errAPI, errAPI.GetHTTPCode())
// 				InternalServerError(c, fmt.Sprint(rec))
// 			} else {
// 				InternalServerError(c, rec)
// 			}
// 			c.Abort()
// 		}
// 		return afterMiddleware(c)
// 	}(c)
// 	c.Next()
// 	return nil
// }

//after doing function
/*
 @param routing.Context c
 @return error
*/
// func afterMiddleware(c *routing.Context) error {
// 	fmt.Println("After Response")

// 	c.Response.Header.SetServer("iBath Backoffice APP")

// 	// fmt.Printf("response :=> %s\n", string(c.Response.Body()))
// 	// fmt.Printf("code :=> %d\n", c.Response.StatusCode())

// 	sha256 := sha256.New()
// 	sha256.Write(c.Response.Body())
// 	sha256.Write([]byte(c.Get("requestID").(string)))
// 	// add signature
// 	c.Response.Header.Add("X-Signature", fmt.Sprintf("%x", sha256.Sum(nil)))

// 	return nil
// }

// func InternalServerError(c *routing.Context, errors ...interface{}) {
// 	errRes := &map[string]interface{}{`Code`: `errorapis.SystemError`, `Message`: "Something went wrong."}
// 	// if len(errors) == 1 {
// 	// 	errRes.Errors = errors[0]
// 	// }
// 	// if len(errors) > 1 {
// 	// 	errRes.Errors = errors
// 	// }
// 	// errType := reflect.TypeOf(rec).String()

// 	js, _ := json.Marshal(errRes)
// 	c.Write(js)
// 	c.SetStatusCode(http.StatusInternalServerError)
// }
