package mysql_model

import (
	"time"

	"github.com/idchats/user_score/pkg/common/constant"
	"github.com/idchats/user_score/pkg/common/db"
	"gorm.io/gorm"
)

func InsertChainUpEventLog(chainUpEventLog *db.ChainUpEventLogs) error {
	chainUpEventLog.CreateTime = time.Now()
	return db.DB.MysqlDB.DefaultGormDB().Create(&chainUpEventLog).Error
}

// 上链事件设置为成功（但是还没确认）
func UpdateChainUpEventSuccess(id, txHash string) error {
	return db.DB.MysqlDB.DefaultGormDB().
		Model(&db.ChainUpEventLogs{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"status":        constant.ChainUpStatusSuccess,
			"tx_hash":       txHash,
			"chain_up_time": time.Now(),
		}).Error
}

// 上链成功，锻造失败
func UpdateChainUpEventReceiptFail(txHash string) error {
	return db.DB.MysqlDB.DefaultGormDB().Transaction(func(tx *gorm.DB) error {
		// 更新上链记录状态
		res := tx.Model(&db.ChainUpEventLogs{}).
			Where("tx_hash = ? and status = ?", txHash, constant.ChainUpStatusSuccess).
			Updates(map[string]interface{}{
				"status":     constant.ChainUpStatusReceiptFail,
				"tx_hash":    txHash,
				"mined_time": time.Now(),
			})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return constant.ErrIdempotent
		}
		// 更新提现状态为失败
		return RollbackWithdrawEventByTx(txHash, tx)
	})
}

// 确认区块打包成功
func ConfirmChainUpEvnetMined(txHash string) error {
	return db.DB.MysqlDB.DefaultGormDB().Transaction(func(tx *gorm.DB) error {
		// 更新上链记录状态
		res := tx.Model(&db.ChainUpEventLogs{}).
			Where("tx_hash=? and status = ?", txHash, constant.ChainUpStatusSuccess).
			Updates(map[string]interface{}{
				"status":     constant.ChainUpStatusMinedSuccess,
				"mined_time": time.Now(),
			})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return constant.ErrIdempotent
		}
		// 更新提现状态为成功
		return UpdateWithdrawEventChainUpSuccessByTx(txHash, tx)
	})
}

// 上链事件设置为失败
func UpdateChainUpEventFail(id, info string) error {
	return db.DB.MysqlDB.DefaultGormDB().
		Model(&db.ChainUpEventLogs{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"status": constant.ChainUpStatusFail,
			"info":   info,
		}).Error
}

// 获取确认中的上链事件
func QueryMinedingChainUpEvent() ([]*db.ChainUpEventLogs, error) {
	var chainUpEventLogs []*db.ChainUpEventLogs
	err := db.DB.MysqlDB.DefaultGormDB().
		Where("status = ?", constant.ChainUpStatusSuccess).
		Find(&chainUpEventLogs).Error
	return chainUpEventLogs, err
}
