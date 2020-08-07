package main

import (
	"fmt"
)

type Queue struct{
	slice []string
}

func (q *Queue) enqueue(i []string) {
	for _, e := range i {
		q.slice = append(q.slice, e)
	}
}

func (q *Queue) dequeue() string {

	ret := fmt.Sprintf("%v", q.slice[0])
	q.slice = q.slice[1:]

	return ret
}

func (q *Queue) isEmpty() bool {
	if len(q.slice) < 1 {
		return true
	}
	return false
}

func isTarget(s string) bool {
	if s == "thom" {
		return true
	}
	return false
}


func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


func breadthFirstSearch(graph map[string][]string, startNode []string) (bool, string) {
	var queue *Queue = new(Queue)
	var searched []string

	queue.enqueue(startNode)

	for {
		if queue.isEmpty() {
			break
		}

		node := queue.dequeue()
		if contains(searched, node) {
			continue
		}
		if isTarget(node) {
			fmt.Println("Tagrget node in graph --> ", node)
			return true, fmt.Sprintf("%v", node)
			break
		} else {
			str := fmt.Sprintf("%v", node)
			queue.enqueue(graph[str])
		}
		searched = append(searched, node)
	}

	return false, "(none)"

}

func main() {

	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	fmt.Println(breadthFirstSearch(graph, graph["you"]))

}
