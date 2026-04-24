package phonebook

import "fmt"

// Contact — отдельная запись в книге.
type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	IsFavourite  bool
	IsBlocked bool
}

// PhoneBook — сама телефонная книга.
type PhoneBook struct {
	Contacts []Contact
	nextID   int
}



func (b *PhoneBook) Add (firstName string, lastname string,numberPh string){
	
	if firstName == ""{
		fmt.Println("Имя не может быть пустым")
		return
	}
	if len([]rune(numberPh)) != 10{
		fmt.Println("Неверный формат номера")
		return
	}
	
	// fmt.Println(len([]rune(numberPh)))
	b.nextID++
	newContact := Contact{
		ID: b.nextID,
		FirstName: firstName,
		LastName: lastname,
		PhoneNumber: numberPh,
	}
	b.Contacts = append(b.Contacts, newContact)
}


func (b *PhoneBook)ChangeName(name string)bool{// меняю имя МАГА на МАГОМЕД
	for i := range b.Contacts{
		if b.Contacts[i].FirstName == "Мага"{
			b.Contacts[i].FirstName = name
			return true
		}

	}
	return false
}

func (b *PhoneBook)Print(){
	for _, r := range b.Contacts{
		fmt.Println("\n",r)
	}
}

func (b *PhoneBook)MarkFavourite (names []string){
	for i := range b.Contacts{
		
		for _, r := range names{
			if b.Contacts[i].FirstName == r{
				b.Contacts[i].IsFavourite = true
			}
		}
	}
}

func (b *PhoneBook) PrintFavourites(){
	for i := range b.Contacts{
		if b.Contacts[i].IsFavourite == true{
			fmt.Print("\n",b.Contacts[i])
		}  
	}
	
}

func (b *PhoneBook)BlockContact (name string){
	for i := range b.Contacts{
		if b.Contacts[i].FirstName == name{
			b.Contacts[i].IsBlocked = true
		} 
	}
}

func (b * PhoneBook)Filter (){
	for i := range b.Contacts{
		if b.Contacts[i].IsBlocked != true{
			fmt.Println(b.Contacts[i])
		}
	}
	
}

