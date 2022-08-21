package externals_test

import (
	"bytes"
	"testing"

	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
)

func TestQueryBuilder(t *testing.T) {
	userRoles := []user.UserRole{
		user.UserRole_KUDAKI_TEAM,
		user.UserRole_ORGANIZER}

	query := "SELECT * FROM kudaki_user.users WHERE role IN"

	var argsBuffer bytes.Buffer

	argsBuffer.WriteString("(?,")
	for i := 1; i < len(userRoles); i++ {
		if i == len(userRoles)-1 {
			argsBuffer.WriteString("?)")
		} else {
			argsBuffer.WriteString("?,")
		}
	}

	query += argsBuffer.String()

	t.Log(query)
}
