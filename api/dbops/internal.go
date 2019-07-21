package dbops

import (
	"database/sql"
	"golang-stream/api/defs"
	"log"
	"sync"
)

func InserSession(sid string, ttl string, uname string) error {
	// 如果ttl 不是string 类型， 还需要转类型库
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id, TTL, login_name) values (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttl, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL, login_name from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	ss.TTL = ttl
	ss.Username = uname

	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Fatalf("%s/n", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttl string
		var login_name string

		if err := rows.Scan(&id, &ttl, &login_name); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}

		ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
		m.Store(id, ss)
		log.Printf("session id: %s, ttl: %s", id, ss.TTL)
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		log.Printf("%s /n", err)
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return nil
	}

	return nil
}
