import flet as ft
import requests as req

import style
import utils

class CartInfo(ft.Container):

    def __init__(self, page: ft.Page, quantity, total_price):
        super().__init__()

        self.page = page
        self.height = 200
        self.padding = 25
        self.border_radius = 30
        self.bgcolor = style.greyColor
        self.col = {"md": 4}

        self.quantity = ft.Text(
            value = quantity
        )
        self.total_price = ft.Text(
            value = total_price
        )

        self.rowQuantity = ft.Row(
            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
            controls = [
                ft.Text(
                    value = 'Товаров',
                ),
                self.quantity
            ]
        )

        self.rowPrice = ft.Row(
            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
            controls = [
                ft.Text(
                    value = 'Итог',
                    weight = ft.FontWeight.BOLD
                ),
                self.total_price
            ]
        )

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text(
                    value = "Ваша корзина",
                    size = 30,
                    weight = ft.FontWeight.BOLD
                ),
                self.rowQuantity,
                self.rowPrice
            ]
        )

class Order(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()
        
        self.page = page
        self.padding = 50

        self.cart = None
        self.getCart()

        cartInfo = CartInfo(
            self.page, 
            self.cart["quantity"],
            self.cart["total_price"],
        )

        self.address = ft.TextField(
            label = "Адрес",
        )

        self.recepient = ft.TextField(
            label = "Получатель",
        )

        self.date = ft.TextField(
            label = "Дата получения",
            read_only = True
        )

        self.time = ft.TextField(
            label = "Время получения",
            read_only = True
        )

        def changeDate(e):
            self.date.value = datePicker.value.strftime('%Y-%m-%d')
            self.update()        
        
        def dismissDate(e):
            self.date.value = datePicker.value.strftime('%Y-%m-%d')
            self.update()

        def changeTime(e):
            self.time.value = timePicker.value
            self.update()        
        
        def dismissTime(e):
            self.time.value = timePicker.value
            self.update()

        datePicker = ft.DatePicker(
            confirm_text = "Подтвердить",
            cancel_text = "Отменить",
            help_text = "Выберите дату получения",
            on_change = changeDate,
            on_dismiss = dismissDate
        )

        timePicker = ft.TimePicker(
            confirm_text = "Подтвердить",
            cancel_text = "Отменить",
            help_text = "Выберите время получения",
            on_change = changeTime,
            on_dismiss = dismissTime
        )

        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
            controls = [
                ft.Text(
                    "Оформление заказа", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Container(
                                height = 500,
                                col = {"md": 6},
                                content = ft.Column(
                                    alignment = ft.MainAxisAlignment.SPACE_AROUND,
                                    controls = [
                                        ft.Row(
                                            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                                            controls = [
                                                self.date,
                                                ft.FilledButton(
                                                    text = "Дата получения",
                                                    style = style.styleGrey,
                                                    icon = ft.icons.CALENDAR_MONTH,
                                                    on_click = lambda e: self.page.open(
                                                        datePicker
                                                    )
                                                )
                                            ]
                                        ),
                                        ft.Row(
                                            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                                            controls = [
                                                self.time,
                                                ft.FilledButton(
                                                    text = "Время получения",
                                                    style = style.styleGrey,
                                                    icon = ft.icons.TIME_TO_LEAVE,
                                                    on_click = lambda e: self.page.open(
                                                        timePicker
                                                    )
                                                ),
                                            ]
                                        ),
                                        self.address,
                                        self.recepient,
                                        ft.Row(
                                            alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                                            controls = [
                                                ft.FilledButton(
                                                    text = "Подтвердить",
                                                    style = style.styleGreen,
                                                    on_click = self.createOrder
                                                ),
                                                ft.FilledButton(
                                                    text = "Отменить",
                                                    style = style.styleOrange,
                                                    on_click = lambda _: self.page.go("/api/cart")
                                                )
                                            ]
                                        )
                                    ]
                                ),
                            ),
                            cartInfo
                        ],
                    ),
                )
            ]
        )

    def getCart(self):

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.get(
            utils.url + '/api/cart',
            headers = headers
        )

        if resp.status_code == 200:
            
            data = resp.json()

            self.cart = data["cart"]
        else:
            print(resp)

    def createOrder(self, e):

        time = str(self.time.value)
        deliveryDate = f"{self.date.value}T{time}Z"

        data = {
            "order_info" : {
                "delivery_date": deliveryDate,
                "address": self.address.value,
                "recepient_name": self.recepient.value,
            }
        }

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.post(
            utils.url + '/api/order',
            headers = headers,
            json = data 
        )

        if resp.status_code == 200:
            self.page.go("/api/cart")