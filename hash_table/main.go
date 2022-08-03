package hash_table

import "hash_example/node"

const TABLE_SIZE = 100

type HashTable struct {
  table [TABLE_SIZE]*node.Node
}

func (hashTable *HashTable) hash(key []byte) (hash uint64) {
  hash = 5381

  for _, value := range key {
    hash = hash * 33 + uint64(value)
  }

  return hash % TABLE_SIZE
}

func (hashTable *HashTable) Set(key string, value string) {
	hashIndex := hashTable.hash([]byte(key))

	currentNode := hashTable.table[hashIndex]

	newNode := node.New(key, value, nil)

	if currentNode == nil {
		hashTable.table[hashIndex] = &newNode
		return
	}

	if currentNode.Key == key {
		currentNode.Value = value
		return
	}

	for currentNode.Next != nil {
		if currentNode.Next.Key == key {
			currentNode.Next.Value = value
			return
		}

		currentNode = currentNode.Next
	}

	currentNode.Next = &newNode
}

func (hashTable *HashTable) Get(key string) *string {
	hashIndex := hashTable.hash([]byte(key))
	currentNode := hashTable.table[hashIndex]

	for currentNode != nil {
		if currentNode.Key == key {
			return &currentNode.Value
		}

		currentNode = currentNode.Next
	}

	return nil
}

func New() (hashTable HashTable){
  return
}
