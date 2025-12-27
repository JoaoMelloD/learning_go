
# 🏷️ Retorno Nomeado em Go

Em Go, você pode dar nomes aos valores de retorno na assinatura da função. Isso não apenas documenta o que cada valor representa, mas também simplifica o corpo da função.

## 1. Como funciona?

Ao definir nomes para os retornos, o Go as trata como variáveis locais que são criadas assim que a função começa a ser executada. Elas são inicializadas com o seu **valor zero** (0 para int, "" para string, etc.).

### Exemplo de Sintaxe:

```go
// Definimos os nomes 'soma' e 'subtracao' diretamente no parêntese de retorno
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
    soma = n1 + n2
    subtracao = n1 - n2
    
    // O return vazio (naked return) sabe que deve retornar as variáveis acima
    return
}

```

---

## 2. O "Naked Return" (Retorno Limpo)

Quando você nomeia os retornos, não precisa especificar as variáveis novamente na linha do `return`. O compilador do Go entende automaticamente que ele deve retornar os valores atuais das variáveis nomeadas.

> **⚠️ Atenção:** Embora o "naked return" seja prático, ele deve ser usado com moderação em funções muito longas, pois pode dificultar a leitura (o programador precisa subir até o topo da função para saber o que está sendo retornado).

---

## 3. Vantagens no Mundo Real

### A. Auto-documentação

Sem nomes, você teria `func buscarDados() (string, string, int)`. O que são essas strings?
Com retornos nomeados, fica claro: `func buscarDados() (nome string, email string, idade int)`.

### B. Inicialização Automática

Você não precisa usar `:=` dentro da função para essas variáveis, pois elas já foram declaradas na assinatura. Basta usar o operador de atribuição `=`.

### C. Clareza em Tratamento de Erros

Em funções complexas, nomear o erro ajuda a identificar o fluxo de saída.

---

## 4. Diferença Visual

| Sem Retorno Nomeado | Com Retorno Nomeado |
| --- | --- |
| Precisa declarar as variáveis no corpo. | Variáveis já nascem prontas para uso. |
| O `return` precisa listar cada variável. | O `return` pode ser vazio (naked). |
| Menos legível em retornos múltiplos. | Funciona como documentação da API. |

---

### Exemplo Real: Divisão e Resto

```go
func divisaoInteira(dividendo, divisor int) (quociente, resto int) {
    quociente = dividendo / divisor
    resto = dividendo % divisor
    return // Retorna automaticamente quociente e resto
}

```
