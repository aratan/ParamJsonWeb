// curl -X POST localhost:8080 -d nombre=Juan -d edad=25 -d ciudad=Madrid

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
)

type Datos struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
    Ciudad string `json:"ciudad"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Obtenemos los valores de la URL
        nombre := r.FormValue("nombre")
        edadStr := r.FormValue("edad")
        ciudad := r.FormValue("ciudad")

        // Convertimos la edad a tipo int
        edad, err := strconv.Atoi(edadStr)
        if err != nil {
            fmt.Println("Error converting edad to int:", err)
            return
        }

        // Creamos un objeto Datos con los valores obtenidos
        datos := Datos{
            Nombre: nombre,
            Edad:   edad,
            Ciudad: ciudad,
        }

        // Guardamos el objeto Datos en un archivo JSON
        jsonData, err := json.Marshal(datos)
        if err != nil {
            fmt.Println(err)
            return
        }

        err = ioutil.WriteFile("datos.json", jsonData, 0644)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Respondemos con un mensaje
        fmt.Fprintf(w, "Los datos se guardaron correctamente")
    })

    http.ListenAndServe(":8080", nil)
}
