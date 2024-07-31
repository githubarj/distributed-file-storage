package p2p

type TCPTransport interface {
  listenAddress string
  listener net.Listener 
  
  mu sync.RWMutex
  peers map[net.Addr]Peer 
}

func NewTCPTransport (listenAddr, string ) *TCPTransport {
  return &TCPTransport {
    listenAddress: listenAddr,

  }
}


