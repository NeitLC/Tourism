package middleware

import (
    "log"
    "fmt"
    "math"
    "math/rand"
    "strings"
    "strconv"
    "net/http"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "server/models"
)

var database *sql.DB

func init () {
    connect()
}

func connect () {
//    db, _   := sql.Open("mysql", "tourism:tourism@tcp(82.209.195.34)/tourism")
    db, _   := sql.Open("mysql", "root:root@/tourism")
    database = db
    if err := database.Ping(); err != nil {
        fmt.Println("Open database fail")
        return 
    }
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func Search (w http.ResponseWriter, r *http.Request) {
    data := search(r)
    json.NewEncoder(w).Encode(data)
}

func search (r *http.Request) ([]models.Routes) {
    var routes = []models.Routes{}

    var data models.Search
    decoder := json.NewDecoder(r.Body)
    err     := decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }

    var route  = []models.DatabaseRoute{}
    rows, err := database.Query("SELECT * FROM `routes` WHERE `from` = '" + data.From + "' AND `to` = '" + data.To + "' AND `date` = '" + data.Date + "' ORDER BY `price` LIMIT 3")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rout := models.DatabaseRoute{}
        err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
        if err != nil {
            log.Fatal(err)
            continue
        }
        route = append(route, rout)
    }

    for _, elem := range route {
        Dep_arp := strings.Split(elem.Dep_arp, ",")
        Arr_arp := strings.Split(elem.Arr_arp, ",")
        Dep_dat := strings.Split(elem.Dep_dat, ",")
        Arr_dat := strings.Split(elem.Arr_dat, ",")
        Airline := strings.Split(elem.Airline, ",")
        Number  := strings.Split(elem.Number, ",")
        Eqp_typ := strings.Split(elem.Eqp_typ, ",")
        Ela_tim := strings.Split(elem.Ela_tim, ",")
        Gro_tim := strings.Split(elem.Gro_tim, ",")
        Acu_tim := strings.Split(elem.Acu_tim, ",")
        Dep_ter := strings.Split(elem.Dep_ter, ",")
        Arr_ter := strings.Split(elem.Arr_ter, ",")

        l       := len(Acu_tim)
        hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
        min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

        var segments []models.Segment
        for j := 0; j < len(Dep_arp); j++ {
            dep := strings.Split(Dep_dat[j], "T")
            arr := strings.Split(Arr_dat[j], "T")

            var segment models.Segment
            segment.Id      = j + 1
            segment.Dep_arp = Dep_arp[j]
            segment.Arr_arp = Arr_arp[j]
            segment.Dep_dat = dep[0]
            segment.Dep_tim = dep[1][0:5]
            segment.Arr_dat = arr[0]
            segment.Arr_tim = arr[1][0:5]
            segment.Airline = Airline[j]
            segment.Number  = Number[j]
            segment.Cabin   = "Economy"
            segment.Eqp_typ = Eqp_typ[j]
            segment.Ela_tim = Ela_tim[j]
            segment.Gro_tim = Gro_tim[j]
            segment.Acu_tim = Acu_tim[j]
            segment.Dep_ter = Dep_ter[j]
            segment.Arr_ter = Arr_ter[j]

            segments = append(segments, segment)
        }

        var out models.ViewRoute
        out.Id         = elem.Id
        out.From       = elem.From
        out.To         = elem.To
        out.Date       = elem.Date
        out.Price      = elem.Price
        out.Currency   = elem.Currency
        out.Passengers = data.Passengers
        out.Duration   = hour * 60 + min
        out.Segments   = segments

        var outbound models.Routes
        outbound.Price      = elem.Price
        outbound.Currency   = elem.Currency
        outbound.Passengers = data.Passengers
        outbound.Id         = strconv.Itoa(elem.Id)
        outbound.Outbound   = append(outbound.Outbound, out)
        routes              = append(routes, outbound)
    }

    if data.Date2 != "" {
        route      = []models.DatabaseRoute{}
        rows, err := database.Query("SELECT * FROM `routes` WHERE `from` = '" + data.To + "' AND `to` = '" + data.From + "' AND `date` = '" + data.Date2 + "' ORDER BY `price` LIMIT 3")
        if err != nil {
            log.Fatal(err)
        }
        for rows.Next() {
            rout := models.DatabaseRoute{}
            err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
            if err != nil {
                log.Fatal(err)
                continue
            }
            route = append(route, rout)
        }

        temp  := routes
        routes = []models.Routes{}
        for i := 0; i < len(temp); i++ {
            var outbound []models.ViewRoute
            outbound = temp[i].Outbound
            for _, elem := range route {
                Dep_arp := strings.Split(elem.Dep_arp, ",")
                Arr_arp := strings.Split(elem.Arr_arp, ",")
                Dep_dat := strings.Split(elem.Dep_dat, ",")
                Arr_dat := strings.Split(elem.Arr_dat, ",")
                Airline := strings.Split(elem.Airline, ",")
                Number  := strings.Split(elem.Number, ",")
                Eqp_typ := strings.Split(elem.Eqp_typ, ",")
                Ela_tim := strings.Split(elem.Ela_tim, ",")
                Gro_tim := strings.Split(elem.Gro_tim, ",")
                Acu_tim := strings.Split(elem.Acu_tim, ",")
                Dep_ter := strings.Split(elem.Dep_ter, ",")
                Arr_ter := strings.Split(elem.Arr_ter, ",")

                l       := len(Acu_tim)
                hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
                min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

                var segments []models.Segment
                for j := 0; j < len(Dep_arp); j++ {
                    dep := strings.Split(Dep_dat[j], "T")
                    arr := strings.Split(Arr_dat[j], "T")

                    var segment models.Segment
                    segment.Id      = j + 1
                    segment.Dep_arp = Dep_arp[j]
                    segment.Arr_arp = Arr_arp[j]
                    segment.Dep_dat = dep[0]
                    segment.Dep_tim = dep[1][0:5]
                    segment.Arr_dat = arr[0]
                    segment.Arr_tim = arr[1][0:5]
                    segment.Airline = Airline[j]
                    segment.Number  = Number[j]
                    segment.Cabin   = "Economy"
                    segment.Eqp_typ = Eqp_typ[j]
                    segment.Ela_tim = Ela_tim[j]
                    segment.Gro_tim = Gro_tim[j]
                    segment.Acu_tim = Acu_tim[j]
                    segment.Dep_ter = Dep_ter[j]
                    segment.Arr_ter = Arr_ter[j]
    
                    segments = append(segments, segment)
                }

                var in models.ViewRoute
                in.Id         = elem.Id
                in.From       = elem.From
                in.To         = elem.To
                in.Date       = elem.Date
                in.Price      = elem.Price
                in.Currency   = elem.Currency
                in.Passengers = data.Passengers
                in.Duration   = hour * 60 + min
                in.Segments   = segments

                var inbound []models.ViewRoute
                inbound = append(inbound, in)

                var r models.Routes
                r.Outbound   = outbound
                r.Inbound    = inbound
                r.Price      = math.Round((outbound[0].Price + elem.Price) * 100) / 100
                r.Currency   = elem.Currency
                r.Passengers = data.Passengers
                r.Id         = strconv.Itoa(outbound[0].Id) + "/" + strconv.Itoa(elem.Id)
                routes       = append(routes, r)
            }
        }
    }

    return routes
}

