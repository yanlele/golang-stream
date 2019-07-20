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
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("yanle", "123")
	if err != nil {
		t.Errorf("Error of addUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("yanle")
	if pwd != "123" || err != nil {
		t.Errorf("Error of getUser")
	}
	log.Println("user info: ", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("yanle", "123")
	if err != nil {
		t.Errorf("Error of delete user: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("yanle")
	if err != nil {
		t.Error("Error of ReGet user", err)
	}

	log.Printf("password %v", pwd)

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
