package businesslogic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Services/common/dao"
	"projects/Services/common/data"
	"projects/Services/common/net"
)

func LoginService(w http.ResponseWriter, r *http.Request) {

	//Get request parameter
	if r.Method != "POST" {
		log.Println("error: Request is not POST method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("error: Request is not JSON format.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var loginInfo data.LoginReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(req, &loginInfo); err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", loginInfo)

	var userInfo data.Users

	userInfo.Id = loginInfo.LoginId
	userInfo.Password = loginInfo.Password

	//Get connection from mysql
	conn := dao.GetConnectionMysql()

	err = conn.QueryRow("select user_type from users where id = ? and password = ?", userInfo.Id, userInfo.Password).Scan(&userInfo.User_type)

	var loginResult data.LoginRes

	//Not found user.
	if err != nil {
		log.Println("error: Not found user.")
		log.Println("error: ", err.Error())

		loginResult.Result = "80"
		loginResult.Code = "NotFoundUser"
		loginResult.ReqId = loginInfo.LoginId
		loginResult.ReqPass = loginInfo.Password

	} else {

		sessionId, err := net.CreateSessionId()

		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}

		userValue := userInfo.Id
		userValue += ","
		userValue += userInfo.User_type
		log.Println("info: userValue>--", userValue)

		err = net.StartSession(sessionId, userValue)

		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}

		loginResult.Result = "00"
		loginResult.Code = "LoginSuccess"
		loginResult.SessionId = sessionId
		loginResult.UserType = userInfo.User_type
		loginResult.ReqId = loginInfo.LoginId
		loginResult.ReqPass = loginInfo.Password
	}
	//Close Mysql connection
	defer dao.CloseConnetionMysql(conn)

	//Set content-type on response
	w.Header().Set("Content-Type", "application/json")

	//Create response data
	jr, err := json.Marshal(loginResult)

	log.Println("info: response data.>--", string(jr))

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(jr)
}
