package domain

type Items struct {
	Items []item `xml:"channel>item"`
}

func (i Items) Len() int {
	return len(i.Items)
}

func (it Items) Swap(i, j int) {
	it.Items[i], it.Items[j] = it.Items[j], it.Items[i]
}
