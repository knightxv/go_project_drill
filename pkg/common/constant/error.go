package constant

import "errors"

// key = errCode, string = errMsg
type ErrInfo struct {
	ErrCode int32
	ErrMsg  string
}

var (
	OK        = ErrInfo{0, ""}
	ErrServer = ErrInfo{500, "server error"}

	ErrParseToken = ErrInfo{700, ParseTokenMsg.Error()}

	ErrTencentCredential = ErrInfo{400, ThirdPartyMsg.Error()}

	ErrTokenExpired             = ErrInfo{701, TokenExpiredMsg.Error()}
	ErrTokenInvalid             = ErrInfo{702, TokenInvalidMsg.Error()}
	ErrTokenMalformed           = ErrInfo{703, TokenMalformedMsg.Error()}
	ErrTokenNotValidYet         = ErrInfo{704, TokenNotValidYetMsg.Error()}
	ErrTokenUnknown             = ErrInfo{705, TokenUnknownMsg.Error()}
	ErrTokenKicked              = ErrInfo{706, TokenUserKickedMsg.Error()}
	ErrTokenDifferentPlatformID = ErrInfo{707, TokenDifferentPlatformIDMsg.Error()}
	ErrTokenDifferentUserID     = ErrInfo{708, TokenDifferentUserIDMsg.Error()}

	ErrAccess                = ErrInfo{ErrCode: 801, ErrMsg: AccessMsg.Error()}
	ErrDB                    = ErrInfo{ErrCode: 802, ErrMsg: DBMsg.Error()}
	ErrArgs                  = ErrInfo{ErrCode: 803, ErrMsg: ArgsMsg.Error()}
	ErrStatus                = ErrInfo{ErrCode: 804, ErrMsg: StatusMsg.Error()}
	ErrCallback              = ErrInfo{ErrCode: 809, ErrMsg: CallBackMsg.Error()}
	ErrSendLimit             = ErrInfo{ErrCode: 810, ErrMsg: "send msg limit, to many request, try again later"}
	ErrMessageHasReadDisable = ErrInfo{ErrCode: 811, ErrMsg: "message has read disable"}
	ErrInternal              = ErrInfo{ErrCode: 812, ErrMsg: "internal error"}
	ErrRPC                   = ErrInfo{ErrCode: 813, ErrMsg: "rpc error"}
	ErrChain                 = ErrInfo{ErrCode: 814, ErrMsg: "chain error"}

	ErrScoreMustBeGreaterMinWithdrawScore = ErrInfo{ErrCode: 815, ErrMsg: "score must be greater min withdraw score"}
	ErrWithdrawCoinNotSupport             = ErrInfo{ErrCode: 816, ErrMsg: "withdraw coin not support"}
	ErrWithdrawChainIdNotSupport          = ErrInfo{ErrCode: 817, ErrMsg: "withdraw chain id not support"}
	ErrUserScoreLessWithdrawScore         = ErrInfo{ErrCode: 818, ErrMsg: "user score less withdraw score"}
	ErrWithdrawalRestriction              = ErrInfo{ErrCode: 819, ErrMsg: "The withdrawal time limit is not less than 1 minute"}
)

var (
	ParseTokenMsg               = errors.New("parse token failed")
	TokenExpiredMsg             = errors.New("token is timed out, please log in again")
	TokenInvalidMsg             = errors.New("token has been invalidated")
	TokenNotValidYetMsg         = errors.New("token not active yet")
	TokenMalformedMsg           = errors.New("that's not even a token")
	TokenUnknownMsg             = errors.New("couldn't handle this token")
	TokenUserKickedMsg          = errors.New("user has been kicked")
	TokenDifferentPlatformIDMsg = errors.New("different platformID")
	TokenDifferentUserIDMsg     = errors.New("different userID")
	AccessMsg                   = errors.New("no permission")
	StatusMsg                   = errors.New("status is abnormal")
	DBMsg                       = errors.New("db failed")
	ArgsMsg                     = errors.New("args failed")
	CallBackMsg                 = errors.New("callback failed")
	InvitationMsg               = errors.New("invitationCode error")

	ThirdPartyMsg = errors.New("third party error")

	// 幂等错误
	ErrIdempotent = errors.New("idempotent error")
)

const (
	NoError              = 0
	FormattingError      = 10001
	HasRegistered        = 10002
	NotRegistered        = 10003
	PasswordErr          = 10004
	GetIMTokenErr        = 10005
	RepeatSendCode       = 10006
	MailSendCodeErr      = 10007
	SmsSendCodeErr       = 10008
	CodeInvalidOrExpired = 10009
	RegisterFailed       = 10010
	ResetPasswordFailed  = 10011
	RegisterLimit        = 10012
	LoginLimit           = 10013
	InvitationError      = 10014
	DatabaseError        = 10002
	ServerError          = 10004
	HttpError            = 10005
	IoError              = 10006
	IntentionalError     = 10007
	SignErrorWeb3        = 10008
	WalletError          = 10009
)

func (e ErrInfo) Error() string {
	return e.ErrMsg
}

func (e *ErrInfo) Code() int32 {
	return e.ErrCode
}
