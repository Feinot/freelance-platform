package user

type DefoltDAata struct {
	Login    string
	Password string
	Acces    string
	Id       int
	Order    bool
}
type OrderList struct {
	OrderID     int
	Users       string
	Description string
	Status      string
	OrderName   string
}
type LogClient struct {
	Log string
	Pwd string
}
