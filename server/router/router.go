package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "server/middleware"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.Use(commonMiddleware)

    router.HandleFunc("/search", middleware.Search).Methods("POST")
    router.HandleFunc("/detail/{id1:[0-9]+}/{id2:[0-9]+}", middleware.DetailRoundtrip).Methods("GET")
    router.HandleFunc("/detail/{id1:[0-9]+}", middleware.DetailOneway).Methods("GET")
    router.HandleFunc("/store", middleware.Store).Methods("POST")
    router.HandleFunc("/pay", middleware.Pay).Methods("POST")
    router.HandleFunc("/book/{id:[0-9]+}", middleware.Book).Methods("GET")
    router.HandleFunc("/retrieve/{id:[0-9]+}", middleware.Retrieve).Methods("GET")
    router.HandleFunc("/retrieve", middleware.RetrieveAll).Methods("GET")
    router.HandleFunc("/checkin/{id:[0-9]+}", middleware.Checkin).Methods("GET")
    router.HandleFunc("/ticket/{id:[0-9]+}", middleware.Ticket).Methods("GET")
    router.HandleFunc("/airports/{code:[a-zA-Z]+}", middleware.Airports).Methods("GET")

    return router
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        w.Header().Add("Context-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
