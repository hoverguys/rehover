package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	port := flag.Int("port", 55020, "Port to connect to")
	loopback := flag.String("ip", "127.0.0.1", "Loopback address (to connect to)")
	interval := flag.Int("ms", 50, "How often to attempt a connection (in milliseconds)")
	debug := flag.Bool("debug", false, "Print all errors for debugging (will get pretty busy!)")
	dumb := flag.Bool("dumb", false, "Just print anything that comes out instead of serving+parsing")
	bind := flag.String("bind", ":7581", "[<address>]:<port> to bind the webserver to")
	html := flag.String("html", "usbcat-html", "Path to static files (html/js/css)")

	flag.Parse()

	var logger Logger
	if *dumb {
		logger = StdoutLogger{}
	} else {
		wslogger := MakeWSLogger()
		go func() {
			log.Println("[WS] Listening on", *bind)
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*html))))
			http.Handle("/ws", wslogger)
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, *html+"/index.html")
			})
			log.Fatalln(http.ListenAndServe(*bind, nil))
		}()
		logger = wslogger
	}

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
		logger.Connected()

		reader := bufio.NewReader(conn)
		for {
			// Read next line
			str, err := reader.ReadString('\n')
			if err != nil {
				if *debug {
					log.Println(err)
				}
				break
			}

			// Check for type
			strtype := "console"
			strmsg := strings.TrimSpace(str)
			if len(str) > 0 && str[0] == ':' {
				stridx := strings.IndexRune(str, ' ')
				if stridx > 0 {
					strtype = str[1:stridx]
					strmsg = str[stridx+1:]
				}
			}

			// Log
			logger.Log(Message{
				Type: strtype,
				Text: strmsg,
			})
		}

		logger.Disconnected()
		conn.Close()
		time.Sleep(time.Millisecond * time.Duration(*interval))
	}
}
