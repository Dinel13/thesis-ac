package grpc

import (
	"context"
	"fmt"

	"github.com/dinel13/thesis-ac/krs/proto"
)

// var url = os.Getenv("URL_AUTH")

// func verifyToken(token string) IsAuth {
// 	conn, err := grpc.Dial(fmt.Sprintf("%s:9091", url), grpc.WithInsecure())
// 	if err != nil {
// 		return IsAuth{
// 			IsAuth: false,
// 			Err:    err,
// 		}
// 	}

// 	defer conn.Close()
// 	c := proto.NewAuthServiceClient(conn)

// 	r, err := c.Verify(context.Background(), &proto.VerifyRequest{
// 		Token: token,
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		return IsAuth{
// 			IsAuth: false,
// 			Err:    err,
// 		}
// 	}
// 	return IsAuth{
// 		IsAuth: r.IsAuth,
// 		Err:    nil,
// 	}
// }

func VerifyToken(ctx context.Context, c proto.AuthServiceClient, token string) (bool, error) {
	// token = token[7:]
	r, err := c.Verify(ctx, &proto.VerifyRequest{
		Token: token,
	})

	if err != nil {
		fmt.Println(err, "ver")
		return false, err
	}
	return r.IsAuth, nil
}
