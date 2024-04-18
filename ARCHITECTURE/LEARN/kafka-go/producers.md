## Constructing a Kafka Producer

```go
cfg := kafka.ConfigMap{"bootstrap.servers": "localhost"}
p, err := kafka.NewProducer(&cfg)
if err != nil {
		//handle error
}
```

## Sending a Message ASynchronously

Produce single message. This is an asynchronous call that enqueues the message on the internal transmit queue, thus returning immediately. The delivery report will be sent on the provided deliveryChan if specified, or on the Producer object's Events() channel if not.  Returns an error if message could not be enqueued.

```go
err = p.Produce(msg, deliveryChan)
if err != nil {
		//handle error
} 
```
## Sending a Message Synchronously
We wait for the delivery result by reading from the delivery channel(blocking). The delivery result indicates whether the message was successfully delivered or if there was an error.
```go
  err = producer.Produce(message, deliveryChan)
    if err != nil {
        fmt.Printf("Failed to produce message: %s\n", err)
        return
    }

    // Wait for message delivery result
    e := <-deliveryChan
```


## Fu

```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Configure the Kafka producer
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        // Add more configuration parameters as needed
    })
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
        return
    }

    // Produce messages synchronously
    topic := "my-topic"
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:            []byte("key"),     // optional key
        Value:          []byte("message"), // message payload
    }
    deliveryChan := make(chan kafka.Event)
    err = producer.Produce(message, deliveryChan)
    if err != nil {
        fmt.Printf("Failed to produce message: %s\n", err)
        return
    }

    // Wait for message delivery result
    e := <-deliveryChan
    m := e.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
    } else {
        fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
            *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
    }

    // Close the producer
    producer.Close()
}


```