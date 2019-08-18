package viewter

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"projects/Services/common/data"
	"projects/Services/common/net"
	"projects/Services/viewer/properties"

	"github.com/gin-gonic/gin"
)

//Display index page.
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func LoginDS(ctx *gin.Context) {

	//	id := ctx.PostForm("loginId")
	//	password := ctx.PostForm("password")
	requestValue := url.Values{}
	requestValue.Set("loginId", ctx.PostForm("loginId"))
	requestValue.Set("password", ctx.PostForm("password"))

	prop := properties.GetProp()
	gatewayUrl := prop.Login.GatewayUrl

	res, err := net.QueryPostRequestSender(gatewayUrl, requestValue)

	if err != nil {
		ctx.Status(404)
	}

	var loginRes data.LoginRes

	if err := json.Unmarshal(res, &loginRes); err != nil {
		log.Println("error: ", err.Error())
	}

	log.Println("info: Parsed response data -->", loginRes)

	ctx.JSON(http.StatusOK, loginRes)

}
