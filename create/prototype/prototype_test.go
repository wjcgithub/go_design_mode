package prototype

import (
	"testing"
	"time"
)

// 给上面代码生成测试用例
func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	// 初始化数据
	words := Keywords{
		"hello": &Keyword{
			word:      "hello",
			visit:     100,
			UpdatedAt: &updateAt,
		},
		"world": &Keyword{
			word:      "world",
			visit:     200,
			UpdatedAt: &updateAt,
		},
	}

	// 生成更新数据
	updateWords := []*Keyword{
		&Keyword{
			word:      "hello",
			visit:     300,
			UpdatedAt: &updateAt,
		},
		&Keyword{
			word:      "world",
			visit:     400,
			UpdatedAt: &updateAt,
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
