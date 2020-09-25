package builder

type complexStruct struct {
	ID   string
	Name string
	Age  string
}

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetName(name string) *Builder {

}
