# addhosts

A simple CLI tool for adding local hosts to `/etc/hosts`.

## Installation

```bash
git clone <repository-url>
cd addhosts
go build -o addhosts

sudo cp addhosts /usr/local/bin/
# or
# go install
```

## Usage

### Add hosts

```bash
# Add a single host
sudo addhosts add example.local

# Add multiple hosts
sudo addhosts add api.local app.local db.local
```

### List current hosts

```bash
addhosts list
```

## Features

- Adds hosts for `127.0.0.1` IP address
- Duplicate checking before adding
- Formatted table output for current local hosts
- Permission check for `/etc/hosts`
- Detailed error messages

## Requirements

- Go 1.21+
- Admin privileges for adding hosts

---

<details>
<summary>🇷🇺 Русский</summary>

## addhosts

Простая командная утилита для добавления локальных хостов в `/etc/hosts`.

### Установка

```bash
git clone <repository-url>
cd addhosts
go build -o addhosts

sudo cp addhosts /usr/local/bin/
# или
# go install
```

### Использование

#### Добавление хостов

```bash
# Добавить один хост
sudo addhosts add example.local

# Добавить несколько хостов
sudo addhosts add api.local app.local db.local
```

#### Просмотр текущих хостов

```bash
addhosts list
```

### Функционал

- Добавление хостов для IP-адреса `127.0.0.1`
- Проверка дубликатов перед добавлением
- Табличный вывод текущих локальных хостов
- Проверка прав доступа к `/etc/hosts`
- Подробные сообщения об ошибках

### Требования

- Go 1.21+
- Права администратора для добавления хостов

</details>
