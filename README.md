# Data Structures & Algorithms in Go

## Linear

- [x] Stack
- [x] Queue
- [x] Double Linked List
- [x] Circular Buffer
- [ ] Deque (segmented array), ⛔ Not possible in Go

## Tree — hierarchical, parent/child relationships

- [x] Binary Search Tree
- [x] AVL Tree
- [x] Heap (min/max)
  - - [x] Priority Queue
- [ ] Trie

## Graph

- [x] Unweighted
- [ ] Weighted
- [ ] Directed

## Sets

- [x] Union Find

## Probabilistic

- [ ] Bloom filter
- [ ] Hyper Log Log

# Algorithms ωψγ

- [x] Kadane's Algorithm
- [x] Euclidean GCD
- [x] Fibonacci
- [x] Heap Sort
- [ ] Modular Arithmetic
- [ ] Sieve of Eratosthenes
- [x] BFS Breadth-First Search
- [ ] DFS Depth-First Search
- [x] Miller Rabin prime test

# Documentation

> linear Structures

```bash
go doc -all ./linear | bat -l go
```

> algorithms

```bash
go doc -all ./algo | bat -l go
```

> Sets

```bash
go doc -all ./sets/ | bat -l go
```

> trees

```bash
for p in ./trees/ ./trees/avl ./trees/heap; do go doc -all "$p"; done | bat -l go
```

> Graphs

```bash
go doc -all ./graphs | bat -l go
```

# Tests

```bash
go test ./tests/...
```

# External references

https://xlinux.nist.gov/dads/

## Disclosure

Unit tests and explanations written by claude
