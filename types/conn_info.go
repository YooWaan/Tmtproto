import (
	"net"
)

type ConnInfo interface {

	// user name
	User() string

	// SessionID
	SessionID() []byte

	// ClientVersion
	ClientVersion() []byte

	// ServerVersion
	//ServerVersion() []byte

	// RemoteAddr
	RemoteAddr() net.Addr

	// LocalAddr
	LocalAddr() net.Addr
}
