package _interface

type MyLearnProxy interface {
	LearnGO(string) (string, error)
}

type MyPlan struct {
	learnEvent MyLearnProxy
}

func NewMyPlan(event MyLearnProxy) *MyPlan {
	return &MyPlan{learnEvent: event}
}

func (m *MyPlan) Learn(res string) (string, error) {
	return m.learnEvent.LearnGO(res)
}
