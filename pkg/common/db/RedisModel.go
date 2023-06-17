package db

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

const (
	accountTempCode               = "ACCOUNT_TEMP_CODE"
	accountTempNonce              = "ACCOUNT_TEMP_NONCE"
	resetPwdTempCode              = "RESET_PWD_TEMP_CODE"
	userIncrSeq                   = "REDIS_USER_INCR_SEQ:" // user incr seq
	appleDeviceToken              = "DEVICE_TOKEN"
	userMinSeq                    = "REDIS_USER_MIN_SEQ:"
	uidPidToken                   = "UID_PID_TOKEN_STATUS:"
	conversationReceiveMessageOpt = "CON_RECV_MSG_OPT:"
	getuiToken                    = "GETUI_TOKEN"
	messageCache                  = "MESSAGE_CACHE:"
	SignalCache                   = "SIGNAL_CACHE:"
	SignalListCache               = "SIGNAL_LIST_CACHE:"
	GlobalMsgRecvOpt              = "GLOBAL_MSG_RECV_OPT"
	GloablAnnouncement            = "GLOBAL_ANNOUNCEMENT"
	FcmToken                      = "FCM_TOKEN:"
	groupUserMinSeq               = "GROUP_USER_MIN_SEQ:"
	groupMaxSeq                   = "GROUP_MAX_SEQ:"
	groupMinSeq                   = "GROUP_MIN_SEQ:"
	sendMsgFailedFlag             = "SEND_MSG_FAILED_FLAG:"
	userBadgeUnreadCountSum       = "USER_BADGE_UNREAD_COUNT_SUM:"
	ThirdStringUser               = "THIRD_PLATFORM_USER:"
	EmailVerifyCodePrefix         = "EmailVerifyCode"
)

func (d *DataBases) SetEmailVerifyCode(emailAddress string, code, ttl int) (err error) {
	key := EmailVerifyCodePrefix + "_" + emailAddress
	return d.RDB.Set(context.Background(), key, code, time.Duration(ttl)*time.Second).Err()
}
func (d *DataBases) GetEmailVerifyCode(emailAddress string) (string, error) {
	key := EmailVerifyCodePrefix + "_" + emailAddress
	return d.RDB.Get(context.Background(), key).Result()
}
func (d *DataBases) DeleteEmailVerifyCode(emailAddress string) (int64, error) {
	key := EmailVerifyCodePrefix + "_" + emailAddress
	return d.RDB.Del(context.Background(), key).Result()
}
func (d *DataBases) ExistVerifyCode(emailAddress string) bool {
	exitValue := d.RDB.Exists(context.Background(), EmailVerifyCodePrefix+"_"+emailAddress)
	fmt.Println(exitValue.Result())
	if nValue, err := exitValue.Result(); err == nil && nValue >= 1 {
		return true
	}
	return false
}

func (d *DataBases) SetUserGlobalAnnouncementRecvOpt(userID string, opt int32) error {
	key := conversationReceiveMessageOpt + userID
	return d.RDB.HSet(context.Background(), key, GloablAnnouncement, opt).Err()
}

// 检查是否存在限制
func (d *DataBases) WithdrawalRestriction(userAddress string) (canWithdraw bool, err error) {
	key := "userscore:WithdrawalRestriction:" + userAddress
	now := time.Now().Unix()
	ctx := context.Background()
	exist, err := d.RDB.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if exist > 0 {
		// 获取上次操作时间
		lastTime, err := d.RDB.HGet(ctx, key, "last_time").Int64()
		if err != nil {
			return false, err
		}
		// 检查是否在限制时间内
		if now-lastTime < 60 {
			return false, nil
		}
	} else {
		// 设置新的限制
		_, err = d.RDB.HMSet(ctx, key, map[string]interface{}{
			"last_time": strconv.FormatInt(now, 10),
		}).Result()
		if err != nil {
			return false, err
		}
		_, err = d.RDB.Expire(ctx, key, time.Minute).Result()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
