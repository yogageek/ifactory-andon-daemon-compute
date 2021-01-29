package logic

import (
	andon_logic "iii/ifactory/compute/pkg/logic/andon_logic"
	"time"
	// . "measure/logic"
)

var (
	timer1    = time.NewTimer(time.Second * 0) //這裡設幾秒 就會等幾秒才開始
	duration1 = time.Second * time.Duration(5)
	timer2    = time.NewTimer(time.Second * 30) //這裡設幾秒 就會等幾秒才開始
	duration2 = time.Minute * time.Duration(10)
)

func RunDaemonLoop() {
	for {
		select {
		case <-timer1.C:
			// wg := sync.WaitGroup{}
			// wg.Add(2)

			andon_logic.Daemon_AbnormalMachineLatest()
			// andon_logic.Daemon_AbnormalMachineHist()

			// wg.Wait()
			timer1.Reset(duration1)

		case <-timer2.C:

			func() {

				timer2.Reset(duration2)
			}()
		}
	}
}
