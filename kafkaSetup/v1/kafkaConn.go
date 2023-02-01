package v1

import (
	"fmt"
	"net"
	"time"
)

// check we are able to ping the kafka url or not
func CheckKafkaConnection() bool {

	const KAFKA_BROKER_URL = "pkc-ymrq7.us-east-2.aws.confluent.cloud:9092"
	timeout := 2 * time.Second

	_, err := net.DialTimeout("tcp", KAFKA_BROKER_URL, timeout)
	if err != nil {
		fmt.Println("Server is unavailable")
		return false
	}

	fmt.Println("Server is available")
	return true
}
