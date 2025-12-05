package dao

import (
	"B00k/logger"
	"B00k/model"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var RContext = context.Background()

func RdbInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	ctx, cancel := context.WithTimeout(RContext, 5*time.Second)
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		logger.DBLog("RDB Connection failed! details: " + err.Error())
		panic("")
	}
	logger.DBLog("RDB Connected successfully")
}

func genSessionKey(token string) string {
	return "RDBKey:Session:" + token
}

func SaveSession(token string, session *model.UserSession, exp time.Duration) error {
	key := genSessionKey(token)

	data, err := json.Marshal(session)
	if err != nil {
		logger.DBLog("Occurred error during saving the session: " + err.Error())
		return err
	}

	return Rdb.Set(RContext, key, data, exp).Err()
}

func GetSession(token string) (*model.UserSession, error) {
	key := genSessionKey(token)

	instance, err := Rdb.Get(RContext, key).Result()
	if err == redis.Nil {
		logger.DBLog("Cannot get the cache")
		return nil, nil
	}
	if err != nil {
		logger.DBLog("Occurred error during saving the session: " + err.Error())
		return nil, err
	}

	session := &model.UserSession{}
	err = json.Unmarshal([]byte(instance), session)
	if err != nil {
		logger.DBLog("Occurred error during getting the session" + err.Error())
		return nil, err
	}

	return session, nil
}

func DeleteSession(token string) error {
	key := genSessionKey(token)

	return Rdb.Del(RContext, key).Err()
}
