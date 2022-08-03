package node

type Node struct {
	Next  *Node
	Key   string
	Value string
}

func New(key string, value string, next *Node) (node Node) {
  node = Node{
    Key: key,
    Value: value,
    Next: next,
  }

  return
}
