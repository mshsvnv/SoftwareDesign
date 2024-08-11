import flet as ft
import requests as req

import style
import utils

class Feedback(ft.AlertDialog):

    def __init__(self, page: ft.Page, racketsID):
        super().__init__()

        self.page = page

        self.title = ft.Text(
            value = "Напишите отзыв к покупке",
            size = 20
        )

        self.racketID = ft.Dropdown(
            options = [
                ft.dropdown.Option(f"{id}") for id in racketsID
            ]
        )

        self.feedback = ft.TextField(
            label = "Отзыв",
            multiline = True

        )
        self.rating = ft.Dropdown(
            options = [
                ft.dropdown.Option(f"{i + 1}") for i in range(5)
            ]
        )

        self.actions_alignment = ft.MainAxisAlignment.SPACE_AROUND
        
        self.content = ft.Container(
            width = 700,
            height = 300,
            content = ft.Column(
                alignment = ft.MainAxisAlignment.SPACE_AROUND,
                controls = [
                    ft.Row(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Text(
                                value = "Ракетка",
                                size = 18    
                            ),
                            self.racketID
                        ]
                    ),
                    ft.Row(
                        alignment = ft.MainAxisAlignment.SPACE_BETWEEN,
                        controls = [
                            ft.Text(
                                value = "Оценка",
                                size = 18
                            ),
                            self.rating
                        ]
                    ),
                    self.feedback,
                    ft.Row(
                        alignment = ft.MainAxisAlignment.END,
                        controls = [
                            ft.ElevatedButton(
                                scale = 1.15,
                                text = "Оценить",
                                style = style.styleGreen,
                                on_click = self.makeFeedback
                            )
                        ]
                    )
                ]
            )
        )

    def makeFeedback(self, e):

        data = {
            "racket_id": int(self.racketID.value),
            "feedback": self.feedback.value,
            "rating": int(self.rating.value)
        }
        
        headers = {
            'Authorization': f'Bearer {self.page.client_storage.get("token")}'
        }

        resp = req.post(
            utils.url + '/api/feedback',
            headers = headers,
            json = data
        )

        return resp.status_code

        
