//Esta línea define al código como el paquete principal.
package main

/*
Se pueden declarar variables fuera de las funciones,
estas son llamadas funciones globales y pueden ser
usadas a lo largo de todo el código.
Sin embargo son consideradas malas prácticas y en su
lugar es mejor usar los argumentos de las funcioes que
declaremos.
*/
//var dog string = "dog name"

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
	El alcance de una variable se refiere al espacio donde existe
	dicha variable.
	Una variable definida dentro de una función es llamada "local"
	y solo existe dentro de esa función.
	*/
	var name string = "Nombre de alguien"
	fmt.Println(name) //Si se intenta llamar esta variable desde otra función que no sea esta, se dará un error pues es local.

	/*
	La declaración defer sirve para que una función especifica no
	se ejecute hasta que la función en la que está contenida se
	ejecute.
	Defer también puede ser usado para la limpieza de recursos
	usados en el código. Esto lo aprendo después.
	*/
	defer bienvenido("Juan") //Esta línea no se ejecutará hasta el final de este bloque de código.

	/*
	Si usamos la declaración defer en múltiples funciones dentro
	de un mismo bloque de código, estas se ejecutan desde la última
	que escribimos hacia la primera.
	De atrás para adelante.
	*/
	fmt.Println("Comienzo")
	for m := 10; m >= 0; m-- {
	    defer fmt.Println(m) //este código imprimer los números del 0 al 10, si no se usara el defer los imprimiría del 10 al 0 en este punto exacto del bloque, sin esperar a que lo finalice
	}
	fmt.Println("Fin")


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
	fmt.Println(num1 / num2)
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
	/*
	Para aceptar input del usuario se usa la función "fmt.Scanln(&[variable])"
	El & antes de la variable se usa para retornar la dirección de la variable.
	*/
	fmt.Println("Ingrese un número: ")
	var input int
	fmt.Scanln(&input) //Scanln puede aceptar diversos tipos de datos.
	fmt.Println(input)
	/*
	Se puede usar la declaración "if" para tomar decisiones y
	hacer que una sección del código corra sólo si cumple una
	condición específica.
	*/
	fmt.Println("Ingrese su edad: ")
	var edad int
	fmt.Scanln(&edad)
	if edad > 18 {
		fmt.Println("Puede ingresar.")
	} else {
		fmt.Println("No puede ingresar.") /*
		Else puede ser usado para ejecutar un código si es que
		la condición de la declaración if resultó falsa.
		*/
	}
	/*
	Cuando se necesita declarar una variable únicamente para la declaración if
	se puede declarar la variable en la línea de la declaración if. Las variables
	declaradas de este modo solo existen dentro del bloque if.
	*/
	if d := 5; d > 18 { //El punto y coma después de la declaración de la variable separa la declaración de la condición.
		fmt.Println("Es mayor de edad.")
	} else {
		fmt.Println("No es mayor de edad.")
	}
	/*
	En caso de que queramos establecer diferentes resultados
	posibles usando if/else, podemos usar la declaración
	"switch" para hacer el código más corto y fácil de leer.
	Igual que en una declaración if se pueden declarar
	variables en la línea misma del switch.
	En caso de que se cumplan dos casos, Go ejecuta solo el
	primero y salta el resto.
	*/
	var d int
	fmt.Println("Ingrese su edad.")
	fmt.Scanln(&d)
	switch d {
	case 1: //Si "d" es 1 imprime.
		fmt.Println("El número es 1.")
	case 2: //Si "d" es 2 imprime.
		fmt.Println("El número es 2.")
	default: //Si "d" no es ninguna de las dos imprime.
		fmt.Println("El número no es ni 1 ni 2.")
	}
	//También se puede usar la declaración switch sin una expresión.
	switch {
	case d > 18 && d < 21:
		fmt.Println("Mayor de edad pero no puede beber.")
	case d < 18:
		fmt.Println("Menor de edad.")
	default:
		fmt.Println("Mayor de edad y puede beber.")
	}
	/*
	Go permite hacer uso de loops que son declaraciones
	que se ejecutan hasta que una condición se vuelva
	verdadera. La única declaración loop que existe en
	Go es "for".
	La declaración "for" se realiza con el formato:
	for [init]; [condición]; [declaración posterior] {[código a ejecutar]}
	El código después de esto escribe los números desde
	15 al 20 mientras que el valor de e sea menor que
	20.
	*/
	for e:=15; e<=20; e++ { //e++ es otra forma de escribir e+=1 o e=e+1
		fmt.Println(e)
	}
	/*
	Se puede omitir las partes con el init y la condición
	de la declaración for. Lo cual la vuelve parecida a
	las declaraciones while de otros lenguajes.
	El código de abajo escribe el resultado de sumar todos
	los números del 1 al 1000 entre sí.
	*/
	sum := 1
	res := 0
	for sum<=1000 {
		res += sum
		sum++
	}
	fmt.Println(res)
	//Imprimir el nombre del número en inglés.
	var f int
	var g int
	var h int
	fmt.Println("Ingrese un número: ")
	fmt.Scanln(&f)
	fmt.Println("Ingrese otro número: ")
	fmt.Scanln(&g)
	fmt.Println("Ingrese otro número: ")
	fmt.Scanln(&h)
	switch {
	case f==0:
		fmt.Println("Zero")
	case f==1:
		fmt.Println("One")
	case f==2:
		fmt.Println("Two")
	case f==3:
		fmt.Println("Three")
	case f==4:
		fmt.Println("Four")
	case f==5:
		fmt.Println("Five")
	case f==6:
		fmt.Println("Six")
	case f==7:
		fmt.Println("Seven")
	case f==8:
		fmt.Println("Eight")
	case f==9:
		fmt.Println("Nine")
	case f==10:
		fmt.Println("Ten")
	default:
		fmt.Println("El número no va de 0-10")
	}
	switch g {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("El número no va de 0-10")
	}
	switch h {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("El número no va de 0-10")
	}
	/*
	Las funciones permiten definir bloques de código y
	asignarles nombres para poder ejecutarlas cuantas
	veces queramos solo con el nombre.
	Un ejemplo de función es "Println".
	Para ejecutar una función usamos el nombre seguido
	de paréntesis.
	*/
	fmt.Println("Chau") //Println es la función, "Chau" es el argumento que le asignamos.
	//Podemos llamar funciones designadas por nosotros.
	bienvenido("Panchito")
	suma (12, 4)
	resultado := adición(12, 39)
	fmt.Println(resultado + 5)
	fmt.Println(división(10, 3))
}


