import flet as ft
import requests as req
import utils
import style


class Cart(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50

        self.cart = None
        self.getCart()

        if self.cart is None or self.cart['total_price'] == 0:
            self.content = ft.Column(
                alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                controls = [
                    ft.Text(
                        "Корзина пуста", 
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
            rackets = []
            cartInfo = CartInfo(
                self.page, 
                self.cart["quantity"],
                self.cart["total_price"],
            )

            for line in self.cart["lines"]:
                rackets.append(RacketInfo(line, cartInfo, self, self.page))
                rackets.append(ft.Divider())

            self.content = ft.Column(
                horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
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
                                        # controls = [
                                        #     RacketInfo(self.page),
                                        #     ft.Divider(),
                                        #     RacketInfo(self.page),
                                        #     ft.Divider()
                                        # ]
                                        controls = rackets
                                    )
                                ),
                                cartInfo
                            ],
                        ),
                    )
                ]
            )

    def deleteRacket(self, cart):

        self.cart = cart
        self.setContent()

        self.update()

    def getCart(self):

        if self.page.client_storage.get("token"):
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

    def setContent(self):

        if self.cart is None or self.cart['total_price'] == 0:
            self.content = ft.Column(
                # horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                controls = [
                    ft.Text(
                        "Корзина пуста", 
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
            rackets = []
            cartInfo = CartInfo(
                self.page, 
                self.cart["quantity"],
                self.cart["total_price"],
            )

            for line in self.cart["lines"]:
                rackets.append(RacketInfo(line, cartInfo, self, self.page))
                rackets.append(ft.Divider())

            self.content = ft.Column(
                horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
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
                                        # controls = [
                                        #     RacketInfo(self.page),
                                        #     ft.Divider(),
                                        #     RacketInfo(self.page),
                                        #     ft.Divider()
                                        # ]
                                        controls = rackets
                                    )
                                ),
                                cartInfo
                            ],
                        ),
                    )
                ]
            )
    
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
                self.rowPrice,
                ft.FilledButton(
                    text = "Оформить заказ",
                    style = style.styleGreen,
                    on_click = lambda _: self.page.go("/api/order")
                )
            ]
        )

    def setNewInfo(self, total_price, quantity):

        self.quantity.value = quantity
        self.total_price.value = total_price

        self.update()


class RacketInfo(ft.Container):

    def __init__(
            self, 
            racket: dict, 
            cartInfo: CartInfo,
            cart: Cart,
            page: ft.Page
    ):
        super().__init__()

        self.page = page
        self.height = 200

        self.cartInfo = cartInfo
        self.racket = racket
        self.cart = cart

        self.quantity = ft.Text(
            value = f"{self.racket['quantity']}",
        )
        self.addButton = ft.FilledButton(
            text = "+",
            style = style.styleGreen,
            on_click = self.addRacket
        )

        self.deleteButton = ft.FilledButton(
            text = "-",
            style = style.styleGreen,
            on_click = self.deleteRacket
        )

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            spans = [
                                ft.TextSpan(
                                    text = f"Код: {self.racket['id']}",
                                    on_click = lambda _: self.page.go(f"/rackets/{self.racket['id']}")
                                )
                            ]
                        ),
                        ft.Row(
                            controls = [
                                ft.FilledButton(
                                    text = "-",
                                    style = style.styleGreen,
                                    on_click = self.deleteRacket
                                ),
                                self.quantity,
                                ft.FilledButton(
                                    text = "+",
                                    style = style.styleGreen,
                                    on_click = self.addRacket
                                )
                            ]
                        )
                    ]
                ),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.Text(
                            value = f"Цена: {self.racket['price']}",
                            size = 15,
                        ),
                        ft.FilledButton(
                            text = "Удалить",
                            style = style.styleOrange,
                            on_click = self.deleteRacketAll
                        )
                    ]
                ),
            ]
        )

    def addRacket(self, e):

        if self.deleteButton.disabled:
            self.deleteButton.disabled = False

        data = {
            "quantity": 1
        }

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.put(
            utils.url + f"/api/cart/{self.racket['id']}",
            headers = headers,
            json = data
        )

        # TODO
        if resp.status_code == 200:

            data = resp.json()
            cart = data["cart"]

            self.racket['quantity'] += 1
            self.quantity.value = f"{self.racket['quantity']}"

            self.cartInfo.setNewInfo(cart['total_price'], cart['quantity'])

            self.update()
        else:
            print(resp)

    def deleteRacket(self, e):

        if self.racket['quantity'] == 1:
            self.deleteButton.disabled = True
            self.update()

            return

        data = {
            "quantity": -1
        }

        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.put(
            utils.url + f"/api/cart/{self.racket['id']}",
            headers = headers,
            json = data
        )

        if resp.status_code == 200:

            data = resp.json()
            cart = data["cart"]

            self.racket['quantity'] -= 1
            self.quantity.value = f"{self.racket['quantity']}"

            self.cartInfo.setNewInfo(cart['total_price'], cart['quantity'])
            
            self.update()
        else:
            print(resp)

    def deleteRacketAll(self, e):

        def handleClose(e):
            self.page.close(dlg)

        def handleOpen(e):

            self.page.close(dlg)

            headers = {
                'Authorization': f'Bearer {self.page.client_storage.get("token")}'
            }

            resp = req.delete(
                utils.url + f"/api/cart/{self.racket['id']}",
                headers = headers,
            )

            if resp.status_code == 200:

                data = resp.json()
                cart = data["cart"]

                self.cartInfo.setNewInfo(cart['total_price'], cart['quantity'])
                self.cart.deleteRacket(cart)
            else:
                print(resp)

        dlg = ft.AlertDialog(
            title = ft.Text(
                value = "Удалить товар",
                style = ft.FontWeight.BOLD
            ),
            content = ft.Text("Вы действительно хотите удалить эту ракетку?"),
            actions = [
                ft.FilledButton(
                    text = "Да", 
                    on_click = handleOpen,
                    style = style.styleGreen
                ),
                ft.FilledButton(
                    text = "Нет", 
                    on_click = handleClose,
                    style = style.styleOrange
                ),
            ]
        )

        self.page.open(dlg)
    
        