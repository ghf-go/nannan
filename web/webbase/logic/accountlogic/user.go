package accountlogic

import "github.com/ghf-go/nannan/web/webbase/logic"

const (
	ROLE_USER = 1
	ROLE_ADMIN = 100

	tb_userid = "tb_user_id"
)

func newUserID(role int) int64  {
	return logic.GetTable(tb_userid).InsertMap(map[string]interface{}{
		"role" :role,
	})
}
func NewUserID() int64 {
	return newUserID(ROLE_USER)
}
func NewAdminID() int64  {
	return newUserID(ROLE_ADMIN)
}