func DetailOneway(w http.ResponseWriter, r *http.Request) {
    data := detail(r)
    json.NewEncoder(w).Encode(data)
}

func DetailRoundtrip(w http.ResponseWriter, r *http.Request) {
    data := detail(r)
    json.NewEncoder(w).Encode(data)
}

func detail(r *http.Request) ([]models.Routes) {
    var routes = []models.Routes{}

    params    := mux.Vars(r)
    var route  = []models.DatabaseRoute{}
    rows, err := database.Query("SELECT * FROM `routes` WHERE `id` = '" + params["id1"] + "'")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rout := models.DatabaseRoute{}
        err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
        if err != nil {
            log.Fatal(err)
            continue
        }
        route = append(route, rout)
    }

    for _, elem := range route {
        Dep_arp := strings.Split(elem.Dep_arp, ",")
        Arr_arp := strings.Split(elem.Arr_arp, ",")
        Dep_dat := strings.Split(elem.Dep_dat, ",")
        Arr_dat := strings.Split(elem.Arr_dat, ",")
        Airline := strings.Split(elem.Airline, ",")
        Number  := strings.Split(elem.Number, ",")
        Eqp_typ := strings.Split(elem.Eqp_typ, ",")
        Ela_tim := strings.Split(elem.Ela_tim, ",")
        Gro_tim := strings.Split(elem.Gro_tim, ",")
        Acu_tim := strings.Split(elem.Acu_tim, ",")
        Dep_ter := strings.Split(elem.Dep_ter, ",")
        Arr_ter := strings.Split(elem.Arr_ter, ",")

        l       := len(Acu_tim)
        hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
        min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

        var segments []models.Segment
        for j := 0; j < len(Dep_arp); j++ {
            dep := strings.Split(Dep_dat[j], "T")
            arr := strings.Split(Arr_dat[j], "T")

            var segment models.Segment
            segment.Id      = j + 1
            segment.Dep_arp = Dep_arp[j]
            segment.Arr_arp = Arr_arp[j]
            segment.Dep_dat = dep[0]
            segment.Dep_tim = dep[1][0:5]
            segment.Arr_dat = arr[0]
            segment.Arr_tim = arr[1][0:5]
            segment.Airline = Airline[j]
            segment.Number  = Number[j]
            segment.Cabin   = "Economy"
            segment.Eqp_typ = Eqp_typ[j]
            segment.Ela_tim = Ela_tim[j]
            segment.Gro_tim = Gro_tim[j]
            segment.Acu_tim = Acu_tim[j]
            segment.Dep_ter = Dep_ter[j]
            segment.Arr_ter = Arr_ter[j]

            segments = append(segments, segment)
        }

        var out models.ViewRoute
        out.Id         = elem.Id
        out.From       = elem.From
        out.To         = elem.To
        out.Date       = elem.Date
        out.Price      = elem.Price
        out.Currency   = elem.Currency
        out.Passengers = 1 // data.Passengers
        out.Duration   = hour * 60 + min
        out.Segments   = segments

        var outbound models.Routes
        outbound.Price      = elem.Price
        outbound.Currency   = elem.Currency
        outbound.Passengers = 1 // data.Passengers
        outbound.Id         = strconv.Itoa(elem.Id)
        outbound.Outbound   = append(outbound.Outbound, out)
        routes              = append(routes, outbound)
    }

    if params["id2"] != "" {
        route      = []models.DatabaseRoute{}
        rows, err := database.Query("SELECT * FROM `routes` WHERE `id` = '" + params["id2"] + "'")
        if err != nil {
            log.Fatal(err)
        }
        for rows.Next() {
            rout := models.DatabaseRoute{}
            err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
            if err != nil {
                log.Fatal(err)
                continue
            }
            route = append(route, rout)
        }

        temp  := routes
        routes = []models.Routes{}
        for i := 0; i < len(temp); i++ {
            var outbound []models.ViewRoute
            outbound = temp[i].Outbound
            for _, elem := range route {
                Dep_arp := strings.Split(elem.Dep_arp, ",")
                Arr_arp := strings.Split(elem.Arr_arp, ",")
                Dep_dat := strings.Split(elem.Dep_dat, ",")
                Arr_dat := strings.Split(elem.Arr_dat, ",")
                Airline := strings.Split(elem.Airline, ",")
                Number  := strings.Split(elem.Number, ",")
                Eqp_typ := strings.Split(elem.Eqp_typ, ",")
                Ela_tim := strings.Split(elem.Ela_tim, ",")
                Gro_tim := strings.Split(elem.Gro_tim, ",")
                Acu_tim := strings.Split(elem.Acu_tim, ",")
                Dep_ter := strings.Split(elem.Dep_ter, ",")
                Arr_ter := strings.Split(elem.Arr_ter, ",")

                l       := len(Acu_tim)
                hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
                min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

                var segments []models.Segment
                for j := 0; j < len(Dep_arp); j++ {
                    dep := strings.Split(Dep_dat[j], "T")
                    arr := strings.Split(Arr_dat[j], "T")

                    var segment models.Segment
                    segment.Id      = j + 1
                    segment.Dep_arp = Dep_arp[j]
                    segment.Arr_arp = Arr_arp[j]
                    segment.Dep_dat = dep[0]
                    segment.Dep_tim = dep[1][0:5]
                    segment.Arr_dat = arr[0]
                    segment.Arr_tim = arr[1][0:5]
                    segment.Airline = Airline[j]
                    segment.Number  = Number[j]
                    segment.Cabin   = "Economy"
                    segment.Eqp_typ = Eqp_typ[j]
                    segment.Ela_tim = Ela_tim[j]
                    segment.Gro_tim = Gro_tim[j]
                    segment.Acu_tim = Acu_tim[j]
                    segment.Dep_ter = Dep_ter[j]
                    segment.Arr_ter = Arr_ter[j]
    
                    segments = append(segments, segment)
                }

                var in models.ViewRoute
                in.Id         = elem.Id
                in.From       = elem.From
                in.To         = elem.To
                in.Date       = elem.Date
                in.Price      = elem.Price
                in.Currency   = elem.Currency
                in.Passengers = 1 // data.Passengers
                in.Duration   = hour * 60 + min
                in.Segments   = segments

                var inbound []models.ViewRoute
                inbound = append(inbound, in)

                var r models.Routes
                r.Outbound   = outbound
                r.Inbound    = inbound
                r.Price      = math.Round((outbound[0].Price + elem.Price) * 100) / 100
                r.Currency   = elem.Currency
                r.Passengers = 1 // data.Passengers
                r.Id         = strconv.Itoa(outbound[0].Id) + "/" + strconv.Itoa(elem.Id)
                routes       = append(routes, r)
            }
        }
    }

    return routes
}

