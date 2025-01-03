package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	if os.Getuid() != 0 {
		fmt.Println("run as root")
		return
	}

	fmt.Printf(
		"running as uid: %v, gid: %v\n",
		os.Getuid(),
		os.Getgid(),
	)

	// pega o usuário nobody
	u, err := user.Lookup("nobody")
	if err != nil {
		fmt.Println(err)
		return
	}

	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)

	// ajusta o id do grupo
	err = syscall.Setgid(gid)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ajusta o id do usuário
	err = syscall.Setuid(uid)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"running as uid: %v, gid: %v\n",
		os.Getuid(),
		os.Getgid(),
	)
}
