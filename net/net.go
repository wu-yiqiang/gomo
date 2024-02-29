package net

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

var target = flag.String("target", "", "target to scan ports")

func scanner(target string, ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				log.Println("Connection timed out")
			} else {
				log.Println("Connection refused", neterr)
			}
			results <- 0
		} else {
			conn.Close()
			results <- port
		}
	}
}

func ScanNet() {
	log.Println("start")
	defer log.Println("end")

	ports := make(chan int, 100)
	results := make(chan int)

	var open_ports []int

	for port := 1; port <= cap(ports); port++ {
		go scanner("192.168.0.1", ports, results)
	}

	go func() {
		for port := 1; port <= 1024; port++ {
			ports <- port
		}
	}()

	for port := 1; port <= 1024; port++ {
		port_status := <-results
		if port_status != 0 {
			open_ports = append(open_ports, port_status)
		}
	}

	close(ports)
	close(results)
	sort.Ints(open_ports)
	for _, port := range open_ports {
		fmt.Printf("Address %s has Port:%d opened\n", *target, port)
	}
}