func Store (w http.ResponseWriter, r *http.Request) {
    data := store(r)
    json.NewEncoder(w).Encode(data)
}

func store (r *http.Request) (string) {
    var data models.Reservation
    decoder := json.NewDecoder(r.Body)
    err     := decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }

    var sql1 = "INSERT INTO reservations (`id_routes`, `firstname`, `lastname`, `birthdate`, `passport`, `email`, `phone`) "
    var sql2 = "VALUES ('" + data.Id_routes + "','" + data.Firstname + "','" + data.Lastname + "','" + data.Birthdate + "','" + data.Passport + "','" + data.Email + "','" + data.Phone + "')"
    var sql  = sql1 + sql2
    _, err   = database.Exec(sql)
    if err != nil {
        log.Fatal(err)
    }

    var id int
    rows, err := database.Query("SELECT `id` FROM `reservations` ORDER BY `id` DESC LIMIT 1")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        err := rows.Scan(&id)
        if err != nil {
            log.Fatal(err)
            continue
        }
    }

    return strconv.Itoa(id)
}

func Pay (w http.ResponseWriter, r *http.Request)  {
    data := pay(r)
    json.NewEncoder(w).Encode(data)
}

func pay (r *http.Request) (string) {
    var data models.Pays

    decoder := json.NewDecoder(r.Body)
    err     := decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }

    payment, _ := json.Marshal(data.Payment)

    _, err = database.Exec("UPDATE `reservations` SET `payment` = '" + string(payment) + "' WHERE `id` = '" + data.Id + "'")
    if err != nil {
        log.Fatal(err)
    }

    return data.Id
}

