package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

/*
 允许两种服务注册方式
1. 直接向etcd注册
2. 走grpc协议 ,向服务器注册
*/
var (
	EtcdClient clientv3.Client
)

const PREFIX = "/apisix/"
func main() {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	EtcdClient, err := clientv3.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	getResp, err := EtcdClient.Get(context.Background(), "/service/", clientv3.WithPrefix())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(getResp.Kvs)
	watchChan := EtcdClient.Watch(context.Background(), "/service/", clientv3.WithPrefix())
	for n := range watchChan {
		for _, event := range n.Events {
			switch event.Type {
			case mvccpb.PUT:
				log.Println("put value :", event.Kv)
			case mvccpb.DELETE:
				log.Println("delete value :", event.Kv)
			}
		}

	}
	select {}
}
