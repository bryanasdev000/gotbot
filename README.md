# Go TBOT! Go!

Aplicação de linha de comando simples baseado no [Golang bindings for the Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) para envio de mensagens.

## Uso

Certifique-se de ter o pacote `go` da sua distribuição instalado.

Compile o projeto: `go build .`

Vai gerar um executável na pasta com o nome de `gotbot`.

Para utilizar é necessário informar duas variáveis de ambientes obrigatórias e uma opcional.

A variável `TELEGRAM_TOKEN` contém o token de autenticação dado pelo [@BotFather](https://telegram.me/BotFather).

A variável `TELEGRAM_TARGET` contém o ChatID da conversa (individual ou grupo), leia [aqui](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id) para entender como conseguir.

A variável `TBOT_DEBUG` controla o modo verboso da API do Telegram, para ativar utilize o valor 1 na mesma, exemplo `TBOT_DEBUG=1`. A omissão da variável ou valores diferentes de 1 são ignorados.

Uso:

```
export TELEGRAM_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXX
export TELEGRAM_TARGET=00000000
./gotbot 'Teste!'
```

> OBS.: Certifique-se de utilizar aspas simples (') para delimitar a mensagem e escapar os espaços, de outro modo o gotbot irá reclamar por ter mais de um parâmetro.