func Book (w http.ResponseWriter, r *http.Request) {
    data := book(r)
    json.NewEncoder(w).Encode(data)
}

func book (r *http.Request) (string) {
    params := mux.Vars(r)

    var locator = RandString(6) // random [A-Z0-9]
    _, err := database.Exec("UPDATE `reservations` SET `locator` = '" + locator + "' WHERE `id` = '" + params["id"] + "'")
    if err != nil {
        log.Fatal(err)
    }

    return params["id"]
}

func Ticket (w http.ResponseWriter, r *http.Request) {
    data := ticket(r)
    json.NewEncoder(w).Encode(data)
}

func ticket (r *http.Request) (string) {
    params := mux.Vars(r)

    var ticket = RandString(13) // random [A-Z0-9]
    _, err := database.Exec("UPDATE `reservations` SET `ticket` = '" + ticket + "' WHERE `id` = '" + params["id"] + "'")
    if err != nil {
        log.Fatal(err)
    }

    return params["id"]
}

func Checkin (w http.ResponseWriter, r *http.Request) {
    data := checkin(r)
    json.NewEncoder(w).Encode(data)
}

func checkin (r *http.Request) (string) {
    params := mux.Vars(r)

    var checkin = RandString(8) // random [A-Z0-9]
    _, err := database.Exec("UPDATE `reservations` SET `checkin` = '" + checkin + "' WHERE `id` = '" + params["id"] + "'")
    if err != nil {
        log.Fatal(err)
    }

    return params["id"]
}

