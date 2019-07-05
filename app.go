    package main
    
    
    import (
        "fmt"
        "net/http"
    )


    func main(){
        http.HandleFunc("/", manejador)
        http.ListenAndServe(":8080", nil)
    }

    func manejador(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w,"Hola, %s, ¡este es un servidor!", r.URL.Path)
      }