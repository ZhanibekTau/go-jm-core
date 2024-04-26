package jRedisHelper

import (
	redis2 "github.com/ZhanibekTau/go-jm-core/pkg/helpers/db/redis"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jLog"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jStructures"
	"github.com/gomodule/redigo/redis"
	"strings"
)

func GetString(pool *redis.Pool, requestData *jStructures.RequestData, key string) (string, error) {
	messageBuilder := strings.Builder{}
	result, err := redis2.GetString(pool, key)

	if err != nil {
		messageBuilder.WriteString(" Error:" + err.Error())
		jLog.PrintErrorJLog(requestData.ServiceName, requestData.RequestMethod, requestData.RequestHost+requestData.RequestUrl, 0, requestData.RequestId, messageBuilder.String())
	}

	return result, err
}

func SetString(pool *redis.Pool, requestData *jStructures.RequestData, key string, data interface{}, ttl int) (interface{}, error) {
	messageBuilder := strings.Builder{}
	reply, err := redis2.SetString(pool, key, data, ttl)

	if err != nil {
		messageBuilder.WriteString(" Error:" + err.Error())
		jLog.PrintErrorJLog(requestData.ServiceName, requestData.RequestMethod, requestData.RequestHost+requestData.RequestUrl, 0, requestData.RequestId, messageBuilder.String())
	}

	return reply, err
}
