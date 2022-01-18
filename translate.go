package main

func translateDataTypeFS(input string) string {
	var output string
	switch input {
	case "0":
		output = "0"
	case "1":
		output = "1"
	case "2":
		output = "5"
	case "3":
		output = "2"
	case "4":
		output = "6"
	case "5":
		output = "3"
	case "6":
		output = "7"
	case "7":
		output = "1"
	case "8":
		output = "3"
	case "9":
		output = "5"
	case "10":
		output = "9"
	case "11":
		output = "TEXT"
	case "12":
		output = "4"
	case "13":
		output = "8"
	case "14":
		output = "10"
	default:
		output = "NaN"
	}

	return output
}

func translateDataTypeAS(dataTypeFS string, DdataTypeAS string) string {
	var output string
	switch dataTypeFS {
	case "7":
		output = "2"
	case "8":
		output = "2"
	case "9":
		output = "2"
	default:
		output = DdataTypeAS
	}

	return output
}

func loytecToModulo5DataTypeAS(registerType string) string {
	var output string
	switch registerType {
	case "INPUT":
		output = "2"
	case "DISCRETE":
		output = "2"
	case "HOLD":
		output = "0"
	default:
		output = "0"
	}

	return output
}

func loytecToModulo5DataTypeFS(dataType string) string {
	var output string
	switch dataType {
	case "bit":
		output = "0"
	case "uint8":
		output = "1"
	case "uint16":
		output = "3"
	case "uint32":
		output = "5"
	case "uint64":
		output = "12"
	case "int8":
		output = "2"
	case "int16":
		output = "4"
	case "int32":
		output = "6"
	case "int64":
		output = "13"
	case "float32":
		output = "10"
	case "float64":
		output = "14"
	default:
		output = "3"
	}

	return output
}

func loytecToModulo5FunctionCode(registerType string, Direction string) string {
	output := "0"

	if registerType == "DISCRETE" {
		output = "2"
	}

	if registerType == "COIL" {
		if Direction == "Input" {
			output = "1"
		} else if Direction == "Output" {
			output = "5"
		} else if Direction == "Value" {
			output = "5"
		}
	}

	if registerType == "INPUT" {
		output = "4"
	}

	if registerType == "HOLD" {
		if Direction == "Input" {
			output = "3"
		} else if Direction == "Output" {
			output = "6"
		} else if Direction == "Value" {
			output = "6"
		}
	}

	return output
}

// Just a mock translation for future implementation
func translateByteOrder(byteOrder string) string {
	return "1"
}

// Just a mock translation for future implementation
func translateWordOrder(byteOrder string) string {
	return byteOrder
}
