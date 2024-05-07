# ППО aka Проектирование программного обеспечения

## Название проекта
Интернет-магазин теннисных ракеток

## Краткое описание идеи проекта
Интернет-магазин специализируется на продаже высококачественных теннисных ракеток, предлагая широкий ассортимент моделей для любого уровня игры и предпочтений. На сайте магазина представлена исчерпывающяя информация о каждой модели: ее характеристики и возможные модификации. Есть возможность просмотра каталога, заказа товара, а также подписки на рассылку.

## Краткое описание предметной области
Теннисная ракетка - спортивное снаряжение, используемое для ударов по теннисному мячу. Она обычно состоит из ракеточной рамы, ручки и натянутой сетки. Ракетки могут быть различной длины, веса, формы и материалов, что позволяет игрокам выбирать подходящую ракетку в соответствии со своим стилем игры и предпочтениями.

## Краткий анализ аналогичных решений
| Название | Наличие истории заказов | Возможность подписаться на рассылку писем | Возможность получать бонусные баллы | 
|-------------|---|---|---|
| [Ракетлон](https://racketlon.ru/)    | + | - | + |
| [Tennis-Store](https://tennis-store.ru/)| + | - | - |
| [Мир-ракеток](https://www.mirraketok.ru/) | + | + | - |
| Предлагаемое решение  | + | + | + |

## Краткое обоснование целесообразности и актуальности проекта
Актуальность проекта интернет-магазина теннисных ракеток может быть обоснована несколькими факторами:
- теннис является популярным видом спорта с множеством поклонников по всему миру, что создает спрос на качественные теннисные ракетки;
- интернет-магазины обладают преимуществами в удобстве покупки, доступности широкого ассортимента товаров и возможности сравнения цен и характеристик различных товаров.

## Краткое описание акторов
|Роль|Описание |
|--|--|
|**Гость**|пользователь, который посещает интернет-магазин без создания учетной записи или входа в систему. Он может просматривать продукты, но не может совершать покупки или получать доступ к каким-либо персонализированным функциям.|
|**Авторизованный клиент**|пользователь, создавший учетную запись и вошедший в интернет-магазин. У них есть доступ к дополнительным функциям, таким как добавление товаров в корзину, совершение покупок, отслеживание статуса заказа и управление данными своей учетной записи.|
|**Продавец**|пользователь, который управляет интернет-магазином и управляет запасами. У них есть доступ к таким функциям, как добавление новых продуктов, обновление информации о продуктах, управление ценами и скидками, обработка заказов.|
|**Администратор**|пользователь с повышенными привилегиями, который контролирует общую работу интернет-магазина. У них есть доступ ко всем функциям и настройкам, включая управление учетными записями пользователей и продавцов. Администраторы также имеют возможность изменять дизайн и макет сайта, а также настраивать функциональность магазина.|
   
## Use-Case - диаграмма
![Диаграмма использования приложения](./schemes/svg/use-case.svg) 

## ER-диаграмма сущностей
![ER-модель в нотации Чена](./schemes/svg/ER-roles.svg)  

## Пользовательские сценарии
1. Сценарий просмотра каталога
   - пользователь заходит в систему;
   - при желении он может либо зарегестироваться, либо авторизоваться;
   - переходит на вкладку "Каталог";
   - просматривает содержимое страницы.
2. Сценарий входа в личный кабинет
   - пользователь заходит в систему;
   - при желении он может либо зарегестироваться, либо авторизоваться;
   - для регистрации вводятся необходимые данные; для авторизации - логин и пароль;
   - данные проверяются на корректность;
   - в случае успешной проверки, пользователь либо зарегистрирован, либо авторизован.
3. Сценарий формирования корзины
   - пользователь заходит в систему;
   - переходит на вкладку "Каталог";
   - просматривает содержимое страницы;
   - выбирает понравившийся товар;
   - для добавления товара в корзину пользователь должен быть авторизован;
   - заполняет необходимую информацию о товаре;
   - информация о товаре отображается в корзине.
4. Сценарий формирования заказа
   - пользователь заходит в личный кабинет;
   - переходит на вкладку "Корзина";
   - выбирает товар из корзины для заказа;
   - переходит на вкладку "Формирование заказа" и вводит необходимые данные;
   - оплачивает заказ.

## Формализация ключевых бизнес-процессов
![](./schemes/svg/BPMN1.svg)
![](./schemes/svg/BPMN2.svg)  

## Тип приложения
Web SPA

## Технологический стек
- Backend: Go
- Frontend: HTML + CSS + JS (Vue)
- Database: PostgreSQL
  
## Верхнеуровневое разбиение на компоненты
Приложение будет состоять из 3 компонентов:
- компонент реализации UI
- Компонент реализации бизнес-логики (Business Logic)
- Компонент доступа к данным (Data Access)
  
![](./schemes/svg/top_level_components.svg)  


## Диаграмма классов (бизнес логикой + доступ к данным)
![](./schemes/svg/UML.svg)  