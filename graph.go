
// 1. Dfs graph traversal

package main

import (
    "fmt"
    "sync"

    "github.com/cheekybits/genny/generic"
)


type Item generic.Type


type Node struct {
    value Item
}

func (n *Node) String() string {
    return fmt.Sprintf("%v", n.value)
}


type ItemGraph struct {
    nodes []*Node
    edges map[Node][]*Node
    lock  sync.RWMutex
}


func (g *ItemGraph) AddNode(n *Node) {
    g.lock.Lock()
    g.nodes = append(g.nodes, n)
    g.lock.Unlock()
}


func (g *ItemGraph) AddEdge(n1, n2 *Node) {
    g.lock.Lock()
    if g.edges == nil {
        g.edges = make(map[Node][]*Node)
    }
    g.edges[*n1] = append(g.edges[*n1], n2)
    g.edges[*n2] = append(g.edges[*n2], n1)
    g.lock.Unlock()
}


func (g *ItemGraph) String() {
    g.lock.RLock()
    s := ""
    for i := 0; i < len(g.nodes); i++ {
        s += g.nodes[i].String() + " -> "
        near := g.edges[*g.nodes[i]]
        for j := 0; j < len(near); j++ {
            s += near[j].String() + " "
        }
        s += "\n"
    }
    fmt.Println(s)
    g.lock.RUnlock()
}
$go run main.go
A -> B C D
B -> A E
C -> A E
D -> A
E -> B C F
F -> E

PASS

// 2.Bfs graph traversal

package graph

import "sync"


type NodeQueue struct {
    items []Node
    lock  sync.RWMutex
}


func (s *NodeQueue) New() *NodeQueue {
    s.lock.Lock()
    s.items = []Node{}
    s.lock.Unlock()
    return s
}

func (s *NodeQueue) Enqueue(t Node) {
    s.lock.Lock()
    s.items = append(s.items, t)
    s.lock.Unlock()
}


func (s *NodeQueue) Dequeue() *Node {
    s.lock.Lock()
    item := s.items[0]
    s.items = s.items[1:len(s.items)]
    s.lock.Unlock()
    return &item
}

func (s *NodeQueue) Front() *Node {
    s.lock.RLock()
    item := s.items[0]
    s.lock.RUnlock()
    return &item
}

func (s *NodeQueue) IsEmpty() bool {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return len(s.items) == 0
}

func (s *NodeQueue) Size() int {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return len(s.items)
}

$go run main.go
A
B
C
D
A
E
F
