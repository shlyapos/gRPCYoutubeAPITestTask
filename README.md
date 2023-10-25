# Test task for Golang Developer position

## Запуск приложения

1. Склонировать репозиторий `git clone git@github.com:shlyapos/EchelonTestTask.git`
2. Перейти в каталог проекта `cd ./EchelonTestTask`
3. В файле `./cmd/configs/config.yml` вставить свой Google API key в поле api_key 
4. Запустить сервер с помощью команды `go run ./cmd/server`
5. Запустить клиент:
  * Чтобы обработать одно видео: `go run ./cmd/client "https://www.youtube.com/watch?v=JPwRrJM4aAQ&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D"`
  * Чтобы обработать несколько видео: `go run ./cmd/client "https://www.youtube.com/watch?v=1O664QWpgUg&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D" "https://www.youtube.com/watch?v=nJ_b4VDbmmk&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D"`
  * Чтобы обработать несколько видео асинхронно: `go run ./cmd/client --async "https://www.youtube.com/watch?v=gGKL3GP2qQU&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D" "https://www.youtube.com/watch?v=egLXHjFL4Lo&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D" "https://www.youtube.com/watch?v=kS2t0kvIMmw&pp=ygUTbmF0aW9uYWwgZ2VvZ3JhcGhpYw%3D%3D"`
