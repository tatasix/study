package implement

type ClassRoomInterface interface {
	LearnC(string) (string, error)
	LearnPHP(string) (string, error)
	LearnGO(string) (string, error)
}

type ClassRoom struct {
}

func NewClassRoom() *ClassRoom {
	return &ClassRoom{}
}

func (c *ClassRoom) LearnC(name string) (string, error) {
	return "on classroom learn c", nil
}

func (c *ClassRoom) LearnPHP(name string) (string, error) {
	return "on classroom learn php", nil
}

func (c *ClassRoom) LearnGO(name string) (string, error) {
	return "on classroom learn go", nil
}
