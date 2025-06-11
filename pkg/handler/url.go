package handler

// import "net/http"

const (
	pathPing                 string = "/ping"
	pathOption               string = "/option"
	pathSingUp               string = "/auth/sign-up"
	pathActionUserSearch     string = "/action/user/search"
	pathActionShutDownServer string = "/action/shutdown"

	patMakeTransaction string = "/transaction/make_transaction"

	pathTestHandler string = "/test/action/user"
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
		Path:    "/v1/action/user",
		Methods: "GET",
	},
	{
		Path:    patMakeTransaction,
		Methods: "PATCH",
	},
	{
		Path:    pathActionShutDownServer,
		Methods: "GET",
	},
	{
		Path:    pathTestHandler,
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
	{
		Path:    pathTestHandler,
		Methods: "GET",
	},
}
