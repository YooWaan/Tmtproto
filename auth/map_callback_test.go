package auth

import (
	"os"
	"net"
	"io/ioutil"
	"types"
	"strings"
	"strconv"
	"testing"
)

type DmyConf struct {
	user string
}

func (d DmyConf) User() string {
	return d.user
}

func (d DmyConf) SessionID() []byte {
	return nil
}

func (d DmyConf) ClientVersion() []byte {
	return nil
}

func (d DmyConf) RemoteAddr() net.Addr {
	return nil
}

func (d DmyConf) LocalAddr() net.Addr {
	return nil
}


func TestAuth(t *testing.T) {
	log := types.NewAppLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	var cb types.AuthCallback = &MapCallback{log:log,
		users:map[string]string{"test":"pass", "user":"pwd"}}

	var datas = map[string]string{ "test":"pass", "user":"pwd", "test:false": "ppp", "nouser:false": "no" }

	for k, v := range datas {
		var ci types.ConnInfo = &DmyConf{user:k}
		var ai types.AuthInfo = &MapUserPassInfo{UserPassInfo{passwd:v}}
		p, e := cb.Authenticate(ci, ai)

		var exp = true
		if strings.Index(k, ":") != -1 {
			exp, _ = strconv.ParseBool( string(k[strings.Index(k, ":") : len(k)]) )
			k = k[:strings.Index(k, ":")]
		}

		//log.Info.Printf("k=%s , exp=%s", k, exp)

		if (e != nil) == exp {
			t.Error("has err ", e)
		} else {
			t.Log("ok --> [", p, "]")
		}

	}

}

