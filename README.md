# Schema-Registry

apache avro depends on the schema to define record types and records
schemas can be strored on centralised location such as schema registry. 

Schema registry i san application that handles the distributuin of  schemas to producer and consumers and storing them for long term availability.

Subject Name strategy :
    In order to extract the exact schema that we need, subject name strategy is achaieving that by categorising the schemas on the topic that they belong to   
    User Tracking Topic --> {topic-name}-key: user-tacking-key 
                            {topic-name}-value:user-tacking-key
                            
                            


Content Type :
preferred format for content types is application/vnd.schemaregistry.v1+json, where v1 is the API version and json is the serialization format.



References:
https://docs.confluent.io/current/schema-registry/develop/api.html
