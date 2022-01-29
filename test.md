ab -k -c 900 -n 80000 127.0.0.1:8080/krs/1

-c: ("Concurrency"). Indicates how many clients (people/users) will be hitting the site at the same time. While ab runs, there will be -c clients hitting the site. This is what actually decides the amount of stress your site will suffer during the benchmark.

-n: Indicates how many requests are going to be made. This just decides the length of the benchmark. A high -n value with a -c value that your server can support is a good idea to ensure that things don't break under sustained stress: it's not the same to support stress for 5 seconds than for 5 hours.

-k: This does the "KeepAlive" funcionality browsers do by nature. You don't need to pass a value for -k as it it "boolean" (meaning: it indicates that you desire for your test to use the Keep Alive header from HTTP and sustain the connection). Since browsers do this and you're likely to want to simulate the stress and flow that your site will have from browsers, it is recommended you do a benchmark with this.


../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.GetKrs   -d '{"id":1}' -n 10000 -c 1000   0.0.0.0:8081


buat handler grpc dan rest pake interface

Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true


protoc proto/krs.proto --go_out=plugins=grpc:.
/home/din/.local/bin/protoc proto/krs.proto --go_out=plugins=grpc:.

// create krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Create   -d '{"token":       "ffdafa","id_mahasiswa":1,"mata_kuliahs" : [{"kode":     "IF-141","nama":"Pemrograman script", "sks":      3, "dosen":    "Dina","semester": "Semester 7"},{"kode":     "IF-101","nama":     "Pemrograman Script","sks":      3,"dosen":    "Dina","semester": "Semester 7"}]}' -n 1 -c 1  0.0.0.0:9090

./ghz/ghz --insecure   --proto ./proto/krs/krs.proto   --call  proto.KrsService.Create   -d '{"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDM0NDgxMTcsImV4cCI6MTY0MzUzNDUxN30.sxU6PoCzsezqQ3tM0vu0TvkZWjy2_a2mbfNLE-oiYKk","id_mahasiswa":1,"mata_kuliahs" : [{"kode":     "IF-141","nama":"Pemrograman script", "sks":      3, "dosen":    "Dina","semester": "Semester 7"},{"kode":     "IF-101","nama":     "Pemrograman Script","sks":      3,"dosen":    "Dina","semester": "Semester 7"}]}' -n 10 -c 1 -O html  172.31.30.48:9090 > report/ghz2.html


//kreate payment
../ghz/ghz --insecure   --proto ./proto/payment.proto   --call  proto.PaymentService.Create   -d '{"id_mahasiswa":1,"jumlah" : 2345, "metode" : "bri"}' -n 1 -c 1   0.0.0.0:9092

//update krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Update   -d '{"token":       "ffdafa","id_mahasiswa":1,"mata_kuliahs" : [{"kode":     "IF-141","nama":"Pemrograman script", "sks":      3, "dosen":    "Dina","semester": "Semester 7"},{"kode":     "IF-101","nama":     "Pemrograman Script","sks":      3,"dosen":    "Dina","semester": "Semester 7"}]}' -n 1 -c 1   0.0.0.0:9090

//read krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Read   -d '{"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDI2ODE5OTJ9.hMVVyeZyWg5Jq7gLiM3uKyV5jA_GbwZastQ2CWJzqek","id_mahasiswa":1}' -n 1 -c 1   0.0.0.0:9090


// delete krs
../ghz/ghz --insecure   --proto ./proto/krs.proto   --call  proto.KrsService.Delete   -d '{"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNheWEiLCJpYXQiOjE2NDI2ODE5OTJ9.hMVVyeZyWg5Jq7gLiM3uKyV5jA_GbwZastQ2CWJzqek","id_mahasiswa":1}' -n 1 -c 1   0.0.0.0:9090


./bin/jmeter  -g grpc.jtl -o reportt

protoc --proto_path=proto --java_out=. proto/krs.proto


./bin/jmeter -n -t krs.jmx -l testresults.jtl

//todo
delete krs by id juga mendelete pay by id karena pake sama id di redis

