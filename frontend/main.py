import flet as ft
import auth
import racket
import user

import utils

def main(page: ft.Page):

    bar = ft.AppBar(
        leading = ft.IconButton(ft.icons.SPORTS_TENNIS, scale = 1.5, on_click = lambda _: page.go("/rackets")),
        title = ft.Text("Racket Shop", size = 25),
        actions = [
            ft.IconButton(ft.icons.SHOPPING_CART, scale = 1.5, on_click = lambda _: page.go("/api/cart")),
            ft.Icon(),
            ft.IconButton(ft.icons.MAN, scale = 1.5, on_click = lambda _: page.go("/api/profile") if page.client_storage.get("token") else page.go("/auth/login"))
        ],
        leading_width = 100
    )

    def route_change(route):

        page.views.clear()

        if page.route == "/rackets":
            page.views.append(
                ft.View(
                    route = "/rackets",
                    controls = [
                        racket.Rackets(page)
                    ],
                    appbar = bar
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
                )
            )

        if page.route == "/api/profile":
            page.views.append(
                ft.View(
                    route = "/api/profile",
                    controls = [
                        user.Profile(page)
                    ],
                    vertical_alignment = ft.CrossAxisAlignment.CENTER,
                    # horizontal_alignment = ft.CrossAxisAlignment.CENTER,
                    appbar = bar,
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

ft.app(target=main, view=ft.AppView.WEB_BROWSER)