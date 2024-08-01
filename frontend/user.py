import flet as ft
import requests as req
import utils
import style

class UserInfo(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.height = 200
        # self.col = {"md": 6}

        self.content = ft.Column(
            alignment = ft.MainAxisAlignment.SPACE_AROUND,
            controls = [
                ft.Text("Имя"),
                ft.Text("Фамилия"),
                ft.Text("Почта"),
                ft.Row(
                    alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                    controls = [
                        ft.FilledButton(
                            text = "Сменить пароль",
                            style = style.styleGreen,
                        ),
                        ft.FilledButton(
                            text = "Выйти",
                            style = style.styleOrange,
                        ),
                    ]
                )

                # ft.Row(
                #     alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                #     controls = [
                #         ft.Text("Название"),
                #         ft.Row(
                #             controls = [
                #                 ft.FilledButton(
                #                     text = "-",
                #                     style = style.styleGreen
                #                 ),
                #                 ft.Text(),
                #                 ft.FilledButton(
                #                     text = "+",
                #                     style = style.styleGreen
                #                 )
                #             ]
                #         )
                #     ]
                # ),
                # ft.Row(
                #     alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                #     controls = [
                #         ft.Text("Цена"),
                #         ft.FilledButton(
                #             text = "Удалить",
                #             style = style.styleOrange
                #         )
                #     ]
                # ),
            ]
        )


class Profile(ft.Container):

    def __init__(self, page: ft.Page):
        super().__init__()

        self.page = page
        self.padding = 50

        self.content = ft.Column(
            horizontal_alignment = ft.CrossAxisAlignment.CENTER,
            controls = [
                ft.Text(
                    "Мой профиль", 
                    size = 40,
                    text_align = ft.TextAlign.CENTER
                ),
                ft.Container(
                    padding = 25,
                    content = ft.ResponsiveRow(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Container(
                                # col = {"md": 6},
                                content = ft.Column(
                                    controls = [
                                        UserInfo(self.page),
                                        ft.Divider()
                                    ]
                                )
                            ),
                        ],
                    ),
                )
            ]
        )

    #     self.data = ft.Column()

    #     self.getProfile()
    #     self.content = self.data

    # def getProfile(self):

    #     url = utils.url + "/api/profile"
    #     headers = {
    #         'Authorization': "Bearer " + self.page.client_storage.get("token")
    #     }
    #     resp = req.get(
    #         url = url,
    #         headers = headers
    #     )

    #     if resp.status_code == 200:
    #         data = resp.json()
            
    #         self.data.controls.append(
    #             ft.Text(
    #                 "Mои данные",
    #                 size = 40,
    #                 weight = ft.FontWeight.W_100,
    #             )
    #         )
    #         self.data.controls.append(
    #             ft.Container(
    #                 height = 300,
    #                 width = 500,
    #                 padding = 20,
    #                 border_radius = ft.border_radius.all(15),
                
    #                 bgcolor = ft.colors.GREY_200,
    #                 content = ft.Column(
    #                     controls = [
    #                         ft.Text("Имя: " + data["user"]["name"]),
    #                         ft.Text("Фамилия: " + data["user"]["surname"]),
    #                         ft.Text("Почта: " + data["user"]["email"]),
                            
    #                         ft.Row(
    #                             controls = [
    #                                 ft.FilledButton(
    #                                     text = "Поменять пароль",
    #                                     style = style.style
    #                                 ),
    #                                 ft.FilledButton(
    #                                     text = "Выйти",
    #                                     style = style.style1
    #                                 )
    #                             ]
    #                         )
    #                     ],
    #                     alignment = ft.MainAxisAlignment.SPACE_EVENLY
    #                 )
    #             )
    #         )