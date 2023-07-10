# Nats JetStream Example

A simple Nats JetStream PubSub example. 

Here is the Medium post: https://medium.com/@ebubekiryigit/distributed-message-streaming-in-golang-using-nats-jetstream-29f28be66dc6

## Dependencies

- Go 1.20
- Nats Server 1.16

https://docs.nats.io/running-a-nats-service/introduction/installation

## Running

```shell
go build -o go_nats_jetstream .
```

```shell
./go_nats_jetstream
```

### Output

```log
2022/06/28 00:39:48 Consumer  =>  Subject: REVIEWS.rateGiven  -  ID: 58c03ac18060197ca0b52d51  -  Author: 58c039018060197ca0b52d4c  -  Rating: 5
2022/06/28 00:39:48 Publisher  =>  Message: I tried this place last week and it was incredible! 
2022/06/28 00:39:50 Consumer  =>  Subject: REVIEWS.rateGiven  -  ID: 58c03af28060197ca0b52d53  -  Author: 58c03ada8060197ca0b52d52  -  Rating: 1
2022/06/28 00:39:50 Publisher  =>  Message: hipsters everywhere
2022/06/28 00:39:50 Consumer  =>  Subject: REVIEWS.rateGiven  -  ID: 58c08780bbb1a51e0d43a050  -  Author: 58c039018060197ca0b52d4c  -  Rating: 5
2022/06/28 00:39:50 Publisher  =>  Message: Always a great spot to grab a coffee with a friend. 
2022/06/28 00:39:51 Consumer  =>  Subject: REVIEWS.rateGiven  -  ID: 58c08796bbb1a51e0d43a051  -  Author: 58c039018060197ca0b52d4c  -  Rating: 3
2022/06/28 00:39:51 Publisher  =>  Message: Great ramen, a little pricey for what you get.
```

