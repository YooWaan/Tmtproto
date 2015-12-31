package types

/*import (
	"errors"
)*/

const (
	PAM uint8 = 1
)

type AuthCallback interface {

	Type() uint8

	Authenticate(coninfo ConnInfo)  (Permissions, error)

}
