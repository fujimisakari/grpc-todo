package domain

type Todo TodoDto

func (t *Todo) ToDto() *TodoDto {
	return (*TodoDto)(t)
}

func (t *TodoDto) ToEntity() (*Todo, error) {
	return (*Todo)(t), nil
}
