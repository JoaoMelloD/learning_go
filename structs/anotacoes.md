As **Structs** são o coração da organização de dados em Go. Como a linguagem não possui classes nem herança tradicional, as structs são a ferramenta principal para criar modelos de dados complexos e reutilizáveis.

Aqui está uma versão robusta e completa das suas anotações:

---

# 🏗️ Structs em Go (Estruturas)

As structs são coleções de campos (campos = atributos/propriedades), onde cada campo possui um nome e um tipo definido. Elas são a base do **encapsulamento** de dados em Go.

## 1. Definição e Instanciação

Para criar uma struct, usamos a palavra-chave `type` seguida pelo nome da estrutura e a palavra `struct`.

```go
type Usuario struct {
    Nome  string
    Idade uint8
    Email string
}

func main() {
    // Forma 1: Atribuindo valores por nome de campo (Recomendado)
    u1 := Usuario{
        Nome:  "Mello",
        Idade: 25,
        Email: "mello@email.com",
    }

    // Forma 2: Atribuição por ordem (Pouco prático para structs grandes)
    u2 := Usuario{"Tunior", 30, "tunior@email.com"}

    // Forma 3: Apenas alguns campos (Os outros recebem o 'Valor Zero')
    u3 := Usuario{Nome: "Ana"} 
}

```

---

## 2. Structs Aninhadas (Composição)

Em Go, você pode colocar uma struct dentro de outra. Isso permite modelar relações complexas entre dados.

```go
type Endereco struct {
    Logradouro string
    Numero     int
}

type Cliente struct {
    Nome     string
    Idade    int
    Endereco Endereco // Struct como campo de outra struct
}

```

---

## 3. "Herança" via Composição (Anonymous Fields)

Go não tem herança de classes. Para reaproveitar campos de uma struct em outra, usamos os **campos anônimos** (ou promoção de campos).

```go
type Pessoa struct {
    Nome      string
    Sobrenome string
}

type Estudante struct {
    Pessoa    // Campo anônimo (Sem nome, apenas o tipo)
    Curso     string
    Faculdade string
}

func main() {
    e1 := Estudante{Pessoa{"João", "Silva"}, "Engenharia", "USP"}
    
    // Podemos acessar os campos de 'Pessoa' diretamente no 'Estudante'
    fmt.Println(e1.Nome) // Em vez de e1.Pessoa.Nome
}

```

---

## 4. Tags de Struct (JSON e Banco de Dados)

Uma funcionalidade muito poderosa das structs são as **Tags**. Elas informam a outros pacotes (como conversores JSON) como aquele campo deve ser tratado.

```go
type Produto struct {
    Nome  string  `json:"nome_produto"` // No JSON, esse campo se chamará 'nome_produto'
    Preco float64 `json:"preco"`
}

```

---

## 5. Diferenças: Structs vs Classes

Se você vem de linguagens como Java ou C#, aqui estão as principais diferenças:

| Característica | Classes (Tradicional) | Structs (Go) |
| --- | --- | --- |
| **Herança** | Sim (Hierarquia) | Não (Apenas Composição) |
| **Construtores** | Sim (Métodos especiais) | Não (Usamos funções comuns) |
| **Métodos** | Definidos dentro da classe | Definidos fora da struct (via receivers) |
| **Ponteiros** | Geralmente abstraídos | Controle total sobre valor vs referência |

---

### Dica de Performance: Ponteiros em Structs

Ao passar uma struct para uma função, o Go faz uma **cópia** de todos os dados. Se a struct for muito grande, prefira passar o **ponteiro** (`*Usuario`) para economizar memória e permitir que a função altere os dados originais.

