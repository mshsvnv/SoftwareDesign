import flet as ft
import requests as req
import utils
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
                    value = "Товаров",
                    size = 20,
                ),
                ft.Text(
                    value = "Итог",
                    size = 20,
                    weight = ft.FontWeight.BOLD
                ),
                ft.FilledButton(
                    text = "Оформить заказ",
                    style = style.styleGreen
                )
            ]
        )


class RacketInfo(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 200
        # self.col = {"md": 6}

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Название"),
                        ft.Row(
                            controls = [
                                ft.FilledButton(
                                    text = "-",
                                    style = style.styleGreen
                                ),
                                ft.Text(),
                                ft.FilledButton(
                                    text = "+",
                                    style = style.styleGreen
                                )
                            ]
                        )
                    ]
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Цена"),
                        ft.FilledButton(
                            text = "Удалить",
                            style = style.styleOrange
                        )
                    ]
                ),
            ]
        )

class Cart(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50
        
        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Корзина", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Container(
                                col = {"md": 6},
                                content = ft.Column(
                                    controls = [
                                        RacketInfo(self.page),
                                        ft.Divider(),
                                        RacketInfo(self.page),
                                        ft.Divider()
                                    ]
                                )
                            ),
                            CartInfo(self.page),
                        ],
                    ),
                )
            ]
        )
    
        