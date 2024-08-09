import flet as ft

import auth
import rackets
import racket
import orders
import user
import cart
import order
import feedback

def main(page: ft.Page):

    bar = ft.AppBar(
        leading = ft.IconButton(ft.icons.SPORTS_TENNIS, scale = 1.5, on_click = lambda _: page.go("/rackets")),
        title = ft.Text("Racket Shop", size = 25),
        actions = [
            ft.IconButton(
                ft.icons.SHOPPING_CART, 
                scale = 1.5, 
                tooltip = "Корзина",
                on_click = lambda _: page.go("/api/cart")
            ),
            ft.Icon(),
            ft.IconButton(
                ft.icons.ALL_INBOX, 
                scale = 1.5, 
                tooltip = "Заказы",
                on_click = lambda _: page.go("/api/orders")
            ),
            ft.Icon(),
            ft.IconButton(
                ft.icons.ACCOUNT_BOX, 
                tooltip = "Войти",
                scale = 1.5, 
                on_click = lambda _: page.go("/api/profile") if page.client_storage.get("token") else page.go("/auth/login")
            ),
            ft.Icon(),
        ],
        leading_width = 100,
        toolbar_height = 80
    )

    def route_change(route):

        page.views.clear()

        if page.route == "/rackets":
            page.views.append(
                ft.View(
                    route = "/rackets",
                    controls = [
                        rackets.Rackets(page)
                    ],
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                    vertical_alignment = ft.MainAxisAlignment.CENTER
                )
            )

        if "/rackets/" in page.route:
            racketID = page.route.split("/")[-1]

            page.views.append(
                ft.View(
                    route = page.route,
                    controls = [
                        racket.Racket(page, racketID),
                    ],
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                    vertical_alignment = ft.MainAxisAlignment.CENTER
                )
            )

        if page.route == "/auth/login":
            page.views.append(
                ft.View(
                    route = "/auth/login",
                    controls = [
                        auth.Login(page)
                    ],
                    vertical_alignment = ft.CrossAxisAlignment.CENTER,
                    horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                )
            )

        if page.route == "/auth/register":
            page.views.append(
                ft.View(
                    route = "/auth/register",
                    controls = [
                        auth.Register(page)
                    ],
                    vertical_alignment = ft.CrossAxisAlignment.CENTER,
                    horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                )
            )
        
        if page.route == "/api/orders":
            page.views.append(
                ft.View(
                    route = "/api/orders",
                    controls = [
                        orders.Orders(page)
                    ],
                    vertical_alignment = ft.CrossAxisAlignment.CENTER,
                    # horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                )
            )

        if page.route == "/api/profile":
            page.views.append(
                ft.View(
                    route = "/api/profile",
                    controls = [
                        user.Profile(page)
                    ],
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                    vertical_alignment = ft.MainAxisAlignment.CENTER
                )
            )

        if page.route == "/api/cart":
            page.views.append(
                ft.View(
                    route = "/api/cart",
                    controls = [
                        cart.Cart(page)
                    ],
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                    vertical_alignment = ft.MainAxisAlignment.CENTER
                )
            )

        if page.route == "/api/order":
            page.views.append(
                ft.View(
                    route = "/api/order",
                    controls = [
                        order.Order(page)
                    ],
                    appbar = bar,
                    scroll = ft.ScrollMode.ADAPTIVE,
                    vertical_alignment = ft.MainAxisAlignment.CENTER
                )
            )

        page.update()

    # def view_pop(view):
    #     page.views.pop()
    #     top_view = page.views[-1]
    #     page.go(top_view.route)

    page.on_route_change = route_change
    # page.on_view_pop = view_pop
    page.go("/rackets")

ft.app(target = main, view = ft.AppView.WEB_BROWSER)
