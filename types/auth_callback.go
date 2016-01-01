package types

/*import (
	"errors"
)*/


type AuthType uint
const (
	PAM AuthType = 1
)

type AuthInfo interface {
	Type() AuthType
}


type AuthCallback interface {

	Type() AuthType

	Authenticate(coninfo ConnInfo, authinfo AuthInfo)  (*Permissions, error)

}
