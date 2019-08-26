package kafka

import (
	"encoding/binary"
	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro"
	"fmt"
)

type AvroProducer struct {
	producer             sarama.SyncProducer
	schemaRegistryClient *SchemaRegistryClient
}

func NewAvroProducer(kafkaServers []string, schemaRegistryServers []string) (*AvroProducer, error) {
	
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer(kafkaServers, config)
	if err != nil {
		return nil, err
	}
	schemaRegistryClient := NewSchemaRegistryClient(schemaRegistryServers)
	return &AvroProducer{producer, schemaRegistryClient}, nil
}

//get schema id from schema-registry service
func (ap *AvroProducer) GetSchemaId(topic string, avroCodec *goavro.Codec) (int, error) {
	schemaId, err := ap.schemaRegistryClient.CreateSubject(topic, avroCodec)
	if err != nil {
		return 0, err
	}
	fmt.Println("**************** I am avro producer ***********************")
	fmt.Println("I have the Schema")
	fmt.Println(avroCodec)
	fmt.Println("I got Id")
	fmt.Println(schemaId)
	fmt.Println("***********************************************************")
	return schemaId, nil
}

func (ap *AvroProducer) Add(topic string, schema string, key []byte, value []byte) error {
	avroCodec, err := goavro.NewCodec(schema)
	schemaId, err := ap.GetSchemaId(topic, avroCodec)
	if err != nil {
		return err
	}
	binarySchemaId := make([]byte, 4)
	binary.BigEndian.PutUint32(binarySchemaId, uint32(schemaId))

	native, _, err := avroCodec.NativeFromTextual(value)
	if err != nil {
		return err
	}

	// Convert native Go form to binary Avro data
	binaryValue, err := avroCodec.BinaryFromNative(nil, native)
	if err != nil {
		return err
	}

	var binaryMsg []byte

	binaryMsg = append(binaryMsg, byte(0))
	//4-byte schema ID as returned by the Schema Registry
	binaryMsg = append(binaryMsg, binarySchemaId...)
	//avro serialized data in Avro’s binary encoding
	binaryMsg = append(binaryMsg, binaryValue...)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(binaryMsg),
	}
	_, _, err = ap.producer.SendMessage(msg)
	return err
}

func (ac *AvroProducer) Close() {
	ac.producer.Close()
}
