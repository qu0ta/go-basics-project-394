### Hexlet tests and linter status:
[![Actions Status](https://github.com/qu0ta/go-basics-project-394/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/qu0ta/go-basics-project-394/actions)
# 🔐 Password Generator & Checker

[![Go Tests](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/go.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPO/actions/workflows/go.yml)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Простое консольное приложение на Go для генерации надёжных паролей и проверки их стойкости.

## 🚀 Возможности

### Генерация паролей
- Настраиваемая длина пароля
- Выбор набора символов:
  - строчные буквы (a-z)
  - прописные буквы (A-Z)
  - цифры (0-9)
  - специальные символы (!@#$%^&*)
- Псевдослучайная генерация на основе линейного конгруэнтного генератора

### Проверка надёжности
Оценка пароля по 5 критериям (от 0 до 5 баллов):
1. Длина ≥ 8 символов
2. Наличие строчных букв
3. Наличие прописных букв
4. Наличие цифр
5. Наличие специальных символов

**Шкала оценок:**
- 0–2 балла → "Слабый"
- 3 балла → "Средний"
- 4 балла → "Надёжный"
- 5 баллов → "Очень надёжный"

## 🛠️ Использование

### API функций

#### GeneratePassword
```go
func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string