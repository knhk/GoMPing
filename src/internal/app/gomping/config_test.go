package gomping

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
)
var confDir = "../../../../test/"

// expected errors
var pathError *os.PathError
var syntaxError *json.SyntaxError
var notargetError = errors.New("no target group")

const (
	configOK int = iota
	pathErr
	syntaxErr
	notargetErr
)

type caseTestNewConfig []struct{
	desc string
	filename string
	result int
}

func getTestNewConfig() caseTestNewConfig{
	return caseTestNewConfig{
		{
			desc: "Perfect config",
			filename: "ok.json",
			result: configOK,
		},{
			desc: "no exist file",
			filename: "gomping.jso",
			result: pathErr,
		},{
			desc: "syntax error file",
			filename: "syntax_error.json",
			result: syntaxErr,
		},{
			desc: "no target group file",
			filename: "no_target.json",
			result: notargetErr,
		},
	}
}

func TestNewConfig(t *testing.T) {
	t.Helper()
	testCase := getTestNewConfig()

	for _, v := range testCase {
		t.Run(v.desc, func(st *testing.T) {
			conf := confDir + v.filename
			c, e := NewConfig(conf)

			if e != nil {
				switch {
				case errors.As(e, &pathError):
					if v.result != pathErr {
						st.Errorf("\x1b[33m[%s] \x1b[31mos.PathError\x1b[37m: %s", v.desc, e)
					} else {
						st.Log("Done: path error")
					}
				case errors.As(e, &syntaxError):
					if v.result != syntaxErr {
						st.Errorf("\x1b[33m[%s] \x1b[31mjson.SyntaxError\x1b[37m: %s", v.desc, e)
					} else {
						st.Log("Done: syntex error")
					}
				case errors.As(e, &notargetError):
					if v.result != notargetErr {
						st.Errorf("%s: %#v", "No target group error: ", e)
					} else {
						st.Log("Done: no target group")
					}
				default:
					st.Errorf("%s: %#v", "Unexpected error: ", e)
				}
			} else if c == nil {
				st.Errorf("%s %#v", "Unexpected value: ", c)
			} else {
				if v.result != configOK {
					st.Error("Unknown Error")
				}
			}
		})
	}
}
