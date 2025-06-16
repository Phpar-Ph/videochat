package webrtc

func RoomConn(c *websocket.Conn, p *Peers){
	var config webrtc.Configuration
    peerConnection, err := webrtc.NewPerConnection(config)
	if err != nil{
		log.Print(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection : peerConnection,
		WebSocket: &ThreadSafeWriter{},
		Conn : c,
		Mutex: sync.Mutex{},

	}
}