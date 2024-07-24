import flet as ft
import requests as req
import utils

class Profile(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page

        self.data = ft.Column()

        self.getProfile()
        self.content = self.data

    def getProfile(self):

        url = utils.url + "/api/profile"
        headers = {
            'Authorization': "Bearer " + self.page.client_storage.get("token")
        }
        resp = req.get(
            url = url,
            headers = headers
        )

        if resp.status_code == 200:
            data = resp.json()
            print(data)

            self.data.controls.append(
                ft.Text(data["user"]["name"] + data["user"]["surname"])
            )
