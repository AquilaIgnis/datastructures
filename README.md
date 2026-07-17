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
- [ ] Trie

## Sets

- [ ] Union Find
- [ ] Bloom filter
- [ ] HyperLogLog

# Algorithms ωψγ

- [x] Kadane's Algorithm
- [x] Euclidean gcd
- [x] Fibonacci
- [x] Heap Sort
- [ ] Modular Arithmetic
- [ ] Sieve of Eratosthenes
- [x] BFS Breadth-First Search
- [ ] DFS Depth-First Search
- [ ] Karatsuba algorithm

# Documentation

```bash
go doc -all ./linear | bat -l go
```

```bash
go doc -all ./algo | bat -l go
```

```bash
for p in ./trees/ ./trees/avl ./trees/heap; do go doc -all "$p"; done | bat -l go
```

# Tests

```bash
go test ./tests/...
```

## Disclosure

Unit tests and explanations written by claude
