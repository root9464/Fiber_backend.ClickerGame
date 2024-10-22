# Документация епт

я просто не ебу как swagger настроить поэтмоу ебашу так, ебать я клоун

## Https enpoints

## http://127.0.0.1:3000/api/auth/register
### Описание
> Используется для регистрации новых пользователей в веб-приложении. Она обрабатывает входящие HTTP-запросы, содержащие данные пользователя (имя, электронная почта и пароль), проверяет корректность и полноту этих данных, а затем, если все проверки пройдены успешно, сохраняет нового пользователя в базу данных. В случае успешной регистрации пользователю возвращается статус 200 с информацией о зарегистрированном пользователе. Если в данных есть ошибки или они неполные, функция возвращает соответствующие сообщения об ошибках с кодом состояния 500. Эта функция обеспечивает критически важный интерфейс для пользователей, желающих создать аккаунт и начать использовать предоставляемые веб-приложением функции.
### Response
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "mysecretpassword"
}
```
### Success request 
```json
{
    "id": 2,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "mysecretpassword"
}
```

## http://127.0.0.1:3000/api/auth/login
### Описание
> Используется для аутентификации пользователей в веб-приложении. Она обрабатывает HTTP-запросы, содержащие электронную почту и пароль пользователя, проверяет наличие этих данных, а затем проверяет, соответствуют ли предоставленные учетные данные какому-либо пользователю в базе данных. Если данные корректны и пользователь найден, функция возвращает статус 200 и данные пользователя. В случае ошибки в данных или если сочетание электронной почты и пароля не найдено, функция возвращает ошибку с кодом состояния 500 и соответствующим сообщением об ошибке. Эта функция является ключевой для процесса входа в систему, позволяя пользователям получить доступ к своим аккаунтам.
### Response
```json
{
  "email": "johndoe@example.com",
  "password": "mysecretpassword"
}
```
### Success request 
```json
{
    "id": 2,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "mysecretpassword"
}
```

## http://127.0.0.1:3000/api/auth/user/:userId
### Описание
> Используется для получения данных пользователя по его уникальному идентификатору в веб-приложении. Она принимает HTTP-запрос, извлекает из него параметр userId, проверяет его наличие и корректность. Если параметр отсутствует, функция возвращает ошибку с сообщением о том, что идентификатор необходим. Если пользователь с таким идентификатором не найден в базе данных, возвращается ошибка с сообщением о том, что пользователь не найден. В случае успешного нахождения пользователя, его данные возвращаются в ответе с кодом статуса 200. Эта функция важна для операций, требующих идентификации пользователя, например, при редактировании профиля или проверке прав доступа.

### Success request 
```json
{
    "id": 2,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "mysecretpassword"
}
```

## http://127.0.0.1:3000/api/improvement/addUserImprovement
### Описание
> Используется для добавления или обновления улучшения пользователя в веб-приложении. Она принимает данные из HTTP-запроса, включая имя пользователя, идентификатор улучшения и значение улучшения. Затем функция начинает транзакцию с базой данных, ищет пользователя по имени и улучшение по идентификатору. Если пользователь или улучшение не найдены, возвращается ошибка с соответствующим сообщением. Если улучшение для пользователя уже существует, обновляется его значение. В противном случае создается новая запись об улучшении пользователя. По завершении всех операций транзакция фиксируется. Функция возвращает сообщение об успешном завершении операции с кодом статуса 200. Эта функция позволяет управлять улучшениями пользователей в системе, что может быть полезно для игровых приложений или систем мотивации. 

### Response
```json
{
  "user_name": "John Doe",
  "improvement_id": 1,
  "value": 3 
}
```

### Success request 
```json
{
  "message": "success"
}
```

## http://127.0.0.1:3000/api/improvement/getImprovements
### Описание
> Используется для получения списка улучшений из базы данных и возвращения их в ответе на HTTP-запрос. Она инициализирует переменную для хранения списка улучшений, затем выполняет запрос к базе данных для извлечения всех улучшений. Если при выполнении запроса возникает ошибка, функция возвращает ошибку с сообщением о проблеме. В противном случае функция возвращает список улучшений с кодом статуса 200. Эта функция полезна для предоставления клиенту информации о доступных улучшениях в системе, например, для отображения пользователю возможных вариантов улучшений или функционала.

### Success request 
```json
[
    {
        "id": 1,
        "name": "Profit per hour",
        "description": "increases your profit by 'n' every hour"
    }
]
```

## http://127.0.0.1:3000/api/saveClicks
### Описание
> Используется для обновления количества кликов (clicks) в модели ProgressClicker в базе данных на основе данных, полученных из HTTP-запроса. Она парсит тело запроса для извлечения данных о кликах. Если происходит ошибка при парсинге, функция возвращает ошибку с соответствующим сообщением. В противном случае функция обновляет количество кликов в записи с идентификатором 1 в таблице ProgressClicker и возвращает сообщение об успешном выполнении с кодом статуса 200. Эта функция может использоваться, например, для отслеживания и сохранения количества кликов пользователя в приложении.

### Response
```json
{
  "id": 2,
  "clicks": 100
}
```

### Success request 
```json
{
  "message": "success"
}
```

## http://127.0.0.1:3000/api/admin/improvements/create
### Описание
> Используется для создания новой записи улучшения (improvement) в базе данных на основе данных, полученных из HTTP-запроса. Она парсит тело запроса для извлечения данных о новом улучшении. В случае ошибки при парсинге, функция возвращает ошибку с соответствующим сообщением. После успешного парсинга и создания записи в базе данных, функция возвращает сообщение о успешном выполнении с кодом статуса 200.

### Response
```json
{
  "name": "Example Improvement",
  "description": "This is an example improvement"
}
```

### Success request 
```json
{
  "message": "success"
}
```

## http://127.0.0.1:3000/api/admin/deleteUserImprovement
### Описание
> Используется для удаления конкретного улучшения, связанного с пользователем, на основе имени пользователя и идентификатора улучшения, предоставленных в теле запроса. Она ищет пользователя по имени и улучшение по идентификатору. Затем удаляет связь пользователя с улучшением и возвращает сообщение об успешном выполнении или сообщение об ошибке в зависимости от результата операции удаления.

### Response
```json
{
  "user_name": "example_username",
  "improvement_id": 1
}
```

### Success request 
```json
{
  "message": "success"
}
```

## http://127.0.0.1:3000/api/admin/deleteUserImprovement
### Описание
> Используется для создания новой записи об улучшении (improvement) в базе данных на основе данных, полученных из HTTP-запроса. Она парсит тело запроса для извлечения данных о новом улучшении. В случае ошибки при парсинге, функция возвращает ошибку с соответствующим сообщением. После успешного парсинга и создания записи в базе данных, функция возвращает сообщение о успешном выполнении с кодом статуса 200.

### Response
```json
{
  "improvement_id": 1
}
```

### Success request 
```json
{
  "message": "success"
}
```

## Модель базы данных
```sql
- User table:
  - Columns: Id (primary key), Name, Email, Password

