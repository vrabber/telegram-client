# Telegram клиент Vrabber

## Репозитории:
- https://github.com/bonefabric/vrabber-protobuf - контракты
- https://github.com/bonefabric/vrabber-agent - клиентские агенты
- https://github.com/bonefabric/vrabber - сервер

## Требования:
- [go 1.23](https://go.dev/dl/)
- [Task](https://taskfile.dev/installation/) (автоматизация, не обязательно)

## Настройка
Настройки передаются через переменные окружения. 
При локальной разработке можно использовать Task с переменными, 
описанными в файле **[.env](.env)** (прим. **[.env.example](.env.example)**)  

- _VAR*_ - обязательное поле
- _VAR_ - необязательное поле, будет использоваться значение по умолчанию 
если указано, иначе - пустая строка
<hr>

- _TG_TOKEN*_ - токен бота, выдается ботом [BotFather](https://t.me/BotFather)
- _VRABBER_HOST_ - хост сервера, по умолчанию **localhost**
- _VRABBER_PORT*_ - порт сервера
- _MESSAGES_BUFFER_ - размер буффера сообщений, по умолчанию **100**
- _RESPONSES_BUFFER_ - размер буффера результатов, по умолчанию **100**
- _LOG_LEVEL_ - размер буффера результатов, enum: **DEBUG/INFO/WARN/ERROR** по умолчанию **INFO**
