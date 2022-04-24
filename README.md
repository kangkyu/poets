# poets

Training with Sarama (Kafka Go library)

+ Producer code
+ Run
```sh
  go build
  KAFKA_URL="..." ./poets
```
+ Send POST requests
```sh
  curl localhost:8080/submit --data-raw '{"message":"Here we go"}' -H 'Content-Type: application/json'
```
