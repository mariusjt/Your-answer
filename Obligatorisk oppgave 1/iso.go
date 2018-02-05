package iso

import "fmt"




func IterateOverExtendedASCIIStringLiteral(ASCII string) {
	for i := 0; i < len(ASCII); i++ {
		fmt.Printf("%X %c %b\n", ASCII[i], ASCII[i], ASCII[i])
	}
}

func GreetingExtendedASCII() string {
	var extASCII []byte
	for i := 0x80; i <= 0xFF; i++ {
		extASCII = append(extASCII, byte(i))
	}

	greeting := "\x22\x53\x61\x6C\x75\x74\x2C\x20\xE7\x61\x20\x76\x61\x20\xB0\x2D\x29\x20\x80\x35\x30\x22"

	return greeting
}
