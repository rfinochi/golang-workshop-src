# Packages

La buena organizacion del codigo es fundamental en cualquier proyecto. Organizar codigo en multiples directorios y multiples archivos permite mantener y probar el codigo mas facilmente. Los directorios son llamados `packages` en Golang. Puede haber uno o mas archivos en un `package`. They are all bound by using the same package name at the beginning of the file. Importing a package is done by providing the path to the directory instead of the file in the package.

La siguiente estructura es la indicada para utilizar `packages`:

```bash
Project Root directory
- bin
- pkg
- src
```

El directorio `bin` contiene los binarios del proyecto que se crean con el comando `go install`.

El directorio `pkg` contiene las version compiladas de los paquetes con extension `.a` (package archives).

El directorio `src` contiene el codigo fuente del proyecto organizado en multiples archivos y directorios.

---
To run the exercise, open the root folder for this exercise in VSCode and it will automatically infer the GOPATH for this project.

Running `GOPATH=$PWD go install app` will compile the code and if all goes well, will add an executable named `app` to the `bin` folder

Use of `GOPATH` is necessary as we are setting the value for the `GOPATH` environment variable just for running go install

Packages for the exercise should be created under the `src` folder without a `main` function as there should be only one `main` function for the whole application since only one entrypoint for an application is permitted.

---

## Referencias

* [Golang standard library packages](https://golang.org/pkg/)
* [Everything you need to know about packages](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
