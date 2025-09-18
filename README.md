# HTTP API Optimization Project

## 📋 Описание
Простое HTTP API для сложения чисел с оптимизациями по CPU и памяти.

## 🚀 Быстрый старт
```bash
go run main.go
# API: POST /add
# pprof: http://localhost:6060/debug/pprof/
```

## 📡 API
```json
POST /add
{"a": 12.3, "b": 67.1}
→ {"result": 79.4}
```

## 📚 История изменений

### 📌 v1.0.0 - Базовая реализация
**handlers/add.go:**
```go
func AddHandler(c *gin.Context) {
    var req AddRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result := req.A + req.B
    c.JSON(http.StatusOK, AddResponse{Result: result})
}
```

**Бенчмарк:**
```
BenchmarkAddHandler    100000    11250 ns/op    1464 B/op    15 allocs/op
```

### 📌 v1.1.0 - Оптимизация с sync.Pool
**Улучшения:**
- Добавлены `sync.Pool` для буферов
- Ручная сериализация JSON
- Оптимизирована обработка ошибок

**handlers/add.go:**
```go
var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func AddHandler(c *gin.Context) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)
    
    // Ручная сериализация вместо c.JSON()
    jsonBuf := append([]byte(`{"result":`), strconv.FormatFloat(result, 'f', -1, 64)...)
    jsonBuf = append(jsonBuf, '}')
    c.Data(http.StatusOK, "application/json", jsonBuf)
}
```

**Бенчмарк:**
```
BenchmarkAddHandler    148000    7890 ns/op    656 B/op    6 allocs/op
```

### 📌 v1.2.0 - Минимизация аллокаций
**Улучшения:**
- Переиспользование байтовых срезов
- Удалены временные строки
- Оптимизированы пулы

**Результаты:**
```
BenchmarkAddHandler    165000    7250 ns/op    512 B/op    4 allocs/op
```

## Итоговое сравнение
```
name         old time/op    new time/op    delta
AddHandler   11.2µs         7.3µs         -34.82%

name         old alloc/op   new alloc/op   delta
AddHandler   1.46kB         0.51kB        -65.07%

name         old allocs/op  new allocs/op  delta
AddHandler   15.0           4.0           -73.33%
```

## 🛠 Инструменты
```bash
# Профилирование
go tool pprof http://localhost:6060/debug/pprof/profile
go tool pprof http://localhost:6060/debug/pprof/heap
