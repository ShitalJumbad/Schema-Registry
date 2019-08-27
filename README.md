# Schema-Registry

apache avro depends on the schema to define record types and records.
schemas can be stored on centralised location such as schema registry. 

Schema registry is an application that handles the distribution of schemas to producer and consumers and storing them for long term availability.

Subject Name strategy :
    In order to extract the exact schema that we need, subject name strategy is achaieving that by categorising the schemas       on the topic that they belong to   
    User Tracking Topic --> {topic-name}-key: user-tacking-key 
                            {topic-name}-value: user-tacking-key
                            
                          

Content Type :
preferred format for content types is application/vnd.schemaregistry.v1+json, where v1 is the API version and json is the serialization format.


* Kafka Producer creates avro record. Which has schemaId and data. With Kafka avro serializer schema is registered if needed and then serializes data and schemaId.
Kafka avro serialzer keeps cache of registered schemas from the schema registery & chema Id 

## Using Confluent
1) curl -O http://packages.confluent.io/archive/5.3/confluent-5.3.0-2.12.zip
2) Unzip the file 
3) setup the zookeeper 
    sudo bin/zookeeper-server-start etc/kafka/zookeeper.properties
4) Setup the broker/server
    sudo bin/kafka-server-start etc/kafka/server.properties
5) Create a topic 
    bin/kafka-topics --create --topic my_topic --zookeeper localhost:2181 --replication-factor 1 --partitions 1
6) Setup the Schema Registry
    sudo bin/schema-registry-start etc/schema-registry/schema-registry.properties 
    tip : Generally Schema Registry listens on port localhost:8081 but I have set it at localhost:8085.
          so next time onwards I am using it at localhost:8085
7) Checking 
    To view all the subjects registered in Schema Registry (assuming Schema Registry is running on the local machine            listening on port 8085):
    
    curl --silent -X GET http://localhost:8085/subjects/ | jq .
    
    op expected :- []
    
8) to define a schema for a new subject
    curl -X POST -v1+json" --data '{"schema": "{\"type\":\"record\",\"name\":\"Payment\",\"namespace\":\"io.confluent.examples.clients.basicavro\",\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"double\"}]}"}' http://localhost:8085/subjects/test-value/versions
    
    I got id :- {"id":41}
    
9) schema id, you can also retrieve the associated schema by querying Schema Registry REST endpoint as follows
    curl --silent -X GET http://localhost:8085/schemas/ids/41 | jq .

10) To view the latest schema for this subject in more detail:
    curl --silent -X GET http://localhost:8085/subjects/test-value/versions/latest | jq .
    
11) Registering a New Version of a Schema Under the Subject "Kafka-key"
    curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" --data '{"schema": "{\"type\": \"string\"}"}' http://localhost:8085/sufka-key/versions

12) Registering a New Version of a Schema Under the Subject "Kafka-value"
    curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" --data '{"schema": "{\"type\": \"string\"}"}' http://localhost:8085/subjects/Kafka-value/versions

12) Listing all the subjects
    curl -X GET http://localhost:8085/subjects
    
13) Registering an Existing Schema to a New Subject Name
    curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" --data "{\"schema\": $(curl -s http://localhost:8085/subjects/Kafka-vaons/latest | jq '.schema')}" http://localhost:8085/subjects/Kafka2-value/versions
    
14) Fetching schema by globally unique id 
    curl -X GET http://localhost:8085/schemas/ids/1
    
15) Listing All Schema Versions Registered Under the Subject "Kafka-value"
    curl -X GET http://localhost:8085/subjects/Kafka-value/versions
    
16) Fetch Version 1 of the Schema Registered Under Subject "Kafka-value" 
    curl -X GET http://localhost:8085/subjects/Kafka-value/versions/1
    
17) Deleting Version 1 of the Schema Registered Under Subject "Kafka-value"
    curl -X DELETE http://localhost:8085/subjects/Kafka-value/versions/1
    
18) Deleting the Most Recently Registered Schema Under Subject "Kafka-value"
    curl -X DELETE http://localhost:8085/subjects/Kafka-value/versions/latest



*** To disable comaptibity in order to register new schema

curl -X PUT -H "Content-Type: application/vnd.schemaregistry.v1+json" --data '{"compatibility": "NONE"}' http://localhost:8085/config

https://docs.confluent.io/current/schema-registry/avro.html




References:
https://docs.confluent.io/current/schema-registry/develop/api.html
https://docs.confluent.io/current/schema-registry/using.html
https://dzone.com/articles/kafka-avro-serialization-and-the-schema-registry
