# 🔄 Estruturas de Repetição em Go

Diferente de quase todas as linguagens modernas, o Go possui apenas uma palavra-chave para loops: **`for`**. No entanto, ele assume diferentes formas para suprir todas as necessidades de iteração.

## 1. As Faces do `for`

### A. O For Clássico (Estilo C)

Utilizado quando você sabe exatamente o início, a condição de parada e o passo de incremento.

```go
for i := 0; i < 10; i++ {
    fmt.Println("Iteração:", i)
}

```

### B. O For como "While"

Como não existe a palavra `while`, usamos apenas a condição.

```go
n := 0
for n < 5 {
    n++
    fmt.Println("Simulando um while...")
}

```

### C. Loop Infinito

Muito usado em servidores que ficam "ouvindo" conexões ou processos em background.

```go
for {
    // Executa para sempre, a menos que haja um 'break' ou 'return'
}

```

---

## 2. O Poder do `range`

O `range` é uma palavra-chave usada para iterar sobre estruturas de dados (Slices, Arrays, Maps e Strings). Ele é mais seguro e limpo porque evita erros de "índice fora do limite" (out of bounds).

### Anatomia do Range

O `range` sempre retorna **dois valores**: o índice (ou chave) e uma cópia do valor do elemento.

```go
nomes := []string{"Mello", "Tunior", "Ana"}

for indice, valor := range nomes {
    fmt.Printf("Posição %d: %s\n", indice, valor)
}

```

### Otimização com o Blank Identifier (`_`)

Em Go, se você declara uma variável, deve usá-la. Se você precisa apenas do valor e não do índice, use o `_` para descartar o que não for necessário.

```go
for _, valor := range nomes {
    fmt.Println("Processando apenas o nome:", valor)
}

```

---

## 3. Iterando sobre Strings e o "Mistério da Rune"

Quando usamos `range` em uma string, o Go itera sobre cada caractere. No entanto, por padrão, ele retorna o valor numérico do caractere na tabela Unicode (**Rune/ASCII**).

```go
for indice, letra := range "GO" {
    fmt.Println(indice, letra) 
}
// Saída: 
// 0 71
// 1 79

```

**✅ Caso de Uso Real:** Para exibir o caractere legível, você deve converter o valor numérico explicitamente para string:

```go
fmt.Println(indice, string(letra)) // Saída: 0 G, 1 O

```

---

## 4. Range em Mapas (Chave e Valor)

O `range` também é a forma padrão de percorrer um dicionário (Map). **Importante:** a iteração em mapas no Go não tem ordem garantida.

```go
usuarios := map[string]string{"login": "mello123", "status": "ativo"}

for chave, valor := range usuarios {
    fmt.Printf("Campo: %s | Valor: %s\n", chave, valor)
}

```

---

### Tabela de Retorno do `range`

| Estrutura | 1º Valor Retornado | 2º Valor Retornado |
| --- | --- | --- |
| **Array/Slice** | Índice (`int`) | Elemento (`valor`) |
| **String** | Índice (`int`) | Caractere (`rune`) |
| **Map** | Chave (`key`) | Valor (`value`) |
| **Channel** | Elemento | (Não possui) |

---

**Qual seria o próximo passo ideal?** Agora que você domina as repetições, seria excelente ver os **Maps** (que citei acima) em detalhes, ou talvez entender como o **Switch** funciona para evitar muitos `if/else` dentro desses loops! O que prefere?