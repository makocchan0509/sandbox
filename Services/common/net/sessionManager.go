package net

import (
	"log"
	"projects/Services/common/dao"

	"github.com/google/uuid"
)

func CreateSessionId() (sessionId string, err error) {
	//Generate UUID.
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	sessionId = u.String()
	return sessionId, err
}

func StartSession(sessionId string, sessionValue string) error {
	//Get connection Redis
	rconn, err := dao.GetConnectionRedis()

	//Set key
	dao.SetRedis(sessionId, sessionValue, rconn)

	expireSec := 180
	//Set expire to key
	dao.ExpireRedis(sessionId, expireSec, rconn)
	log.Println("info: Create key into CVS.", sessionId)

	//Close Redis connection.
	defer dao.CloseConnectionRedis(rconn)
	return err
}
