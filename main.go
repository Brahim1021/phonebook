package main

import (
	"app/phonebook"
	"fmt"
)
func main() {
	//contact := phonebook.Contact{}

	book := phonebook.PhoneBook{}

	book.Add("Ахмед", "Абакаров", "9380990202")
	book.Add("Ваха", "Садаев", "9954520010")
	book.Add("Халид", "Джамалдаев", "9380990202")
	book.Add("Расул", "IntoCode", "9954520010")
	book.Add("Мага", "Ваша", "9380990202")
	book.Add("Муса", "Абу-Даби", "9954520010")

	book.Print()

	book.ChangeName("Магомед")
	fmt.Println("После изменения МАГОМЕД --------")
	book.Print()

	favouirites := []string{
		"Расул",
		"Халид",
	}
	book.MarkFavourite(favouirites)
	fmt.Println("После объявления избранных--------")
	book.Print()

	fmt.Println("Вывод избранных--------")
	book.PrintFavourites()

	book.BlockContact("Ваха")


	fmt.Println("После фильтра\n Ожидаю все контакты кроме Вахи")
	book.Filter()
	fmt.Println("new")
	

}