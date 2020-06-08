package main

import "github.com/etcd-io/etcd/raft"

func main() {
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}
/*	// Set peer list to the other nodes in the cluster.
	// Note that they need to be started separately as well.
	n := raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})*/

	// Create storage and config as shown above.
	// Set peer list to itself, so this node can become the leader of this single-node cluster.
	peers := []raft.Peer{{ID: 0x01}}
	n := raft.StartNode(c, peers)
}

func main1() {
	storage := raft.NewMemoryStorage()

	// Recover the in-memory storage from persistent snapshot, state and entries.
	storage.ApplySnapshot(snapshot)
	storage.SetHardState(state)
	storage.Append(entries)

	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	// Restart raft without peer information.
	// Peer information is already included in the storage.
	n := raft.RestartNode(c)
}