package main

import "github.com/apache/pulsar-client-go/pulsar"

func main(){

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	defer client.Close()

}
