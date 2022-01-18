package grpc

import (
	"context"
	"fmt"

	"github.com/dinel13/thesis-ac/krs/proto"
	"google.golang.org/grpc"
)

func VerifyToken(token string) (bool, error) {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer conn.Close()
	c := proto.NewAuthServiceClient(conn)

	r, err := c.Verify(context.Background(), &proto.VerifyRequest{
		Token: token,
	})

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return r.IsAuth, nil
}