func Retrieve (w http.ResponseWriter, r *http.Request) {
    data := retrieve(r)
    json.NewEncoder(w).Encode(data)
}

func retrieve (r *http.Request) (models.ViewResevation) {
    params := mux.Vars(r)

    var res    = models.Reservation{}
    rows, err := database.Query("SELECT `id`, `locator`, `id_routes`, `firstname`, `lastname`, `birthdate`, `passport`, `email`, `phone` FROM `reservations` WHERE `id` = '" + params["id"] + "'")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        err := rows.Scan(&res.Id, &res.Locator, &res.Id_routes, &res.Firstname, &res.Lastname, &res.Birthdate, &res.Passport, &res.Email, &res.Phone)
        if err != nil {
            log.Fatal(err)
            continue
        }
    }

    ids       := strings.Split(res.Id_routes, "/")
    var routes = []models.Routes{}
    var route  = []models.DatabaseRoute{}
    rows, err  = database.Query("SELECT * FROM `routes` WHERE `id` = '" + ids[0] + "'")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rout := models.DatabaseRoute{}
        err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
        if err != nil {
            log.Fatal(err)
            continue
        }
        route = append(route, rout)
    }

    for _, elem := range route {
        Dep_arp := strings.Split(elem.Dep_arp, ",")
        Arr_arp := strings.Split(elem.Arr_arp, ",")
        Dep_dat := strings.Split(elem.Dep_dat, ",")
        Arr_dat := strings.Split(elem.Arr_dat, ",")
        Airline := strings.Split(elem.Airline, ",")
        Number  := strings.Split(elem.Number, ",")
        Eqp_typ := strings.Split(elem.Eqp_typ, ",")
        Ela_tim := strings.Split(elem.Ela_tim, ",")
        Gro_tim := strings.Split(elem.Gro_tim, ",")
        Acu_tim := strings.Split(elem.Acu_tim, ",")
        Dep_ter := strings.Split(elem.Dep_ter, ",")
        Arr_ter := strings.Split(elem.Arr_ter, ",")

        l       := len(Acu_tim)
        hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
        min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

        var segments []models.Segment
        for j := 0; j < len(Dep_arp); j++ {
            dep := strings.Split(Dep_dat[j], "T")
            arr := strings.Split(Arr_dat[j], "T")

            var segment models.Segment
            segment.Id      = j + 1
            segment.Dep_arp = Dep_arp[j]
            segment.Arr_arp = Arr_arp[j]
            segment.Dep_dat = dep[0]
            segment.Dep_tim = dep[1][0:5]
            segment.Arr_dat = arr[0]
            segment.Arr_tim = arr[1][0:5]
            segment.Airline = Airline[j]
            segment.Number  = Number[j]
            segment.Cabin   = "Economy"
            segment.Eqp_typ = Eqp_typ[j]
            segment.Ela_tim = Ela_tim[j]
            segment.Gro_tim = Gro_tim[j]
            segment.Acu_tim = Acu_tim[j]
            segment.Dep_ter = Dep_ter[j]
            segment.Arr_ter = Arr_ter[j]

            segments = append(segments, segment)
        }

        var out models.ViewRoute
        out.Id         = elem.Id
        out.From       = elem.From
        out.To         = elem.To
        out.Date       = elem.Date
        out.Price      = elem.Price
        out.Currency   = elem.Currency
        out.Passengers = 1 // data.Passengers
        out.Duration   = hour * 60 + min
        out.Segments   = segments

        var outbound models.Routes
        outbound.Price      = elem.Price
        outbound.Currency   = elem.Currency
        outbound.Passengers = 1 // data.Passengers
        outbound.Id         = strconv.Itoa(elem.Id)
        outbound.Outbound   = append(outbound.Outbound, out)
        routes              = append(routes, outbound)
    }

    if ids[1] != "" {
        route      = []models.DatabaseRoute{}
        rows, err := database.Query("SELECT * FROM `routes` WHERE `id` = '" + ids[1] + "'")
        if err != nil {
            log.Fatal(err)
        }
        for rows.Next() {
            rout := models.DatabaseRoute{}
            err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
            if err != nil {
                log.Fatal(err)
                continue
            }
            route = append(route, rout)
        }

        temp  := routes
        routes = []models.Routes{}
        for i := 0; i < len(temp); i++ {
            var outbound []models.ViewRoute
            outbound = temp[i].Outbound
            for _, elem := range route {
                Dep_arp := strings.Split(elem.Dep_arp, ",")
                Arr_arp := strings.Split(elem.Arr_arp, ",")
                Dep_dat := strings.Split(elem.Dep_dat, ",")
                Arr_dat := strings.Split(elem.Arr_dat, ",")
                Airline := strings.Split(elem.Airline, ",")
                Number  := strings.Split(elem.Number, ",")
                Eqp_typ := strings.Split(elem.Eqp_typ, ",")
                Ela_tim := strings.Split(elem.Ela_tim, ",")
                Gro_tim := strings.Split(elem.Gro_tim, ",")
                Acu_tim := strings.Split(elem.Acu_tim, ",")
                Dep_ter := strings.Split(elem.Dep_ter, ",")
                Arr_ter := strings.Split(elem.Arr_ter, ",")

                l       := len(Acu_tim)
                hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
                min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

                var segments []models.Segment
                for j := 0; j < len(Dep_arp); j++ {
                    dep := strings.Split(Dep_dat[j], "T")
                    arr := strings.Split(Arr_dat[j], "T")

                    var segment models.Segment
                    segment.Id      = j + 1
                    segment.Dep_arp = Dep_arp[j]
                    segment.Arr_arp = Arr_arp[j]
                    segment.Dep_dat = dep[0]
                    segment.Dep_tim = dep[1][0:5]
                    segment.Arr_dat = arr[0]
                    segment.Arr_tim = arr[1][0:5]
                    segment.Airline = Airline[j]
                    segment.Number  = Number[j]
                    segment.Cabin   = "Economy"
                    segment.Eqp_typ = Eqp_typ[j]
                    segment.Ela_tim = Ela_tim[j]
                    segment.Gro_tim = Gro_tim[j]
                    segment.Acu_tim = Acu_tim[j]
                    segment.Dep_ter = Dep_ter[j]
                    segment.Arr_ter = Arr_ter[j]
    
                    segments = append(segments, segment)
                }

                var in models.ViewRoute
                in.Id         = elem.Id
                in.From       = elem.From
                in.To         = elem.To
                in.Date       = elem.Date
                in.Price      = elem.Price
                in.Currency   = elem.Currency
                in.Passengers = 1 // data.Passengers
                in.Duration   = hour * 60 + min
                in.Segments   = segments

                var inbound []models.ViewRoute
                inbound = append(inbound, in)

                var r models.Routes
                r.Outbound   = outbound
                r.Inbound    = inbound
                r.Price      = math.Round((outbound[0].Price + elem.Price) * 100) / 100
                r.Currency   = elem.Currency
                r.Passengers = 1 // data.Passengers
                r.Id         = strconv.Itoa(outbound[0].Id) + "/" + strconv.Itoa(elem.Id)
                routes       = append(routes, r)
            }
        }
    }

    var reserv       = models.Reservation{} 
    reserv.Id        = res.Id
    reserv.Locator   = res.Locator
    reserv.Id_routes = res.Id_routes
    reserv.Email     = res.Email
    reserv.Phone     = res.Phone
    reserv.Firstname = res.Firstname
    reserv.Lastname  = res.Lastname
    reserv.Birthdate = res.Birthdate
    reserv.Passport  = res.Passport

    var rqw         = models.ViewResevation{}
    rqw.Reservation = reserv
    rqw.Routes      = routes

    return rqw
}

