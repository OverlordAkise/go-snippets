package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
)

//Taken from: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("WARNING: COULDN'T DIAL UDP! LISTENING ON 0.0.0.0 !")
		return net.ParseIP("0.0.0.0")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	var clientip string
	flag.StringVar(&clientip, "clientip", "", "If set: Only this IPv4 can connect to the proxy")
	var proxyip string
	flag.StringVar(&proxyip, "proxyip", "", "Manually set the IP the proxy server should listen on, if not set will automatically select the outbound one")
	var proxyport int
	flag.IntVar(&proxyport, "port", 21000, "On what port should the proxy server listen")
	var target string
	flag.StringVar(&target, "target", "", "Where should the traffic be forwarded to? Format IP+Port")

	flag.Parse()

	if target == "" {
		fmt.Println("ERROR: Please provide a target!")
		fmt.Println("Example: ./luctusproxy -target '1.2.3.4:27015'")
		flag.PrintDefaults()
		return
	}

	_, splitPort, err := net.SplitHostPort(target)
	if err != nil {
		panic(err)
	}
	TargetPort, err := strconv.Atoi(splitPort)
	if err != nil {
		panic(err)
	}

	listenip := net.ParseIP(proxyip)
	if proxyip == "" {
		listenip = GetOutboundIP()
	}
	fmt.Println("Listening on:", listenip, ":", proxyport)

	add := net.UDPAddr{
		Port: proxyport,
		IP:   listenip,
	}
	svdst, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		panic(err)
	}

	sock, err := net.ListenUDP("udp", &add)
	if err != nil {
		panic(err)
	}
	defer sock.Close()

	recvbuffer := make([]byte, 65000)

	var cldst *net.UDPAddr
	clientFound := false

	go func() {
		for {
			n, udpaddr, err := sock.ReadFromUDP(recvbuffer)
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}
				panic(err)
			}
			if !clientFound {
				if clientip != "" && udpaddr.IP.String() != clientip {
					continue
				}
				cldst = udpaddr
				clientFound = true
				fmt.Println("Proxy client is:", cldst)
			}
			if udpaddr.Port == TargetPort {
				_, err = sock.WriteTo(recvbuffer[:n], cldst)
				if err != nil {
					panic(err)
				}
			} else {
				_, err = sock.WriteTo(recvbuffer[:n], svdst)
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	fmt.Println("Press CTRL+C to quit")
	<-sigchan
	fmt.Println("Caught CTRL+C, exiting...")
	sock.Close()
}
