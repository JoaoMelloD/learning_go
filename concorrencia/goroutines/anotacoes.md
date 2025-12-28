A concorrência é o "superpoder" do Go. Enquanto a maioria das linguagens lida com threads pesadas do sistema operacional, o Go introduziu as **Goroutines**, que são extremamente leves e permitem que você execute milhares (ou milhões) de tarefas simultaneamente com um custo de memória baixíssimo.

Aqui estão suas anotações definitivas sobre Concorrência em Go:

---

# 🚀 Concorrência e Goroutines

Em Go, concorrência não é o mesmo que paralelismo. Concorrência é sobre **lidar** com várias coisas ao mesmo tempo (estrutura), enquanto paralelismo é sobre **fazer** várias coisas ao mesmo tempo (execução em múltiplos núcleos de CPU).

## 1. O que são Goroutines?

Uma Goroutine é uma função que é executada de forma independente e assíncrona.

* **Leveza:** Uma thread do SO ocupa cerca de 1MB de memória. Uma Goroutine começa com apenas **2KB**.
* **Gerenciamento:** O próprio *Go Runtime* gerencia as Goroutines (através de um escalonador próprio), e não o Sistema Operacional. Isso torna a troca de contexto muito rápida.
* **Sintaxe:** Basta adicionar a palavra-chave `go` antes da chamada de uma função.

```go
func tarefa(nome string) {
    fmt.Println("Executando tarefa:", nome)
}

func main() {
    go tarefa("Assíncrona") // Inicia uma nova Goroutine
    tarefa("Síncrona")      // Executa na main thread
}

```

---

## 2. WaitGroups: Sincronizando o Final

Um problema comum é que a função `main` não espera as Goroutines terminarem; se a `main` acaba, o programa morre e leva as Goroutines junto. Para resolver isso, usamos o pacote `sync.WaitGroup`.

```go
var wg sync.WaitGroup

func processo(id int) {
    defer wg.Done() // Avisa que terminou ao sair da função
    fmt.Printf("Processo %d finalizado\n", id)
}

func main() {
    wg.Add(2) // Dizemos que vamos esperar 2 tarefas
    
    go processo(1)
    go processo(2)
    
    wg.Wait() // Bloqueia a execução até que o contador chegue a zero
}

```

---

## 3. Canais (Channels): A Comunicação Segura

A filosofia do Go é: *"Não comunique compartilhando memória; em vez disso, compartilhe memória comunicando"*. Canais são os "tubos" por onde as Goroutines enviam e recebem dados.

* **Sincronização Nativa:** Enviar ou receber dados de um canal bloqueia a execução até que o outro lado esteja pronto.

```go
func enviarDados(canal chan string) {
    canal <- "Olá do Canal!" // Envia dado
}

func main() {
    canal := make(chan string)
    go enviarDados(canal)
    
    mensagem := <-canal // Recebe dado (bloqueia aqui até chegar algo)
    fmt.Println(mensagem)
}

```

---

## 4. Casos de Uso Reais

### A. Web Scrapers / Crawlers

Você pode disparar uma Goroutine para cada URL que deseja baixar dados, reduzindo o tempo total de horas para segundos.

### B. Processamento de Background (Workers)

Em um servidor Web, você pode responder ao usuário imediatamente e disparar uma Goroutine para enviar um e-mail ou processar uma imagem em segundo plano.

### C. APIs de Alta Performance

Lidar com milhares de conexões simultâneas (como chats ou sistemas de streaming) de forma eficiente.

---

## 5. Boas Práticas e Recomendações da Comunidade

### 🛡️ Evite Race Conditions (Condições de Corrida)

Nunca deixe duas Goroutines tentarem alterar a mesma variável ao mesmo tempo. Use **Canais** ou **Mutexes** (`sync.Mutex`) para proteger dados compartilhados.

* *Dica:* Rode seus testes com `go test -race` para detectar esses problemas.

### 🛑 Não crie Goroutines "Zumbis"

Sempre saiba como uma Goroutine vai terminar. Se você inicia uma Goroutine que fica esperando algo em um canal que nunca será fechado, você tem um **Memory Leak**.

### 🏗️ Padrão Worker Pool

Em vez de criar Goroutines infinitas, crie um número fixo de "trabalhadores" (ex: 10 Goroutines) que ficam lendo de um canal de tarefas. Isso evita que sua aplicação consuma todos os recursos da máquina sob carga pesada.

### ⏳ Use Select para Timeouts

Sempre use o comando `select` para evitar que uma Goroutine fique travada esperando um canal para sempre.

```go
select {
case res := <-canal:
    fmt.Println(res)
case <-time.After(time.Second * 3):
    fmt.Println("Timeout: A operação demorou demais!")
}

```

---

### Resumo Técnico

| Recurso | Função |
| --- | --- |
| **`go`** | Cria a execução concorrente. |
| **`sync.WaitGroup`** | Espera um grupo de tarefas terminar. |
| **`chan`** | Canal de comunicação e sincronização entre Goroutines. |
| **`select`** | Controla múltiplos canais simultaneamente. |

---
