package domain

type Contact struct {
	Name string
	Number string
}

func (c *Contact) GetName() (string) {
    return c.Name
}

func (c *Contact) GetNumber() (string) {
    return c.Number
}
