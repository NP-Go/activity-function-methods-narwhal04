package main

//Declare a struct

type custStruct struct {
	fName   string
	lName   string
	uName   string
	passwd  string
	email   string
	phone   int
	address string
}

//Create Methods
func (cs custStruct) userCredentials() (string, string) {
	return cs.uName, cs.passwd
}

func (cs custStruct) userAddress() string {
	return cs.address
}

func (cs custStruct) allUserInfor() (string, string, string, string, string, int, string) {
	return cs.fName, cs.lName, cs.uName, cs.passwd, cs.email, cs.phone, cs.address
}
