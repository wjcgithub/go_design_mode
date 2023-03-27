package prototype

import (
	"encoding/json"
	"time"
)

type Prototype interface {
}

// Keyword is the keyword
type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone is the clone of keyword
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword

	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords is the map of keyword
type Keywords map[string]*Keyword

// Clone is the clone of keywords
func (words *Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range *words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
