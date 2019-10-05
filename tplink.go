package smarthome

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"

	"github.com/golang/glog"
)

//type Info struct {
//	System     string     `json:"system"`
//	GetSysInfo System     `json:"get_sysinfo"`
//	ActiveMode GetSysInfo `json:"active_mode"`
//}
type Tplink struct{}

const TPLINK_API_PORT = "9999"
const TPLINK_API_PORT_INT = 9999

func (t Tplink) encrypt(unenc string) []byte {
	key := 171
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(len(unenc)))
	for _, c := range unenc {
		a := key ^ int(c)
		key = a
		result = append(result, byte(a))
	}
	return result
}

func (t Tplink) decrypt(enc []byte) string {
	key := 171
	result := []byte("")
	for _, c := range enc {
		a := key ^ int(c)
		key = int(c)
		result = append(result, byte(a))
	}
	return string(result)
}

func (t Tplink) Send(ip string, command string) (string, error) {
	address := net.JoinHostPort(ip, TPLINK_API_PORT)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	glog.Info(conn, string(t.encrypt(command)))

	result, err := ioutil.ReadAll(conn)
	if err != nil || len(result) == 0 {
		return "", errors.New("Could not read data back")
	}

	json.Marshal(t.decrypt(result[4:]))

	return t.decrypt(result[4:]), nil
}