package handler

// import "net/http"

const (
	pathPing                    string = "/ping"
	pathOption                  string = "/option"
	pathSingUp                  string = "/auth/sign-up"
	pathActionUserSearch        string = "/action/user/search"
	pathActionUserGetByID       string = "/action/user/get-by-ID"
	pathActionUserGetByUsername string = "/action/user/get-by-username"
	pathActionShutDownServer    string = "/action/shutdown"
)

const (
	pathMakeTransaction    string = "/transaction/make-transaction"
	pathTransactionGetByID string = "/transaction/get-by-ID"
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
		Path:    pathActionUserGetByID,
		Methods: "GET",
	},
	{
		Path:    pathActionUserGetByUsername,
		Methods: "GET",
	},
	{
		Path:    pathMakeTransaction,
		Methods: "PATCH",
	},
	{
		Path:    pathTransactionGetByID,
		Methods: "GET",
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
