package dbops

func InserSession(sid string, ttl string, uname string) error {
	// 如果ttl 不是string 类型， 还需要转类型库
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id, TTL, login_name) values (?, ?, ?, ?)")
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

