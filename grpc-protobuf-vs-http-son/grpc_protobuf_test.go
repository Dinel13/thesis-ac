package benchmarks

import (
	"testing"
	"time"

	grpcprotobuf "github.com/dinel13/thesis-ac/protobuf-vs-json/grpc-protobuf"
	"github.com/dinel13/thesis-ac/protobuf-vs-json/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	client := proto.NewAPIClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client proto.APIClient, b *testing.B) {
	resp, err := client.CreateKrs(context.Background(), &proto.Krs{
		IdMahasiswa: 1,
		MataKuliahs: []*proto.MataKuliah{
			{
				Kode:     "A",
				Nama:     "Algoritma",
				Sks:      3,
				Dosen:    "Dosen A",
				Semester: "Semester 1",
			},
			{
				Kode:     "B",
				Nama:     "Basis Data",
				Sks:      3,
				Dosen:    "Dosen B",
				Semester: "Semester 1",
			},
			{
				Kode:     "C",
				Nama:     "Pemrograman Web",
				Sks:      3,
				Dosen:    "Dosen C",
				Semester: "Semester 1",
			},
		},
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.MataKuliahs == nil {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
