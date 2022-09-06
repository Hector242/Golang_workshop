package mypackage

import "fmt"

// Para poder exportar este struct Tenemos que declararlo en Mayuscula
// Al igual que todos sus componentes
// CarPublic va a ser un struct de acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// Para poder exportar una funcion es el mismo proceso
func PrintMessagePublic() {
	fmt.Println("This is your public message from mypackage")
}
