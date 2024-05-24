package implement

type BiliBiliInterface interface {
	LearnC(string) (string, error)
	LearnPHP(string) (string, error)
	LearnGO(string) (string, error)
}

type BiliBili struct {
}

func NewBiliBili() *BiliBili {
	return &BiliBili{}
}
func (b *BiliBili) LearnC(content string) (string, error) {

	return "on bilibili learn c", nil
}

func (b *BiliBili) LearnPHP(content string) (string, error) {

	return "on bilibili learn php", nil
}

func (b *BiliBili) LearnGO(content string) (string, error) {

	return "on bilibili learn go", nil
}
