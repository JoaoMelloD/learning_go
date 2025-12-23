# 📦 Arrays e Slices

Em Go, tanto Arrays quanto Slices são coleções de elementos do mesmo tipo, mas a grande diferença reside na **flexibilidade** e no **tamanho**.

## 1. Arrays (Tamanho Fixo)

O Array é uma estrutura rígida. Uma vez definido o seu tamanho, ele nunca mais muda.

* **Estaticidade:** O tamanho faz parte do tipo do array. Um `[5]int` é um tipo diferente de um `[10]int`.
* **Uso:** Menos comum em projetos reais, usado apenas quando você tem certeza absoluta da quantidade de itens (ex: dias da semana).

### Exemplos de Declaração:

```go
// Declarando sem valores (preenche com valor zero)
var array1 [5]int 

// Declarando com valores
array2 := [3]string{"A", "B", "C"}

// Uso dos "..." (Tamanho fixo baseado na quantidade de itens inseridos)
array3 := [...]int{1, 2, 3, 4, 5} // O compilador define como [5]int

```

---

## 2. Slices (Tamanho Dinâmico)

O Slice é a estrutura mais utilizada em Go. Ele funciona como uma "janela" ou "fatia" que aponta para um array por baixo dos panos.

* **Flexibilidade:** Você pode adicionar ou remover itens.
* **Componentes:** Um slice possui **ponteiro** (para o array real), **tamanho** (len) e **capacidade** (cap).

### Exemplo de Uso Prático:

```go
slice := []int{10, 20, 30} // Note que não há número dentro dos colchetes

// Adicionando itens com a função append
// O append retorna um NOVO slice com o item adicionado
slice = append(slice, 40) 

```

---

## 3. O Slice como "Fatia" de um Array

Como o nome sugere, um slice pode ser literalmente um pedaço de um array existente. Isso é extremamente eficiente em termos de memória, pois o Go não copia os dados, apenas aponta para eles.

```go
arrayBase := [5]int{1, 2, 3, 4, 5}

// Criando um slice a partir do índice 1 até o 2 (o índice 3 é exclusivo)
fatia := arrayBase[1:3] 

fmt.Println(fatia) // Resultado: [2, 3]

```

---

## 4. Comparativo Direto

| Característica          | Array                                         | Slice                           |
| ----------------------- | --------------------------------------------- | ------------------------------- |
| **Tamanho**             | Fixo (Imutável)                               | Dinâmico (Pode crescer)         |
| **Definição**           | `[n]tipo`                                     | `[]tipo`                        |
| **Alocação**            | Memória estática/Pilha                        | Memória dinâmica (Heap)         |
| **Passagem em funções** | Passado por cópia (lento para listas grandes) | Passado por referência (rápido) |

---

### ⚠️ Dica de Ouro: O `append`

Muitas pessoas esquecem que o `append` não altera o slice original "no lugar". Ele pode precisar criar um novo array se a capacidade acabar. Por isso, sempre faça:
`meuSlice = append(meuSlice, novoValor)`

