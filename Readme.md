Написать web -приложение (или аналогичное desktop приложение) на свой выбор с использованием front-end и back-end.
Приложение должно использовать базу данных из нескольких связанных таблиц (СУБД - PosgreSQL)
и отображать данные в табличной форме, с возможностью добавления, редактирования и удаления данных.
В front-end должны использоваться - JavaScript, CSS и ReactJS.
В back-end должны использоваться запросы к базе данных.
Поставить приложение под контроль версий (GIT)
Выслать ссылку на репозиторий


В приложении используется тематика Салон красоты, понятия Клиент, Услуга, Заказ


План выполнения

1. Postgres 
    + инит, 3 таблицы
    - формат времени
2. back 
    + crud на 3 таблицы (клиент, вид услуг, заказ) (back/internal/service/postgres)
    + бизнес логика (back/internal/core)
        + клиент - имя
        + вид услуги - название и стоимость
        + заказ - связь между клиентом, запрошенной услугой и временем оказания услуги
    + эндпоинты (back/internal/server)
        - проработать коды ответов, ответы 
3. front 
    + настроить dev-сервер, webpack, подключить react, попытаться отобразить helloworld
      Dev-сервер просто отдаёт собранный spa-фронт
    - страницы. CRUD в каждой странице. После редактировании данных информацию на странице следует обновить
        + список пользователей
        - таблица услуг
        - список заказов, сортировка по времени
        - страница с информацией о всех заказах клиента, открывается при клике на поле, связанное с клиентом
4. NGINX
    + проксирование на бэк, дев сервер фронта
5. Docker
    + каждая подсистема упакована в контейнер
    + оркестрация с помощью docker-compose, достаточная для локальной разработки

После копирования репозитория

docker-compose build

docker-compose up

После запуска приложение доступно на http://localhost:8080