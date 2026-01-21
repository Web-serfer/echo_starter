# Echo + Templ Starter

Это стартовый проект на Go с использованием веб-фреймворка Echo и шаблонизатора Templ, с поддержкой Tailwind CSS через CDN.

## Структура проекта

```
echo_starter/
├── cmd/
│   └── app/
│       └── main.go                 # Основной файл приложения
├── internal/
│   ├── handlers/                   # Обработчики HTTP-запросов
│   ├── middleware/                 # Пользовательские middleware
│   └── render/                     # Вспомогательные функции рендеринга
├── static/
│   └── css/
│       └── style.css               # Локальные стили (дополняющие Tailwind)
├── view/
│   ├── components/
│   │   └── layout/
│   │       ├── header/
│   │       │   └── header.templ    # Компонент шапки сайта
│   │       └── footer/
│   │           └── footer.templ    # Компонент подвала сайта
│   ├── layout/
│   │   └── base.templ              # Базовый шаблон страницы
│   └── pages/
│       └── index.templ             # Главная страница
├── go.mod                          # Модуль Go
├── go.sum                          # Суммы зависимостей
├── Makefile                        # Скрипты для автоматизации
├── .air.toml                       # Конфигурация для Air (горячая перезагрузка)
└── README.md                       # Документация (этот файл)
```

## Возможности

- **Echo v5**: Современный и быстрый веб-фреймворк для Go
- **Templ**: Типизированный шаблонизатор для Go
- **Tailwind CSS**: Подключен через CDN для быстрой стилизации
- **Горячая перезагрузка**: Поддержка автоматической перезагрузки при разработке
- **Статические файлы**: Поддержка раздачи статических файлов

## Установка и запуск

1. Убедитесь, что у вас установлен Go 1.25+
2. Установите необходимые зависимости:
   ```bash
   go mod tidy
   ```
3. Установите templ CLI:
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```
4. Сгенерируйте шаблоны:
   ```bash
   templ generate
   ```
5. Запустите приложение:
   ```bash
   go run cmd/app/main.go
   ```

## Использование

### Добавление новых страниц

Создайте новый файл в директории `view/pages/` с расширением `.templ`, например `about.templ`, и опишите компонент страницы:

```templ
package pages

import "echo_starter/view/layout"

templ About() {
  @layout.Base("О нас") {
    <h1>О нашем проекте</h1>
    <p>Это страница о нас</p>
  }
}
```

### Добавление компонентов

Создайте новый файл в директории `view/components/` и опишите компонент:

```templ
package components

templ Button(text string, onClick func()) {
  <button onclick={ onClick } class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
    { text }
  </button>
}
```

### Использование Tailwind CSS

Все классы Tailwind CSS доступны через CDN. Просто используйте их в ваших шаблонах:

```templ
<div class="container mx-auto p-4 bg-gray-100">
  <h1 class="text-2xl font-bold text-blue-600">Заголовок</h1>
</div>
```

## Makefile команды

- `make deps` - установка зависимостей
- `make generate` - генерация шаблонов
- `make run` - запуск приложения с генерацией
- `make build` - сборка приложения
- `make test` - запуск тестов
- `make dev` - запуск в режиме разработки с горячей перезагрузкой

## Автор

Проект создан как стартовый шаблон для веб-приложений на Go с использованием Echo и Templ.