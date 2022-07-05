package ex25

import (
	"context"
	"time"
)

//sleep с помощью after, ожидает истечения продолжительности d, а затем отправляет текущее время по возвращаемому каналу
func SleepAfter(d time.Duration) {
	<-time.After(d)
}

//sleep с помощью tick, ожидание, пока не вернется значение в канал через duration
func SleepTick(duration time.Duration) {
	<-time.Tick(duration)
}

//sleep с помощью context c timeout duration
func SleepTimeout(duration time.Duration) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()

	<-ctx.Done()
}

//sleep с помощью context c deadline duration
func SleepDeadline(duration time.Duration) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(duration))
	defer cancelFunc()

	<-ctx.Done()
}
