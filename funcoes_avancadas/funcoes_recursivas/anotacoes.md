As **funções recursivas** são um conceito fundamental na ciência da computação e no Go. Elas permitem resolver problemas complexos dividindo-os em subproblemas menores da mesma natureza.

Aqui estão suas anotações reestruturadas, com foco na lógica de parada e exemplos práticos:

---

# 🔄 Funções Recursivas

Uma função é considerada recursiva quando **chama a si mesma** dentro do seu próprio bloco de código. Ela é ideal para processar estruturas de dados repetitivas (como árvores de diretórios) ou cálculos matemáticos baseados em definições recursivas.

## 1. A Anatomia da Recursão

Toda função recursiva precisa de dois componentes obrigatórios para não causar um erro de estouro de memória (*stack overflow*):

1. **Caso Base (Condição de Parada):** É o momento em que a função para de chamar a si mesma.
2. **Caso Recursivo:** A chamada da própria função com um argumento modificado, aproximando-a do caso base.

---

## 2. Exemplo Clássico: Sequência de Fibonacci

A sequência de Fibonacci é o exemplo perfeito, pois cada número é a soma dos dois anteriores: .

```go
func fibonacci(posicao uint) uint {
    // CASO BASE: Se a posição for 0 ou 1, paramos a recursão.
    if posicao <= 1 {
        return posicao
    }

    // CASO RECURSIVO: A função chama a si mesma duas vezes
    return fibonacci(posicao-2) + fibonacci(posicao-1)
}

```

---

## 3. Casos de Uso no Mundo Real

Embora o Fibonacci seja ótimo para aprender, na vida real as funções recursivas são usadas em cenários como:

* **Navegação em Pastas (File System):** Para listar todos os arquivos de uma pasta e suas subpastas.
* **Menus Multinível:** Carregar categorias e subcategorias em um e-commerce.
* **Algoritmos de Ordenação:** Algoritmos como *QuickSort* e *MergeSort* utilizam recursão para dividir listas grandes em partes menores.
* **Estruturas em Árvore:** Processar arquivos XML ou JSON onde um objeto pode conter outros objetos dentro de si.

---

## 4. Cuidados e Performance

* **Stack Overflow:** Se você esquecer o caso base ou se a recursão for profunda demais, o programa consumirá toda a memória da pilha (*stack*) e irá travar.
* **Custo de Processamento:** Funções como o Fibonacci recursivo simples são ineficientes para números grandes, pois recalculam o mesmo valor muitas vezes. Em Go, às vezes é preferível usar um loop `for` (iterativo) para ganhar performance.

---

### Resumo Comparativo

| Abordagem | Como funciona | Quando usar |
| --- | --- | --- |
| **Recursiva** | Chama a si mesma. | Estruturas de dados complexas e ramificadas. |
| **Iterativa** | Usa loops (`for`). | Cálculos matemáticos simples e alta performance. |

---

**Qual o próximo passo?** Agora que fechamos o ciclo de Funções, você gostaria de ver como o Go gerencia erros de forma elegante usando o **Defer, Panic e Recover**? São conceitos que salvam vidas quando uma função recursiva ou complexa dá erro!