package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/Services/common/data"
	"projects/Services/common/net"
	"projects/Services/routes/properties"

	"github.com/gin-gonic/gin"
)

//Display index page.
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

//Login service.
func Login(ctx *gin.Context) {
	id := ctx.PostForm("loginId")
	password := ctx.PostForm("password")

	//Create JSON format
	var loginInfo data.LoginReq
	loginInfo.LoginId = id
	loginInfo.Password = password

	input, err := json.Marshal(loginInfo)

	log.Println("info: Received parameter", loginInfo)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	prop := properties.GetProp()
	//TODO parameter
	url := prop.Login.Url

	log.Println("info: Redirect URL --->", url)

	//Request http POST
	res, err := net.JsonPostRequestSender(url, input)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	var loginRes data.LoginRes

	if err := json.Unmarshal(res, &loginRes); err != nil {
		log.Println("error: ", err.Error())
	}

	log.Println("info: Parsed response data -->", loginRes)

	ctx.JSON(http.StatusOK, loginRes)

}

//Login and display screen service.
func LoginDS(ctx *gin.Context) {

}
