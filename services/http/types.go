package http

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/spf13/afero"
)

type WebSocketHandler interface {
	In()
	Out()
}

type Context interface {
	HandleWith(handler Handler) error
	HandleAuth(Handler) error
	HandleCleanup(Handler) error
	Request() *http.Request
	Writer() http.ResponseWriter

	RawResponse() bool
	SetRawResponse(val bool)

	Variables() map[string]interface{}
	SetVariable(key string, val interface{})

	Body() []byte
	SetBody([]byte)

	GetStringVariable(key string) (string, error)
	GetStringArrayVariable(key string) ([]string, error)
	GetStringMapVariable(key string) (map[string]interface{}, error)
	GetIntVariable(key string) (int, error)
}

type RouteDefinition struct {
	Host        string
	Path        string
	Vars        Variables
	Scope       []string
	Auth        RouteAuthHandler
	Handler     Handler
	RawResponse bool
}

type Handler func(ctx Context) (interface{}, error)

type RouteAuthHandler struct {
	Validator Handler
	GC        Handler
}

type RawRouteDefinition struct {
	Host        string
	Path        string
	PathPrefix  string
	Vars        Variables
	Scope       []string
	Auth        RouteAuthHandler
	Handler     Handler
	RawResponse bool
}

type Variables struct {
	Required []string
	Optional []string
}

type Service interface {
	Context() context.Context
	Start()
	Stop()
	Wait() error
	Error() error
	/********************************/
	GET(*RouteDefinition)
	PUT(*RouteDefinition)
	POST(*RouteDefinition)
	DELETE(*RouteDefinition)
	PATCH(*RouteDefinition)
	ALL(*RouteDefinition)

	Raw(*RawRouteDefinition) *mux.Route
	LowLevel(*LowLevelDefinition) *mux.Route

	WebSocket(*WebSocketDefinition)

	ServeAssets(*AssetsDefinition)

	AssetHandler(*HeadlessAssetsDefinition, Context) (interface{}, error)
	LowLevelAssetHandler(*HeadlessAssetsDefinition, http.ResponseWriter, *http.Request) error

	GetListenAddress() (*url.URL, error)
}

type WebSocketDefinition struct {
	Host       string
	Path       string
	Vars       Variables
	Scope      []string
	Auth       RouteAuthHandler
	NewHandler func(ctx Context, conn *websocket.Conn) WebSocketHandler
}

type HeadlessAssetsDefinition struct {
	FileSystem            afero.Fs
	Directory             string
	SinglePageApplication bool

	BeforeServe func(w http.ResponseWriter)
}

type AssetsDefinition struct {
	Host                  string
	Path                  string
	Vars                  Variables
	Scope                 []string
	Auth                  RouteAuthHandler
	FileSystem            afero.Fs
	Directory             string
	SinglePageApplication bool

	BeforeServe func(w http.ResponseWriter)
}

type LowLevelDefinition struct {
	Path       string
	PathPrefix string
	Handler    func(w http.ResponseWriter, r *http.Request)
}

type Redirect string
type TemporaryRedirect Redirect
type PermanentRedirect Redirect

type RawData struct {
	ContentType string
	Data        []byte
}

type RawStream struct {
	ContentType string
	Stream      io.ReadCloser
}