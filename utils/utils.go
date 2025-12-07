package utils
import (
	"os"
)

func ReadFile(location string)(string,error){
	contents,err := os.ReadFile(location)
	if err != nil {
		return "",err
	}
	return string(contents),nil
}
