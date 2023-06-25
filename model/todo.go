package model

type Todo struct {
	ID					string 		`json:"id"`
	Title 			string 		`json:"title"`
	Completed 	bool 			`json:"isComplete"`
}
