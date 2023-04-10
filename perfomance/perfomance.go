package perfomance

// go test -bench
// go test -bench='xxx$' .
// go test -bench -cpu=2,4 .  指定使用的cpu核心数
// go test -bench -benchtime=5s . 指定执行时间为5s
// go test -bench -benchtime=30x . 指定执行次数为30次
// go test -bench -count=3 . 指定执行三轮（默认一轮）
// go test -bench -benchmem 查看内存分配情况
// go test -bench -cpuprofile=cpu.pprof . 生成cpu profile文件，后续可以用go tool pprof -http=:9999 /cpu.pprof 来分析
// go test -bench -memprofile=mem.pprof,-memprofilerate=N . 生成mem profile. 调整记录速率为原来的1/N

// 输出结果： BenchmarkXxxx-n m xns/op    -n表示使用的cpu核心数，m 表示执行的次数，x 表示每个操作的耗时（单位ns）

// b.ResetTimer() 重置定时器，屏蔽上面耗时操作的影响
// b.StopTimer()-b.StartTimer() 暂定和继续计时器