/*
Además de las funciones definidas en los paquetes
de Go podemos declarar las nuestras usando la
palabra clave "func" seguida del nombre que le
vayamos a asignar junto a unos paréntesis.
También podemos asignarle argumentos a las funciones
para que tomen input.
Por ejemplo se puede agregar el argumento "nombre"
a la función que creamos para que se use en el
output.
*/
func bienvenido(nombre string) {
	fmt.Println("Bienvenido, "+nombre)
}
/*
La función declarada arriba permite a la función
tomar un argumento del tipo string.
El argumento es utilizado como si fuera una variable
dentro del cuerpo de la función y usa su valor en
el output.
Para llamar a la función con sus argumentos pasamos
un valor string dentro de los paréntesis al invocarla.
*/

/*
También se puede hacer que una función tome más de
un argumento usando comas para separarlas al momento
de declarar la función.
*/
func suma(a int, b int) {
	fmt.Println(a+b)
}
/*
Si los argumentos son del mismo tipo podemos solo
especificar el tipo del segundo.
por ejemplo func suma(a, b int) en lugar de lo
escrito arriba.
*/

/*
Se puede usar la declaración return al final de la
función que declaremos para asignar el valor que
obtengamos.
El tipo de return debe ser escrito después de los
argumentos que definamos.
*/
func adición(a, b int) int {
    return a+b //Asigna el valor de a+b a la función.
} // Podemos asignar el valor retornado de esta función a una variable.

/*
Go permite retornar varios valores con una sola
función.
Útil para imprimir el resultado de una operación
y los errores que se puedan haber dado.
*/
func división(a,b int) (int, int) {
    return a/b, a%b
}
