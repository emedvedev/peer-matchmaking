package main

import (
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"regexp"
)

type VyOS struct {
	session *gexpect.ExpectSubprocess
}

func (vyos *VyOS) GetConnectedRoutes() string {
	vyos.session.SendLine("show ip route connected")
	result, _ := vyos.session.ReadUntil('$')

	re := regexp.Mustcompile

	match := re.FindAllString(result, -1)

	return string(result)
}

func (vyos *VyOS) GetBGPNeighbors() string {
	return "hello"
}

func (vyos *VyOS) SessionAuth() {
	vyos.session.ReadUntil(':')
	vyos.session.SendLine("vagrant")
}

func CreateVyOS(ipAddress string) *VyOS {
	connectionString := fmt.Sprintf("ssh vagrant@%v", ipAddress)
	session, _ := gexpect.Spawn(connectionString)
	vyos := &VyOS{session: session}

	vyos.SessionAuth()

	return vyos
}

func main() {
	vyos := CreateVyOS("172.30.90.2")
	fmt.Print(vyos.session.ReadUntil('$'))

	fmt.Print(vyos.GetConnectedRoutes())
}
