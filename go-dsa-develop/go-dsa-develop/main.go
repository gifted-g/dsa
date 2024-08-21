package main

import "fmt"

func main() {
	sb := NewStringBuilder()

	sb.Append("Hello ")
	sb.Append("World")

	fmt.Println(sb.ToString())

	fmt.Println("***************************NewStringBuilder end***************************")

	ht := NewHashTable()

	ht.Put("wewe", "1")
	ht.Put("w", "2")
	ht.Put("helwelo", "3")
	ht.Put("helwewewelo", "4")
	ht.Put("hewewello", "5")
	ht.Put("hedllo", "6")
	ht.Put("helldo", "7")
	ht.Put("hello", "8")
	ht.Put("helflo", "9")
	ht.Put("heallo", "10")
	ht.Put("hefllo", "11")
	ht.Put("heldflo", "12")
	ht.Put("hellvdo", "13")
	ht.Put("helerrrwlo", "14")

	fmt.Println(ht.Get("wewe"))
	fmt.Println(ht.Get("w"))
	fmt.Println(ht.Get("helwelo"))
	fmt.Println(ht.Get("helwewewelo"))
	fmt.Println(ht.Get("hewewello"))
	fmt.Println(ht.Get("hedllo"))
	fmt.Println(ht.Get("helldo"))
	fmt.Println(ht.Get("hello"))
	fmt.Println(ht.Get("helflo"))
	fmt.Println(ht.Get("heallo"))
	fmt.Println(ht.Get("hefllo"))
	fmt.Println(ht.Get("heldflo"))
	fmt.Println(ht.Get("hellvdo"))
	fmt.Println(ht.Get("helerrrwlo"))

	fmt.Println("***************************NewHashTable end***************************")

	al := NewArrayList()

	al.Add("NewArrayList: Hello")
	al.Add(" ")
	al.Add(1)
	al.Add("NewArrayList: world")

	for _, val := range al.GetAll() {
		fmt.Println(val)
	}

	str := "obrgcs"

	fmt.Println("Is unique", IsUnique(str))
}
