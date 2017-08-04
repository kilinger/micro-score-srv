package main

import (
	// registry
	_ "github.com/micro/go-plugins/registry/etcdv3"

	// transport
	_ "github.com/micro/go-plugins/transport/grpc"
	_ "github.com/micro/go-plugins/transport/rabbitmq"

	// broker
	_ "github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/redis"
)
