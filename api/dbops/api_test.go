package dbops

import (
	"log"
	"testing"
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)

func ClearTables() {
	_, _ = dbConn.Exec("truncate user")
	_, _ = dbConn.Exec("truncate video_info")
	_, _ = dbConn.Exec("truncate comments")
	_, _ = dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	ClearTables()
	m.Run()
	ClearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("Get", TestGetUser)
	t.Run("Del", TestDeleteUser)
	t.Run("ReGet", TestReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("yanle", "123")
	if err != nil {
		t.Errorf("Error of addUser: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	pwd, err := GetUserCredential("yanle")
	if pwd != "123" || err != nil {
		t.Errorf("Error of getUser")
	}
	log.Println("user info: ", pwd)
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("yanle", "123")
	if err != nil {
		t.Errorf("Error of delete user: %v", err)
	}
}

func TestReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("yanle")
	if err != nil {
		t.Error("Error of ReGet user", err)
	}

	log.Printf("password %v", pwd)

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
