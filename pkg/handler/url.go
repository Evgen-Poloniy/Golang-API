package handler

// import "net/http"

const (
	pathPing                 string = "/ping"
	pathOption               string = "/option"
	pathSingUp               string = "/auth/sign-up"
	pathActionUserSearch     string = "/action/user/search"
	pathActionUserByID       string = "/action/user/get-by-ID"
	pathActionUserByUsername string = "/action/user/get-by-username"
	pathActionShutDownServer string = "/action/shutdown"

	pathMakeTransaction string = "/transaction/make-transaction"
)

const (
	methodNotAllowed string = "method not allowed"
)

type handlerParametrs struct {
	Path    string
	Methods string
}

var urlsParametrs = []handlerParametrs{
	{
		Path:    pathPing,
		Methods: "GET",
	},
	{
		Path:    pathOption,
		Methods: "GET",
	},
	{
		Path:    pathSingUp,
		Methods: "POST",
	},
	{
		Path:    pathActionUserSearch,
		Methods: "GET, POST",
	},
	{
		Path:    pathActionUserByID,
		Methods: "GET",
	},
	{
		Path:    pathActionUserByUsername,
		Methods: "GET",
	},
	{
		Path:    pathMakeTransaction,
		Methods: "PATCH",
	},
	{
		Path:    pathActionShutDownServer,
		Methods: "GET",
	},
}

var urlsParametrsDebugWithoutDB = []handlerParametrs{
	{
		Path:    pathPing,
		Methods: "GET",
	},
	{
		Path:    pathOption,
		Methods: "GET",
	},
	{
		Path:    pathActionShutDownServer,
		Methods: "GET",
	},
}
