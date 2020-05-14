package postgreDB

import(
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "postgres"
  password = "ugohuche"
  dbname = "bookdata"
)

func CreateTable(tablename string)  {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement :=`
  CREATE TABLE $1 (
    id SERIAL PRIMARY KEY,
    name TEXT,
    author TEXT,
    published_at TEXT
  );`
  _, err = db.Exec(sqlStatement, tablename)
  if err != nil {
    panic(err)
  }
}

func Connect() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  //defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected")
}

func Insert(name string, author string, published_at string) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement := `
  INSERT INTO books (name, author, published_at)
  VALUES ($1, $2, $3)
  RETURNING id`
  id := 0
  err = db.QueryRow(sqlStatement, name, author, published_at).Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)

}

func Update(name string, author string, id int) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement := `
  UPDATE books
  SET name = $2, author = $3
  WHERE id = $1
  RETURNING id, published_at;`
  var id int
  var published_at string
  err = db.QueryRow(sqlStatement, id, name, author).scan(&id, &published_at)
  if err != nil {
    panic(err)
  }
  fmt.Println(id, email)
}

func Delete(id int)  {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement := `
  DELETE FROM books
  WHERE id = $1;`
  _,err = db.Exec(sqlStatement, id)
  if err != nil {
    panic(err)
  }
}
