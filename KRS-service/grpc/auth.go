package grpc

import (
	"context"
	"fmt"
	"os"

	"github.com/dinel13/thesis-ac/krs/proto"
	"google.golang.org/grpc"
)

var url = os.Getenv("URL_AUTH")

func verifyToken(token string) IsAuth {
	conn, err := grpc.Dial(fmt.Sprintf("%s:9091", url), grpc.WithInsecure())
	if err != nil {
		return IsAuth{
			IsAuth: false,
			Err:    err,
		}
	}

	defer conn.Close()
	c := proto.NewAuthServiceClient(conn)

	r, err := c.Verify(context.Background(), &proto.VerifyRequest{
		Token: token,
	})

	if err != nil {
		fmt.Println(err)
		return IsAuth{
			IsAuth: false,
			Err:    err,
		}
	}
	return IsAuth{
		IsAuth: r.IsAuth,
		Err:    nil,
	}
}
