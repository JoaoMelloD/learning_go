

# 👷 Worker Pools em Go

Um **Worker Pool** consiste em manter um número fixo de Goroutines lendo de um canal comum de tarefas. Isso evita o custo de criar e destruir milhares de Goroutines e, mais importante, evita que sua máquina esgote os recursos (CPU/RAM) ao tentar fazer tudo de uma vez.

## 1. O Conceito

Imagine uma agência de correios:

* **O Canal de Tarefas:** É a esteira onde os pacotes chegam.
* **Os Workers:** São os funcionários fixos no balcão.
* **O Canal de Resultados:** É a caixa onde os pacotes processados são colocados.

Mesmo que cheguem 1.000 pacotes, se você tem apenas 4 funcionários, o trabalho será feito de forma organizada e constante, sem causar caos na agência.

---

## 2. Anatomia do Código (Análise do seu Exemplo)

No seu código, você implementou o cálculo de Fibonacci. Esta é uma tarefa **CPU-bound** (consome muito processador).

```go
// tarefas <-chan int : Canal de APENAS LEITURA (Recebe ordens)
// resultados chan<- int : Canal de APENAS ESCRITA (Entrega o trabalho)
func worker(tarefas <-chan int, resultados chan<- int) {
    for numero := range tarefas {
        resultados <- fibonacci(numero)
    }
}

```

### Por que usar Canais com Direção?

Note o uso de `<-chan` e `chan<-`. Isso é uma **boa prática de segurança**:

* Garante que o `worker` não feche acidentalmente o canal de tarefas.
* Garante que o `worker` não tente ler do canal onde ele deveria apenas escrever os resultados.

---

## 3. Quando usar Worker Pools?

Você deve implementar este padrão quando:

1. **Limitação de Recursos:** Você quer limitar quantos núcleos de CPU ou conexões de rede estão sendo usados simultaneamente.
2. **Volume de Dados:** Você tem milhares de itens para processar, mas não quer disparar uma Goroutine para cada um (o que poderia causar lentidão no escalonador do Go).
3. **Controle de Taxa (Throttling):** Você está consumindo uma API externa que tem um limite de requisições por segundo.

---

## 4. Casos de Uso Reais

### A. Processamento de Imagens

Se um usuário faz upload de 50 fotos, você cria um pool de 4 workers para redimensioná-las. Isso garante que o servidor ainda tenha CPU sobrando para atender outros usuários.

### B. Web Scrapers / Crawlers

Ao varrer um site com 10.000 páginas, um pool de workers garante que você não seja bloqueado por fazer acessos demais ao mesmo tempo, mantendo um fluxo constante de requisições.

### C. Importação de Grandes Arquivos (CSV/Logs)

Ler um arquivo de 1GB e inserir no banco de dados. O pool de workers pode processar as linhas em paralelo enquanto a `main` lê o arquivo.

---

## 5. Recomendações da Comunidade e Boas Práticas

* **Tamanho do Pool:** Para tarefas de CPU (como Fibonacci), o número ideal de workers costuma ser o número de núcleos (CPUs) da máquina (`runtime.NumCPU()`). Para tarefas de I/O (rede/banco de dados), o pool pode ser muito maior.
* **Fechamento de Canais:** Sempre feche o canal de tarefas (`close(tarefas)`) para sinalizar aos workers que não há mais trabalho. Caso contrário, eles ficarão travados esperando no `range`.
* **Canais com Buffer:** Use buffers nos canais de tarefas e resultados para evitar que a `main` trave enquanto os workers ainda estão pegando os itens iniciais.
* **WaitGroups vs Canais:** No seu exemplo, você usou a leitura do canal de resultados para sincronizar o fim. Em sistemas mais complexos, é comum usar um `sync.WaitGroup` dentro da main para garantir que todos os workers terminaram antes de fechar o canal de resultados.

---

### Exemplo de Melhoria (Sincronização Profissional)

```go
func main() {
    numTarefas := 45
    tarefas := make(chan int, numTarefas)
    resultados := make(chan int, numTarefas)

    // Criando o pool de acordo com a capacidade da máquina
    for w := 1; w <= 4; w++ {
        go worker(tarefas, resultados)
    }

    // Enviando tarefas
    for i := 0; i < numTarefas; i++ {
        tarefas <- i
    }
    close(tarefas) // Importantíssimo!

    // Coletando resultados
    for i := 0; i < numTarefas; i++ {
        fmt.Println(<-resultados)
    }
}

```
