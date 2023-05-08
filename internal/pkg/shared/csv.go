package shared

import (
	"encoding/csv"
	"io"
	"os"

	accountModel "eirc.app/internal/v1/structure/accounts"
)

//為了導入EMAIL用的(暫時寫死)
func CsvToUser(filename string,created_account string) (output []*accountModel.UpdatedCsv,err error){
	filename = "C:/Users/User/Go/src/hta_apis/"+filename
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	result1 := make([][]string, 0)
	result2 := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		result1 = append(result1, record)
	}

	for k, _ := range result1 {
		if result1[k][0] == "Account" && result1[k][1] == "Email"{
			result2 = append(result1[:k], result1[k+1:]...)
		}
	}
	for _, v := range result2 {
		user := &accountModel.UpdatedCsv{
			AccountID: "00000000-0000-0000-0000-000000000000",
			CreatedAccount: created_account,
			//Dep: "00000000-0000-0000-0000-000000000000",
			Account: &v[0],
			Email: v[1],
			Status: true,
		}
		output = append(output, user)
		// user := &User{
		// 	Name: v[0],
		// 	Email: v[1],
		// }
		// Users = append(Users, user)
	}
	return output,nil
}
