# ZipDownloadUrl

## Как запустить

### 1. Склонировать репозиторий 

```bash
git clone https://github.com/Genrih/ZipDownloadUrl.git
```

### Уставинить зависимости, если не установились при копировании

```bash
go get go get github.com/gorilla/mux
```

### 2. Запуск сервера

#### 2.1 Способ напрямую

```bash
go run ./cmd
```

#### 2.2 Зайти в папку cmd и запустить main.gp

```bash
cd cmd
```

```bash
go run main.go
```
### 3. Проверка работоспобности API и сервера

#### 3.1 Через postman

#### 3.1.1 POST-запрос для создания задачи

```bash
http://localhost:8080/task/create
```

#### 3.1.1 Или же Использовать в другом окне терминала команду

```bash
go run main.go create
```
task_id, который выдает после команды, надо скопировать из терминала, чтобы после использовать для других команд

#### 3.1.2 POST-запрос для добавления файла в задачу, но также надо указать JSON

```bash
http://localhost:8080/task/add
```
вписать в postman JSON, при примеру ниже

{
    "task_id": "1",
    "url": "https://github.com/Genrih/zipUrl/blob/master/ZipDownloadUrl/README.md?raw=true"
}

task_id - id созданной задачи, url - ссылка на скачиваемый файл, также самое главное, чтобы в конце было расширение .pdf, .jpg, .jpeg иначе запрос не пройдет


#### 3.1.2 Или же Использовать в другом окне терминала команду,

```bash
go run main.go add task_id url
```
task_id - id созданной задачи, url - ссылка на скачиваемый файл, также самое главное, чтобы в конце было расширение .pdf, .jpg, .jpeg иначе запрос не пройдет

#### 3.2.1 GET-запросы для получения одной задачи


Вписать в postman ниже

```bash
http://localhost:8080/task/statusFile/task_id
```

task_id - id созданной задачи

#### 3.2.1 Или использовать в другом окне терминала команду

```bash
go run main.go status task_id
```

task_id - id созданной задачи

#### 3.3.1 GET-запросы для получения всех задач

```bash
http://localhost:8080/tasks
```

#### 3.3.1 Или использовать в другом окне терминала команду

```bash
go run main.go status-all
```