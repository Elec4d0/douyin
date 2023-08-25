package sensitiveWord

import (
	"bufio"
	"github.com/Chain-Zhang/pinyin"
	"log"
	"os"
	"regexp"
)

var trie *SensitiveTrie

// HansCovertPinyin 中文汉字转拼音
func HansCovertPinyin(contents []string) []string {
	pinyinContents := make([]string, 0)
	for _, content := range contents {
		chineseReg := regexp.MustCompile("[\u4e00-\u9fa5]")
		if !chineseReg.Match([]byte(content)) {
			continue
		}

		// 只有中文才转
		pin := pinyin.New(content)
		pinStr, err := pin.Convert()
		if err == nil {
			pinyinContents = append(pinyinContents, pinStr)
		}
	}
	return pinyinContents
}

func InitWords() {
	file, err := os.Open("./sensitiveWord.txt")
	if err != nil {
		log.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	var sensitiveWords []string
	for fileScanner.Scan() {
		sensitiveWords = append(sensitiveWords, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()

	// 汉字转拼音
	pinyinContents := HansCovertPinyin(sensitiveWords)
	trie = NewSensitiveTrie()
	trie.AddWords(sensitiveWords)
	trie.AddWords(pinyinContents) // 添加拼音敏感词
}

func ToInsensitive(content string) string {
	_, replaceText := trie.Match(content)
	return replaceText
}
