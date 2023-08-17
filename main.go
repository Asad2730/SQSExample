package main

import (
	"fmt"

	"github.com/Asad2730/SQSExample/controlers"
)

func main() {

	queName := "MY_Queue"

	err := controlers.CreateQueue(queName)

	if err != nil {
		panic(err.Error())
	}

	queUrl, err := controlers.GetQueueURL(queName)

	if err != nil {
		panic(err.Error())
	}

	maxMsgs := 13

	res, err := controlers.ReceiveMessages(queUrl, maxMsgs)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, i := range res.Messages {

		fmt.Println("Body", i.Body)
	}
}
