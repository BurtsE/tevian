# Запуск

make docker-up 

(необходимо  установить python-dotenv и определить переменные окружения в .env)

либо 

docker-compose -f deploy/compose.yml up

# Список необходимых переменных окружения:

* FACE_CLOUD_LOGIN="new-user@example.com"
* FACE_CLOUD_PASSWORD="123"
* FACE_CLOUD_URL="https://backend.facecloud.tevian.ru"
* POSTGRES_USER="admin"
* POSTGRES_PASSWORD="123"
* POSTGRES_DB="tevian"

* LOGIN="abc"
* PASSWORD="123"
