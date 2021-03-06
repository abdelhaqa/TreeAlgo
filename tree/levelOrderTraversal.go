package tree

import (
	"fmt"
)

// Perform Level Order traversal and print content on console
// Uses queue to push Tree nodes
func levelOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	queue := New()
	queue.enqueue(root)

	for queue.isEmpty() != true {
		n := queue.dequeue()
		fmt.Printf("%d ", n.val)

		if n.left != nil {
			queue.enqueue(n.left)
		}

		if n.right != nil {
			queue.enqueue(n.right)
		}

	}
}
