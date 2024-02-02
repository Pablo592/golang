package users

import (
	"fmt"
	"golang/modelos"
)

func AltaUsuario() {
	u := new(modelos.Usuario)
	u.AddUser(1, "Eduar", true)
	fmt.Println(u)

}
