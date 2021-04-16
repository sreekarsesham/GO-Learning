package main
import(

	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"

)


func main(){

	var userPassword 	[]byte;
	var err 			error;
	var hash			string;
	
	fmt.Println("Enter a password to Hash");
	_,err = fmt.Scan(&userPassword);
	
	
	if err != nil {
        log.Println(err)
    }
	
	hash = hashIt(userPassword);
	fmt.Println("The hashed password is :",hash);

}

func hashIt(userPwd []byte) string{
	
	var hashedPwd 	[]byte;
	var err 		error;
	
	hashedPwd,err = bcrypt.GenerateFromPassword(userPwd,bcrypt.MinCost)// can also use bcrypt.MinCost or bcrypt.MaxCost or bcrypt.DefaultCost
	
	if err != nil {
        log.Println(err)
    }
    
     return string(hashedPwd)
	
}




//references : https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
//			   https://pkg.go.dev/golang.org/x/crypto/bcrypt#Cost
