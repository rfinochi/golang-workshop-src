# Instalacion

1. Descargar la version para tu OS desde [https://golang.org/dl](https://golang.org/dl).
1. Abrir el instalador y seguir las [instrucciones](https://golang.org/doc/install).
    
    Go quedara instalado por default en:
    
        Unix: /usr/local/go
        Windows: C:\Go

1. Workspace

    La variable de entorno GOPATH especifica la ubicacion del entorno de trabajo.

    Por defecto es:
    
        Unix: $HOME/go
        Windows: %USERPROFILE%\go

    El workspace contiene:

        GOPATH
            |_ BIN: Contiene los archivos ejecutables
            |_ SRC: Contiene los fuentes organizados en paquetes
            |_ PKG: Contiene los paquetes en formato binario

    Es conveniente agregar [GOPATH](https://golang.org/doc/code.html#GOPATH)/bin al PATH.

1. Para chequear que la instalacion esta correctar tipear en la consola 
    
        $ go version

1.  Para chequear que todas las variables de entorno estan correctas tipear en la consola
    
        $ go env 

        GOROOT = [Directorio de instalacion]
        GOPATH = [Directorio del Workspace, tiene que ser distinto al de instalacion]

## Otras Herramientas

- [Visual Studio Code](https://code.visualstudio.com/download)
- [Visual Studio Code official Go extension](https://code.visualstudio.com/docs/languages/go)
- [Postman](https://www.getpostman.com)
- [MongoDB](https://www.mongodb.com)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/)
