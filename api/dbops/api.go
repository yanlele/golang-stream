package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into user (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	_, _ = stmtIns.Exec(loginName, pwd)
	_ = stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from user where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string

	// 只查询一行
	_ = stmtOut.QueryRow(loginName).Scan(&pwd)
	_ = stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from user where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
	}
	_, _ = stmtDel.Exec(loginName, pwd)
	_ = stmtDel.Close()
	return nil
}
