package searching

import (
	"ds_algo_course/pkg/data_structures/queues"
	"ds_algo_course/pkg/data_structures/trees"
	"log"
)

func bfs(root *trees.BSTNode) []int {
	var traversedNodes []int

	queue := queues.NewLinkedListQueue()
	queue.Enqueue(root)

	for queue.Length() > 0 {
		dequeued, err := queue.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		currentNode := dequeued.(*trees.BSTNode)
		traversedNodes = append(traversedNodes, currentNode.Value)

		if currentNode.Left != nil {
			queue.Enqueue(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.Enqueue(currentNode.Right)
		}
	}

	return traversedNodes
}
