package main

import (
	"fmt"

	kafka "booking-service/kafkaSetup/v1"
)

// func testDbConnection() {
// 	dbClient, err := db.GetDbClient()
// 	if err != nil {
// 		fmt.Println("error occured while initializing db")
// 		panic(err)
// 	}
// 	fmt.Println(dbClient)
// 	fmt.Println("test database table creation....")
// 	db.CheckTableCreation(dbClient)
// }

func testKafkaConnection() {
	kafka.WriteMessageOnTopic()
}

func main() {
	fmt.Println("App started...!")
	// testDbConnection()
	testKafkaConnection()
	fmt.Println("closing application")
}

// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// func main() {
// 	// host := "google.com"
// 	timeout := 2 * time.Second
// 	const KAFKA_BROKER_URL = "pkc-ymrq7.us-east-2.aws.confluent.cloud:9092"

// 	_, err := net.DialTimeout("tcp", KAFKA_BROKER_URL, timeout)
// 	if err != nil {
// 		fmt.Println("Server is unavailable")
// 		return
// 	}

// 	fmt.Println("Server is available")
// }
