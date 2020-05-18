package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// RPC registry
type Rpcs struct {
	UnimplementedServerRegistryGrpcServer
}

func (Rpcs) ServerRegistry(ctx context.Context, request *ServerRequest) (*ServerResult, error) {
	// 先向 upsteam注册
	upsteam := &Upstream{}
	upsteam.Name = request.Name
	upsteam.Desc = request.Desc
	upsteam.Nodes = make([]*_nodes, 0)
	for _, v := range request.Nodes {
		_n := &_nodes{
			Addr:  v.Addr,
			Weigh: v.Weigh,
		}
		upsteam.Nodes = append(upsteam.Nodes, _n)
	}
	toUpstream, err := upsteam.MarshalJSON()
	if err != nil {
		return nil, err
	}
	_, err = EtcdClient.Put(ctx, PREFIX, string(toUpstream))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func StartRpc() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	RegisterServerRegistryGrpcServer(grpcServer, &Rpcs{})
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
