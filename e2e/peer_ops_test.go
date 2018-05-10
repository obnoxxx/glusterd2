package e2e

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddRemovePeer(t *testing.T) {
	r := require.New(t)

	g1, err := spawnGlusterd("./config/1.toml", true)
	r.Nil(err)
	defer g1.Stop()
	r.True(g1.IsRunning())

	g2, err := spawnGlusterd("./config/2.toml", true)
	r.Nil(err)
	defer g2.Stop()
	r.True(g2.IsRunning())

	g3, err := spawnGlusterd("./config/3.toml", true)
	r.Nil(err)
	defer g3.Stop()
	r.True(g3.IsRunning())

	client := initRestclient(g1.ClientAddress)

	_, err2 := client.PeerAttach(g2.PeerAddress)
	r.Nil(err2)

	time.Sleep(6 * time.Second)

	// add peer: ask g1 to add g3 as peer
	_, err3 := client.PeerAttach(g3.PeerAddress)
	r.Nil(err3)

	time.Sleep(6 * time.Second)

	// list and check you have 3 peers in cluster
	peers, err4 := client.Peers()
	r.Nil(err4)
	r.Len(peers, 3)

	// remove peer: ask g1 to remove g2 as peer
	err5 := client.PeerDetach(g2.PeerID())
	r.Nil(err5)
}
