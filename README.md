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

- _VAR_* - обязательное поле
- _VAR_ - необязательное поле, будет использоваться значение по умолчанию 
если указано, иначе - пустая строка
<hr>

- _TG_TOKEN_* - токен бота, выдается ботом [BotFather](https://t.me/BotFather)

