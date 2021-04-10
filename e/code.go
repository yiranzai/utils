package e

var (
	// 100开头的为系统错误
	Success               = NewErr(10000, "Success")
	ServerError           = NewErr(10001, "Internal server error")
	NoRouteError          = NewErr(10002, "No router")
	NoMethodError         = NewErr(10003, "No method")
	PermissionDeniedError = NewErr(10004, "Permission denied")
	NotFoundError         = NewErr(10005, "Not Found")
	BadRequestError       = NewErr(10006, "Bad Request")
)
