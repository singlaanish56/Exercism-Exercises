package gross


// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit := make(map[string]int)
	unit["quarter_of_a_dozen"] =3
	unit["half_of_a_dozen"] = 6
	unit["dozen"] = 12
	unit["small_gross"] = 120
	unit["gross"] = 144
	unit["great_gross"] = 1728

	return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
	if(!exists){
		return false
	}

	_, e := bill[item]
	if(e){
		bill[item]+=val
	}else{
		bill[item]=val
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	b, eb := bill[item]
	u, eu := units[unit]

	if(!eb || !eu){
		return false
	}

	if(b-u<0){
		return false
	}else if((b-u)==0){ 
		delete(bill, item)
	}else{
		bill[item]=b-u
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	a, b := bill[item]

	return a, b
}
