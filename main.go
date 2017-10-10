package main
import (
	"crypto/rand"
	"github.com/vsergeev/btckeygenie/btckey"
	"strings"
	"fmt"
)

type keyChannel chan *btckey.PrivateKey
func getKey(ch keyChannel) {
	key, err :=btckey.GenerateKey(rand.Reader)
	if err != nil {
		ch <- nil
		return
	}
	ch <- &key
}
const THREAD_COUNT int = 8

func main()  {
	resChannel := make(keyChannel)

	for i:= 0; i < THREAD_COUNT; i++ {
		go getKey(resChannel)
	}
	var key *btckey.PrivateKey
	var addr string

	for {
		key = <-resChannel
		go getKey(resChannel)
		if key == nil {
			continue
		}
		addr = key.ToAddress()
		if strings.Contains(strings.ToLower(addr), "meerkat") {
			fmt.Println(addr, key.ToWIFC())
		}
	}
}
