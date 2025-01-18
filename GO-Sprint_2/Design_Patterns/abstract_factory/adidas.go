package main

type AdidasShoe struct {
	logo string
	size int
}

func (a *AdidasShoe) setLogo(logo string) {
	a.logo = logo
}

func (a *AdidasShoe) setSize(size int) {
	a.size = size
}

func (a *AdidasShoe) getLogo() string {
	return a.logo
}

func (a *AdidasShoe) getSize() int {
	return a.size
}

type AdidasShirt struct {
	logo string
	size int
}

func (a *AdidasShirt) setLogo(logo string) {
	a.logo = logo
}

func (a *AdidasShirt) setSize(size int) {
	a.size = size
}

func (a *AdidasShirt) getLogo() string {
	return a.logo
}

func (a *AdidasShirt) getSize() int {
	return a.size
}
