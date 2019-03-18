package main
 
import "time"
 
type Item struct {
	Id		  int 		`json:"id"`
    Name      string    `json:"name"`
    Checked   bool      `json:"checked"`
    Due       time.Time `json:"due"`
    Quantity  int   	`json:"quantity"`
}
 
type Items []Item