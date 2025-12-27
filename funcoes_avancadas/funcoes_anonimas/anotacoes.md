As **funções anônimas** são ferramentas poderosas em Go, especialmente quando você precisa de uma lógica rápida que não será reutilizada em outros lugares ou quando quer criar **Closures** (funções que capturam variáveis do ambiente ao redor).


---

# 🎭 Funções Anônimas

Uma função anônima é, como o nome diz, uma função sem um identificador (nome). Ela pode ser definida e executada no mesmo instante ou armazenada em uma variável para uso posterior.

## 1. A Sintaxe da Execução Imediata (IIFE)

Para que uma função anônima seja executada no momento em que é definida, precisamos adicionar um par de parênteses `()` logo após o fechamento das chaves `{}`.

```go
func() {
    fmt.Println("Executando imediatamente!")
}() // Esses parênteses "chamam" a função

```

---

## 2. Passagem de Parâmetros

Se a função anônima exigir parâmetros, você os define na declaração e passa os valores reais nos parênteses de execução ao final.

### Exemplo Real: Processamento de Texto

```go
func(texto string) {
    fmt.Println("Mensagem recebida:", texto)
}("Olá, Go!") // O valor "Olá, Go!" entra no parâmetro 'texto'

```

---

## 3. Retornando Valores para Variáveis

Você pode usar uma função anônima para processar algo e armazenar o resultado diretamente em uma variável. Isso é muito útil para cálculos complexos que você quer isolar.

```go
resultado := func(n1, n2 int) int {
    return n1 * n2
}(10, 5)

fmt.Println(resultado) // Saída: 50

```

---

## 4. Closures: Funções que "Lembram"

Uma função anônima pode acessar e modificar variáveis que foram declaradas fora dela. Isso é chamado de **Closure**.

### Exemplo Real: Contador Isolado

```go
func main() {
    texto := "Variável externa"

    func() {
        // A função anônima tem acesso ao escopo superior
        fmt.Println(texto)
    }()
}

```

---

## 5. Quando usar Funções Anônimas?

* **Go Routines:** Muito comuns ao iniciar processos paralelos com a palavra-chave `go func() { ... }()`.
* **Defer:** Usadas com `defer` para agrupar múltiplas ações de limpeza no final de uma função.
* **Lógica Única:** Quando uma função é tão específica que não faz sentido dar um nome a ela e poluír o escopo global.

---

### Resumo Técnico

| Componente | Função |
| --- | --- |
| `func(...) { ... }` | Declaração da lógica e tipos. |
| `(args)` (no final) | Gatilho que executa a função imediatamente. |
| **Escopo** | Pode acessar variáveis do "pai" (Closures). |

---

**Qual o próximo passo da sua jornada?** Agora que terminamos o módulo de Funções, seria o momento perfeito para falarmos sobre o **Defer**, ou se preferir, podemos entrar no mundo dos **Ponteiros avançados** aplicados a Structs! O que você acha?