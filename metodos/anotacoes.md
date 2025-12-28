Os **métodos** são a forma como o Go implementa o comportamento de objetos, mesmo não sendo uma linguagem orientada a objetos (OOP) tradicional. Eles permitem "anexar" funções a tipos específicos, criando uma relação clara entre o dado e a ação.

Aqui estão suas anotações reestruturadas com foco em arquitetura e boas práticas:

---

# 🛠️ Métodos em Go

Um método é essencialmente uma função com um argumento especial chamado **Receiver** (receptor). Esse receptor é posicionado entre a palavra-chave `func` e o nome do método, vinculando a função a um tipo (geralmente uma `struct`).

## 1. Estrutura e Sintaxe

O receptor dá ao método acesso aos campos da struct, funcionando de forma semelhante ao `this` ou `self` de outras linguagens.

```go
type Usuario struct {
    Nome  string
    Idade uint8
}

// Método com Receiver de Valor (Cópia)
func (u Usuario) saudar() {
    fmt.Printf("Olá, meu nome é %s\n", u.Nome)
}

```

---

## 2. Métodos com Ponteiros vs. Valor

Esta é a decisão mais importante ao criar um método:

### A. Receiver de Valor (`u Usuario`)

* **O que faz:** Cria uma cópia da struct para uso dentro do método.
* **Uso:** Apenas leitura de dados ou operações que não devem alterar o original.
* **Vantagem:** Segurança de que o objeto original permanecerá imutável.

### B. Receiver de Ponteiro (`u *Usuario`)

* **O que faz:** Passa o endereço de memória da struct.
* **Uso:** Quando o método precisa **modificar** um atributo ou para evitar cópias de structs muito grandes.
* **Vantagem:** Performance e capacidade de alteração de estado.

```go
func (u *Usuario) fazerAniversario() {
    u.Idade++ // Altera o valor na struct original
}

```

---

## 3. Casos de Uso Reais

### A. Formatação de Dados (Getter/Presentation)

Muitas vezes usamos métodos para transformar dados da struct em algo legível ou formatado.

```go
type Produto struct {
    Nome  string
    Preco float64
}

func (p Produto) PrecoFormatado() string {
    return fmt.Sprintf("R$ %.2f", p.Preco)
}

```

### B. Lógica de Negócio e Validação

Métodos são ideais para encapsular regras que pertencem àquele dado específico.

```go
type ContaBancaria struct {
    Saldo float64
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.Saldo {
        return errors.New("saldo insuficiente")
    }
    c.Saldo -= valor
    return nil
}

```

---

## 4. Diferenças Cruciais: Funções vs. Métodos

| Característica | Função Comum | Método |
| --- | --- | --- |
| **Chamada** | `salvar(usuario)` | `usuario.salvar()` |
| **Vínculo** | Independente. | Ligado a um Tipo específico. |
| **Organização** | Espalhada pelo pacote. | Agrupa comportamento junto ao dado (Encapsulamento). |
| **Interfaces** | Não satisfaz interfaces sozinho. | Essencial para implementar interfaces. |

---

### 💡 Dica de Design: Consistência de Receivers

A comunidade Go recomenda que, se um dos métodos da sua struct precisa de um receptor de ponteiro (`*`), todos os outros métodos dessa mesma struct também devem ser de ponteiro, mesmo os que apenas leem dados. Isso evita confusão sobre como o tipo se comporta.

---