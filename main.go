package main

import (
	"flag"
	"fmt"
	"github.com/kr/beanstalk"
	"chefkoch.de/recommender/lib"
)

func main() {
	var host string
	var port int

	flag.StringVar(&host, "h", "0.0.0.0", "Specify beanstalkd host. Defaults to 0.0.0.0")
	flag.IntVar(&port, "p", 11300, "Specify beanstalkd port. Defaults to 11300")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	manager := lib.NewManager(conn, 10)
	manager.Run()
}
