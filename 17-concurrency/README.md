# Concurrency

En Go, una tarea que se ejecuta de forma independiente se conoce como `goroutine`. Podemos iniciar en nuestro programa tantas `goroutines` como se desee. La comunicación entre `goroutines` se realiza utilizando `channels`. 
Las `goroutine` son similares a las corutinas, fibras, procesos o hilos en otros lenguages, aunque no son exactamente iguales a ninguno de ellos. Son muy eficientes de crear y Go facilita la coordinación de muchas operaciones simultáneas. 

## Go Routines

Iniciar una `goroutine` es tan fácil como llamar a una función. Todo lo que se necesita es usar la palabra clave `go` frente a la invocación.
Cuando la función principal regresa, todas las `goroutines` en el programa se detienen inmediatamente, por lo que debemos utilzar algun mecanismo de espera para que esto no suceda. En el siguiente ejemplo utilizaremos `Sleep` para poder realizar una prueba básica:

```go
func main() {
	go f()
	time.Sleep(4 * time.Second)
}

func f() {
	time.Sleep(3 * time.Second)
	fmt.Println("go rutine")
}
```

## Channels

Se puede utilizar un `channel` para enviar valores de forma segura de una `goroutine` a otra. 
Para crear un `channel`, se utiliza la funcion `make` igual que para crear `maps` o `slices`. Los `channels` tienen un tipo que se especifica cuando se crean y no puede ser cambiado. 
Por ejemplo, el siguiente `channel` se crea para manejar mensajes de tipo `int`, o sea que solamente se puede utilizar para enviar o recibir valores enteros:

```go
c := make(chan int)
```

Una vez que tengamos un `channel`, podemos enviarle valores y recibir los valores que se le envíen. Para esto se usa el operador `left arrow` (<-).
Para enviar un valor, apuntamos la flecha hacia la expresión del `channel`, como si la flecha fuera indicando que el valor de la derecha fluya hacia el `channel`.

```go
c <- 99
```

La operación de envío esperará hasta que "alguien" (otra  `goroutine`) intente recibir un valor en el mismo un `channel`. Mientras espera, el remitente no puede hacer nada más, aunque todas las demás `goroutines` seguirán ejecutándose libremente.

Para recibir un valor de un `channel`, la flecha apunta en dirección opuesta al `channel`:

```go
r := <-c
```

En el codigo anterior recibimos un valor del `channel` `c` y lo asignamos a la variable `r`. Al intentar recibir un valor desde el `channel` el receptor esperará hasta que otro `goroutine` intente enviar un valor en el mismo `channel`.

El siguiente ejemplo creamos un `channel` y lo pasamos a cinco `goroutines` creadas con la invocación a la funcion `f`.
Luego vamos a esperar recibir cinco mensajes, uno por cada `goroutine` que se ha iniciado. Cada `goroutines` duerme y luego envía un valor identificándose. Cuando la ejecución alcanza el final de la función principal, vamos a poder tener la certeza que todas las invocaciones a `f` finalizaron correctamente.

```go
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go f(i, c)
	}

	for i := 0; i < 5; i++ {
		routineId := <-c
		fmt.Println("go routine with Id ", routineId, " has finished")
	}
}
func f(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("go routine with Id ", id, " is running")
	c <- id
}
```