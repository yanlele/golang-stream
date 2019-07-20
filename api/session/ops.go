package session

import (
	"time"
	"sync"
	"golang-stream/api/defs"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {

}

func GenerateNewSessionId(un string) string {

}

func IsSessionExpired(sid string) (string, bool) {

}
