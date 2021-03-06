# Tipos

Los siguientes tipos de datos estan disponibles en Golang:

* bool
* Numeric Types:
  * int8, int16, int32, int64, int
  * uint8, uint16, uint32, uint64, uint
  * float32, float64
  * complex64, complex128
  * byte
  * rune
* string

`bool` representa valores boleanos `true` o `false`.

Para tipos numericos, el numero que sigue al tipo indica el numero de bits que se usa para representar el valor en memoria. El numero de bits determina que tan grande el numero puede ser.

---
`int8`: 8 bit signed integers

`int16`: 16 bit signed integers

`int32`: 32 bit signed integers

`int`: 32 or 64 bit integers depending on the underlying platform.

`uint` : 32 or 64 bit unsigned integersdepending on the underlying platform.

`float32`: 32 bit floating point numbers

`float64`: 64 bit floating point numbers

`byte` is an alias of `uint8`

`rune` is an alias of `int32`

---

`complex64`: complex numbers which have float32 real and imaginary parts

`complex128`: complex numbers with float64 real and imaginary parts

La funcion `complex` se utiliza para construir un numero complejo con partes reales e imaginarias:

```go
func complex(r, i FloatType) ComplexType
```
---

`string` es una coleccion de caracteres encerrados entre comillas.

```go
first := "Allen"
last := "Varghese"
name := first +" "+ last
fmt.Println("My name is",name)
```

---

El tipo de una variable se puede imprimir usando el especificador de formato `%T` disponible en el metodo `Printf`.

```go
var a = 10
fmt.Printf("a is of type %T, value is ", a, a)
```

---

## Conversion de tipos

No hay conversion automatica de tipos en Golang. Se utilizan funciones de tipo para hacer una conversion.

```go
int(<float value>)
float64(<integer value>)
```

## Constantes

Las constantes son valores que no cambian una vez que se asignan a una variable. Se usa el keyword `const` en vez de `var` para declarar un constante.

```go
const a bool = true
const b int32 = 32890
const c string = "Something"
```
