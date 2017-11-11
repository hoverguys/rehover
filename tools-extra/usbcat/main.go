package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 55020, "Port to connect to")
	loopback := flag.String("ip", "127.0.0.1", "Loopback address (to connect to)")
	interval := flag.Int("ms", 50, "How often to attempt a connection (in milliseconds)")
	debug := flag.Bool("debug", false, "Print all errors for debugging (will get pretty busy!)")
	flag.Parse()

	for {
		log.Println("Waiting for virtual USBGecko...")
		var conn net.Conn
		var err error
		for {
			conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", *loopback, *port))
			if err != nil {
				if *debug {
					log.Println(err)
				}
				continue
			}
			break
		}
		log.Println("Connected to virtual USBGecko!")
		fmt.Fprintln(os.Stderr, "==============================")
		_, err = io.Copy(os.Stdout, conn)
		fmt.Fprintln(os.Stderr, "==============================")
		if *debug {
			log.Println(err)
		}
		conn.Close()
		time.Sleep(time.Millisecond * time.Duration(*interval))
	}
}
