package main
import "fmt"
func main(){
  const angka uint8 = 5;
  if angka % 2 == 0{
    fmt.Println("Bilangan genap")
  }else if angka % 2 != 0{
    fmt.Println("Bilangan ganjil")
  }else{
    fmt.Println ("error")
}
}
