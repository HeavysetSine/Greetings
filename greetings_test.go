package greetings

import (
	"regexp"
	"testing"
)

// dos tipos de test usamos
func TestHelloName(t *testing.T) { //testear la varianle que le pasamos
	name := "Jaime"                                //variable creada para el test
	want := regexp.MustCompile(`\b` + name + `\b`) //busca la condicidencia en la variable
	mensaje, err := Hello("Jaime")                 //uso de la funcion hello usando los dos valores mensaje y el err
	//want rectifica que la variable juan concida en algun lugar del codigo
	//si se presenta un error se ejecuta el iff
	if !want.MatchString(mensaje) || err != nil {
		t.Fatalf(`Hello("juan") = %q,%v, quiere coincidencia para %#q, nil`, mensaje, err, want) //regresamos un objeto con el error

	}
}

// Comprobar que nos envie un error,forzamos el error
func TestHelloEmpty(t *testing.T) {
	mensaje, err := Hello("")
	if mensaje != "" || err != nil {
		t.Fatalf(`Hello("") = %q,%v, quiere "", eror`, mensaje, err)
	}
}
