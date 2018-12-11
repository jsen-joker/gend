package boot

import (
	"context"
	pb "gend/exporter/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)
type ClientService struct {
	connect *grpc.ClientConn
	gendRpcClient *pb.GendRpcClient
	//context *context.Context
	//cancelFunc *context.CancelFunc
}

func (cs *ClientService) init(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	cs.connect = conn

	c := pb.NewGendRpcClient(conn)
	cs.gendRpcClient = &c
}

func (cs *ClientService) GenId() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := (*cs.gendRpcClient).GenId(ctx, &pb.Request{})

	if err != nil {
		log.Printf("could not greet: %v", err)
		return 0
	}
	return r.Id
}

func (cs *ClientService) ExpId(id int64) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := (*cs.gendRpcClient).ExpId(ctx, &pb.RequestLong{Id: id})

	if err != nil {
		log.Printf("could not greet: %v", err)
		return ""
	}
	return r.Json
}

var client *ClientService
func GetClientService(address string) *ClientService {
	if client == nil {
		client = &ClientService{}
		client.init(address)

	}
	return client
}
