package main

import "github.com/thealamu/papilo/pkg/papilo"

func lowerCmpt(p *papilo.Pipe) {
	for !p.IsClosed { // read for as long as the pipe is open
		// p.Next returns the next data in the pipe
		d, _ := p.Next()
		byteData, ok := d.([]byte)
		if !ok {
			// we did not receive a []byte, we can be resilient and move on
			continue
		}
		// Write to next pipe
		p.Write(bytes.ToLower(byteData))
	}
}

func main() {
	p := papilo.New()
	m := &papilo.Pipeline{
		Components: []papilo.Component{lowerCmpt},
	}
	p.Run(m)
}
