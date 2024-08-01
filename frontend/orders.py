import flet as ft
import style

class Order(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        # self.page = page
        self.height = 200

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text("Дата"),
                        ft.Text("Цена"),
                    ]
                ),
                ft.Text("Доставка"),
                ft.Text("Товары"),
                ft.FilledButton(
                    text = "Оценить товар",
                    style = style.styleGrey

                )
            ]
        )

class Orders(ft.Container):

    def __init__(self):
        super().__init__()

        self.padding = 50
        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Мои заказы", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.ResponsiveRow(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        Order(None),
                        ft.Divider(),
                        Order(None),
                        ft.Divider()
                    ]
                )
            ]
        )
