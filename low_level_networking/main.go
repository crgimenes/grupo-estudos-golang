package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleLine(conn net.Conn) error {
	defer conn.Close()
	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(conn, "echo:%s", strings.TrimSpace(line))
	return err
}
