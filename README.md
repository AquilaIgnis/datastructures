# Data Structures & Algorithms in Go

## Linear

- [x] Stack
- [x] Queue
- [x] Double Linked List
- [x] Circular Buffer
- [ ] Deque (segmented array), ⛔ Not possible in Go

## Tree — hierarchical, parent/child relationships

- [ ] Binary Search Tree
- [ ] AVL Tree
- [ ] Heap (min/max)
- [ ] Trie

## Hash Based — key/value

- [ ] Hash Map
- [ ] Hash Set

# Each category solves different problems:

- Linear — ordered data, undo/redo, scheduling
- Tree — searching, sorting, hierarchical data like file systems
- Graph — networks, maps, social connections, dependencies
- Hash — fast lookups, caching, counting
- Set — membership testing, deduplication

# Algorithms ωψγ

- [x] Kadane's Algorithm
- [x] Euclidean gcd
- [x] Fibonacci
- [ ] Modular Arithmetic
- [ ] Sieve of Eratosthenes
- [ ] BFS Breadth-First Search
- [ ] DFS Depth-First Search
- [ ] Karatsuba algorithm

# Documentation

```bash
go doc -all ./linear | bat -l go
```

```bash
go doc -all ./algo | bat -l go
```

# Tests

```bash
go test ./tests/...
```
