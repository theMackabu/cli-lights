package helpers

import (
	"bufio"
	"fmt"
	"net"
)

const (
	Power   = `{"id":1,"method":"setState","params":{"state":%s}}`
	Dimming = `{"id":1,"method":"setPilot","params":{"dimming":%s}}`
	Speed   = `{"id":1,"method":"setPilot","params":{"speed":%s}}`
	Warm    = `{"id":1,"method":"setPilot","params":{"w":%s}}`
	Cold    = `{"id":1,"method":"setPilot","params":{"c":%s}}`
	Scene   = `{"id":1,"method":"setPilot","params":{"sceneId":%s}}`
	Temp    = `{"id":1,"method":"setPilot","params":{"temp":%s}}`
	RGB     = `{"id":1,"method":"setPilot","params":{"r":%s,"g":%s,"b":%s}}`
)

var Scenes = [32]string{
	"ocean",
	"romance",
	"sunset",
	"party",
	"fire place",
	"cozy",
	"forest",
	"pastel",
	"morning",
	"bed time",
	"warm",
	"day light",
	"cool",
	"night light",
	"focus",
	"relax",
	"true colors",
	"tv time",
	"plant growth",
	"spring",
	"summer",
	"fall",
	"deep dive",
	"jungle",
	"mojito",
	"club",
	"christmas",
	"halloween",
	"candle light",
	"golden white",
	"pulse",
	"steam punk",
}

func Write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}
	return number, err
}

func RunClient(addr string, cmd string) string {
	conn, err := net.Dial("udp", addr)
	recvBuf := make([]byte, 1024)

	if err != nil {
		return fmt.Sprintf("can't connect to server: %s\n", err)
	}

	Write(conn, cmd)
	conn.Read(recvBuf[:])

	return string(recvBuf)
}
