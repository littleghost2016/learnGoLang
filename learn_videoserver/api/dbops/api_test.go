package dbops

import (
	"testing"
)

func clearTables()  {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed. pwd : %s", pwd)
	}
}