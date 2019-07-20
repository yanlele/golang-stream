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

// video api 测试
var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	ClearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of add new video: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of get video info: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of delete video info: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of re get video info %v", err)
	}
}
