import flet as ft
import requests as req
import utils
import style

class Login(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()
        
        self.page = page

        self.border = ft.border.all(2, ft.colors.GREY)
        self.border_radius = ft.border_radius.all(15)

        self.width = 500
        self.height = 500

        self.padding = 50

        self.email = ft.TextField(label = "Почта")
        self.password = ft.TextField(
            label = "Пароль", 
            password = True, 
            can_reveal_password = True
        )
        
        self.content = ft.Column(
            controls = [
                ft.Text("Авторизация"),
                self.email,
                self.password,
                ft.Row(
                    controls = [
                        ft.Text("Нет аккаунта?"),
                        ft.TextButton(
                            text = "Зарегестрироваться", 
                            on_click = lambda _: self.page.go("/auth/register"),
                            on_hover = False
                        ),
                        ft.FilledButton(
                            text = "Войти",
                            on_click = self.login,
                            style = style.style
                        )
                    ]
                )
            ],
            alignment = ft.MainAxisAlignment.CENTER,
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

        self.border = ft.border.all(2, ft.colors.GREY)
        self.border_radius = ft.border_radius.all(15)

        self.width = 500
        self.height = 500

        self.padding = 50

        self.name = ft.TextField(label = "Имя")
        self.surname = ft.TextField(label = "Фамилия")

        self.email = ft.TextField(label = "Почта")
        self.password = ft.TextField(
            label = "Пароль", 
            password = True, 
            can_reveal_password = True
        )

        self.content = ft.Column(
            controls = [
                ft.Text("Регистрация"),
                self.name,
                self.surname,
                self.email,
                self.password,
                ft.Row(
                    expand = True,
                    controls = [
                        ft.Text("Есть аккаунта?"),
                        ft.TextButton(
                            text = "Войти", 
                            on_click = lambda _: self.page.go("/auth/login")
                        ),
                        ft.FilledButton(
                            text = "Зарегестрироваться",
                            on_click = self.register,
                            style = style.style
                        )
                    ]
                )
            ],
            alignment = ft.MainAxisAlignment.CENTER,
        )

    def register(self, e):
        print("login")
        self.page.go("/api/profile")

