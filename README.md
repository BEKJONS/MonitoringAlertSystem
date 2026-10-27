# Monitoring Alert System

A small Go service that samples CPU, memory and disk I/O on a fixed interval, logs every sample to `logs/monitoring.log`, and sends a Telegram alert when a value crosses its threshold.

## Run

```sh
export TELEGRAM_BOT_TOKEN=<token from @BotFather>
go run ./cmd
```

Thresholds and the sampling interval live in `config/config.yaml`. The bot sends alerts to every chat that has recently messaged it.
