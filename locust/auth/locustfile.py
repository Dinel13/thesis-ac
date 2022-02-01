import json
from locust import HttpUser, task

class Signup(HttpUser):
    @task
    def hello_world(self):
        response = self.client.post("/signup", data=json.dumps({'username': 'salahuddin', 'password': 'Password123'}), headers={'content-type': 'application/json'})
        print("Response status code:", response.status_code)
        print("Response text:", response.text)


class Login(HttpUser):
    @task
    def hello_world(self):
        response = self.client.post("/login", data=json.dumps({"username": "salahuddin", "password": "Password123"}), headers={'content-type': 'application/json'})
        print("Response status code:", response.status_code)
        print("Response text:", response.text)

class Verify(HttpUser):
    @task
    def hello_world(self):
        response = self.client.post("/verify", {}, headers={'content-type': 'application/json', 'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNhbGFodWRkaW4iLCJpYXQiOjE2NDM3Mjg1MzksImV4cCI6MTY0MzgxNDkzOX0.yrcVEGkOhUl16H3uUFbgaPHUo-wtjoSfOKQg-1JYtCc'})
        print("Response status code:", response.status_code)
        print("Response text:", response.text)