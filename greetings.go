package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//funcion hola a una persona especifica

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(random(), name)
	return message, nil
}

// para obtener diferentes nombres
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Mensajes aleatorios , saludos aleatorios
func random() string {

	formatos := []string{
		"Hola, %v Bienvenido",
		"Que bueno verte, %v",
		"Saludo, %v",
	}
	return formatos[rand.Intn(len(formatos))]
}
