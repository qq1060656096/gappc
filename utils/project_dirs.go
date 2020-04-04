package utils

const (
	CommandDir = "command"
	EventDir = "event"
	ApiDir = "api"
	AppDir = "app"
	RequestDir = "request"
	ListenerDir = "listener"
	ModelDir = "model"
	TransformerDir = "transformer"
	BootstrapDir = "bootstrap"
	ConfigDir = "config"
	MiddlewareDir = "middleware"
	ThemesDir = "themes"
	RoutersDir = "routers"
	LogDir = "log"
)
var ProjectDirs = map[string]string{
	CommandDir : "app/command",
	EventDir: "app/event",
	ApiDir: "app/http/controller/api/v1",
	AppDir: "app/http/controller/app/v1",
	RequestDir: "app/http/request",
	ListenerDir: "app/listener",
	ModelDir: "app/model",
	TransformerDir: "app/transformer",
	BootstrapDir: "bootstrap",
	ConfigDir: "config",
	MiddlewareDir: "middleware",
	ThemesDir: "resources/themes",
	RoutersDir: "routers",
	LogDir: "runtime/log",
}