package idGeneration

import (
	"log"
	"testing"
)

func TestGeneratelID(t *testing.T) {

	for i := 0; i < 10; i++ {
		id, err := GeneratelID()

		if err != nil {
			log.Println("生成失败！")
		}
		log.Println("id", i, ": ", id)
	}
}
