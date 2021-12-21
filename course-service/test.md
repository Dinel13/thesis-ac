ab -k -c 900 -n 80000 127.0.0.1:8080/course/1

-c: ("Concurrency"). Indicates how many clients (people/users) will be hitting the site at the same time. While ab runs, there will be -c clients hitting the site. This is what actually decides the amount of stress your site will suffer during the benchmark.

-n: Indicates how many requests are going to be made. This just decides the length of the benchmark. A high -n value with a -c value that your server can support is a good idea to ensure that things don't break under sustained stress: it's not the same to support stress for 5 seconds than for 5 hours.

-k: This does the "KeepAlive" funcionality browsers do by nature. You don't need to pass a value for -k as it it "boolean" (meaning: it indicates that you desire for your test to use the Keep Alive header from HTTP and sustain the connection). Since browsers do this and you're likely to want to simulate the stress and flow that your site will have from browsers, it is recommended you do a benchmark with this.


../ghz/ghz --insecure   --proto ./proto/course.proto   --call  proto.CourseService.GetCourse   -d '{"id":1}' -n 10000 -c 1000   0.0.0.0:8081


buat handler grpc dan rest pake interface

Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true


protoc proto/course.proto --go_out=plugins=grpc:.