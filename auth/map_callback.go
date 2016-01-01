package auth

import (
	"types"
	"errors"
)

type MapUserPassInfo struct {
	UserPassInfo
}

func (m *MapUserPassInfo) Type() types.AuthType {
	return types.MAP
}


type MapCallback struct {
	log *types.AppLogger
	users map[string]string
}

func (c *MapCallback) Type() types.AuthType {
	return types.MAP
}


func (c *MapCallback) Authenticate(coninfo types.ConnInfo, authinfo types.AuthInfo) (*types.Permissions, error) {
	if authinfo.Type() != types.MAP {
		return nil, errors.New("unsupported auth type");
	}
	userinfo := authinfo.(UserPassAuthInfo)
	pw, ok := c.users[coninfo.User()]
	if ok == false || userinfo.Password() != pw {
		c.log.Err.Printf("invalid user=%s", coninfo.User())
		return nil, errors.New("unknown user")
	}
	return &types.Permissions{Options:nil, Extensions:nil}, nil
}



