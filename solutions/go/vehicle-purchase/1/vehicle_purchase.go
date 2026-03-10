package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	//panic("NeedsLicense not implemented")
    if(kind == "car" || kind == "truck"){
        return true
    }
    return false
    
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	//panic("ChooseVehicle not implemented")
    // Se a opção 1 vier antes no alfabeto (for "menor")
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	
	// Se não entrou no if acima, a opção 2 é a vencedora
	return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//panic("CalculateResellPrice not implemented")
    if(age < 3){
        return float64(originalPrice*0.8)
    } else if (age >= 3 && age < 10){
        return float64(originalPrice*0.7)
    } else{
        return float64(originalPrice*0.5)
    }
}
