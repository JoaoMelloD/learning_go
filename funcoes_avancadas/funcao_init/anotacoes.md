A função `init` é uma ferramenta poderosa e única do Go para preparar o terreno antes que a lógica principal da sua aplicação comece a rodar. Ela resolve o problema de "quem configura o que?" de forma automática e organizada.

# 🏁 Função `init`: O Setup Automático

A função `init` é uma função especial do Go que é executada **automaticamente** assim que o pacote é inicializado. Ela não precisa (e não pode) ser chamada manualmente.

## 1. Características Principais

* **Precedência:** Ela é executada antes da função `main`.
* **Múltiplas Inits:** Diferente da `main`, você pode ter várias funções `init` espalhadas por diferentes arquivos do mesmo pacote. Elas serão executadas na ordem em que os arquivos aparecem para o compilador.
* **Sem Parâmetros ou Retornos:** Ela deve ser sempre declarada como `func init() { ... }`.
* **Escopo de Inicialização:** Ela serve para garantir que, quando a sua `main` começar, todas as variáveis e conexões necessárias já estejam prontas.

---

## 2. Ordem de Execução (Fluxo de Inicialização)

O Go segue uma hierarquia rígida para garantir a segurança dos dados:

1. **Pacotes Importados:** Primeiro, o Go inicializa todos os pacotes que você importou.
2. **Variáveis de Pacote:** Em seguida, ele inicializa as variáveis globais (no nível do pacote).
3. **Funções `init`:** Depois, ele executa todas as funções `init()`.
4. **Função `main`:** Por fim, a execução entra na `main`.

---

## 3. Casos de Uso Reais

### A. Configuração de Variáveis de Ambiente

Ideal para carregar chaves de API ou variáveis do sistema antes que o servidor suba.

```go
var conexaoDB string

func init() {
    conexaoDB = os.Getenv("DB_URL")
    if conexaoDB == "" {
        fmt.Println("Erro: DB_URL não configurada!")
    }
}

```

### B. Inicialização de Drivers (Side Effects)

Alguns pacotes (como drivers de banco de dados ou imagens) usam a `init` para se registrarem no sistema apenas ao serem importados.

```go
import _ "github.com/go-sql-driver/mysql" 
// O '_' faz com que o pacote seja importado apenas para executar sua função init()

```

### C. Validação de Pré-requisitos

Garantir que o sistema operacional seja compatível ou que pastas essenciais existam no disco.

---

## 4. Diferenças Cruciais

| Característica | Função `init` | Função `main` |
| --- | --- | --- |
| **Ponto de entrada** | Inicialização de pacotes. | Início da execução do programa. |
| **Chamada** | Automática (nativa). | Automática (apenas no pacote main). |
| **Quantidade** | Pode haver várias por pacote/arquivo. | Apenas uma por aplicação. |
| **Dependência** | Roda antes da `main`. | Roda após todas as `init` terminarem. |

---

## ⚠️ Dica de Ouro: Evite Lógicas Complexas

Embora a `init` seja útil, evite colocar lógicas muito pesadas ou que dependam de outras funções externas complexas nela. Como ela roda antes da `main`, se uma `init` travar ou entrar em loop, sua aplicação nem chegará a começar, dificultando o debug.

---