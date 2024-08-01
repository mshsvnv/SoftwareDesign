import flet as ft
import requests as req
import utils
import style

class Login(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()
        
        self.page = page

        self.border = ft.border.all(3, ft.colors.GREY)
        self.border_radius = ft.border_radius.all(30)

        self.width = 500
        self.height = 400
        self.padding = 50
        self.margin = 50

        self.email = ft.TextField(
            label = "Почта",
            border_radius = 30
        )
        self.password = ft.TextField(
            label = "Пароль", 
            password = True, 
            can_reveal_password = True,
            border_radius = 30
        )
        
        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text(
                    "Авторизация", 
                    size = 30,
                    text_align = ft.TextAlign.CENTER
                ),
                self.email,
                self.password,
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Row(
                            controls = [
                                ft.Text("Нет аккаунта?"),
                                ft.Text(
                                    spans = [
                                        ft.TextSpan(
                                            text = "Зарегестрироваться",
                                            on_click = lambda _: self.page.go("/auth/register"),
                                            style = ft.TextStyle(
                                                weight = ft.FontWeight.BOLD
                                            )
                                        )
                                    ]
                                ),
                            ]
                        ),
                        ft.FilledButton(
                            text = "Войти",
                            on_click = self.login,
                            style = style.styleGreen
                        )
                    ]
                )
            ],
        )

    def login(self, e):
        data = {
            "email": self.email.value,
            "password": self.password.value
        }

        resp = req.post(utils.url + "/auth/login", json = data)

        if resp.status_code == 200:
            data = resp.json()

            self.page.client_storage.set("token", data["token"])
            self.page.go("/api/profile")
        else:
            print(resp)

class Register(ft.Container):

    def __init__(self, page: ft.Page):

        super().__init__()

        self.page = page

        self.border = ft.border.all(3, ft.colors.GREY)
        self.border_radius = ft.border_radius.all(30)

        self.width = 500
        self.height = 500
        self.padding = 50
        self.margin = 50

        self.name = ft.TextField(
            label = "Имя",
            border_radius = 30
        )
        self.surname = ft.TextField(
            label = "Фамилия",
            border_radius = 30
        )

        self.email = ft.TextField(
            label = "Почта",
            border_radius = 30
        )

        self.password = ft.TextField(
            label = "Пароль", 
            password = True, 
            can_reveal_password = True,
            border_radius = 30
        )

        self.content = ft.Column(
            horizontal_alignment= ft.CrossAxisAlignment.CENTER,
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text(
                    "Регистрация", 
                    size = 25,
                    text_align = ft.TextAlign.END
                ),
                self.name,
                self.surname,
                self.email,
                self.password,
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Row(
                            controls = [
                                ft.Text("Есть аккаунт?"),
                                ft.Text(
                                    spans = [
                                        ft.TextSpan(
                                            text = "Войти",
                                            on_click = lambda _: self.page.go("/auth/login"),
                                            style = ft.TextStyle(
                                                weight = ft.FontWeight.BOLD
                                            )
                                        )
                                    ]
                                ),
                            ]
                        ),
                        ft.FilledButton(
                            text = "Зарегестрироваться",
                            on_click = self.register,
                            style = style.styleGreen
                        )
                    ]
                )
            ],
        )

    def register(self, e):
        data = {
            "name": self.name.value,
            "surname": self.surname.value,
            "email": self.email.value,
            "password": self.password.value
        }

        resp = req.post(utils.url + "/auth/register", json = data)

        if resp.status_code == 200:
            data = resp.json()

            self.page.client_storage.set("token", data["token"])
            self.page.go("/api/profile")
        else:
            print(resp)

