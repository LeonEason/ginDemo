package dao

var dataBase = map[string]string{
	"vita": "123",
	"tipo": "456",
}

func AddUser(username, password string) {
	dataBase[username] = password
}
func SelectUser(username string) bool {
	if dataBase[username] == "" {
		return false
	} else {
		return true
	}
}
func GetPassword(username string) string {
	return dataBase[username]
}
