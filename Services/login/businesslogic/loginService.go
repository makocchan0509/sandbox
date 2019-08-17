package businesslogic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Services/common/dao"
	"projects/Services/common/data"
)

func LoginService(w http.ResponseWriter, r *http.Request) {

	//Get request parameter
	if r.Method != "POST" {
		log.Println("error: Request is not POST method")
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("error: Request is not JSON format.")
		w.WriteHeader(http.StatusBadRequest)
	}

	var loginInfo data.LoginReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := json.Unmarshal(req, &loginInfo); err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", loginInfo)

	var loginResult data.LoginRes

	//Get connection from mysql
	conn := dao.GetConnectionMysql()

	var name string
	err = conn.QueryRow("select name from users where id = ? and password = ?", loginInfo.LoginId, loginInfo.Password).Scan(&name)

	//Not found user.
	if err != nil {
		log.Println("error: Not found user.")
		log.Println("error: ", err.Error())

		loginResult.Result = "80"
		loginResult.Code = "NotFoundUser"
		loginResult.ReqId = loginInfo.LoginId
		loginResult.ReqPass = loginInfo.Password

	} else {

		//Get connection Redis
		rconn := dao.GetConnectionRedis()

		//Check exists key(login key)
		i := dao.ExistsRedis(loginInfo.LoginId, rconn)

		if i == 1 {
			log.Println("info: Exists key", loginInfo.LoginId)
			log.Println("info: Skip set key")
		} else {
			//Set key
			sessionValue := loginInfo.Password
			dao.SetRedis(loginInfo.LoginId, sessionValue, rconn)

			//Set expire time to key
			dao.ExpireRedis(loginInfo.LoginId, 120, rconn)
			log.Println("info: Create key into CVS.", loginInfo.LoginId)
		}
		//Close Redis connection.
		defer dao.CloseConnectionRedis(rconn)

		loginResult.Result = "00"
		loginResult.Code = "LoginSuccess"
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
	}

	w.Write(jr)
}
