package DB

//
//import (
//	"context"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/jackc/pgx/v5"
//	"log"
//	"os"
//)
//
//func main() {
//
//	conn, err := pgx.Connect(context.Background(), "postgresql://localhost/postgres?user=postgres&password=1234")
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//	rows, err := conn.Query(context.Background(), "select * from public.orders")
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//
//	// iterate through the rows
//	for rows.Next() {
//		values, err := rows.Values()
//		if err != nil {
//			log.Fatal("error while iterating dataset")
//		}
//
//		defer conn.Close(context.Background())
//		vals := values[0].([16]byte)
//		id, _ := uuid.FromBytes(vals[:])
//		firstName := values[1].(string)
//		lastName := values[2]
//		dateOfBirth := values[3]
//		log.Println("id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth)
//		sqlString := fmt.Sprintf("select * from public.orders where id = ('%s'::uuid)", id)
//		row, err := conn.Query(context.Background(), sqlString)
//
//		for row.Next() {
//			value11, err := rows.Values()
//			if err != nil {
//				log.Fatal("error while iterating dataset")
//			}
//			firstName := value11[1].(string)
//			lastName := value11[2]
//			dateOfBirth := value11[3]
//			log.Println("id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth)
//		}
//
//	}
//
//	var greeting string
//	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Println(greeting)
//}
