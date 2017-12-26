package main

import (
	"github.com/nats-io/go-nats"
	"time"
	"fmt"
	"os"
)

func t()  {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)
}


func main()  {


	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)


	nc, _ := nats.Connect(argsWithoutProg[0])

	nc.Subscribe("send", func(m *nats.Msg){
		fmt.Printf("Received a message: %s\n", string(m.Data))
		t();
		nc.Publish("receive",m.Data)
	})
	// Simple Publisher

	for{
		time.Sleep(time.Second)
	}
}
