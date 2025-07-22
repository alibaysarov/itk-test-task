# ITK тестовое задание "Wallet API"

## Запуск проекта
1. Проставить переменные в .env из .env.example
2. Запусить docker compose

    <code>docker compose up -d</code>
3. Накатить миграции 

    <code>make migrate</code>

## Swagger
/swagger/index.html

## Миграции
#### Создать миграцию
<pre>
make create_migration MIGRATION_NAME=название
</pre>
#### Применить миграцию
<pre>
make migrate
</pre>