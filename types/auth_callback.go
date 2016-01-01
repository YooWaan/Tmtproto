package types

/*import (
	"errors"
)*/


type AuthType uint
const (
	UNKNOWN AuthType = 0
	MAP     AuthType = 1
	PAM     AuthType = 2
)

type AuthInfo interface {
	Type() AuthType
}


type AuthCallback interface {

	Type() AuthType

	Authenticate(coninfo ConnInfo, authinfo AuthInfo)  (*Permissions, error)

}
