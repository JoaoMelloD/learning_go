

---

# 🚀 Funções em Go (Functions)

Em Go, funções são muito mais do que blocos de código; elas são **tipos de dados**. Você pode armazená-las em variáveis, passá-las como argumentos e até retorná-las de outras funções.

## 1. Estrutura Básica

Uma função é definida pela palavra reservada `func`, seguida por seu nome, parâmetros e o tipo de retorno.

```go
func somar(n1 int, n2 int) int {
    return n1 + n2
}

```

---

## 2. Funções com Múltiplos Retornos

Diferente de C ou Java, o Go permite retornar vários valores nativamente. Isso é muito usado para retornar o **resultado** e um **erro** simultaneamente.

### Exemplo Real: Divisão Segura

```go
func dividir(dividendo, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, errors.New("não é possível dividir por zero")
    }
    return dividendo / divisor, nil
}

```

### O Operador Blank Identifier (`_`)

Se uma função retorna dois valores, mas você só precisa de um, você **deve** usar o `_` para descartar o outro. O Go não permite declarar variáveis que não serão usadas.

```go
resultado, _ := dividir(10, 2) // Ignora o erro propositalmente
fmt.Println(resultado)

```

---

## 3. Funções como Tipos e Variáveis

Você pode atribuir uma função a uma variável. Isso é útil para criar comportamentos dinâmicos.

### Exemplo: Variável do tipo função

```go
var f = func(texto string) {
    fmt.Println("Executando:", texto)
}

f("Minha Função")

```

---

## 4. Funções Anônimas e Aninhadas

Você pode declarar uma função dentro de outra (aninhamento). Se ela não tiver nome, é chamada de **função anônima**.

### Exemplo Real: Processamento de Dados

```go
func processar() {
    // Função anônima sendo executada imediatamente (IIFE)
    func() {
        fmt.Println("Limpando banco de dados...")
    }() 
}

```

---

## 5. Ordem e Assinatura

**Importante:** No Go, a assinatura de uma função (parâmetros e retornos) define o seu tipo.

* `func(int, int) int` é um tipo.
* `func(string) string` é outro tipo completamente diferente.

> **Atenção:** A ordem dos parâmetros e dos retornos altera a assinatura. Se você inverter a ordem dos tipos de retorno, o Go considerará uma função diferente e seu código não compilará se estiver esperando o formato original.

---

### Tabela de Resumo: Anatomia da Função

| Componente     | Descrição                                                      |
| -------------- | -------------------------------------------------------------- |
| `func`         | Palavra-chave obrigatória.                                     |
| **Parâmetros** | Entradas (sempre com tipo definido).                           |
| **Retorno**    | Pode ser único, múltiplo ou inexistente.                       |
| **Escopo**     | Se começar com Letra Maiúscula, é pública para outros pacotes. |

---

