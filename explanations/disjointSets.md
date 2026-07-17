# Disjoint-Set (Union-Find) — Visual Explanations

A disjoint-set structure tracks a collection of elements split into
non-overlapping groups. It answers one question extremely fast:

> **"Are these two elements in the same group?"**

while allowing groups to be merged at any time.

Only two operations exist:

| Operation     | Meaning                                                      |
| ------------- | ------------------------------------------------------------ |
| `Find(x)`     | Which group does `x` belong to? (returns the group's _root_) |
| `Union(a, b)` | Merge the groups containing `a` and `b`                      |

Two elements are in the same group **iff** `Find(a) == Find(b)`.

---

## 1. How groups are stored: parent pointers

There is no list of groups anywhere. Each element stores exactly one
thing: **its parent**. A group is a tree, and the tree's **root** (the
one element that is its own parent) acts as the group's name.

```
parent array:   index:   0  1  2  3  4  5
                value:  [0, 0, 1, 3, 3, 5]
```

This array encodes three trees, i.e. three groups:

```
      0            3          5
     /            |
    1             4
    |
    2

  {0,1,2}       {3,4}        {5}
```

- `parent[0] == 0` → 0 is a root (its own parent)
- `parent[2] == 1`, `parent[1] == 0` → 2's chain leads to root 0
- The **shape** of the tree carries no meaning. Only "which root do
  you reach" matters. This fact is what makes path compression legal.

### Fresh initialization

`New(6)` makes every element its own root — six groups of one:

```
parent: [0, 1, 2, 3, 4, 5]

  0    1    2    3    4    5      (six separate trees)
```

---

## 2. Find — walking to the root

`Find(4)` on the arrangement above:

```
      3
      |
      4   ← start here, parent[4] = 3
```

Walk: `4 → 3`. `parent[3] == 3`, so 3 is the root. `Find(4)` returns 3.

`Find` never compares values or searches anything — it only follows
parent links upward until it finds a node that points at itself.

---

## 3. Union — root adoption

`Union(a, b)` does **not** link `a` and `b` directly. It:

1. finds the **root** of each,
2. points one root at the other.

Example: `Union(2, 4)` on our three trees.

```
Step 1: Find(2) → root 0        Find(4) → root 3

Step 2: point root 3 at root 0

BEFORE                      AFTER
      0        3                  0
     /         |                / | \
    1          4               1  .  3
    |                          |     |
    2                          2     4

parent: [0,0,1,3,3,5]       parent: [0,0,1,0,3,5]
                                            ↑
                                only ONE value changed: parent[3]
```

One write merged the entire groups: every element under 3 now reaches
root 0, because their chains pass through 3.

### Why roots, never elements directly

If `Union(2, 4)` had set `parent[4] = 2` (element to element):

```
      0
      |
      1
      |
      2
      |
      4        ← chain keeps growing with every careless union
```

Chains are the enemy: `Find` cost equals chain length. Linking
root-to-root keeps trees shallow; the next two techniques keep them
_very_ shallow.

---

## 4. Union by rank — shorter tree goes under taller

When merging two roots, we get to choose who adopts whom. The rule:
**attach the shorter tree under the taller root.**

```
Tall tree T (height 3)     Short tree S (height 1)

        T                        S
       /|\                       |
      . . .                      s
      |
      .
```

Option A — short under tall (what we do):

```
        T            height stays 3
       /|\ \         (S sits at depth 1, reaches depth 2 ≤ 3)
      . . . S
      |     |
      .     s
```

Option B — tall under short (what we avoid):

```
        S            height becomes 4!
        | \          (the whole tall tree got pushed one level deeper)
        s  T
          /|\
         . . .
           |
           .
```

Each root stores a `rank` — an upper bound on its tree height — used
only for this comparison. The one growth case: when both trees have
**equal** rank, the merged tree is necessarily one taller, so the
winning root's rank increments. Ranks never decrease, and ranks of
non-roots become stale garbage that is simply never read again.

With union by rank alone, tree height stays ≤ log n.

---

## 5. Path compression — Find repairs the tree as a side effect

The key idea: **`Find` is not a read-only query.** While answering
"what is x's root?", it rewires every node it walked past to point
_directly_ at that root. The walk happens once; afterwards those nodes
answer in a single hop.

### Worked example

A degenerate chain (parent: `[1, 2, 3, 4, 4]`):

```
        4    ← root
        ↑
        3
        ↑
        2
        ↑
        1
        ↑
        0    ← Find(0) starts here: 4 hops to the root
```

Call `Find(0)`. It walks `0 → 1 → 2 → 3 → 4`, learns the root is 4,
and on the way back **overwrites each node's parent with 4**:

```
BEFORE  parent: [1, 2, 3, 4, 4]        AFTER  parent: [4, 4, 4, 4, 4]

        4                                        4
        ↑                                     ↗ ↑ ↑ ↖
        3                                    0  1  2  3
        ↑
        2               ONE call to Find(0) flattened
        ↑               the entire chain into a star.
        1
        ↑
        0
```

Nothing about membership changed — all five elements are still one
group with root 4. But the _next_ `Find(0)` is 1 hop instead of 4.
The expensive walk prepaid for every future query on this path.

### Why this is legal

Because the tree shape is not information (section 1). Any shape that
preserves "everyone reaches root 4" is equivalent, so `Find` is free
to pick the flattest one.

### The two-pass mechanics

You can't write the answer before you know it. So:

```
Pass 1 (read-only):  walk up from x, discover the root.

        4  ← found it
        ↑
        3
        ↑          nothing written yet;
        2          the chain is still intact
        ↑
        1
        ↑
        0  ← started here

Pass 2 (writes):     re-walk the same chain from x,
                     paving each parent link with the root.

   visit 0:  parent[0] = 4
   visit 1:  parent[1] = 4
   visit 2:  parent[2] = 4
   visit 3:  parent[3] = 4   (was already 4)
```

```go
func (uf *UnionFind) Find(element int) int {
    // Pass 1: find the root (read-only)
    root := element
    for uf.parent[root] != root {
        root = uf.parent[root]
    }

    // Pass 2: point everyone on the path at the root
    current := element
    for uf.parent[current] != root {
        next := uf.parent[current]  // where the old chain goes
        uf.parent[current] = root   // pave over it
        current = next
    }

    return root
}
```

The recursive version hides the same two phases: calls going _down_
are pass 1; the assignments while unwinding are pass 2.

### Path halving — one-pass variant

Instead of two passes, skip each node to its **grandparent** while
walking. The path halves in length on every traversal:

```
BEFORE                 AFTER ONE Find(0)

  4                        4
  ↑                       ↑ ↑
  3                      3  2
  ↑                     ↑   ↑
  2                    (…)  0     ← 0 skipped over 1
  ↑                         ↑
  1                    every node now points
  ↑                    at its old grandparent
  0
```

```go
func (uf *UnionFind) Find(element int) int {
    for uf.parent[element] != element {
        uf.parent[element] = uf.parent[uf.parent[element]] // skip to grandparent
        element = uf.parent[element]
    }
    return element
}
```

Less thorough per call, but repeated calls flatten just as fast, and
it's the shortest non-recursive form — the common choice in practice.

---

## 6. Cost: why it's nearly free

| Variant                         | Cost per operation       |
| ------------------------------- | ------------------------ |
| Naive (no rank, no compression) | O(n) worst case — chains |
| Union by rank only              | O(log n)                 |
| Rank **+** path compression     | **amortized O(α(n))**    |

α is the inverse Ackermann function: **α(n) ≤ 4 for any n that fits
in the physical universe.** Effectively constant.

"Amortized" matters: one `Find` on a fresh tall tree can still cost
O(log n) — but that call flattens the path, making future calls
1-hop. Averaged over any operation sequence, each costs α(n).

```
Find(0) #1 on a 1,000,000-chain:   ~1,000,000 hops   (pays)
Find(0) #2:                         1 hop            (collects)
Find(500000):                       1 hop            (collects)
```

Compare with re-running BFS/DFS per connectivity query: O(n) _every_
time, with no memory of previous work. For q queries that's O(q·n) vs
Union-Find's O(q·α(n)) — the entire reason this structure exists.

---

## 7. Mental model: lazy signposts

You're at house 0 looking for the town hall. Each house only knows
"ask next door." You walk 0 → 1 → 2 → 3 → 4; house 4 says "I'm the
town hall." On your way back you update every signpost you passed:

```
before:  [→1] [→2] [→3] [→4] [HALL]
after:   [→4] [→4] [→4] [→4] [HALL]
```

Nobody ever does the long walk again. `Find` = asking directions and
fixing the signs; `Union` = one town hall deferring to another.

---

## 8. Cheat sheet

- A group = a tree; the root = the group's identity.
- `parent[x] == x` ⟺ x is a root.
- `Union` links **root to root**, never element to element.
- Tree shape carries no meaning → compression is always safe.
- `rank` is internal balancing bookkeeping, only valid on roots — not
  a ranking of your data.
- Merges only — no splitting. If you need "un-union", this is the
  wrong structure.
- `Union` returning false = "already connected" = **cycle detected**
  (this is the heart of Kruskal's algorithm).