func RetrieveAll (w http.ResponseWriter, r *http.Request) {
    data := retrieveAll(r)
    json.NewEncoder(w).Encode(data)
}

func retrieveAll (r *http.Request) ([]models.ViewResevation) {
    var rqws = []models.ViewResevation{}

    var res    = []models.Reservation{}
    rows, err := database.Query("SELECT `id`, `locator`, `id_routes`, `firstname`, `lastname`, `birthdate`, `passport`, `email`, `phone`, `ticket`, `checkin` FROM `reservations` WHERE `locator` IS NOT NULL")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        re    := models.Reservation{}
        err := rows.Scan(&re.Id, &re.Locator, &re.Id_routes, &re.Firstname, &re.Lastname, &re.Birthdate, &re.Passport, &re.Email, &re.Phone, &re.Ticket, &re.Checkin)
        if err != nil {
            log.Fatal(err)
            continue
        }
        res = append(res, re)
    }

    for _, eee := range res {
        ids       := strings.Split(eee.Id_routes, "/")
        var routes = []models.Routes{}
        var route  = []models.DatabaseRoute{}
        rows, err  = database.Query("SELECT * FROM `routes` WHERE `id` = '" + ids[0] + "'")
        if err != nil {
            log.Fatal(err)
        }
        for rows.Next() {
            rout := models.DatabaseRoute{}
            err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
            if err != nil {
                log.Fatal(err)
                continue
            }
            route = append(route, rout)
        }

        for _, elem := range route {
            Dep_arp := strings.Split(elem.Dep_arp, ",")
            Arr_arp := strings.Split(elem.Arr_arp, ",")
            Dep_dat := strings.Split(elem.Dep_dat, ",")
            Arr_dat := strings.Split(elem.Arr_dat, ",")
            Airline := strings.Split(elem.Airline, ",")
            Number  := strings.Split(elem.Number, ",")
            Eqp_typ := strings.Split(elem.Eqp_typ, ",")
            Ela_tim := strings.Split(elem.Ela_tim, ",")
            Gro_tim := strings.Split(elem.Gro_tim, ",")
            Acu_tim := strings.Split(elem.Acu_tim, ",")
            Dep_ter := strings.Split(elem.Dep_ter, ",")
            Arr_ter := strings.Split(elem.Arr_ter, ",")

            l       := len(Acu_tim)
            hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
            min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

            var segments []models.Segment
            for j := 0; j < len(Dep_arp); j++ {
                dep := strings.Split(Dep_dat[j], "T")
                arr := strings.Split(Arr_dat[j], "T")

                var segment models.Segment
                segment.Id      = j + 1
                segment.Dep_arp = Dep_arp[j]
                segment.Arr_arp = Arr_arp[j]
                segment.Dep_dat = dep[0]
                segment.Dep_tim = dep[1][0:5]
                segment.Arr_dat = arr[0]
                segment.Arr_tim = arr[1][0:5]
                segment.Airline = Airline[j]
                segment.Number  = Number[j]
                segment.Cabin   = "Economy"
                segment.Eqp_typ = Eqp_typ[j]
                segment.Ela_tim = Ela_tim[j]
                segment.Gro_tim = Gro_tim[j]
                segment.Acu_tim = Acu_tim[j]
                segment.Dep_ter = Dep_ter[j]
                segment.Arr_ter = Arr_ter[j]

                segments = append(segments, segment)
            }

            var out models.ViewRoute
            out.Id         = elem.Id
            out.From       = elem.From
            out.To         = elem.To
            out.Date       = elem.Date
            out.Price      = elem.Price
            out.Currency   = elem.Currency
            out.Passengers = 1 // data.Passengers
            out.Duration   = hour * 60 + min
            out.Segments   = segments

            var outbound models.Routes
            outbound.Price      = elem.Price
            outbound.Currency   = elem.Currency
            outbound.Passengers = 1 // data.Passengers
            outbound.Id         = strconv.Itoa(elem.Id)
            outbound.Outbound   = append(outbound.Outbound, out)
            routes              = append(routes, outbound)
        }

        if ids[1] != "" {
            route      = []models.DatabaseRoute{}
            rows, err := database.Query("SELECT * FROM `routes` WHERE `id` = '" + ids[1] + "'")
            if err != nil {
                log.Fatal(err)
            }
            for rows.Next() {
                rout := models.DatabaseRoute{}
                err  := rows.Scan(&rout.Id, &rout.From, &rout.To, &rout.Date, &rout.Price, &rout.Currency, &rout.Dep_arp, &rout.Arr_arp, &rout.Dep_dat, &rout.Arr_dat, &rout.Airline, &rout.Number, &rout.Eqp_typ, &rout.Ela_tim, &rout.Gro_tim, &rout.Acu_tim, &rout.Dep_ter, &rout.Arr_ter)
                if err != nil {
                    log.Fatal(err)
                    continue
                }
                route = append(route, rout)
            }

            temp  := routes
            routes = []models.Routes{}
            for i := 0; i < len(temp); i++ {
                var outbound []models.ViewRoute
                outbound = temp[i].Outbound
                for _, elem := range route {
                    Dep_arp := strings.Split(elem.Dep_arp, ",")
                    Arr_arp := strings.Split(elem.Arr_arp, ",")
                    Dep_dat := strings.Split(elem.Dep_dat, ",")
                    Arr_dat := strings.Split(elem.Arr_dat, ",")
                    Airline := strings.Split(elem.Airline, ",")
                    Number  := strings.Split(elem.Number, ",")
                    Eqp_typ := strings.Split(elem.Eqp_typ, ",")
                    Ela_tim := strings.Split(elem.Ela_tim, ",")
                    Gro_tim := strings.Split(elem.Gro_tim, ",")
                    Acu_tim := strings.Split(elem.Acu_tim, ",")
                    Dep_ter := strings.Split(elem.Dep_ter, ",")
                    Arr_ter := strings.Split(elem.Arr_ter, ",")

                    l       := len(Acu_tim)
                    hour, _ := strconv.Atoi(Acu_tim[l - 1][0:2])
                    min, _  := strconv.Atoi(Acu_tim[l - 1][2:4])

                    var segments []models.Segment
                    for j := 0; j < len(Dep_arp); j++ {
                        dep := strings.Split(Dep_dat[j], "T")
                        arr := strings.Split(Arr_dat[j], "T")

                        var segment models.Segment
                        segment.Id      = j + 1
                        segment.Dep_arp = Dep_arp[j]
                        segment.Arr_arp = Arr_arp[j]
                        segment.Dep_dat = dep[0]
                        segment.Dep_tim = dep[1][0:5]
                        segment.Arr_dat = arr[0]
                        segment.Arr_tim = arr[1][0:5]
                        segment.Airline = Airline[j]
                        segment.Number  = Number[j]
                        segment.Cabin   = "Economy"
                        segment.Eqp_typ = Eqp_typ[j]
                        segment.Ela_tim = Ela_tim[j]
                        segment.Gro_tim = Gro_tim[j]
                        segment.Acu_tim = Acu_tim[j]
                        segment.Dep_ter = Dep_ter[j]
                        segment.Arr_ter = Arr_ter[j]

                        segments = append(segments, segment)
                    }

                    var in models.ViewRoute
                    in.Id         = elem.Id
                    in.From       = elem.From
                    in.To         = elem.To
                    in.Date       = elem.Date
                    in.Price      = elem.Price
                    in.Currency   = elem.Currency
                    in.Passengers = 1 // data.Passengers
                    in.Duration   = hour * 60 + min
                    in.Segments   = segments

                    var inbound []models.ViewRoute
                    inbound = append(inbound, in)

                    var r models.Routes
                    r.Outbound   = outbound
                    r.Inbound    = inbound
                    r.Price      = math.Round((outbound[0].Price + elem.Price) * 100) / 100
                    r.Currency   = elem.Currency
                    r.Passengers = 1 // data.Passengers
                    r.Id         = strconv.Itoa(outbound[0].Id) + "/" + strconv.Itoa(elem.Id)
                    routes       = append(routes, r)
                }
            }
        }

        var reserv       = models.Reservation{} 
        reserv.Id        = eee.Id
        reserv.Locator   = eee.Locator
        reserv.Id_routes = eee.Id_routes
        reserv.Email     = eee.Email
        reserv.Phone     = eee.Phone
        reserv.Firstname = eee.Firstname
        reserv.Lastname  = eee.Lastname
        reserv.Birthdate = eee.Birthdate
        reserv.Passport  = eee.Passport
        reserv.Ticket    = eee.Ticket
        reserv.Checkin   = eee.Checkin
        
        var rqw         = models.ViewResevation{}
        rqw.Reservation = reserv
        rqw.Routes      = routes
        rqws            = append(rqws, rqw)
    }

    return rqws
}

func Airports (w http.ResponseWriter, r *http.Request) {
    data := airports(r)
    json.NewEncoder(w).Encode(data)
}

func airports (r *http.Request) ([]models.Airports) {
    params := mux.Vars(r)

    var airports = []models.Airports{}
    rows, err := database.Query("SELECT `name` AS `city`,`country`,`code` FROM `airports` WHERE `code` LIKE '%" + params["code"] + "%'")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        temp := models.Airports{}
        err  := rows.Scan(&temp.City, &temp.Country, &temp.Code)
        if err != nil {
            log.Fatal(err)
            continue
        }
        airports = append(airports, temp)
    }

    return airports
}
