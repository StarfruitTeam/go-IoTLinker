package NetService

import "../nats/nats_client"
import "./connecter"

type NetServcie struct {
	connecters map[string]connecter
	client nats_client
}

// Net Service 는 Connecter 을 가진다 
