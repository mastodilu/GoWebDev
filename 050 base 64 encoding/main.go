package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "La prima formazione stabile del gruppo risale al 1972, al tempo della registrazione del primo disco di Springsteen, anche se il nome E Street Band ha cominciato ad essere utilizzato pubblicamente nell'estate del 1974. Alcuni dei musicisti però suonavano già con Springsteen in precedenza, in quella che era nota come Bruce Springsteen Band e prima ancora, partendo dagli Steel Mill nel 1969, in altre formazioni dalla vita breve che anticiparono l'epoca del successo del cantautore del New Jersey."
	fmt.Printf("base 64 encoding of:\n\n%v\n\n", s)

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
}
