package main

import (
	"flag"
	"log"
	"net/url"
	"strings"

	"github.com/fastbyt3/go-lb/lb"
)

func main() {
	var port int
	var serverList string

	flag.StringVar(&serverList, "servers", "", "List of backend addresses comma separated")
	flag.IntVar(&port, "port", 9999, "Port to run load balancer")
	flag.Parse()

	if len(serverList) == 0 {
		log.Fatalln("No backend server address were provided...")
	}

	var serverPool lb.ServerPool

	addresses := strings.Split(serverList, ",")
	for _, address := range addresses {
		validUrl, err := url.Parse(address)
		if err != nil {
			log.Fatalln("Failed to parse url", address, err)
		}

		be := lb.NewBackedServer(validUrl)
		serverPool.AddServer(&be)
	}

	loadBalancer := lb.LoadBalancer{
		Servers: serverPool,
		Port: port,
	}

	if err := loadBalancer.Start(); err != nil {
		log.Fatalln("Error occured in load balancer", err)
	}
}
