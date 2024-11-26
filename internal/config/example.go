package config

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func CreateExample() {

	writer := GetKafkaWriter()
	defer writer.Close()

	const LIMIT = 1000

	var messagesArray [LIMIT]kafka.Message
	nextDomain := roundRobin()

	for i := 0; i < LIMIT; i++ {
		testMessage := fmt.Sprintf(`{
            "u_id": 4663,
            "c_id": "1728556244.450",
            "ac": "6707acd467af5",
            "sip_d": "192.168.1.1",
            "s_wid": 1512,
            "url": "%s",
            "hm": "POST",
            "hdr": {
                "api_token": "12345678",
                "content-type": "application/json"
            },
            "w_type": 0,
            "re": 0,
            "cid_num": "7505064723",
            "call_num": "+911244637984",
            "pyld": {
                "uuid": "6707acd467af5",
                "call_to_number": "+911244637984",
                "caller_id_number": "7505064723",
                "start_stamp": "2024-10-10 16:00:44",
                "call_id": "1728556244.450",
                "billing_circle": {
                    "operator": "Reliance Mobile GSM",
                    "circle": "UP (East)"
                },
                "customer_no_with_prefix ": "+917505064723"
            },
            "meta": {},
            "ch": "inbound",
            "partition": 0
        }`,
			nextDomain(),
		)
		messagesArray[i] = kafka.Message{
			Key:   []byte("key"),
			Value: []byte(testMessage),
			Topic: "webhook",
		}
	}

	writer.WriteMessages(context.Background(), messagesArray[:]...)

}
