package idGeneration

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	node *snowflake.Node
	once sync.Once
)

func GetInstance() *snowflake.Node {
	once.Do(func() {
		node, _ = snowflake.NewNode(1)
	})
	return node
}

func GeneratelID() (int64, error) {
	// Generate a snowflake ID.
	id := GetInstance().Generate()

	return int64(id), nil
}
