package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang-stream/api/dbops"
	"golang-stream/api/defs"
	"golang-stream/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.UserName)
	su := &defs.SingedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		SendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user_name")
	_, _ = io.WriteString(w, username)
}
