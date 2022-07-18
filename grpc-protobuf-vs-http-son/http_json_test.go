package benchmarks

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	httpjson "github.com/dinel13/thesis-ac/protobuf-vs-json/http-json"
	"github.com/mailru/easyjson"
)

func init() {
	go httpjson.Start()
	time.Sleep(time.Second)
}

func BenchmarkHTTPJSON(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doPost(client, b)
	}
}

func doPost(client *http.Client, b *testing.B) {
	u := &httpjson.Krs{
		IdMahasiswa: 1,
		MataKuliahs: []*httpjson.MataKuliah{
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
	}

	// buf := new(bytes.Buffer)
	// json.NewEncoder(buf).Encode(u)
	byte, _ := easyjson.Marshal(u)
	reader := bytes.NewReader(byte)
	resp, err := client.Post("http://127.0.0.1:60001/", "application/json", reader)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target httpjson.Response
	decodeErr := easyjson.UnmarshalFromReader(resp.Body, &target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.MataKuliahs == nil {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
