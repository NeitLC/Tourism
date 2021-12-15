package models

type Search struct {
    From       string  `json:"from"`
    To         string  `json:"to"`
    Date       string  `json:"date"`
    Date2      string  `json:"date2"`
    Passengers int     `json:"passengers"`
}

type DatabaseRoute struct {
    Id       int     `json:"id"`
    From     string  `json:"from"`
    To       string  `json:"to"`
    Date     string  `json:"date"`
    Price    float64 `json:"price"`
    Currency string  `json:"currency"`
    Dep_arp  string  `json:"dep_arp"`
    Arr_arp  string  `json:"arr_arp"`
    Dep_dat  string  `json:"dep_dat"`
    Arr_dat  string  `json:"arr_dat"`
    Airline  string  `json:"airline"`
    Number   string  `json:"number"`
    Eqp_typ  string  `json:"eqp_typ"`
    Ela_tim  string  `json:"ela_tim"`
    Gro_tim  string  `json:"gro_tim"`
    Acu_tim  string  `json:"acu_tim"`
    Dep_ter  string  `json:"dep_ter"`
    Arr_ter  string  `json:"arr_ter"`
}

type Segment struct {
    Id       int     `json:"id"`
    Dep_arp  string  `json:"dep_arp"`
    Arr_arp  string  `json:"arr_arp"`
    Dep_dat  string  `json:"dep_dat"`
    Dep_tim  string  `json:"dep_tim"`
    Arr_dat  string  `json:"arr_dat"`
    Arr_tim  string  `json:"arr_tim"`
    Airline  string  `json:"airline"`
    Number   string  `json:"number"`
    Cabin    string  `json:"cabin"`
    Eqp_typ  string  `json:"eqp_typ"`
    Ela_tim  string  `json:"ela_tim"`
    Gro_tim  string  `json:"gro_tim"`
    Acu_tim  string  `json:"acu_tim"`
    Dep_ter  string  `json:"dep_ter"`
    Arr_ter  string  `json:"arr_ter"`
}

type ViewRoute struct {
    Id         int       `json:"id"`
    From       string    `json:"from"`
    To         string    `json:"to"`
    Date       string    `json:"date"`
    Price      float64   `json:"price"`
    Currency   string    `json:"currency"`
    Passengers int       `json:"passengers"`
    Duration   int       `json:"duration"`
    Segments   []Segment `json:"segments"`
}

type Routes struct {
    Id         string      `json:"id"`
    Price      float64     `json:"price"`
    Currency   string      `json:"currency"`
    Passengers int         `json:"passengers"`
    Outbound   []ViewRoute `json:"outbound"`
    Inbound    []ViewRoute `json:"inbound"`
}

type Reservation struct {
    Id         int    `json:"id"`
    Locator    string `json:"locator"`
    Id_routes  string `json:"id_routes"`
    Firstname  string `json:"firstname"`
    Lastname   string `json:"lastname"`
    Birthdate  string `json:"birthdate"`
    Passport   string `json:"passport"`
    Email      string `json:"email"`
    Phone      string `json:"phone"`
    Ticket     string `json:"ticket"`
    Checkin    string `json:"checkin"`
}

type ViewResevation struct {
    Reservation Reservation `json:"reservation"`
    Routes []Routes         `json:"routes"`
}

type Card struct {
    Number string `json:"number"`
    Holder string `json:"holder"`
    Expiry string `json:"expiry"`
    Cvc    string `json:"cvc"`
}

type Pays struct {
    Id      string `json:"id"`
    Payment Card   `json:"payment"`
}

type Airports struct {
    City    string `json:"city"`
    Country string `json:"country"`
    Code    string `json:"code"`
}
