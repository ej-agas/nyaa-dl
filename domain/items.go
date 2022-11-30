package domain

type ItemCollection struct {
	Items []item `xml:"channel>item"`
}

func (i ItemCollection) Len() int {
	return len(i.Items)
}

func (it ItemCollection) Swap(i, j int) {
	it.Items[i], it.Items[j] = it.Items[j], it.Items[i]
}
