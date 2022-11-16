package variablelengthquantity

func clearFirst(i byte) byte {
	return i ^ 128
}

func setFirst(i byte) byte {
	return i | 128
}

func EncodeVarint(input []uint32) []byte {
	panic("Please implement the EncodeVarint function")
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the EncodeVarint function")
}
