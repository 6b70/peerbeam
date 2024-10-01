package conn

import (
	"fmt"
	"github.com/pion/webrtc/v4"
	"log/slog"
)

func (c *Session) AddRemote(remoteSDP *webrtc.SessionDescription) error {
	err := c.Conn.SetRemoteDescription(*remoteSDP)
	if err != nil {
		return err
	}
	return nil
}

func (c *Session) SetupPeerConn() error {
	conn, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{iceServers[0]},
	})
	if err != nil {
		return err
	}
	c.Conn = conn
	c.monitorState()
	go c.connClose()

	return nil
}

func (c *Session) DataChHandler(ch *webrtc.DataChannel) {
	c.DataCh = ch
	ch.OnOpen(func() {
		c.DataChOpen <- struct{}{}
	})
	c.DataCh.OnClose(func() {
		c.CtxCancel()
	})

	c.DataCh.OnMessage(func(msg webrtc.DataChannelMessage) {
		c.MsgCh <- &msg
	})

}

func (c *Session) connClose() {
	<-c.Ctx.Done()

	if c.DataCh != nil {
		err := c.DataCh.GracefulClose()
		if err != nil {
			slog.Error("Error closing control channel:", err)
		}
	}
	if c.Conn == nil {
		err := c.Conn.GracefulClose()
		if err != nil {
			slog.Error("Error closing peer connection:", err)
		}
	}
}

func (c *Session) monitorState() {
	c.Conn.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		switch connectionState {
		case webrtc.ICEConnectionStateFailed:
			c.CtxCancel()
		default:
		}
	})
}

func (c *Session) PeerInfoStr() (string, error) {
	selectedPair, err := c.DataCh.Transport().Transport().ICETransport().GetSelectedCandidatePair()
	if err != nil {
		return "", err
	}
	remote := selectedPair.Remote
	return fmt.Sprintf("%s %s:%d", remote.Protocol, remote.Address, remote.Port), nil
}
