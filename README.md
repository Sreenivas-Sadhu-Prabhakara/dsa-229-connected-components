# Number of Connected Components

> NeetCode 150 · Graphs · Blind 75 — the component count is how many times the outer loop has to launch a fresh DFS. Go, stdlib only, tested.

## Problem

How many connected components does an undirected graph on n labeled nodes have? (Isolated nodes count — the edgeless graph has n components.)

## Approach

The outer scan-and-launch loop that every traversal repo in this series carries for disconnected graphs *is* the algorithm here: each time the scan meets an unvisited node, that's a new component — increment and flood it. Number-of-islands (dsa-161) is this exact computation with the adjacency implicit in the grid. The union-find alternative shines when edges stream in or components must merge dynamically (see dsa-230, where that's the whole point); for a static edge list, DFS is less machinery.

**Complexity:** O(n + e) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
