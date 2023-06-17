package mysql_model

import (
	"time"

	"github.com/idchats/user_score/pkg/common/db"
	"gorm.io/gorm"
)

func InsertRewardEventLog(rewardEventLog *db.RewardEventLogs) error {
	rewardEventLog.CreateTime = time.Now()
	return db.DB.MysqlDB.DefaultGormDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("reward_event_logs").Create(&rewardEventLog).Error; err != nil {
			return err
		}
		return CreateOrUpdateUserScoreByTx(&db.UserScore{
			UserId: rewardEventLog.UserId,
			Score:  rewardEventLog.Reward,
		}, tx)
	})
}

// 分页获取奖励日志
func GetRewardEventLogsByPage(userId, rewadType string, page, pageSize int) ([]*db.RewardEventLogs, error) {
	// 如果 rewadType 为空，则查询所有类型的奖励日志
	if rewadType == "" {
		rewadType = "%"
	}
	var rewardEventLogs []*db.RewardEventLogs
	err := db.DB.MysqlDB.DefaultGormDB().Table("reward_event_logs").
		Where("user_id=? and reward_type like ?", userId, rewadType).
		Order("create_time desc").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&rewardEventLogs).Error
	if err != nil {
		return nil, err
	}
	return rewardEventLogs, nil
}

func DelRewardEventLogByUserId(userId string) error {
	return db.DB.MysqlDB.DefaultGormDB().Where("user_id=?", userId).Delete(&db.RewardEventLogs{}).Error
}

func TestClearUserScore() error {
	DelRewardEventLogByUserId("0x1")
	DelRewardEventLogByUserId("0x2")
	db.DB.MysqlDB.DefaultGormDB().Where("user_id=?", "0x1").Delete(&db.UserScore{})
	db.DB.MysqlDB.DefaultGormDB().Where("user_id=?", "0x2").Delete(&db.UserScore{})
	return nil
}
