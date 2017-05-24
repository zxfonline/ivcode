package ivcode

import (
	"errors"

	. "github.com/zxfonline/bases"
)

const (
	BASEID = int64(1000000000)
)

//根据玩家唯一id获取玩家的唯一邀请码
func GetInviteCode(userid int64) (invitecode string) {
	invitecode = ToBase(int(userid+BASEID), 58)
	return
}

//根据邀请码获取的玩家的唯一id
func GetUserId(invitecode string) (int64, error) {
	userid := int64(FromBase(invitecode, 58)) - BASEID
	if userid < 0 {
		return -1, errors.New("invalid invite code")
	}
	return userid, nil
}
