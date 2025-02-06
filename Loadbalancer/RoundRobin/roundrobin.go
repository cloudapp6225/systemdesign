package roundrobin

// RoundRobin takes a list of servers and returns a function that will return the next server in the list each time it is called.
var count int

type RoundRobin struct {
	servers []string
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) RoundRobin(servers []string) string {
	// Implement the RoundRobin function
	size := len(servers)
	// simple algorithm that counts the number of requests and count % size will give the index of the server
	// to which the request should be forwarded
	server := servers[count%size]
	count++
	return server
}
