package main

import (
	"encoding/json"
	"testing"
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

// 给上面代码生成测试用例
func TestKeywords_Clone(t *testing.T) {
	// 初始化数据
	words := Keywords{
		"hello": &Keyword{
			word:  "hello",
			visit: 100,
		},
		"world": &Keyword{
			word:  "world",
			visit: 200,
		},
	}

	// 生成更新数据
	updateWords := []*Keyword{
		&Keyword{
			word:  "hello",
			visit: 300,
		},
		&Keyword{
			word:  "world",
			visit: 400,
		},
	}

	// 拷贝数据
	newWords := words.Clone(updateWords)

	// 检查是否拷贝成功
	if newWords["hello"].visit != 300 {
		t.Error("拷贝数据失败")
	}

	if newWords["world"].visit != 400 {
		t.Error("拷贝数据失败")
	}
}
