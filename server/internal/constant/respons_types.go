package resType

type ResType string

const (
	AccountCreated ResType = "ACCOUNT_CREATED"
	Login          ResType = "LOGIN"
	Refresh        ResType = "REFRESH"
	Logout         ResType = "LOGOUT"
	GetDisks       ResType = "GETDISKS"
	GetDirInfo     ResType = "GETDIRINFO"
)
