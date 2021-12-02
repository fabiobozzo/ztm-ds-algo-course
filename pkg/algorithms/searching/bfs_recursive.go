package searching

import (
	"ds_algo_course/pkg/data_structures/queues"
	"ds_algo_course/pkg/data_structures/trees"
	"log"
)

func recursiveBfs(queue queues.Queue, traversedNodes []int) []int {
	if queue.Length() == 0 {
		return traversedNodes
	}

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

	return recursiveBfs(queue, traversedNodes)
}
