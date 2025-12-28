As **Interfaces** são um dos conceitos mais refinados do Go. Elas permitem que você escreva um código flexível, genérico e focado no **comportamento** (o que o objeto faz) em vez da **identidade** (o que o objeto é).



# 🔌 Interfaces em Go

Em Go, uma interface é um **conjunto de assinaturas de métodos**. Ela define um "contrato": qualquer tipo que possua todos os métodos listados na interface, automaticamente a implementa.

## 1. O Conceito: Implementação Implícita

Diferente de Java ou C#, em Go você não usa uma palavra-chave como `implements`.

* **"Duck Typing":** Se ele anda como um pato e quacka como um pato, então é um pato. Se a sua `struct` tem o método `area() float64`, o Go entende que ela **é** uma `forma`, sem que você precise declarar isso explicitamente.

---

## 2. Exemplo Prático: Abstraindo Formas Geométricas

O código que você escreveu é o exemplo clássico de polimorfismo. Note como a função `escreverArea` não se importa se está recebendo um círculo ou um retângulo, desde que ambos saibam calcular sua área.

```go
type forma interface {
    area() float64
}

// Qualquer struct que tiver o método abaixo 'satisfaz' a interface forma
func escreverArea(f forma) {
    fmt.Printf("A área é %0.2f\n", f.area())
}

```

---

## 3. Casos de Uso Reais

### A. Desacoplamento de Banco de Dados

Você pode criar uma interface `Repositorio` e ter implementações diferentes para MySQL, MongoDB ou até um Mock para testes unitários.

```go
type SalvarDados interface {
    Salvar(usuario Usuario) error
}

```

### B. Plugins e Extensões

Permite que o seu sistema aceite novos tipos de componentes sem que você precise alterar o código principal (Open/Closed Principle).

### C. Abstração de I/O

As interfaces `io.Reader` e `io.Writer` da biblioteca padrão do Go são os melhores exemplos. Elas permitem ler dados de um arquivo, de uma conexão de rede ou de uma string usando exatamente o mesmo código.

---

## 4. Recomendações da Comunidade (Best Practices)

A comunidade Go segue alguns princípios sólidos para manter o código limpo:

* **Interfaces Pequenas:** Prefira interfaces com apenas um ou dois métodos. Interfaces gigantes são difíceis de implementar e manter (ex: `io.Reader` tem apenas o método `Read`).
* **Nomenclatura:** É comum dar nomes que terminam em "er" para interfaces de um único método.
* Método `Read` -> Interface `Reader`
* Método `Write` -> Interface `Writer`


* **Aceite Interfaces, Retorne Structs:** * Sempre que possível, peça uma interface como parâmetro de função (isso torna sua função mais flexível).
* Retorne o tipo concreto (`struct`) para que quem chama a função saiba exatamente com o que está trabalhando.


* **Não crie interfaces antes de precisar:** No Go, como a implementação é implícita, você pode criar a interface apenas quando precisar abstrair dois ou mais tipos. Não tente "adivinhar" o futuro.

---

## 5. A Interface Vazia: `interface{}` ou `any`

Uma interface que não tem nenhum método é satisfeita por **qualquer tipo** (int, string, struct, etc.).

* A partir do Go 1.18, usamos o alias `any`.
* **Cuidado:** Use apenas quando você realmente não sabe que tipo de dado irá receber, pois você perde a segurança da tipagem forte do Go.

---

### Resumo Comparativo

| Característica | Struct | Interface |
| --- | --- | --- |
| **Foco** | Dados e Estado. | Comportamento e Métodos. |
| **Tipo** | Concreto. | Abstrato. |
| **Uso** | Representar objetos do mundo real. | Definir contratos e permitir polimorfismo. |

---

**Qual o próximo passo?** Agora que você entende Interfaces, você está pronto para o nível "Ninja" do Go: **Concorrência (Goroutines e Channels)** ou talvez queira ver como lidar com **Generics**, que foi introduzido recentemente na linguagem! O que prefere?