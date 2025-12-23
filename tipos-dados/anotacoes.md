

# 📊 Tipos de Dados em Go

Go é uma linguagem que oferece controle fino sobre como os dados são armazenados na memória. Abaixo, os principais tipos e suas características:

## 1. Números Inteiros

Os inteiros são divididos pelo espaço que ocupam na memória (em bits):

* **Com sinal (Podem ser negativos):** `int8`, `int16`, `int32`, `int64`.
* **Sem sinal (`uint` - Unsigned):** `uint8`, `uint16`, `uint32`, `uint64`. (Não aceitam valores negativos).
* **Tipo Genérico (`int`):** É o mais comum. O Go define o tamanho dele com base na **arquitetura do seu computador** (se seu PC for 64 bits, o `int` será `int64`).

---

## 2. Números Reais (Decimais)

Para números com ponto flutuante, o Go oferece duas precisões:

* `float32`
* `float64` (É o padrão usado pela maioria das funções matemáticas nativas).

---

## 3. Strings e o "Mistério" do Char

Em Go, **não existe o tipo `char**`. O que acontece por baixo dos panos é curioso:

* **Aspas Duplas (`" "`):** Criam uma `string` normal (cadeia de caracteres).
* **Aspas Simples (`' '`):** Criam uma **Rune** (que na verdade é um `int32`). Ela representa o número daquele caractere na tabela Unicode/ASCII.

```go
char := 'B'
fmt.Println(char) // Resultado: 66 (O código da letra B na tabela)

```

---

## 4. O Conceito de "Valor Zero"

Toda variável em Go, se não for inicializada com um valor, recebe automaticamente um **Valor Zero** (padrão). Não existe "undefined":

| Tipo de Dado        | Valor Zero (Default) |
| ------------------- | -------------------- |
| `int` / `float`     | `0`                  |
| `string`            | `""` (vazio)         |
| `bool`              | `false`              |
| `error` / Ponteiros | `nil`                |

---

## 5. Boleanos

Usado para lógica condicional.

* `bool`: Aceita apenas `true` ou `false`.

---

## 6. O Tipo `error` (Tratamento de Erro)

O Go não usa `try/catch`. O tratamento de erro é feito através de um tipo específico chamado `error`.

* **Valor Zero:** O valor padrão de um erro é `nil` (significa que não há erro).
* **Criação de Erros:** Para criar uma mensagem de erro personalizada, usamos o pacote nativo `errors`.

```go
import "errors"

// Criando um erro manualmente
var erro error = errors.New("Erro Interno no Servidor")

```

> **Nota sobre o `nil`:** É o equivalente ao "nulo" de outras linguagens, mas em Go ele é o valor zero para tipos de referência como erros, ponteiros, interfaces, slices e maps.

---

