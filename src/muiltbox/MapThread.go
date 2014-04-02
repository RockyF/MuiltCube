/**
 * User: RockyF
 * Date: 13-12-5
 * Time: 上午10:02
 */
package muiltbox

import (
	"fmt"
	//"time"
)

type MapThread struct{
	fps			int
	frameChan	chan int
}

func CreateMapThread(fps int) * MapThread{
	mp := &MapThread{fps: fps}
	mp.frameChan = make(chan int)

	return mp
}

func (this *MapThread) start(){
	go this.handler()
}

func (this *MapThread) handler(){
	for{
		//time.Sleep(time.Millisecond * (6000 / this.fps))
		fmt.Println("frame start.")

		fmt.Println("frame finish.")
	}
}
