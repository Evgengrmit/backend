Задания:

1. API Пополнения счета (Пост запрос с указанием имени, кол-ва денег и валюты)
2. Апи снятия денег со счета (пут запрос с указанием имени, кол-ва денег и валюты)
3. Апи получения кол-ва денег (гет запрос с указанием имени)
4. Апи отправки денег другому пользователю (пост запрос от кого и кому по имени)

* Если пользователь 1 имеет валюту в рублях, то при отправки денег пользователю 2 домнажать на 100
* Логировать все действия с помощью fmt или log
* Проверки, что счет не отрицательный, если снимаем или переводим больше имееющегося
* Информативные ошибки

Обязательное:

1. Выложить в гит
2. Закоммитить в мастер/мейн текущее состояние
3. Создать ветку таск_1 и реализовывать задание в ней
4. Осмысленные коммиты
5. После окончания работ или когда требуется провести ревью - создается пулл реквест на ветку мастер

Мой ник в гитхаб we2beast

Структурирование проекта:
1. Controller
2. Service
3. Repository
4. Entity
5. Model

Domain Driven Design. Есть книга на русском Предметно‑ориентированное проектирование 

Задачи:

1. Подключить Gin +
2. Разделить работу с пользователями и их счетами на отдельные папки +
3. Расширить кол-во полей у пользователя, чтобы он принимал почту и пароль 
4. При запрашивании пользователя, не возвращать пароль+
5. При сохранении пользователя пароль шифровать (почитать про шифрование). При успешном создании отвечать 201 статусом+
6. Сделать ручку логин, которая принимает почту и пароль и сверяет с тем, что хранится. 
Важно, чтобы статусы ответов были правильные. Если данные не сошлись, то 403 статус.+
7. Поставить docker
8. Попробовать в докере запустить postgresql
9. GORM просто прочитать
10. Начать принимать переменные DB_URI, DB_USERNAME, DB_PASSWORD через переменные окружение

* Попробовать подключить го приложение к базе данных, используя GORM

Вопросы на подумать:
1. В БД у каждой записи есть уникальные идентификаторы, какие стоит использовать в распределенной системе?
2. Какими способами можно запретить изменение строчек/столбоцов в БД?
3. Изучить подходы, при которых внешняя система или БД недоступно?
4. Изучить что такое очередь сообщений и какие есть варианты? Примеры их использования?
5. Изучить кеши для бекенда и какие инструменты для этого есть?


Задания:
1. Миграции, чтобы выполнялись при запуске приложения. https://stackoverflow.com/questions/1766046/postgresql-create-table-if-not-exists
2. Создать две таблицы
3. Написать запросы по работе с базой
4. https://metanit.com/go/tutorial/10.3.php


Таблица токенов хранит в себе
1. Айди пользователя
2. Сам выданный токен
3. поле отозван или нет
4. Дата создания токена

Модифицируешь миддлвар проверки токена.
ПОсле парсинга проверяешь, что он не отозван

Таблица с хранением попыток авторизации
1. Айди пользователя
2. Успешен ли вход
3. Время, когда был осуществлен вход
4. Айпи пользователя, с которого пришел запрос
5. 1 Валюта  но по требованию обращаться к курсу валют и выводить в нужном

Почитать про сваггер, что это такое и зачем нужно.
Может сможешь подключить к своей системе