package main

import (
    "fmt"
    "net/http"
    "os" // Importar paquete para leer variables de entorno
    "log" // Importar el paquete log
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola, mundo")
}

func main() {
    http.HandleFunc("/", helloHandler)

    // Obtén el puerto desde la variable de entorno
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Puerto por defecto para local
    }

    fmt.Println("Servidor escuchando en puerto:", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("Error al iniciar el servidor:", err)
    }
    
    }

