package helpers

import (
	"fmt"
	"net"
)

const (
	Power   = `{"id":1,"method":"setState","params":{"state":%s}}`
	Dimming = `{"id":1,"method":"setPilot","params":{"dimming":%s}}`
	Warm    = `{"id":1,"method":"setPilot","params":{"w":%s,"dimming":100}}`
	Cold    = `{"id":1,"method":"setPilot","params":{"c":%s,"dimming":100}}`
	Scene   = `{"id":1,"method":"setPilot","params":{"sceneId":%s,"dimming":100}}`
	Temp    = `{"id":1,"method":"setPilot","params":{"temp":%s,"dimming":100}}`
	RGB     = `{"id":1,"method":"setPilot","params":{"r":%s,"g":%s,"b":%s,"dimming":100}}`
)

func RunClient(addr string, cmd string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}

	conn.Write([]byte(cmd))
}
