package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"sync"
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "123"
	dbname   = "postgres"
)

var (
	db       *sql.DB
	complite = "complite"
	during   = "during"
	defolt   = "not done"
)

func Init() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	dbOpen, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	db = dbOpen
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	// body  { "login": "deadthr","password": "124", "acces": "ss","order":true}
	users := DeserializeRequest(w, r)
	exist := checkloginExist(users.Login)
	if exist {
		http.Error(w, "login already exists", 400)
		return
	}
	_, err := db.Exec("insert into account(login,password,acces)values($1,$2,$3)", users.Login, users.Password, users.Acces)
	if err != nil {
		fmt.Errorf("can't add %d", err)

	}
	fmt.Fprintf(w, "complite add user", users.Login)

}
func LogInSuchUser(w http.ResponseWriter, r *http.Request) {

	users := DeserializeRequest(w, r)
	exist := checkUsersExist(users.Login, users.Password)
	if exist {
		http.Error(w, "account not found", 400)
	} else {
		//

	}

}
func checkloginExist(login string) bool {
	var result bool

	if err := db.QueryRow("select exists(select login from account where login = $1)", login).Scan(&result); err != nil {
		fmt.Errorf("error%d", err)
	}
	res := result

	return res

}
func checkUsersExist(login string, password string) bool {
	var result bool

	if err := db.QueryRow("select exists(select login from account where (login = $1 && password =$2))", login, password).Scan(&result); err != nil {
		fmt.Errorf("error%d", err)
	}

	res := result

	return res

}

func canPurchase(login string, db *sql.DB) (string, error) {
	var result string

	if err := db.QueryRow("select password from account where login = $1", login).Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return result, fmt.Errorf("not found")
		}

		return result, fmt.Errorf("not found")

	}
	return result, nil
}

func WorkSpace(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	mutex.Lock()
	//We look up which method was received, and then we call the corresponding functions
	switch r.Method {
	case http.MethodPost:
		NewUser(w, r)

	case http.MethodGet:

	default:
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Please send a correct request!", 405)
	}
	mutex.Unlock()
}

func RecOrder(w http.ResponseWriter, r *http.Request) {
	var req = new(OrderList)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "err %q\n", err.Error())
	} else {

		err = json.Unmarshal(body, &req)

		if err != nil {
			fmt.Println(w, "can't unmarshal: ", err.Error())
		}
	}
	switch req.OrderName {
	case "":
		rows, err := db.Query("select * from orderList where users = $1", req.Users)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			p := OrderList{}
			err := rows.Scan(&p.OrderID, &p.OrderName, &p.Description, &p.Status, &p.Users)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Fprintf(w, "Name orders:%s status:%s creater:%s\n", p.Users, p.Status, p.OrderName)
		}
	default:
		rows, err := db.Query("select * from orderList where orderName = $1", req.OrderName)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			p := OrderList{}
			err := rows.Scan(&p.OrderID, &p.OrderName, &p.Description, &p.Status, &p.Users)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Fprintf(w, "Name orders:%s status:%s creater:%s\n", p.Users, p.Status, p.OrderName)
		}

	}

}

func GiveOrder(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from orderList")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := OrderList{}
		err := rows.Scan(&p.OrderID, &p.OrderName, &p.Description, &p.Status, &p.Users)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Fprintf(w, "Name orders:%s status:%s creater:%s\n", p.Users, p.Status, p.OrderName)
	}

}

func OrderWorkSpace(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	mutex.Lock()

	switch r.Method {
	case http.MethodPost:
		RecOrder(w, r)

	case http.MethodGet:
		GiveOrder(w, r)

	default:
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Please send a correct request!", 405)
	}
	mutex.Unlock()
}

func DeserializeRequest(w http.ResponseWriter, r *http.Request) *DefoltDAata {
	var req = new(DefoltDAata)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "err %q\n", err.Error())
	} else {

		err = json.Unmarshal(body, &req)

		if err != nil {
			fmt.Println(w, "can't unmarshal: ", err.Error())
		}
	}
	return req

}
