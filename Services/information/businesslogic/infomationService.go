package businesslogic

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Services/common/dao"
	"projects/Services/common/data"
	"projects/Services/common/net"
	"projects/Services/information/sqls"
)

func GetInfoService(w http.ResponseWriter, r *http.Request) {

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

	var infoReq data.InfoReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(req, &infoReq); err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", infoReq)

	sessionId := infoReq.SessionId

	//Get uvo.
	res, err := net.GetSessionByJson(sessionId)

	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	var uvo data.UserValueObject

	if err := json.Unmarshal(res, &uvo); err != nil {
		log.Println("error: ", err.Error())
		return
	}

	//Get connection from mysql
	conn := dao.GetConnectionMysql()

	//Get SQL
	sqlprop := sqls.GetProp()

	selectInfoListSQL := sqlprop.Sql.SelectInfoList

	rows, err := conn.Query(selectInfoListSQL, uvo.UserType)

	var entityInfos []data.EntityInfos

	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	for rows.Next() {
		entity := data.EntityInfos{}
		if err := rows.Scan(&entity.Information_id, &entity.Title, &entity.Contents, &entity.Issue_user, &entity.Update_date); err != nil {
			log.Println("error:", err.Error())
		}
		entityInfos = append(entityInfos, entity)
	}
	err = rows.Err()
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}

	//Close rows
	defer rows.Close()

	//Close Mysql connection
	defer dao.CloseConnetionMysql(conn)

	var infoRes data.InfoRes

	infoRes.Result = "00"
	infoRes.Code = "Success"
	infoRes.Informartions = entityInfos

	//Set content-type on response
	w.Header().Set("Content-Type", "application/json")

	//Create response data
	jr, err := json.Marshal(infoRes)

	log.Println("info: response data.>--", string(jr))

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(jr)
}

func EditInfoService(w http.ResponseWriter, r *http.Request) {

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

	var editInfoReq data.EditInfoReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(req, &editInfoReq); err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", editInfoReq)

	sessionId := editInfoReq.SessionId

	//Get uvo.
	res, err := net.GetSessionByJson(sessionId)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var uvo data.UserValueObject

	if err := json.Unmarshal(res, &uvo); err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Get connection from mysql
	conn := dao.GetConnectionMysql()

	defer dao.CloseConnetionMysql(conn)

	//Get SQL
	sqlprop := sqls.GetProp()

	if editInfoReq.EditFlg == "0" {

		// Insert
		sql := sqlprop.Sql.InsertInfo

		count, err := execAddInfo(sql, conn, editInfoReq, uvo)

		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("info: affected count by insert", count)

	} else if editInfoReq.EditFlg == "1" {
		// Update
		sql := sqlprop.Sql.UpdateInfo

		count, err := execEditInfo(sql, conn, editInfoReq, uvo)

		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("info: affected count by update", count)

	} else if editInfoReq.EditFlg == "2" {
		// Delete
		sql := sqlprop.Sql.DeleteInfo

		count, err := execDeleteInfo(sql, conn, editInfoReq, uvo)

		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("info: affected count by delete", count)

	} else {
		//Error
		log.Println("error: unknown edit flg.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var editInfoRes data.EditInfoRes

	editInfoRes.Code = "success"
	editInfoRes.Result = "00"

	//Set content-type on response
	w.Header().Set("Content-Type", "application/json")

	//Create response data
	jr, err := json.Marshal(editInfoRes)

	log.Println("info: response data.>--", string(jr))

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(jr)

}

func execEditInfo(sqlQuery string, conn *sql.DB, editInfoReq data.EditInfoReq, uvo data.UserValueObject) (int64, error) {
	data, err := dao.TransactAndReturnData(conn, func(tx *sql.Tx) (interface{}, error) {

		stmt, err := tx.Prepare(sqlQuery)
		if err != nil {
			return nil, err
		}
		result, err := stmt.Exec(editInfoReq.Title, editInfoReq.Contents, uvo.UserId, editInfoReq.InformationId)
		if err != nil {
			return nil, err
		}
		count, _ := result.RowsAffected()

		return count, err
	})
	if err != nil {
		return 0, err
	}
	return data.(int64), nil
}

func execAddInfo(sqlQuery string, conn *sql.DB, editInfoReq data.EditInfoReq, uvo data.UserValueObject) (int64, error) {
	data, err := dao.TransactAndReturnData(conn, func(tx *sql.Tx) (interface{}, error) {

		stmt, err := tx.Prepare(sqlQuery)
		if err != nil {
			return nil, err
		}
		result, err := stmt.Exec(editInfoReq.Title, editInfoReq.Contents, uvo.UserId, uvo.UserType)
		if err != nil {
			return nil, err
		}
		count, _ := result.RowsAffected()

		return count, err
	})
	if err != nil {
		return 0, err
	}
	return data.(int64), nil
}

func execDeleteInfo(sqlQuery string, conn *sql.DB, editInfoReq data.EditInfoReq, uvo data.UserValueObject) (int64, error) {
	data, err := dao.TransactAndReturnData(conn, func(tx *sql.Tx) (interface{}, error) {

		stmt, err := tx.Prepare(sqlQuery)
		if err != nil {
			return nil, err
		}
		result, err := stmt.Exec(editInfoReq.InformationId)
		if err != nil {
			return nil, err
		}
		count, _ := result.RowsAffected()

		return count, err
	})
	if err != nil {
		return 0, err
	}
	return data.(int64), nil
}
