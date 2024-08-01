import flet as ft
import style

class CartInfo(ft.Container):

    def __init__(self, page: ft.Page):
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
                    value = "Цена",
                    size = 20,
                    weight = ft.FontWeight.BOLD
                ),
                ft.FilledButton(
                    text = "В корзину",
                    style = style.styleGreen
                )
            ]
        )

class RacketInfo(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 200
        self.col = {"md": 6}

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text("Фирма"),
                ft.Text("Вес"),
                ft.Text("Баланс"),
                ft.Text("Размер головы")
            ]
        )

class Feedback(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 200

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Имя"),
                        ft.Text("Дата"),
                    ]
                ),
                ft.Text("Текст отзыва")
            ]
        )

class Racket(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50

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
                            RacketInfo(self.page),
                            CartInfo(self.page),
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
                        controls = [
                            Feedback(page),
                            ft.Divider(),
                            Feedback(page),
                            ft.Divider(),
                        ]
                    ),
                )
            ]
        )

        self.ink = True
        self.ink_color = style.greyColor