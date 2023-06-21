# 关注列表


首先版本1数据都在一个表中，
type UserFollow struct {
	ID            uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null" json:"id"`
	FromUserID    string    `gorm:"primaryKey;uniqueIndex:followindex;column:from_user_id;type:varchar(64);not null" json:"fromUserID"`
	FollowUserID  string    `gorm:"primaryKey;uniqueIndex:followindex;column:follow_user_id;type:varchar(64);not null" json:"followUserID"`
	HandleResult  int       `gorm:"column:handle_result;type:int(11);default:null" json:"handleResult"`
	Follow        int8      `gorm:"column:follow;type:tinyint(4);default:null" json:"follow"`
	Remark        string    `gorm:"column:remark;type:varchar(255);default:null" json:"remark"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime(3);default:null" json:"createTime"`
	HandlerUserID string    `gorm:"column:handler_user_id;type:varchar(64);default:null" json:"handlerUserId"`
	HandleMsg     string    `gorm:"column:handle_msg;type:varchar(255);default:null" json:"handleMsg"`
	HandleTime    time.Time `gorm:"column:handle_time;type:datetime(3);default:null" json:"handleTime"`
	Ex            string    `gorm:"column:ex;type:varchar(1024);default:null" json:"ex"`
}

一共包含两个接口，查看关注者和查看粉丝列表

// 然后编写api进行增删改查
// 版本1 /v1/user_follow
// rpc getFriendListByDbQuery
// 通过mysql数据库库进行关注列表和粉丝列表



// 版本2 /v2/friendlist /v2/addfriend(redis也设置一份)
当数据量比较大的时候，那么该如何进行分页。
首先按照 from_user_id 进行分页，这样的好处是方便查看用户的关注列表。但是查看粉丝列表的时候就会比较麻烦，因为需要对每个用户进行分页，然后再进行合并。
// 考虑到业务关系，因为查看用户的关注列表比查询粉丝列表要求比较实时，
// 还是按照 from_user_id 的方式进行分页，然后在数据库存储的时候，
// 再进行一次额外的存储操作，把数据存一份到 es 中，这样就可以实现实时查询了

// 版本3 /v3/friendlist /v1/addfriend（然后使用cdc的尝试捕获数据库的变动）
// rpc getFriendListByLocalCacheQuery

// 中间件 cdc kafka
// 压测工具
// vegeta