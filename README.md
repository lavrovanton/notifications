# Notifications
## Deploy
```
cp .env.example .env
docker compose up -d
```

## Swagger
http://localhost:9000/docs/index.html

## RabbitMQ
http://localhost:15672

Username: guest\
Password: guest

## Notification example for API and RabbitMQ
RabbitMQ Queue: notify
```
{
    "sender_email": "foo@email.com",
    "receiver_email": "bar@email.com",
    "text": "hello"
}
```