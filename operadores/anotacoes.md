

# ⚖️ Operadores em Go

Go é extremamente rigoroso com tipos de dados para evitar erros de precisão e comportamentos inesperados na memória.

## 1. Operadores Aritméticos

Os operadores básicos são os mesmos de outras linguagens: `+`, `-`, `*`, `/` e `%`. Porém, existe uma regra de ouro em Go:

> **Regra de Tipagem:** Você **não pode** realizar operações entre tipos diferentes, mesmo que ambos sejam números.

### O Erro de Tipagem

```go
var numeroA int16 = 10
var numeroB int8 = 5

// soma := numeroA + numeroB // ❌ ERRO: tipos diferentes (int16 vs int8)

```

**Como resolver?** É necessário fazer a conversão explícita (Casting):

```go
soma := numeroA + int16(numeroB) // ✅ Agora funciona!

```

---

## 2. Operadores Relacionais

Sempre comparam dois valores e retornam um **Booleano** (`true` ou `false`).

* `==` (Igual a)
* `!=` (Diferente de)
* `>` , `<` (Maior ou Menor que)
* `>=` , `<=` (Maior/Menor ou igual a)

---

## 3. Operadores Lógicos

Usados para combinar condições booleanas:

* `&&` (AND): Verdadeiro se **ambos** forem verdadeiros.
* `||` (OR): Verdadeiro se **pelo menos um** for verdadeiro.
* `!` (NOT): Inverte o valor booleano (**Negação**).

---

## 4. Operadores Unários

Usados para aumentar ou diminuir o valor de uma variável existente.

* `valor++` (Incremento: soma 1)
* `valor--` (Decremento: subtrai 1)
* `valor += 10` (Atribuição com soma: soma 10 ao valor atual)

**Atenção:** Em Go, `valor++` é um *statement* (instrução) e não uma expressão. Isso significa que você não pode fazer algo como `x = y++`.

---

## 5. O Caso do Operador Ternário

Uma das perguntas mais comuns de quem vem do JavaScript ou C# é: "Onde está o operador `condicao ? true : false`?"

* **Em Go, NÃO existe operador ternário.**
* **Por quê?** Os criadores do Go decidiram que o código fica mais legível e menos propenso a erros usando a estrutura clara de `if/else`.

### Exemplo Real: Substituindo o Ternário

**Em outras linguagens:**
`status := (idade >= 18) ? "Maior" : "Menor"`

**Em Go:**

```go
var status string
if idade >= 18 {
    status = "Maior"
} else {
    status = "Menor"
}

```

---

### Resumo Visual

| Tipo            | Operadores              | Exemplo      |
| --------------- | ----------------------- | ------------ |
| **Aritméticos** | `+`, `-`, `*`, `/`, `%` | `10 / 2`     |
| **Relacionais** | `==`, `!=`, `>`, `<`    | `a > b`      |
| **Lógicos**     | `&&`, `                 |              |
| **Unários**     | `++`, `--`, `+=`, `-=`  | `contador++` |

---
