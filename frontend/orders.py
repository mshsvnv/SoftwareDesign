import flet as ft
import requests as req

import style
import utils
import feedback

class CustomButton(ft.TextButton):

    def __init__(self, page: ft.Page, i: int):
        super().__init__()

        self.style = style.styleGrey
        self.text = f'{i}'

        self.on_click = lambda _: self.page.go(f'/rackets/{i}')


class Order(ft.Container):

    def __init__(self, page: ft.Page, order: dict):
        super().__init__()

        self.page = page
        self.height = 350
        self.padding = 50

        self.order = order
        self.ink = True
        self.ink_color = ft.colors.RED

        self.racketsID = []
        self.getRackets()

        self.goods = ft.Row()
        
        for i in self.racketsID:
            self.goods.controls.append(
                CustomButton(self.page, i)
            )

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = f"Заказ от {self.order['creation_date'][:10]}",
                            size = 25,
                            weight = ft.FontWeight.BOLD
                        ),
                        ft.Row(
                            controls = [
                                ft.Text(
                                    "оплачено",
                                    size = 16
                                ),
                                ft.Text(
                                    f"{self.order['total_price']} ₽",
                                    size = 16,
                                    weight = ft.FontWeight.BOLD
                                )
                            ]
                        )
                    ]
                ),
                ft.Row(
                    controls = [
                        ft.Text(
                            value = "Статус",
                            size = 18,
                            weight = ft.FontWeight.BOLD
                        ),
                        ft.ElevatedButton(
                            text = "В пути" if self.order["status"] == 'InProgress' else "Получен",
                            style = style.styleGrey,
                            scale = 1.15
                        )
                    ]
                ),
                ft.Text(
                    value = f"Дата доставки {self.order['order_info']['delivery_date'][:10]} в {self.order['order_info']['delivery_date'][11:19]}",
                    size = 16
                ),
                ft.Text(
                    value = f"Получатель {self.order['order_info']['recepient_name']}",
                    size = 16
                ),
                ft.Row(
                    controls = [
                        ft.Text(
                            value = "Товары",
                            size = 18,
                            weight = ft.FontWeight.BOLD
                        ),
                        self.goods
                    ]
                ),
                ft.ElevatedButton(
                    text = "Оценить товар",
                    scale = 1.15,
                    style = style.styleGrey,
                    on_click = self.evaluateRacket
                )
            ]
        )

    def getRackets(self):

        for line in self.order["lines"]:
            self.racketsID.append(line["racket_id"])

    def evaluateRacket(self, e):
        
        dlg = feedback.Feedback(self.page, self.racketsID)

        self.page.open(dlg)

    def print(self, i):
        print(i)

class Orders(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page

        self.padding = 50
        self.orders = None
        self.getOrders()

        if self.orders is None:
            self.content = ft.Column(
                alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                controls = [
                    ft.Text(
                        "Заказов нет", 
                        size = 40
                    ),
                    ft.ElevatedButton(
                        text = "Начать покупки",
                        scale = 1.15,
                        style = style.styleGreen,
                        on_click = lambda _: self.page.go("/rackets")
                    )
                ]
            )
        else:
            orders = []
            for order in self.orders:
                orders.append(
                    Order(
                        page,
                        order
                    ),
                )
                orders.append(ft.Divider())

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
                        controls = orders
                    )
                ]
            )
    

    def getOrders(self):

        if self.page.client_storage.get("token"):
            
            headers = {
                'Authorization': f'Bearer {self.page.client_storage.get("token")}'
            }

            resp = req.get(
                utils.url + f"/api/orders",
                headers = headers,
            )

            if resp.status_code == 200:

                data = resp.json()
                self.orders = data['orders']
            else:
                print(resp)
