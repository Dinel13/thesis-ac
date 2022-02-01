self.client.post("/login", json={"username":"foo", "password":"bar"}))


locust --headless --users 10 --spawn-rate 1 -H http://localhost:8080
locust --headless --users 10 --spawn-rate 1 -H http://localhost:8081

locust Signup --headless --users 10 --spawn-rate 1 -H http://127.0.0.1:8081


r = l.client.post("/post/endpoint", data=json.dumps(payload), headers=headers, catch_response=True)