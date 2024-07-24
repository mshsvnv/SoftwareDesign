import flet as ft

url = "http://localhost:8080"
token = ""

bar = ft.AppBar(
    leading = ft.IconButton(ft.icons.SPORTS_TENNIS, scale = 1.5),
    # leading_width = 100,
    title = ft.Text("Racket Shop", size = 25),
    actions = [
        ft.IconButton(ft.icons.SHOPPING_CART, scale = 1.5),
        ft.IconButton(ft.icons.MAN, scale = 1.5)
    ],
)