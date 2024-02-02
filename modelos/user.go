package modelos

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	CreatedAt time.Time
	status    bool
}

func (u *Usuario) AddUser(id int, nombre string, status bool) {
	u.Id = id
	u.Nombre = nombre
	u.status = status
	u.CreatedAt = time.Now()
}
