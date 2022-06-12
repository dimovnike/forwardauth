package forward

const (
	XForwardedProto        = "X-Forwarded-Proto"
	XForwardedFor          = "X-Forwarded-For"
	XForwardedHost         = "X-Forwarded-Host"
	XForwardedPort         = "X-Forwarded-Port"
	XForwardedServer       = "X-Forwarded-Server"
	XRealIp                = "X-Real-Ip"
	Connection             = "Connection"
	KeepAlive              = "Keep-Alive"
	ProxyAuthenticate      = "Proxy-Authenticate"
	ProxyAuthorization     = "Proxy-Authorization"
	Te                     = "Te" // canonicalized version of "TE"
	Trailers               = "Trailers"
	TransferEncoding       = "Transfer-Encoding"
	Upgrade                = "Upgrade"
	ContentLength          = "Content-Length"
	SecWebsocketKey        = "Sec-Websocket-Key"
	SecWebsocketVersion    = "Sec-Websocket-Version"
	SecWebsocketExtensions = "Sec-Websocket-Extensions"
	SecWebsocketAccept     = "Sec-Websocket-Accept"
)
