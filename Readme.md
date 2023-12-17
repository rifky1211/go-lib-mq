# Go Lib MQ

This library is to help you to do publish / subscribe with
RabbitMQ with exchange. I hope this library
can be useful for you

## Installation
```
go get github.com/rifky1211/go-lib-mq
```

## Usage
```go
ps, err := rabbitmq.New(cfg.PubSub.Address)
if err != nil {
    return
}
defer ps.Close()
```

Good Luck Have Fun
