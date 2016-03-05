package protoapi

import (
	"net"

	"github.com/octavore/naga-boilerplate/grpc/proto/protoapi/api"
	"github.com/octavore/naga/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const grpcAddr = "localhost:8888"

// Module provides the protoapi GRPC server
type Module struct {
	server *grpc.Server
}

// Init implements service.Module.Init
func (m *Module) Init(c *service.Config) {
	c.Setup = func() error {
		m.server = grpc.NewServer()
		api.RegisterThingServiceServer(m.server, m)
		return nil
	}

	c.Start = func() {
		l, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			panic(err)
		}
		err = m.server.Serve(l)
		if err != nil {
			panic(err)
		}
	}
}

// GetThing implements the proto GetThing method
func (m *Module) GetThing(ctx context.Context, req *api.GetThingRequest) (*api.GetThingResponse, error) {
	return &api.GetThingResponse{
		Things: []int32{1, 2, 3},
	}, nil
}
