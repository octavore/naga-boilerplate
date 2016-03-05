package protoapi

import (
	"reflect"
	"testing"

	"github.com/octavore/naga-boilerplate/grpc/proto/protoapi/api"
	"github.com/octavore/naga/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestGetThing(t *testing.T) {
	m := &Module{}
	stopper := service.New(m).StartForTest()
	defer stopper()

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	c := api.NewThingServiceClient(conn)
	req := &api.GetThingRequest{}
	res, err := c.GetThing(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res.GetThings(), []int32{1, 2, 3}) {
		t.Fatalf("unexpected response %v", res.GetThings())
	}
}
