package dbops

import (
	"testing"
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)

func ClearTables() {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {

}

func TestUserWorkFlow(t *testing.T) {

}

func TestAddUser(t *testing.T) {

}

func TestGerUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}

func TestReGetUser(t *testing.T)  {

}
