package main

var items = []Item{}

type InMemory struct {
}

func (InMemory) CreateItem(newItem Item) {
	items = append(items, newItem)
}

func (InMemory) UpdateItem(updatedItem Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

func (InMemory) GetItems() []Item {
	return items
}

func (InMemory) GetItem(id int) Item {
	var result Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}

	return result
}

func (InMemory) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
