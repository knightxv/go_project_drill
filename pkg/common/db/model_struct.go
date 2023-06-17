package db

import (
	"time"
)

type UserScore struct {
	UserId     string    `gorm:"column:user_id;index:idx_user_id;type:varchar(60);not null" json:"userId"` // 用户id
	Score      uint64    `gorm:"column:score;type:bigint(20);default:0" json:"score"`
	Status     int8      `gorm:"column:status;type:tinyint(4);default:0" json:"status"` // 0：正常 1：封禁
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:null" json:"createdTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:null" json:"updateTime"` // 更新时间
	Ex         string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (m *UserScore) TableName() string {
	return "user_scores"
}

type RewardEventLogs struct {
	Id         string    `gorm:"primaryKey;column:id;type:varchar(255);not null" json:"id"`                // id索引
	UserId     string    `gorm:"column:user_id;index:idx_user_id;type:varchar(60);not null" json:"userId"` // 用户id
	RewardType string    `gorm:"column:reward_type;type:varchar(50)" json:"rewardType"`                    // 奖励类型
	Reward     uint64    `gorm:"column:reward;type:bigint(20);default:0" json:"reward"`                    // 奖励数量
	CreateTime time.Time `gorm:"column:create_time;type:datetime;index;default:null" json:"createdTime"`
	Info       string    `gorm:"column:info;type:varchar(1024)" json:"info"`
	Ex         string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}

func (m *RewardEventLogs) TableName() string {
	return "reward_event_logs"
}

type WithdrawEventLogs struct {
	Id                int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20);not null;" json:"id"`   // id索引
	UserId            string    `gorm:"column:user_id;index:idx_user_id;type:varchar(60);not null" json:"userId"`      // 用户id
	Score             uint64    `gorm:"column:score;type:bigint(20);default:0" json:"score"`                           // 提现积分
	ScoreWhenWithdraw uint64    `gorm:"column:score_when_withdraw;type:bigint(20);default:0" json:"scoreWhenWithdraw"` // 提现时的积分
	Amount            uint64    `gorm:"column:amount;type:bigint(20);default:0" json:"amount"`                         // 提现金额
	Commission        uint64    `gorm:"column:commission;type:bigint(20);default:0" json:"commission"`                 // 手续费
	Rate              int64     `gorm:"column:rate;type:bigint(20);default:0" json:"rate"`                             // 积分汇率
	Decimal           int       `gorm:"column:decimal;type:int(11);default:0" json:"decimal"`                          // 小数位数
	Status            string    `gorm:"column:status;type:varchar(10);default:'UN_START';" json:"status"`              // UN_START：提现中 SUCCESS：提现成功 CONFIRMING：确认中 FAIL：提现失败 ROLLBACK：提现被回滚
	Remark            string    `gorm:"column:remark;type:varchar(1024)" json:"remark"`                                // 提现备注
	Coin              string    `gorm:"column:coin;type:varchar(20)" json:"coin"`                                      // 提现币种 idc usdt
	ChainId           int64     `gorm:"column:chain_id;type:int(11);default:0" json:"chainId"`                         // 链id
	TxHash            string    `gorm:"column:tx_hash;type:varchar(100);index;" json:"txHash"`                         // 交易hash
	CreateTime        time.Time `gorm:"column:create_time;type:datetime;index;default:null" json:"createdTime"`        // 提现时间
	Ex                string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`                                        // 扩展字段
}

func (m *WithdrawEventLogs) TableName() string {
	return "withdraw_event_logs"
}

// 存在的目的就是为了记录上链的信息顺便做分布式锁（可能不止一个提现需要用到，后续的上链都可以记录到这）
type ChainUpEventLogs struct {
	Id          string    `gorm:"primaryKey;column:id;type:varchar(255);not null" json:"id"`              // id索引
	TxHash      string    `gorm:"column:tx_hash;type:varchar(100);index;" json:"txHash"`                  // 交易hash
	ChainId     int64     `gorm:"column:chain_id;type:int(11);default:0" json:"chainId"`                  // 链id
	Endpoint    string    `gorm:"column:endpoint;type:varchar(255);not null" json:"endpoint"`             // 链节点
	Type        string    `gorm:"column:type;type:varchar(20)" json:"type"`                               // 类型
	Status      string    `gorm:"column:status;type:varchar(12);default:'UPING'" json:"status"`           // UPING：上链中 SUCCESS：上链成功 FAIL：上链失败 RECEIPT_FAIL：执行失败，交易被回滚
	Info        string    `gorm:"column:info;type:varchar(1024)" json:"info"`                             // 上链信息
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;index;default:null" json:"createdTime"` // 创建时间
	ChainUpTime time.Time `gorm:"column:chain_up_time;type:datetime;default:null" json:"chainUpTime"`     // 上链时间
	MinedTime   time.Time `gorm:"column:mined_time;type:datetime;default:null" json:"minedAt"`            // 确认时间
	Ex          string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`                                 // 扩展字段
}

func (m *ChainUpEventLogs) TableName() string {
	return "chain_up_event_logs"
}
