package auth

import (
	"errors"
	"fmt"
	"teastore/api/utils"

	"github.com/gomodule/redigo/redis"
)

// RedisConnect ...
// Reusable function to establish connection with redis
func RedisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Error connecting to Redis")
	}
	return conn
}

// CreateSession is for creating a new user session
func CreateSession(uid uint64, utype string) (string, error) {
	conn := RedisConnect()
	defer conn.Close()

	HashedString := utils.GenerateHash(32)
	_, err := conn.Do("HMSET", "session_id:"+HashedString, "uid", uid, "type", utype)

	if err != nil {
		return "", err
	}

	return HashedString, nil
}

// CheckSession tries to find the user in the redis database and returns an error if not found
func CheckSession(sessionID string) (uint64, string, error) {
	conn := RedisConnect()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("EXISTS", "session_id:"+sessionID))

	if err == nil && reply > 0 {
		uid, err := redis.Uint64(conn.Do("HGET", "session_id:"+sessionID, "uid"))
		if err != nil {
			return 0, "", err
		}
		utype, err := redis.String(conn.Do("HGET", "session_id:"+sessionID, "type"))
		if err != nil {
			return 0, "", err
		}
		return uid, utype, nil
	}

	return 0, "", errors.New("Session not found")
}
