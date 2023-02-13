package crypto

import (
	"crypto/md5"
	"fmt"
)

func HashSolution() {

	var str = "my"
	hash := md5.Sum([]byte(str))
	//md5Hash := murmurhash3.New3A()
	//hash1 := md5Hash.Sum([]byte(str))
	fmt.Printf("%x\n", hash)
	//fmt.Printf("%x\n", hash1)
	//fmt.Printf("%s\n", string(hash[:]))

}
