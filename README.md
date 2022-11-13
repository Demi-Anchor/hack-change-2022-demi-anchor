# hack-change-2022-demi-anchor

## Видео-демонстрация работы всего проекта https://youtu.be/BiUrRpVx65E

https://radiant-wildwood-88953.herokuapp.com

#### Локальный запуск

1. Создать файл ".env" с необходимыми переменными окружения рядом с файлом "Makefile"

   *.env example*

   ```
   DB_HOST=dbhost.ru
   DB_PORT=5432
   DB_NAME=beautifulname
   DB_USER=thebestuser
   DB_PASSWORD=verystrongpassword
   ```

2.  Выполнить make run в консоли

#### API

*Добавление доната*

POST https://radiant-wildwood-88953.herokuapp.com/api/v1/donations

   ```
   Request
   
   {
     "streamer_id": 222,
     "author": "Dog",
     "money": 7000,
     "comment": "For cat",
     "time": "2022-11-05T09:52:36Z"
   }
   ```

*Получение суммы и количества донатов по дням за определенный период*

POST https://radiant-wildwood-88953.herokuapp.com/api/v1/donations/daily

   ```
   Request
   
   {
     "streamer_id": 222,
     "author": "Dog",
     "money": 7000,
     "comment": "For cat",
     "time": "2022-11-05T09:52:36Z"
   }
   
   Response
   
   [
       {
           "sum": 30656000,
           "count": 9,
           "date": "2022-11-04T00:00:00Z"
       },
       {
           "sum": 7000,
           "count": 1,
           "date": "2022-11-05T00:00:00Z"
       }
   ]
   
   ```