As **Closures** são um conceito avançado e extremamente útil em Go. Elas permitem que uma função "capture" e "carregue" variáveis do ambiente onde foram criadas, mesmo depois que esse ambiente já terminou sua execução.

---

# 🔒 Closures em Go

Uma **Closure** (fechamento) ocorre quando uma função anônima referencia variáveis que estão fora do seu corpo. A função "fecha" o escopo em volta dessas variáveis, mantendo-as vivas enquanto a função existir.

## 1. Definição Técnica

Em Go, as funções são "cidadãos de primeira classe", o que significa que podem retornar outras funções. Uma Closure é essa função retornada que continua tendo acesso às variáveis locais da função "pai", mesmo após a função pai ter retornado.

---

## 2. Anatomia de uma Closure

No seu exemplo, observe como a variável `text` "sobrevive" dentro da `funcaoNova`:

```go
func closure() func() {
    text := "Valor capturado"

    // Esta função anônima é uma closure pois usa 'text'
    funcao := func() {
        fmt.Println(text)
    }

    return funcao // Retornamos a lógica E o acesso à variável 'text'
}

func main() {
    funcaoNova := closure()
    funcaoNova() // Saída: "Valor capturado"
}

```

---

## 3. Quando usar Closures?

Você deve considerar o uso de closures quando precisar de:

1. **Isolamento de Dados:** Manter uma variável protegida sem torná-la global.
2. **Geradores:** Criar funções que mantêm um estado interno (ex: contadores).
3. **Configuração de Funções:** Criar funções personalizadas com base em parâmetros iniciais.

---

## 4. Casos de Uso Reais

### A. Geradores de Sequência (Contadores)

As closures são perfeitas para criar contadores que não podem ser resetados externamente.

```go
func novoContador() func() int {
    i := 0
    return func() int {
        i++ // A closure incrementa a variável do escopo pai
        return i
    }
}

func main() {
    contar := novoContador()
    fmt.Println(contar()) // 1
    fmt.Println(contar()) // 2
    
    outroContador := novoContador() // Começa do zero, de forma independente
    fmt.Println(outroContador()) // 1
}

```

### B. Middlewares e Decorators

Muito usado em servidores Web para "envolver" uma função com lógica extra (como logs ou autenticação).

```go
func criarSaudacao(saudacao string) func(string) string {
    return func(nome string) string {
        return saudacao + ", " + nome
    }
}

func main() {
    bomDia := criarSaudacao("Bom dia")
    boaNoite := criarSaudacao("Boa noite")

    fmt.Println(bomDia("Mello"))   // Bom dia, Mello
    fmt.Println(boaNoite("Tunior")) // Boa noite, Tunior
}

```

---

## 5. Resumo Visual: Escopo Global vs. Closure

| Característica | Variável Global | Variável em Closure |
| --- | --- | --- |
| **Acesso** | Qualquer parte do programa. | Apenas pela função que a capturou. |
| **Segurança** | Baixa (pode ser alterada por erro). | Alta (encapsulada/protegida). |
| **Estado** | Compartilhado por todos. | Único para cada instância da closure. |

---

### Dica de Performance

Closures são muito poderosas, mas lembre-se: como elas mantêm variáveis vivas na memória (no *Heap*), o uso excessivo e sem controle pode aumentar o consumo de memória da sua aplicação. Use-as para encapsular lógica, não para guardar grandes volumes de dados desnecessariamente.
