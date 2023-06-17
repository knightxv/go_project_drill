package constant

const (

	//group admin
	//	OrdinaryMember = 0
	//	GroupOwner     = 1
	//	Administrator  = 2
	//group application
	//	Application      = 0
	//	AgreeApplication = 1

	//friend related
	BlackListFlag         = 1
	ApplicationFriendFlag = 0
	FriendFlag            = 1
	RefuseFriendFlag      = -1

	//Websocket Protocol
	WSGetNewestSeq     = 1001
	WSPullMsgBySeqList = 1002
	WSSendMsg          = 1003
	WSSendSignalMsg    = 1004
	WSPushMsg          = 2001
	WSKickOnlineMsg    = 2002
	WsLogoutMsg        = 2003
	WSDataError        = 3001

	///ContentType
	//UserRelated
	Text                         = 101
	Picture                      = 102
	Voice                        = 103
	Video                        = 104
	File                         = 105
	AtText                       = 106
	Merger                       = 107
	Card                         = 108
	Location                     = 109
	Custom                       = 110
	Revoke                       = 111
	HasReadReceipt               = 112
	Typing                       = 113
	Quote                        = 114
	GroupHasReadReceipt          = 116
	AdvancedText                 = 117
	AdvancedRevoke               = 118 //影响前者消息
	CustomNotTriggerConversation = 119
	CustomOnlineOnly             = 120
	Announcement                 = 121

	Common             = 200
	GroupMsg           = 201
	SignalMsg          = 202
	CustomNotification = 203

	// 准备提现
	WithdrawStatusUnStart = "UN_START"
	// 提现上链成功，确认中
	WithdrawStatusConfirming = "CONFIRMING"
	// 提现成功
	WithdrawStatusSuccess = "SUCCESS"
	// 提现失败，提现回滚
	WithdrawStatusFail = "FAIL"
	// 确认失败，提现回滚
	WithdrawStatusRollback = "ROLLBACK"

	// 上链中
	ChainUpStatusUnStart = "UPING"
	// 上链成功
	ChainUpStatusSuccess = "UP_SUCCESS"
	// 上链失败
	ChainUpStatusFail = "FAIL"
	// 上链成功，回执失败
	ChainUpStatusReceiptFail = "RECEIPT_FAIL"
	// 上链成功，执行成功（真正的成功）
	ChainUpStatusMinedSuccess = "MINED"

	ChainUpTypeWithdraw = "withdraw"

	UserScoreInfoUpdatedNotification = 1305

	//SessionType
	SingleChatType             = 1
	GroupChatType              = 2
	SuperGroupChatType         = 3
	NotificationChatType       = 4
	NotificationOnlinePushType = 5
)

const (
	WriteDiffusion = 0
	ReadDiffusion  = 1
)

const (
	AtAllString       = "AtAllTag"
	AtNormal          = 0
	AtMe              = 1
	AtAll             = 2
	AtAllAtMe         = 3
	GroupNotification = 4
)

var ContentType2PushContent = map[int64]string{
	Picture:      "[图片]",
	Voice:        "[语音]",
	Video:        "[视频]",
	File:         "[文件]",
	Text:         "你收到了一条文本消息",
	AtText:       "[有人@你]",
	GroupMsg:     "你收到一条群聊消息",
	Common:       "你收到一条新消息",
	SignalMsg:    "音视频通话邀请",
	Announcement: "公告信息",
}

const (
	FieldRecvMsgOpt    = 1
	FieldIsPinned      = 2
	FieldAttachedInfo  = 3
	FieldIsPrivateChat = 4
	FieldGroupAtType   = 5
	FieldIsNotInGroup  = 6
	FieldEx            = 7
	FieldUnread        = 8
	FieldBurnDuration  = 9
)

const (
	AppOrdinaryUsers = 1
	AppAdmin         = 2

	GroupOrdinaryUsers = 1
	GroupOwner         = 2
	GroupAdmin         = 3

	GroupResponseAgree  = 1
	GroupResponseRefuse = -1

	FriendResponseAgree  = 1
	FriendResponseRefuse = -1

	Male   = 1
	Female = 2
)

const (
	UnreliableNotification    = 1
	ReliableNotificationNoMsg = 2
	ReliableNotificationMsg   = 3
)

const (
	ApplyNeedVerificationInviteDirectly = 0 // 申请需要同意 邀请直接进
	AllNeedVerification                 = 1 //所有人进群需要验证，除了群主管理员邀请进群
	Directly                            = 2 //直接进群
)

const (
	GroupRPCRecvSize = 30
	GroupRPCSendSize = 30
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"

const BigVersion = "v2"

const LogFileName = "UserScore.log"

const StatisticsTimeInterval = 60

const MaxNotificationNum = 100

const CurrentVersion = "v2.3.4-rc0"
