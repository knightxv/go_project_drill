package mysql_model

import (
	"errors"
	"time"

	"github.com/idchats/user_score/pkg/common/db"

	"gorm.io/gorm"
)

var (
	ErrChatFinishChatWithSameOne = errors.New("not chat user with other one")
	ErrTodayIsFinished           = errors.New("today has finished")
	ErrTimeHasNotBeenReached     = errors.New("time has not been reached")
)

func CreateOrUpdateUserScoreByTx(userScore *db.UserScore, tx *gorm.DB) error {
	userScore.CreateTime = time.Now()
	// 没有就创建,有就增加分数
	return tx.Table("user_scores").
		Where("user_id=?", userScore.UserId).
		// Assign("score", gorm.Expr("score + ?", userScore.Score)).
		// Assign("update_time", time.Now()).
		Assign(map[string]interface{}{
			"update_time": time.Now(),
			"score":       gorm.Expr("score + ?", userScore.Score),
		}).
		FirstOrCreate(&userScore).Error
}

func GetUserScore(userId string) (*db.UserScore, error) {
	var userScore db.UserScore
	err := db.DB.MysqlDB.DefaultGormDB().Table("user_scores").Where("user_id=?", userId).First(&userScore).Error
	if err != nil {
		return nil, err
	}
	return &userScore, nil
}
