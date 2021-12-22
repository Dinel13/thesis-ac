package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/krs/proto"
	"google.golang.org/grpc"
)

func Create() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Create(context.Background(), &proto.CreateUpdateKRSRequest{
		Token:       "ffdafa",
		IdMahasiswa: int32(1),
		MataKuliahs: []*proto.MataKuliah{
			{
				Kode:     "IF-101",
				Nama:     "Pemrograman Script",
				Sks:      int32(3),
				Dosen:    "Dina",
				Semester: "Semester 7",
			},
			{
				Kode:     "IF-102",
				Nama:     "Pemrograman Berbasis Objek",
				Sks:      int32(3),
				Dosen:    "udin",
				Semester: "Semester 7",
			},
		},
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Krs: %s", r.GetMataKuliahs()[0].GetNama())

}

func Read() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Read(context.Background(), &proto.ReadKRSRequest{
		Token:       "ffdafa",
		IdMahasiswa: int32(1),
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Krs: %s", r.GetMataKuliahs()[0].GetNama())

}

func Update() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Update(context.Background(), &proto.CreateUpdateKRSRequest{
		Token:       "ffdafa",
		IdMahasiswa: int32(1),
		MataKuliahs: []*proto.MataKuliah{
			{
				Kode:     "IF-101",
				Nama:     "Pemrograman Scripts",
				Sks:      int32(3),
				Dosen:    "Dina",
				Semester: "Semester 7",
			},
			{
				Kode:     "IF-102",
				Nama:     "Pemrograman Berbasis Objek",
				Sks:      int32(3),
				Dosen:    "udin",
				Semester: "Semester 7",
			},
		},
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Krs: %s", r.GetMataKuliahs()[0].GetNama())

}

func Delete() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Delete(context.Background(), &proto.DeleteKRSRequest{
		Token:       "ffdafa",
		IdMahasiswa: int32(1),
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Krs: %s", r.GetStatus())

}
