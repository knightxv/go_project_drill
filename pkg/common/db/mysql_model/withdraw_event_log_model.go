package mysql_model

import (
	"time"

	"github.com/idchats/user_score/pkg/common/constant"
	"github.com/idchats/user_score/pkg/common/db"
	"gorm.io/gorm"
)

func InsertWithdrawEventLog(withdrawEventLogs *db.WithdrawEventLogs) error {
	withdrawEventLogs.CreateTime = time.Now()
	return db.DB.MysqlDB.DefaultGormDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("withdraw_event_logs").Create(&withdrawEventLogs).Error; err != nil {
			return err
		}
		res := tx.Table("user_scores").
			Where("user_id=? and score >= ?", withdrawEventLogs.UserId, withdrawEventLogs.Score).
			Updates(map[string]interface{}{
				"score":       gorm.Expr("score - ?", withdrawEventLogs.Score),
				"update_time": time.Now(),
			})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
func GetWithdrawEventLogById(id int64) (*db.WithdrawEventLogs, error) {
	var withdrawEventLogs db.WithdrawEventLogs
	err := db.DB.MysqlDB.DefaultGormDB().Table("withdraw_event_logs").Where("id=?", id).First(&withdrawEventLogs).Error
	if err != nil {
		return nil, err
	}
	return &withdrawEventLogs, nil
}
func GetWithdrawEventLogByTxHash(txHash string) (*db.WithdrawEventLogs, error) {
	var withdrawEventLogs db.WithdrawEventLogs
	err := db.DB.MysqlDB.DefaultGormDB().Table("withdraw_event_logs").Where("tx_hash=?", txHash).First(&withdrawEventLogs).Error
	if err != nil {
		return nil, err
	}
	return &withdrawEventLogs, nil
}
func UpdateWithdrawEventChainUpFail(withdrawId int64, errMsg string) error {
	withdrawEventLogs, err := GetWithdrawEventLogById(withdrawId)
	if err != nil {
		return err
	}
	if withdrawEventLogs.Status != constant.WithdrawStatusUnStart {
		return constant.ErrIdempotent
	}
	// 提现失败，自动返还积分
	return db.DB.MysqlDB.DefaultGormDB().Transaction(func(tx *gorm.DB) error {
		// 更新提现状态为失败
		res := tx.Table("withdraw_event_logs").
			Where("id=? and status = ?", withdrawId, constant.WithdrawStatusUnStart).
			Update("status", constant.WithdrawStatusFail)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		res = tx.Table("user_scores").
			Where("user_id=?", withdrawEventLogs.UserId).
			Updates(map[string]interface{}{
				"score":       gorm.Expr("score + ?", withdrawEventLogs.Score),
				"update_time": time.Now(),
				"ex":          errMsg,
			})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
func RollbackWithdrawEventByTx(txHash string, tx *gorm.DB) error {
	// 更新提现状态为失败
	res := tx.Table("withdraw_event_logs").
		Where("tx_hash = ? and status = ?", txHash, constant.WithdrawStatusConfirming).
		Update("status", constant.WithdrawStatusRollback)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	// 返还积分
	withdrawEventLogs, err := GetWithdrawEventLogByTxHash(txHash)
	if err != nil {
		return err
	}
	return tx.Table("user_scores").
		Where("user_id=?", withdrawEventLogs.UserId).
		Updates(map[string]interface{}{
			"score":       gorm.Expr("score + ?", withdrawEventLogs.Score),
			"update_time": time.Now(),
		}).Error
}

// 提现成功并更新上链交易txHash
func UpdateWithdrawEventChainUpSuccessByTx(txHash string, tx *gorm.DB) error {
	return tx.
		Model(db.WithdrawEventLogs{}).
		Where("tx_hash=? and status = ?", txHash, constant.WithdrawStatusConfirming).
		Updates(map[string]interface{}{
			"status": constant.WithdrawStatusSuccess,
		}).Error
}

// 提现上链中 txHash，等待确认
func UpdateWithdrawEventChainUpConfirming(withdrawId int64, txHash string) error {
	return db.DB.MysqlDB.DefaultGormDB().Table("withdraw_event_logs").
		Where("id= ? and status = ?", withdrawId, constant.WithdrawStatusUnStart).
		Updates(map[string]interface{}{
			"status":  constant.WithdrawStatusConfirming,
			"tx_hash": txHash,
		}).Error
}

func GetWithdrawEventLogsByPage(userId string, page, pageSize int) ([]*db.WithdrawEventLogs, error) {
	var withdrawEventLogs []*db.WithdrawEventLogs
	err := db.DB.MysqlDB.DefaultGormDB().Table("withdraw_event_logs").
		Where("user_id=?", userId).
		Order("create_time desc").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&withdrawEventLogs).Error
	if err != nil {
		return nil, err
	}
	return withdrawEventLogs, nil
}
