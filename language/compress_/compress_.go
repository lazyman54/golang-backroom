package compress_

import (
	"encoding/json"
	"fmt"
	"github.com/lazymank54/golang-backroom/language/encode"
	"time"
)

var holder = make(map[string]ICompress)

func CompressSolution() {
	times := 50000
	originData := getOriginStrData()
	//fmt.Println(string(originData))
	for key, _ := range holder {
		//if key != "snappy" {
		//	continue
		//}
		time1 := time.Now()
		for i := 0; i < times; i++ {
			doTimeCompress(key, originData)
		}
		dura := time.Now().Sub(time1).Milliseconds()
		fmt.Printf("%s compress cost:%dms\n", key, dura)
	}
	//doSolution("zstd", originData)
}

func doSolution(name string, data []byte) {
	compressObj := holder[name]

	compressedData, _ := compressObj.compress(data)
	fmt.Printf("%s Solution compress. byte len: %d\n", name, len(compressedData))
	compressObj.deCompress(compressedData)
	//fmt.Printf("originData byte len:%d\n", len(originData))
}
func doTimeCompress(name string, data []byte) {
	compressObj := holder[name]

	compressObj.compress(data)

}

func register(compressObj ICompress) {
	holder[compressObj.name()] = compressObj
}

func getOriginStrData() []byte {
	battleInfo := encode.DataForJson()
	battleByte, _ := json.Marshal(battleInfo)
	fmt.Printf("origin data byte len:%d\n", len(battleByte))
	return battleByte
}
