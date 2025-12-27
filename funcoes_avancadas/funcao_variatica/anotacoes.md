As **funções variáticas** são extremamente úteis quando você não sabe antecipadamente quantos argumentos serão passados para uma função. O exemplo mais clássico que você já usa no dia a dia é o próprio `fmt.Println()`, que aceita qualquer quantidade de argumentos.

Aqui estão suas anotações reestruturadas e aprofundadas para um nível profissional:

---

# 🔗 Funções Variáticas em Go

Uma função é chamada de **variática** quando ela pode ser invocada com um número variável de argumentos. Em Go, isso é indicado pelo prefixo de reticências (`...`) antes do tipo do parâmetro.

## 1. Como Funciona por Baixo dos Panos?

Quando você define um parâmetro como `...int`, o Go recebe esses valores e os converte automaticamente em um **Slice** do tipo especificado dentro da função.

### Exemplo: Soma de N Números

```go
// numeros ...int indica que a função recebe de 0 a N inteiros
func soma(numeros ...int) {
    total := 0
    // Como 'numeros' é tratado como um Slice, podemos usar o range
    for _, numero := range numeros {
        total += numero
    }
    fmt.Println("Total da soma:", total)
}

```

---

## 2. Regras e Restrições Importantes

Existem duas regras de ouro que você deve seguir ao usar funções variáticas em Go:

1. **Apenas um parâmetro variático:** Você não pode ter dois conjuntos de `...` na mesma função.
2. **O variático deve ser o último:** Se a função receber outros parâmetros comuns, o parâmetro variático deve ser obrigatoriamente o último da lista.

### Exemplo Real: Prefixo de Mensagem

```go
func logComMensagem(prefixo string, mensagens ...string) {
    for _, msg := range mensagens {
        fmt.Printf("[%s]: %s\n", prefixo, msg)
    }
}

func main() {
    logComMensagem("INFO", "Sistema iniciado", "Conexão estabelecida")
}

```

---

## 3. Passando um Slice para uma Função Variática

Um truque muito útil: se você já tem um Slice pronto e quer passá-lo para uma função variática, você pode "desempacotar" o slice usando as reticências no final da variável.

```go
meusNumeros := []int{10, 20, 30}

// soma(meusNumeros)     // ❌ ERRO: esperando int, recebeu []int
soma(meusNumeros...)    // ✅ OK: desempacota o slice em valores individuais

```

---

## 4. Diferenças e Casos de Uso

| Característica | Função Comum | Função Variática |
| --- | --- | --- |
| **Flexibilidade** | Rígida (exige número exato de argumentos) | Total (aceita 0, 1 ou muitos) |
| **Tratamento interno** | Variável simples | Tratada como um **Slice** |
| **Uso Comum** | Lógica de negócio fixa | Formatadores, Logs, Cálculos matemáticos |

---

### Dica de Performance

Lembre-se que, como o Go cria um slice para agrupar os argumentos variáticos, há uma pequena alocação de memória envolvida. Em funções críticas de performance chamadas milhões de vezes, use com consciência.

---

**Qual seria o próximo passo?** Agora que você entende como passar múltiplos valores, que tal explorarmos o **Defer**? Ele é essencial para garantir que processos sejam encerrados corretamente, não importa o que aconteça na função! Gostaria de ver como ele funciona?