type Search func(query string) string

func First(query string, replicas ...Search) string {
	c := make(chan string)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []string) {
	c := make(chan string)
	go func() { c <- First(query, WebSearch1, WebSearch2, WebSearch3) }()
	go func() { c <- First(query, ImageSearch1, ImageSearch2) }()
	go func() { c <- First(query, VideoSearch1, VideoSearch2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}
	return results
}
