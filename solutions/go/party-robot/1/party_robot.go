package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	//panic("Please implement the Welcome function")
	str := fmt.Sprintf("Welcome to my party, %s!", name)
    return str	    
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	//panic("Please implement the HappyBirthday function")
    str := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return str
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	//panic("Please implement the AssignTable function")
    str := fmt.Sprintf("Welcome to my party, %s!",name)
    str = str + "\n"
    str = str + fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.",table, direction, distance)
    str = str + "\n"
    str = str + fmt.Sprintf("You will be sitting next to %s.", neighbor)
    return str
}
