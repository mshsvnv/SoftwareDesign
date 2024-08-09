import flet as ft
import requests as req

import style
import utils

class CartInfo(ft.Container):

    def __init__(self, page: ft.Page, price, addRacket):
        super().__init__()

        self.page = page
        self.height = 200
        self.padding = 25
        self.border_radius = 30
        self.bgcolor = style.greyColor
        self.col = {"md": 4}

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text(
                    value = f"Цена: {price}",
                    weight = ft.FontWeight.BOLD
                ),
                ft.FilledButton(
                    text = "В корзину",
                    style = style.styleGreen,
                    on_click = addRacket
                )
            ]
        )

class RacketInfo(ft.Container):

    def __init__(self, page: ft.Page, racket: dict):
        super().__init__()

        self.page = page
        self.height = 200
        self.col = {"md": 6}

        self.racket = racket

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Фирма"),
                        ft.Text(
                            value = self.racket['brand']
                        )
                    ]    
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Вес"),
                        ft.Text(
                            value = self.racket['weight']
                        )
                    ]    
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Баланс"),
                        ft.Text(
                            value = self.racket['balance']
                        )
                    ]    
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [ 
                        ft.Text("Размер головы"),
                        ft.Text(
                            value = self.racket['headsize']
                        )
                    ]   
                )
            ]
        )

    def addRacket(self, e):

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

        if resp.status_code == 200:
            print("ok")
        else:
            print(resp)

class Feedback(ft.Container):

    def __init__(self, page: ft.Page, feedback: dict):
        super().__init__()

        self.page = page
        self.height = 200

        name = self.getUserByID(feedback["user_id"])
        
        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = name
                        ),
                        ft.Row(
                            alignment = ft.MainAxisAlignment.END,
                            controls = [
                                ft.Text(
                                    value = f"{feedback['date'][:10]}",
                                ),
                                *[
                                    ft.Icon(
                                        name = ft.icons.SPORTS_BASEBALL,
                                        color = style.greenColor
                                    ) for _ in range(int(feedback['rating']))]
                            ]
                        ),
                    ]
                ),
                ft.Text(feedback["feedback"])
            ]
        )

    def getUserByID(self, id: int):
        
        resp = req.get(utils.url + f"/user/{id}")

        if resp.status_code == 200:
            data = resp.json()

            return data["user"]["name"] + " " + data["user"]["surname"]
        else:
            print(resp)

class Racket(ft.Container):

    def __init__(self, page: ft.Page, racketID: str):
        super().__init__()

        self.page = page
        self.padding = 50

        racket = self.getRacket(racketID)
        racketInfo = RacketInfo(self.page, racket)

        controls = []
        feedbacks = self.getFeedbacks(racketID)

        if feedbacks is not None:
            for feedback in feedbacks:
                controls.append(Feedback(page, feedback))
                controls.append(ft.Divider())
        else:
            controls.append(
                ft.Text(
                    value = "Отзывов нет!"
                )
            )

        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Товар", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            racketInfo,
                            CartInfo(
                                self.page,
                                racket['price'],
                                racketInfo.addRacket
                            ),
                        ],
                    ),
                ),
                ft.Text(
                    "Отзывы", 
                    size = 30,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = controls
                    ),
                )
            ]
        )

        self.ink = True
        self.ink_color = style.greyColor

    def getRacket(self, racketID: str):

        resp = req.get(utils.url + f"/rackets/{racketID}")

        if resp.status_code == 200:
            data = resp.json()
            return data["racket"]
        else:
            print(resp)

    def getFeedbacks(self, racketID: str):
        
        resp = req.get(utils.url + f"/feedbacks/{racketID}")

        if resp.status_code == 200:
            data = resp.json()
            return data["feedbacks"]
        else:
            print(resp)