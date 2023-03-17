// @program:     LearnGo
// @file:        mockretriever.go
// @create:      2022-09-27 14:15
// @description:

package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
