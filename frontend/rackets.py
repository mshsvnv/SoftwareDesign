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
        # self.col = 6

        self.column = ft.Column(
            alignment = ft.MainAxisAlignment.END,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = "Цена",
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
        self.on_click = lambda _: self.page.go("/rackets/1")

    def addRacketToCart(self):
        return
    
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
        
        for i in range(21):

            self.rackets.controls.append(
                Racket(None, self.page)
            )

    #     resp = req.get(utils.url + '/rackets')
    #     data = resp.json()

    #     if resp.status_code == 200:

    #         for racket in data["rackets"]:
                
    #             self.rackets.controls.append(
    #                 Racket(racket)
    #             )

            # self.update()
        # return