- ProgressClicker table:
  - Columns: Id (primary key), UserId (foreign key referencing User table), Clicks, UserName

- Improvement table:
  - Columns: Id (primary key), Name, Description

- UserImprovement table:
  - Columns: Id (primary key), UserId (foreign key referencing User table), UserName, ImprovementId (foreign key referencing Improvement table), Value
```

## Визуализация модели базы данных
### Users Table
| id                    | name             | email            | password   | 
| --------------------- | ---------------- | ---------------- | ---------- |
| PRIMARY KEY, NOT NULL | NOT NULL, UNIQUE | NOT NULL, UNIQUE |            |

### ProgressClickers Table
| id                    | user_id          | user_name        | clicks     | 
| --------------------- | ---------------- | ---------------- | ---------- |
| PRIMARY KEY, NOT NULL | NOT NULL, INDEX  |                  |            |

### Improvements Table
| id                    | name             | description      |
| --------------------- | ---------------- | ---------------- |
| PRIMARY KEY, NOT NULL |                  |                  |

### UserImprovements Table
| id                    | user_id          | user_name        | improvement_id     |  value         |
| --------------------- | ---------------- | ---------------- | ------------------ |----------------|
| PRIMARY KEY, NOT NULL | NOT NULL, INDEX  |                  |                    |                |

## Примечание
>Хук AfterCreate - это специальный метод, который автоматически вызывается GORM после создания новой записи в базе данных. В этом случае, мы определили хук AfterCreate для структуры User, который создает новую запись в таблице ProgressClicker после создания новой записи в таблице User.
Когда мы создаем новую запись в таблице User с помощью database.DB.Create(&user), GORM автоматически вызывает хук AfterCreate для этой структуры. Это происходит потому, что мы определили хук AfterCreate для структуры User, и GORM знает, что он должен быть вызван после создания новой записи.
В нашем примере, хук AfterCreate создает новую запись в таблице ProgressClicker с default значениями (UserId равен ID только что созданной записи в таблице User, а Clicks равен 0). Это происходит автоматически, без необходимости явно вызывать хук в коде.
Таким образом, хук AfterCreate работает как бы "в тени", автоматически выполняя дополнительные действия после создания новой записи в базе данных. Это позволяет упростить код и сделать его более эффективным.

