
from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def baidu_index1(self):
       self.client.post("/ping",data={'exp':'1+2*5-6/3'})