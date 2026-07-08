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
- [ ] Heap (min/max)
- [ ] Trie

## Hash Based — key/value

- [ ] Hash Map
- [ ] Hash Set

# Algorithms ωψγ

- [x] Kadane's Algorithm
- [x] Euclidean gcd
- [x] Fibonacci
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
go doc -all ./trees/ | bat -l go && go doc -all ./trees/avl  | bat -l go
```

# Tests

```bash
go test ./tests/...
```

## Disclosure

Unit tests and explanations written by claude
