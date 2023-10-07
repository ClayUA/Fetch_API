package Handler_Functions
 
import (
		"net/http"
		"encoding/json"
		"time"
		"Fetch_API/types"
		"fmt"
)

func addHandler(w http.ResponseWriter,r http.Request){

	// Making sure the correct method is used by the user

	if r.Method != http.MethodPost {
		http.Error(w,"Invalid Method",http.StatusMethodNotAllowed)
		return
	}

	// Creating a Transaction using our type defined in types.go
	//decoding json contents and making sure they are valid with built in functions from encoding/json

	var TransactionData types.Transaction
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decoder(&TransactionData)

	if err != nil{
		http.Error(w,"Invalid Request",http.StatusBadRequest)
		return
	}
	// validing our data after it has been decoded
	// using a helper function
	err := validateJSON(TransactionData)
	if err != nil {
		http.Error(w, "",http.StatusBadRequest)
	}



}

func validateJSON(t types.Transaction){
	if t.Points <= 0 {
		return fmt.Errorf("Invalid Point size. Must be greater than 0")
	}
	if t.Payer == "" {
		return fmt.Errorf("Payer name must be included")
	}
	format := "2022-11-02T14:00:00Z"
	_,err := time.Parse(format,t.Time)
	if err != nil {
		return fmt.Errorf("Invalid Time Format",)
	}
	return nil
}



}







