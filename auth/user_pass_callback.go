package auth

// 
import (
	"os/user"
	"types"
	"errors"
	"github.com/msteinert/pam"
)

type UserPassAuthInfo interface {
	types.AuthInfo
	Password() string
}

type UserPassInfo struct {
	passwd string
}

func (u *UserPassInfo) Type() types.AuthType {
	return types.PAM
}

func (u *UserPassInfo) Password() string {
	return u.passwd
}

type PAMHandler struct {
	username string
	password string
}


func (p PAMHandler) RespondPAM(style pam.Style, msg string) (string, error) {
	switch style {
	case pam.PromptEchoOn:
		return p.username, nil
	case pam.PromptEchoOff:
		return p.password, nil
	}
	return "", errors.New("unexpected")
	//return p.password, nil
}


type UserPassCallback struct {
	log *types.AppLogger
}

func (c *UserPassCallback) Authenticate(coninfo types.ConnInfo, authinfo types.AuthInfo) (*types.Permissions, error) {
	if authinfo.Type() != types.PAM {
		return nil, errors.New("unsupported auth type");
	}

	user, err := user.Lookup(coninfo.User())
	if err != nil {
		c.log.Err.Printf("user not found %s", coninfo.User())
		return nil, err
	}

	userinfo := authinfo.(UserPassAuthInfo)
	err = c.PamAuth(user.Username, userinfo.Password())
	if err != nil {
		c.log.Err.Printf("can't user authenticate");
		return nil, err
	}
	return &types.Permissions{Options:nil, Extensions:nil}, nil
}

func (c *UserPassCallback) PamAuth(user string, pass string) error {
	/*
	c.log.Info.Printf("user=%s, pw=%s", user, pass);
	tx, err := pam.StartFunc("", "", func (s pam.Style, msg string) (string, error) {

		c.log.Info.Printf("style=%s, msg=%s", s, msg)

		switch s {
		case pam.PromptEchoOn:
			return user, nil
		case pam.PromptEchoOff:
			return pass, nil
		}
		return "", errors.New("unexpected")

		return pass, nil
	})
    */
	tx, err := pam.Start("passwd", user, PAMHandler{username:user, password:pass})
	if err != nil {
		// TODO connection info
		c.log.Err.Printf("can't start user transaction");
		return  err
	}
	defer tx.CloseSession(pam.Silent)
	return tx.Authenticate(pam.Silent)
}
