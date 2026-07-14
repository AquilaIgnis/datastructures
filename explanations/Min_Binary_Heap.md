# Explanations by AI using my code as example

## Min Binary Heap

> Source: `trees/heap/minHeap.go`

A binary heap is a complete binary tree flattened into a single slice
(`MinHeap.Array`). No node pointers — the tree shape lives in the index math.

### The invariant

Every parent is **≤** both its children. The smallest value therefore always sits
at the root (`Array[0]`). This is a *partial* order: siblings are unordered, so the
array is "loosely sorted", which is cheaper to maintain than a fully sorted array.

### Index math (implicit tree)

For a node at index `i`:

| Relation     | Index       |
| ------------ | ----------- |
| left child   | `2*i + 1`   |
| right child  | `2*i + 2`   |
| parent       | `(i-1) / 2` |

### The two repair operations

Both walk one root-to-leaf path, so both are **O(log n)**.

- **`siftUp(i)`** — used by `Insert`. A value that may be *too small* for its position
  bubbles **up**, swapping with its parent while it's smaller than the parent.
- **`siftDown(i)`** — used by `PopMin` and `Heapify`. A value that may be *too large*
  sinks **down**, repeatedly swapping with its **smallest** child until both children
  are ≥ it (or it hits a leaf).

### Methods

| Method       | What it does                                                    | Cost         |
| ------------ | -------------------------------------------------------------- | ------------ |
| `NewMinHeap` | returns an empty heap                                           | O(1)         |
| `Insert`     | appends to the end, then `siftUp` to restore the invariant     | O(log n)     |
| `PopMin`     | returns the root; see the swap-and-sink dance below            | O(log n)     |
| `Heapify`    | bulk-builds a heap from an arbitrary slice (empty heap only)   | O(n)         |

### PopMin, step by step

You can't just delete `Array[0]` — that leaves a hole. Instead:

1. Save the root (`Array[0]`) as the return value.
2. Move the **last** element into the root slot.
3. Zero the old last slot and shrink the slice by one (order matters — clear the slot
   *before* reslicing, or the index is out of range).
4. `siftDown(0)` to push that moved-up value back to its rightful depth.

Returns `(zero, false)` when the heap is empty.

> The zeroing (`Array[last] = zero`) only matters when `T` holds a pointer
> (e.g. `string`, which `cmp.Ordered` allows) — it lets the GC reclaim the dropped
> element instead of keeping it alive in the backing array.

### Heapify: why it's O(n), not O(n log n)

Inserting `n` items one by one would be O(n log n). `Heapify` is faster: it copies the
slice, then calls `siftDown` on every **non-leaf** node, from the last parent up to the
root:

```go
for i := len(h.Array)/2 - 1; i >= 0; i-- {
    h.siftDown(i)
}
```

Starting at `len/2 - 1` skips the leaves (they're already trivially valid heaps of
size 1). Working bottom-up means each `siftDown` sinks into subtrees that are *already*
valid heaps. Most nodes are near the bottom and barely move, so the total work sums to
O(n), not O(n log n).

It refuses to run on a non-empty heap (returns an error) to avoid clobbering existing
data.

### Summary

- **Array + index math** = a tree with no pointers and cache-friendly memory.
- **Invariant (parent ≤ children)** = the min is always `Array[0]`, peek is O(1).
- **`siftUp` / `siftDown`** = the O(log n) repairs that keep the invariant after an
  insert or a pop.
- **`Heapify`** = build the whole heap in O(n) by sinking non-leaves bottom-up.
