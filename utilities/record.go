package utilities

import (
	"os"
	"strconv"
)

//Record ...
type Record struct {
	items []Item
}

//Item ...
type Item struct {
	item   int
	values []string
}

//InsertItem ...
//need a pointer here since we'll be modifying the receiever
func (record *Record) InsertItem(item int, value string) {
	index := record.indexOf(item)
	if index != -1 {
		itemValues := record.items[index].values
		itemValues = append(itemValues, value)
	} else {
		itemValues := []string{value}
		record.items = append(record.items, Item{item: item, values: itemValues})
	}
}

//indexOf ...
//private function (not capitalized)
func (record Record) indexOf(item int) int {
	for index, value := range record.items {
		if value.item == item {
			return index
		}
	}
	return -1
}

//WriteToFile ...
func (record Record) WriteToFile(file *os.File) {
	for _, items := range record.items {
		for _, itemVal := range items.values {
			file.WriteString(strconv.Itoa(items.item) + "," + itemVal + "\n")
		}
	}
}
