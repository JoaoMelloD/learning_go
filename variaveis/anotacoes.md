# 🐹 Dominando Variáveis em Go

Go é uma linguagem **estaticamente e fortemente tipada**. Isso significa que o tipo da variável é verificado em tempo de compilação e não muda durante a execução.

## 1. Regras de Ouro (O "Código Limpo" Nativo)

Diferente de linguagens como Python ou JavaScript, o compilador do Go é rigoroso para evitar o acúmulo de "lixo" no código:

* **Zero Desperdício de Importação:** Se você importar um pacote e não usá-lo, o código não compila.
* **Variáveis Obrigatórias:** Se você declarar uma variável (especialmente a de forma implícita), **deve** utilizá-la. Caso contrário, o Go retornará um erro de compilação.
* **Eficiência em Escala:** Essas restrições garantem que binários grandes permaneçam leves e que o código seja legível para qualquer pessoa da equipe.

---

## 2. Formas de Declaração

Existem duas maneiras principais de criar variáveis em Go:

### A. Declaração Explícita (`var`)

Utilizada quando você quer deixar claro o tipo da variável ou quando vai declarar variáveis sem atribuir um valor imediato (o Go atribuirá o "valor zero" do tipo).

```go
var nome string = "Mello"

```

### B. Declaração Implícita (Operador Curto `:=`)

O Go infere (adivinha) o tipo com base no valor que você atribui. É a forma mais comum dentro de funções.

```go
name := "Tunior" // O Go entende automaticamente que 'name' é uma string
fmt.Println(name) // Lembre-se: declarou, tem que usar!

```

---

## 3. Atribuição em Massa e Truques Úteis

O Go oferece sintaxes elegantes para lidar com múltiplas variáveis ao mesmo tempo.

### Agrupamento de Variáveis (`var block`)

Ideal para organizar constantes ou variáveis globais no topo do arquivo.

```go
var (
    variavelA string = "Testosterona"
    variavelB string = "Texto"
)

```

### Declaração Curta Múltipla

```go
variavelC, variavelD := "textoA", "textoB"

```

### O "Truque" da Inversão (Swap)

Em outras linguagens, você precisaria de uma variável temporária. No Go, você pode inverter valores em uma única linha:

```go
a, b = b, a

```

> **Como funciona:** O Go avalia todos os valores do lado direito primeiro e depois faz as atribuições para o lado esquerdo, permitindo a troca direta.

---

### Dica Extra: Valores Zero

Se você declarar uma variável e não der valor a ela, o Go atribui um padrão:

* `int`: `0`
* `string`: `""` (vazio)
* `bool`: `false`

---
