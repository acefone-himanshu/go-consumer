package config

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const testMessage = `{
	"u_id": 4663,
	"c_id": "1728556244.450",
	"ac": "6707acd467af5",
	"sip_d": "192.168.1.1",
	"s_wid": 1512,
	"url": "https://webhook.site/3ec53ee7-6595-4df6-889e-bf0e91347406",
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
}`

func CreateExample() {

	writer := GetKafkaWriter()

	var messagesArray [1000]kafka.Message

	for i := 0; i < 1000; i++ {
		messagesArray[i] = kafka.Message{
			Key:   []byte("key"),
			Value: []byte(testMessage),
			Topic: "webhook",
		}
	}

	writer.WriteMessages(context.Background(), messagesArray[:]...)

}
