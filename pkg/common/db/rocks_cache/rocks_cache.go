package rocksCache

import (
	"context"
	"fmt"

	"github.com/idchats/user_score/pkg/common/db"
	"github.com/idchats/user_score/pkg/common/log"
	"github.com/idchats/user_score/pkg/utils"
)

const (
	userInfoCache             = "USER_INFO_CACHE:"
	friendRelationCache       = "FRIEND_RELATION_CACHE:"
	FollowFriendCache         = "Follow_Friend_CACHE:" //用户相互关注的列表
	blackListCache            = "BLACK_LIST_CACHE:"
	groupCache                = "GROUP_CACHE:"
	groupInfoCache            = "GROUP_INFO_CACHE:"
	groupOwnerIDCache         = "GROUP_OWNER_ID:"
	joinedGroupListCache      = "JOINED_GROUP_LIST_CACHE:"
	groupMemberInfoCache      = "GROUP_MEMBER_INFO_CACHE:"
	groupAllMemberInfoCache   = "GROUP_ALL_MEMBER_INFO_CACHE:"
	allFriendInfoCache        = "ALL_FRIEND_INFO_CACHE:"
	allDepartmentCache        = "ALL_DEPARTMENT_CACHE:"
	allDepartmentMemberCache  = "ALL_DEPARTMENT_MEMBER_CACHE:"
	joinedSuperGroupListCache = "JOINED_SUPER_GROUP_LIST_CACHE:"
	groupMemberListHashCache  = "GROUP_MEMBER_LIST_HASH_CACHE:"
	groupMemberNumCache       = "GROUP_MEMBER_NUM_CACHE:"
	conversationCache         = "CONVERSATION_CACHE:"
	conversationIDListCache   = "CONVERSATION_ID_LIST_CACHE:"
	userEmailInfo             = "UserEmail:"
	userSpaceGroupInfo        = "USER_GROUP_INFO:"
)

func DelKeys() {
	fmt.Println("init to del old keys")
	for _, key := range []string{groupCache, friendRelationCache, blackListCache, userInfoCache, groupInfoCache, groupOwnerIDCache, joinedGroupListCache,
		groupMemberInfoCache, groupAllMemberInfoCache, allFriendInfoCache} {
		fName := utils.GetSelfFuncName()
		var cursor uint64
		var n int
		for {
			var keys []string
			var err error
			keys, cursor, err = db.DB.RDB.Scan(context.Background(), cursor, key+"*", 3000).Result()
			if err != nil {
				panic(err.Error())
			}
			n += len(keys)
			// for each for redis cluster
			for _, key := range keys {
				if err = db.DB.RDB.Del(context.Background(), key).Err(); err != nil {
					log.NewError("", fName, key, err.Error())
					err = db.DB.RDB.Del(context.Background(), key).Err()
					if err != nil {
						panic(err.Error())
					}
				}
			}
			if cursor == 0 {
				break
			}
		}
	}
}

// func GetSpaceInfoByUser(userid string) (*db.Group, error) {
// 	getGroupInfo := func() (string, error) {
// 		groupInfo, err := imdb.GetOneGroupInfoByUserID(userid)
// 		if err != nil {
// 			return "", utils.Wrap(err, "")
// 		}
// 		bytes, err := json.Marshal(groupInfo)
// 		if err != nil {
// 			return "", utils.Wrap(err, "")
// 		}
// 		return string(bytes), nil
// 	}

// 	groupInfoStr, err := db.DB.Rc.Fetch(userSpaceGroupInfo+userid, time.Minute*5, getGroupInfo)
// 	if err != nil {
// 		return nil, utils.Wrap(err, "")
// 	}
// 	groupInfo := &db.Group{}
// 	err = json.Unmarshal([]byte(groupInfoStr), groupInfo)
// 	return groupInfo, utils.Wrap(err, "")
// }
