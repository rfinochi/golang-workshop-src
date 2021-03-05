# Error Handling

## Errores vs Excepciones

Los errores en Golang no son excepciones, por convención, si una función puede fallar deber retornar un tipo `error`.
El tipo error contiene la infomacion del problema y si es `nil` quiere decir que no hubo errores

```go
func calculate(a, b int) (int, error) { }

result, err := caculate(a, b)

if err != nil {
    // handle the error
}
// continue
```

## Defer

Una declaracion `defer` guarda la llamada a la funcion en una lista. Una vez finalizada la ejecucíon de la funcion circundante, se ejecutan todas las llamadas a las funciones diferidas que se guardaron previamente en esa lista.
`defer` se usa comunmente para simplificar acciones de limpieza, o combinado con `recover` para recuperarnos de un error insalbable (`panic`)

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

## Panic y Recover

`panic` detiene el flujo de control ordinario y comienza un proceso de pánico `(panicking)`. 
Cuando la funcion F llama a `panic`, la ejecucion de F se detiene, las funciones diferidas (que utilizaron `defer`) en F se ejecutan normalmente y luego F vuelve a su llamador.
Para el llamador, F se comporta como una llamada al pánico, y vuelve a su llamador, el proceso continua hasta que todas las funciones retornan, momento en el cual el programa falla.

Para evitar que el pánico termine el programa, Go proporciona una función de recuperación `recover` la cual recupera el control de una funcion en pánico.
La recuperación solo es útil dentro de las funciones diferidas, Durante la ejecucion normal, una llamada a `recover` devolvera `nil` y no tendrá otro efecto.
Si la función actual está en pánico, una llamada a recuperar capturara el valor que se especifico al crear el pánico y reanudara la ejecucion normal.

```go
defer func() {
    //Recovers from panic
    if e := recover(); e != nil {
        //Prints PANIC ATTACK
        fmt.Println(e)
    }
}()

//Causes panic
panic("PANIC ATTACK")
```

## Rob Pike, Go Proverbs
- Errors are values.
- Don’t just check errors, handle them gracefully.
- Don’t panic.
- Make the zero value useful.
- The bigger the interface, the weaker the abstraction.
- interface{} says nothing.
- Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.
- Documentation is for users.
- A little copying is better than a little dependency.
- Clear is better than clever.
- Concurrency is not parallelism.
- Don’t communicate by sharing memory, share memory by communicating.
- Channels orchestrate; mutexes serialize.