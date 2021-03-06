package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/Services/common/data"
	util "projects/Services/common/net"
	"projects/Services/routes/properties"

	"github.com/gin-gonic/gin"
)

//Handling OPTIONS method
func Options(ctx *gin.Context) {

	log.Println("info: Called OPTONS method.")
	net.CORSForGin(ctx)
	ctx.Status(200)
}

//Login service
func Login(ctx *gin.Context) {

	util.CORSForGin(ctx)

	var loginInfo data.LoginReq

	ctx.BindJSON(&loginInfo)

	log.Println("info: Received parameter", loginInfo)

	input, err := json.Marshal(loginInfo)

	log.Println("info: json.marshal", loginInfo)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	prop := properties.GetProp()
	//TODO parameter
	url := prop.Service.LoginUrl

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

//Get information lists
func GetInfoLists(ctx *gin.Context) {

	util.CORSForGin(ctx)

	var infoReq data.InfoReq

	ctx.BindJSON(&infoReq)

	log.Println("info: Received parameter", infoReq)

	input, err := json.Marshal(infoReq)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	prop := properties.GetProp()

	url := prop.Service.InfoUrl

	log.Println("info: Redirect URL --->", url)

	res, err := net.JsonPostRequestSender(url, input)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	var infoRes data.InfoRes

	if err := json.Unmarshal(res, &infoRes); err != nil {
		log.Println("error: ", err.Error())
	}

	log.Println("info: Parsed response data -->", infoRes)
	ctx.JSON(http.StatusOK, infoRes)
}

func CheckSession(ctx *gin.Context) {
	util.CORSForGin(ctx)

	var checkSessionReq data.CheckSessionReq

	ctx.BindJSON(&checkSessionReq)

	log.Println("info: Received parameter", checkSessionReq)

	input, err := json.Marshal(checkSessionReq)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	prop := properties.GetProp()

	url := prop.Service.CheckSessionUrl

	log.Println("info: Redirect URL --->", url)

	res, err := net.JsonPostRequestSender(url, input)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	var checkSessionRes data.CheckSessionRes

	if err := json.Unmarshal(res, &checkSessionRes); err != nil {
		log.Println("error: ", err.Error())
	}

	log.Println("info: Parsed response data -->", checkSessionRes)
	ctx.JSON(http.StatusOK, checkSessionRes)
}

func EditInfo(ctx *gin.Context) {

	util.CORSForGin(ctx)

	var editInfoReq data.EditInfoReq

	ctx.BindJSON(&editInfoReq)

	log.Println("info: Received parameter", editInfoReq)

	input, err := json.Marshal(editInfoReq)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	prop := properties.GetProp()

	url := prop.Service.EditInfoUrl

	log.Println("info: Redirect URL --->", url)

	res, err := net.JsonPostRequestSender(url, input)

	if err != nil {
		log.Println("error: ", err.Error())
		ctx.Status(404)
	}

	var editInfoRes data.EditInfoRes

	if err := json.Unmarshal(res, &editInfoRes); err != nil {
		log.Println("error: ", err.Error())
	}

	log.Println("info: Parsed response data -->", editInfoRes)
	ctx.JSON(http.StatusOK, editInfoRes)
}
