package goreds

import (
	"testing"
)

func BenchmarkSplitEnFromCh(b *testing.B) {
	teststr := "日本双语幼儿园「Kids Duo International Center 南」导入了俄罗斯Universal Terminal Systems开发的AR沙场「iSandBOX」，今年4月AR沙场将作为学校的一部分课程计划开始实施。通过iSandBOX，孩子们可以创作自己喜欢的风景，观察自然现象和生态，也可以能掌握团队合作与协调性。另外值得关注的是，这家学校在上这门课时将全部使用英语实践，在实践过程中还能提高孩子们的词汇量和英语口语。 "

	for i := 0; i < b.N; i++ {
		SplitEnFromCh(teststr)
	}
}
