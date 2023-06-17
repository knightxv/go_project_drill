# 好友列表的获取


// 然后编写api进行增删改查
// 版本1 /v1/friendlist
// rpc getFriendListByDbQuery
// 通过mysql数据库库进行好友列表的获取

// 版本2 /v2/friendlist /v2/addfriend(redis也设置一份)
// rpc getFriendListByRedisQuery

// 版本3 /v3/friendlist /v1/addfriend（然后使用cdc的尝试捕获数据库的变动）
// rpc getFriendListByLocalCacheQuery

// 中间件 cdc kafka
// 压测工具
// vegeta