package main
import "fmt"
type NameAge struct{
  Name string
  Age int
}
func main(){
  var nameAgeSlice []NameAge
  nameAges := map[string]int{
    "Michael": 30,
    "John": 25,
    "Jessica": 26,
    "Ali": 18,
  }
  for key, value := range nameAges{
    nameAgeSlice = append(nameAgeSlice, NameAge {key, value})
  }

  fmt.Println(nameAgeSlice)

}