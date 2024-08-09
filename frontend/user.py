import flet as ft
import requests as req
import utils
import style

class UserInfo(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 300
        self.padding = 25

        self.name = None
        self.surname = None
        self.email = None

        self.getUser()

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = "Имя",
                            size = 15
                        ),
                        ft.Text(
                            value = self.name,
                            size = 15
                        )
                    ]
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = "Фамилия",
                            size = 15
                        ),
                        ft.Text(
                            value = self.surname,
                            size = 15
                        )
                    ]
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = "Почта",
                            size = 15
                        ),
                        ft.Text(
                            value = self.email,
                            size = 15
                        )
                    ]
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.END,
                    controls = [ 
                        ft.FilledButton(
                            text = "Выйти",
                            style = style.styleOrange,
                            on_click = self.exit
                        ),
                    ]
                )
            ]
        )

    def getUser(self):

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.get(utils.url + '/api/profile', headers = headers)

        if resp.status_code == 200:
            data = resp.json()

            self.name = data["user"]["name"]
            self.surname = data["user"]["surname"]
            self.email = data["user"]["email"]
        else:
            # error
            print(resp)

    def exit(self, e):

        self.page.client_storage.remove("token")
        self.page.go("/rackets")

class Feedback(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 200
        self.padding = 25

        feedbacks = self.getFeedbacks()

        if feedbacks is None:
            self.content = ft.Text(
                value = "Отзывов нет",
                size = 15,
                text_align = ft.TextAlign.LEFT
            )
        else:
            controls = []
            for feedback in feedbacks:
                controls.append(
                    ft.Row(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Text(
                                spans = [
                                    ft.TextSpan(
                                        text = f"Ракетка {feedback['racket_id']}",
                                        style = ft.TextStyle(
                                            size = 25,
                                            weight = ft.FontWeight.BOLD
                                        ),
                                        on_click = lambda _: self.page.go(f"/rackets/{feedback['racket_id']}")
                                    )
                                ]
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
                                        ) for _ in range(int(feedback['rating']))
                                    ]
                                ]
                            ),
                        ]
                    ),
                )

                controls.append(
                    ft.Row(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Text(feedback["feedback"]),
                            ft.FilledButton(
                                text = "Удалить",
                                style = style.styleOrange,
                                on_click = lambda _: self.deleteFeedback(feedback['racket_id']),
                            )
                        ]
                    )
                )
            
                self.content = ft.Column(
                    alignment = ft.MainAxisAlignment.SPACE_AROUND,
                    controls = controls
                )

    def getFeedbacks(self):

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.get(
            utils.url + '/api/feedbacks', 
            headers = headers
        )

        if resp.status_code == 200:
            data = resp.json()
            return data["feedbacks"]
        else:
            # error
            print(resp)

    def deleteFeedback(self, racketID: int):

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.delete(
            utils.url + f'/api/feedback/{racketID}',
            headers = headers,
        )

        print(resp)

class Profile(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50

        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Мой профиль", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Container(
                                content = ft.Column(
                                    controls = [
                                        UserInfo(self.page),
                                        ft.Divider()
                                    ]
                                )
                            ),
                        ],
                    ),
                ),
                ft.Text(
                    "Мои отзывы", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                Feedback(page)
            ]
        )
