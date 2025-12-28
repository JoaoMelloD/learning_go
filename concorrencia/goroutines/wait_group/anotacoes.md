O `sync.WaitGroup` é uma ferramenta essencial para o controle de fluxo em Go. Ele funciona como um **contador de tarefas** que impede que a função principal (`main`) encerre antes que todas as tarefas paralelas (Goroutines) terminem seu trabalho.

Aqui estão suas anotações reestruturadas para um nível profissional e didático:

---

# ⏱️ Sincronização com sync.WaitGroup

O **WaitGroup** faz parte do pacote `sync` e é utilizado para coordenar a finalização de múltiplas Goroutines. Sem ele, a função `main` terminaria sua execução e fecharia o programa antes mesmo das Goroutines terem a chance de completar.

## 1. O Trio Fundamental de Métodos

Para utilizar um WaitGroup, você precisa dominar estes três métodos que manipulam um contador interno:

1. **`Add(int)`**: Incrementa o contador. Você informa quantas Goroutines o grupo deve esperar.
2. **`Done()`**: Decrementa o contador em 1. Deve ser chamado obrigatoriamente quando a Goroutine finaliza sua tarefa.
3. **`Wait()`**: Bloqueia a execução do código (geralmente na `main`) até que o contador chegue a zero.

---

## 2. Anatomia do Código Sincronizado

Observe como o ciclo de vida da sincronização se comporta no seu exemplo:

```go
func main() {
    var waitGroup sync.WaitGroup

    // 1. Preparamos o grupo para esperar 2 Goroutines
    waitGroup.Add(2)

    go func() {
        escrever("Olá Mundo")
        // 2. Avisamos que esta rotina terminou (Contador: 2 -> 1)
        waitGroup.Done() 
    }()

    go func() {
        escrever("Outro Texto")
        // 2. Avisamos que esta rotina terminou (Contador: 1 -> 0)
        waitGroup.Done()
    }()

    // 3. A main fica "travada" aqui até que o contador seja 0
    waitGroup.Wait() 
}

```

---

## 3. Boas Práticas da Comunidade

### 🛡️ Use o `defer` para garantir o `Done()`

Em funções complexas, pode ocorrer um erro ou um `return` precoce que impeça a chamada do `Done()`, causando um *deadlock* (o programa espera para sempre). O uso do `defer` garante que o contador seja decrementado assim que a função sair.

```go
go func() {
    defer waitGroup.Done() // Executa no último milissegundo da função
    escrever("Execução segura")
}()

```

### 📏 Chame o `Add()` fora da Goroutine

Sempre chame o `Add()` na thread principal antes de iniciar a Goroutine com o comando `go`. Se você chamar dentro da Goroutine, existe o risco dela não ser escalonada a tempo e o `Wait()` ser ultrapassado porque o contador ainda estava em zero.

---

## 4. Casos de Uso Reais

* **Processamento em Lote (Batch):** Você tem 100 imagens para redimensionar. Você cria 100 Goroutines, dá um `Add(100)` e espera todas terminarem antes de enviar um e-mail de confirmação ao usuário.
* **Agregadores de Dados:** Disparar chamadas para 3 APIs diferentes simultaneamente e esperar o retorno de todas para montar a resposta final.
* **Shutdown Gracioso:** Antes de desligar um servidor, usar um WaitGroup para garantir que todas as requisições em andamento sejam finalizadas primeiro.

---

## 5. Diferença: WaitGroup vs Channels

| Recurso | Objetivo Principal |
| --- | --- |
| **WaitGroup** | Apenas sincronização de finalização (esperar os outros). |
| **Channels** | Comunicação e transferência de dados entre as rotinas. |

---

**Dica de Ouro:** Se o seu `Wait()` nunca termina e o programa trava, verifique se a quantidade de `Add()` é exatamente igual à quantidade de `Done()` executados. Se o `Add` for maior que o `Done`, você terá um erro de **deadlock**.

---
