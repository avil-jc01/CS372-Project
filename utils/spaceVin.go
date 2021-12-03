package utils

import (

)

func SpacedVin(vin string) string {
	var returnVin string

	for i, char := range vin {
		returnVin = returnVin + string(char) + "   "

		if i % 2 == 0 && i > 0 {
			returnVin = returnVin + " "

		}
		if i % 8 == 0 && i > 0 {
			returnVin = returnVin[:len(returnVin)-1]

		}
		if i % 9 == 0 && i > 0 {
			returnVin = returnVin + " "

		}

		
	}

	if len(returnVin) > 73 {
		returnVin = returnVin[:73]
	}
	return returnVin
}
