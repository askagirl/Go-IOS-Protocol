package network

import (
	"net"
	"sync"

	"github.com/iost-official/Go-IOS-Protocol/common/mclock"
	"github.com/iost-official/Go-IOS-Protocol/network/discover"
)

// Peer manages connections with other nodes.
type Peer struct {
	conn      net.Conn
	blockConn net.Conn
	local     string
	remote    string
	created   mclock.AbsTime
	closed    chan struct{}
}

// Disconnect disconnects a connection.
func (p *Peer) Disconnect() {
	if p != nil && p.conn != nil {
		p.conn.Close()
	}
	if p != nil && p.blockConn != nil {
		p.blockConn.Close()
	}
}

func newPeer(conn net.Conn, blockConn net.Conn, local, remote string) *Peer {
	return &Peer{
		conn:      conn,
		blockConn: blockConn,
		local:     local,
		remote:    remote,
		created:   mclock.Now(),
		closed:    make(chan struct{}),
	}
}

// peerSet represents the collection of active peers.
type peerSet struct {
	peers  map[string]*Peer
	lock   sync.Mutex
	closed bool
}

// Get returns a connection with a node.
func (ps *peerSet) Get(node *discover.Node) *Peer {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	if ps.peers == nil {
		return nil
	}
	peer, ok := ps.peers[node.String()]
	if !ok {
		return nil
	}
	return peer
}

// SetAddr stores a peer in peerSet.
func (ps *peerSet) SetAddr(addr string, p *Peer) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	p.remote = addr
	if ps.peers == nil {
		ps.peers = make(map[string]*Peer)
	}
	ps.peers[addr] = p
	return
}

// Set stores a peer in peerSet.
func (ps *peerSet) Set(node *discover.Node, p *Peer) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	p.remote = node.String()
	if ps.peers == nil {
		ps.peers = make(map[string]*Peer)
	}
	ps.peers[node.String()] = p
	return
}

// RemoveByNodeStr removes a peer in peerSet by nodeStr.
func (ps *peerSet) RemoveByNodeStr(nodeStr string) {
	node, _ := discover.ParseNode(nodeStr)
	ps.Remove(node)
}

// Remove removes a peer in peerSet.
func (ps *peerSet) Remove(node *discover.Node) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	ps.peers[node.String()].Disconnect()
	delete(ps.peers, node.String())
	return
}
