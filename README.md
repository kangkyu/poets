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
  make test
```

+ Consume (will run in foreground)
```sh
  go build
  KAFKA_URL="..." ./poets consume
```
