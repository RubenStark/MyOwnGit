package command

type Command struct {
	Name string
	Args []string
}

func (c Command) String() string {
	return c.Name
}
