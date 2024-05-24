package _interface

type Learn interface {
	LearnC(string) (string, error)
	LearnPHP(string) (string, error)
	LearnGO(string) (string, error)
}
