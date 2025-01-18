// nike.go
package main

type NikeShoe struct {
	logo string
	size int
}

func (n *NikeShoe) setLogo(logo string) {
	n.logo = logo
}

func (n *NikeShoe) setSize(size int) {
	n.size = size
}

func (n *NikeShoe) getLogo() string {
	return n.logo
}

func (n *NikeShoe) getSize() int {
	return n.size
}

type NikeShirt struct {
	logo string
	size int
}

func (n *NikeShirt) setLogo(logo string) {
	n.logo = logo
}

func (n *NikeShirt) setSize(size int) {
	n.size = size
}

func (n *NikeShirt) getLogo() string {
	return n.logo
}

func (n *NikeShirt) getSize() int {
	return n.size
}
