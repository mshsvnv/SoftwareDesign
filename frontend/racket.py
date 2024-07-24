import flet as ft
import utils
import requests as req

class Racket(ft.Container):

    def __init__(self, racket: dict):
        super().__init__()

        self.col = {"sm": 6, "md": 4, "xl": 2}
        self.bgcolor = ft.colors.GREY_300

        self.height = 400
        self.width = 200
        self.margin = 10

        self.content = ft.Column(
            alignment=ft.MainAxisAlignment.CENTER,
            controls = [
                ft.Text(value = "Цена: " + str(racket["Price"])),
                ft.Text(value = "Бренд: " + racket["Brand"])
            ]
        )

        self.ink = True
        self.on_click = lambda _: self.page.go(f"/rackets/:{racket['ID']}")

class Rackets(ft.Container):
    
    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page

        self.padding = 100

        self.rackets = ft.ResponsiveRow()
        self.getRackets()

        self.content = self.rackets

    def getRackets(self):

        resp = req.get(utils.url + '/rackets')
        data = resp.json()

        if resp.status_code == 200:

            for racket in data["rackets"]:
                
                self.rackets.controls.append(
                    Racket(racket)
                )

            # self.update()
        # return
