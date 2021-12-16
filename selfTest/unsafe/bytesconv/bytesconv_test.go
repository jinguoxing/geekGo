package bytesconv

import "testing"

var testString = "Albert Einstein: Logic will get you from A to B. Imagination will take you everywhere."
var testBytes = []byte(testString)


func rawBytesToStr(b []byte) string {

	return string(b)
}

func rawStrToBytes(s string)[]byte {
	return []byte(s)
}

func BenchmarkStrToBytesRaw(b *testing.B) {

	for i := 0; i <= b.N; i++ {
		rawStrToBytes(testString)
	}
}

func BenchmarkStrToBytes(b *testing.B){

	for i:=0;i<=b.N;i++{
		StringToBytes(testString)
	}

}



func BenchmarkRawBytesToStr(b *testing.B){

	for i:=0 ;i<=b.N;i++{
		rawBytesToStr(testBytes)
	}
}


func BenchmarkBytesToString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BytesToString(testBytes)
	}
}






