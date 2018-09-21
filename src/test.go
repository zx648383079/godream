package main

import (
	"fmt"
)

type Hub struct {
	width, height int
}

func (r Hub) size() int  {
	return r.width * r.height
}

func (r *Hub) bound() int  {
	return r.width * r.height
}

func (r Hub) setX(x int)  {
	r.width = x
}

func (r *Hub) setY(y int)  {
	r.height = y
}

func main()  {
	hub := Hub{width: 1, height: 2}

	fmt.Println(hub)
	fmt.Println(hub.size())
	fmt.Println(hub.bound())

	hub.setX(10)
	hub.setY(20)

	fmt.Println(hub)
	fmt.Println(hub.size())
	fmt.Println(hub.bound())
}

