package conditions

func Auth(username string, password string) bool {
	if username == "bilal" && password == "123" {
		return true
	}
	return false
}
