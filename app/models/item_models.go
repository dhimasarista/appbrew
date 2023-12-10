package models

type Item struct {
	ID    int
	Name  string
	Price int64
}

var Items = []Item{
	{
		ID:    1023,
		Name:  "Cappucino",
		Price: 18000,
	}, {
		ID:    1024,
		Name:  "Matcha Latte",
		Price: 16000,
	},
}

// func (item *Item) FindAll()
