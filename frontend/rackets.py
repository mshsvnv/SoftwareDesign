import flet as ft
import utils
import requests as req
import style

class Racket(ft.Container):

    def __init__(self, racket: dict, page: ft.Page):

        super().__init__()

        self.page = page
        self.racket = racket

        self.bgcolor = ft.colors.GREY_300

        self.height = 250
        self.padding = 25
        self.margin = 25
        self.col = {"sm": 6, "md": 4}

        self.column = ft.Column(
            alignment = ft.MainAxisAlignment.END,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = f"Цена: {self.racket['price']}",
                        ),
                        ft.FilledButton(
                            text = "В корзину",
                            style = style.styleGreen,
                            on_click = self.addRacketToCart
                        )
                    ]
                )
            ]
        )

        self.content = self.column

        self.ink = True
        self.on_click = lambda _: self.page.go(f"/rackets/{self.racket['id']}")

    def addRacketToCart(self, e):

        data = {
            "racket_id": self.racket["id"],
            "quantity": 1
        }

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.post(
            utils.url + '/api/cart', 
            headers = headers,
            json = data
        )

        # TODO
        if resp.status_code == 200:
            print("success")
        else:
            print(resp)

class Rackets(ft.Container):
    
    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50
        
        self.rackets = ft.ResponsiveRow()
        self.getRackets()

        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Каталог", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                self.rackets
            ]
        )

    def getRackets(self):

        resp = req.get(utils.url + '/rackets')
        
        if resp.status_code == 200:
            data = resp.json()

            for racket in data["rackets"]:
                
                self.rackets.controls.append(
                    Racket(racket, self.page)
                )

            self.update()
        else:
            print(resp.error)
        
