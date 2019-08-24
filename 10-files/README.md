# Archivos

## Lectura

La funcion `ReadFile` del paquete `ioutil` permite leer informacion.

El siguiente ejemplo lee el contenido del archivo `sample.txt`, chequea que si se producieron errores durante la operacion de lectura, e imprime el contenido del archivo.

```go
import (  
    "fmt"
    "io/ioutil"
)

func main() {  
    data, err := ioutil.ReadFile("sample.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

La funcion `ReadFile` recibe el path completo al archivo como argumento. Como el archivo esta en el mismo directorio que el programa `main.go`, solo se indica el nombre del archivo. 

La funcion `ReadFile` retorna dos valores, el contendio del archivo, y un objeto con la informacion detallada si se hubiera producido un error, si no hay error el objeto retornado es `nil`.

La informacion se lee como una arreglo de bytes. Para convertir ese arreglo de bytes a un string se utiliza la funcion `string`.

## Escritura

Se pueden escribir datos con la funcion `WriteFile` del paquete `ioutil`. La informacion se debe convertir a una arreglo de bytes antes de escribirla al archivo:

```go
data := []byte("This is some information!")
err := ioutil.WriteFile("write_data.txt", data, 0666)
if err != nil {
    fmt.Println("There has been an error:", err)
    return
}
```

En el ejemplo anterior los datos son un string que es convertido a un arreglo de bytes con la funcion `byte`. La funcion `WriteFile` requiere especificcar el modo de escritura para chequear si el usuario tiene permisos suficientes en el sistema de archivos.

El valor de retorno de la funcion `WriteFile` es un onjeto que confiente la informacion detalla si se hubiera producido un error, o `nil` si no lo hubiera.