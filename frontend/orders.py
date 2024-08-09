import flet as ft
import requests as req

import style
import utils
import feedback

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
                        ft.Text(f"отплачено {self.order['total_price']} ₽")
                    ]
                ),
                ft.Row(
                    controls = [
                        ft.Text("Статус"),
                        ft.FilledButton(
                            text = "В пути" if self.order["status"] == 'InProgress' else "Получен",
                            style = style.styleGrey,
                        )
                    ]
                ),
                ft.Text(f"Дата доставки {self.order['order_info']['delivery_date'][:10]} в {self.order['order_info']['delivery_date'][11:19]}"),
                ft.Text(f"Получатель {self.order['order_info']['recepient_name']}"),
                ft.Row(
                    controls = [
                        ft.Text("Товары"),
                        ft.Text(
                            spans = [
                                ft.TextSpan(
                                    text = f"{id}",
                                    on_click = lambda _: self.page.go(f"/rackets/{id}")
                                ) for id in self.racketsID
                            ]
                        )
                    ]
                ),
                ft.FilledButton(
                    text = "Оценить товар",
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
                    ft.FilledButton(
                        text = "Начать покупки",
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
