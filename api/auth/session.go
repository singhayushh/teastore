package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

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

// A local utility function to generate unsigned hash string of @length bits
func getHash(length int) string {
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789" + "({[$.`!~-/&#%@]})" + os.Getenv("secret_key"))
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(characters[rand.Intn(len(characters))])
	}

	return b.String()
}

// CreateSession is for creating a new user session
func CreateSession(email string) (string, error) {
	conn := RedisConnect()
	defer conn.Close()

	HashedString := getHash(32)
	_, err := conn.Do("MSET", "session_id:"+HashedString, "email:"+email)

	if err != nil {
		return "", err
	}

	return HashedString, nil
}

// CheckSession tries to find the user in the redis database and returns an error if not found
func CheckSession(sessionID string) (string, error) {
	conn := RedisConnect()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("EXISTS", "session_id:"+sessionID))

	if err == nil && reply > 0 {
		return "", nil
	}

	return "Good", errors.New("Session not found")
}
