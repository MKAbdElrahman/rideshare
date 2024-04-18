## Kafka Producer
### Initialization 
```golang
import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

p, err := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "host1:9092,host2:9092",
    "client.id": socket.gethostname(),
    "acks": "all"})

if err != nil {
    fmt.Printf("Failed to create producer: %s\n", err)
    os.Exit(1)
}
```
**Imports:**

* `github.com/confluentinc/confluent-kafka-go/kafka`: This line imports the necessary package for interacting with Kafka using the Confluent Kafka Go client.

**Producer Creation:**

* `p, err := kafka.NewProducer(&kafka.ConfigMap{...})`: This line attempts to create a new Kafka producer instance (`p`) with a specific configuration (`kafka.ConfigMap`).

**Configuration:**

* `"bootstrap.servers": "host1:9092,host2:9092"`: This property specifies the list of broker addresses (host and port) for your Kafka cluster. You can provide multiple comma-separated addresses for a high availability setup.
* `"client.id": socket.gethostname()"`: This sets a unique identifier for this producer instance. It retrieves the hostname of the machine using `socket.gethostname()`.
* `"acks": "all"`: This setting defines the acknowledgment level for messages sent by the producer. `"all"` ensures that the message is written successfully to all replicas before considering it sent. There are other options like `leader` and `none` for different levels of durability.

**Error Handling:**

* `if err != nil {...}`: This block checks if there was an error (`err`) during producer creation.
    * `fmt.Printf("Failed to create producer: %s\n", err)`: If there's an error, it prints an informative message with the error details.
    * `os.Exit(1)`: The program exits with an error code (1) indicating the failure.


