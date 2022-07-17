//Esta línea define al código como el paquete principal.
package main

/*
Estas líneas se usan para importar paquetes con distintas
funciones. fmt es usado para que el programa acepte imput
y output. math es usado para operaciones aritméticas.
*/
import (
	"fmt"
	//"math"
)

/*
func main() no acepta argumentos, es la función que corre
cuando ejecutamos el programa.
*/
func main(){
	/*
	Para usar funciones específicas de algún paquete importado
	debemos poner el nombre del paquete seguido de un punto y
	luego el nombre de la función. Además el nombre de la función
	debe iniciar con mayúscula.
	*/
	fmt.Println("Hello, world!")
	/*
	Se pueden declarar variables usando la palabra "var" seguida
	de el nombre que le queremos asignar a la variable y luego
	el tipo de variable que le asignamos y luego podemos pasarla
	por un output.
	En el ejemplo que pongo después declaro la variable num como
	un entero y le asigno el valor 10.
	*/
	var num int = 10
	fmt.Println(num)
	/*
	Puedo declarar varias variables en una sola línea con un solo
	"var". Además, si le asigno un valor puedo evitar asignarle
	un tipo. El proceso de asignar valores a una variable se le
	denomina inicialización.
	*/
	var num1, num2 = 20, 30
	fmt.Println(num1, num2)
	/*
	Con GO se puede usar ":=" para declarar una variable, darle
	un valor y establecer su tipo sin necesidad de escribir todo
	el proceso de inicialización.
	*/
	num3 := 40
	fmt.Println(num3)
	/*
	GO acepta diferentes tipos de valores para las variables.
	float32 = un número decimal con 8 valores como máximo
	después del punto.
	float64 = un número decimal con 16 valores como máximo
	después del punto.
	string = texto
	bool = un valor booleano (verdadero o falso)
	Las valores que son declaradas sin un valor toman el valor
	correspondiente a 0 de su tipo de dato.
	*/
	var txt string = "Esto es una variable de texto."
	var dcm float32 = 3.141516
	var online bool
	fmt.Println(txt, dcm, online)
	//Las variables pueden cambiar de valor a lo largo del código.
	var x = 8
	fmt.Println(x)
	x = 10
	fmt.Println(x)
	/*Si necesitamos que una variable mantenga un solo valor en
	todo el código entonces declaramos una constante en lugar de
	una variable usando "const"
	*/
	const pi = 3.14
	fmt.Println(pi)
	//Adición
	fmt.Println(num + num2)
	//Sustracción
	fmt.Println(num3 - num)
	//Multiplicación
	fmt.Println(num1 * num2)
	//División
	fmt.Println(num1 / num)
	//Módulo (resto de la divisón)
	fmt.Println(num % num2)
	/*
	En GO se puede asignar operaciones aritméticas usan operadores
	acortados con formato "[operador]=".
	*/
	z := 15
	w := 32
	fmt.Println(z)
	z += w //z = z + w
	fmt.Println(z)
	/*
	Los operadores de relación se usan para comparar dos valores y
	dar como resultado un booleano: verdadero si la condición se
	cumple y falso si no lo hace.
	*/
	a := 1
	b := 2
	fmt.Println(a == b) //igual que
	fmt.Println(a != b) //diferente de
	fmt.Println(a > b) //mayor que
	fmt.Println(a < b) //menor que
	fmt.Println(a >= b) //mayor igual que
	fmt.Println(a <= b) //menor igual que
	/*
	Los operadores lógicos son usados para combinar dos o más
	condiciones.
	*/
	fmt.Println(a == b && a != b) //&& Ambas condiciones deben ser verdaderas para que retorne verdadero.
	fmt.Println(a == b || a != b) //||Si una o ambas condiciones son verdaderas entonces retorna verdadero.
	fmt.Println(!(a == b)) //! Si las condiciones son falsas, retorna verdadero.
	//Se pueden combinar varias condiciones y usar paréntesis para agruparlas.
	fmt.Println((x > 0 && x < 100) || x == 42)

}
