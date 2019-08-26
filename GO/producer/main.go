package main

import (
	"fmt"
	"kafka"
	"time"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8085"}
var topic = "schematry"

func main() {
	
	schema := /*`{
				"type": "record",
				"name": "PayementField",
				"fields": [
					{"name": "Name", "type": "string"},
					{"name": "Account_number", "type": "int"},
					{"name": "Val", "type": "string"}
				]
			}`*/
			`{
				"type":"record",
				"name":"Payment",
				"fields":[
					{"name":"id","type":"string"},
					{"name":"amount","type":"double"},
					{"name":"region","type":"string","default":""}
				]
			}`
	producer, err := kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("Could not create avro producer: %s", err)
	}
	addMsg(producer, schema)
}

func addMsg(producer *kafka.AvroProducer, schema string) {
	value := `{
		"id": "ID_Mech_1024",
		"amount": 987654,
		"region": "Aerospace"
	}`
	key := time.Now().String()
	err := producer.Add(topic, schema, []byte(key), []byte(value))
	
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}
