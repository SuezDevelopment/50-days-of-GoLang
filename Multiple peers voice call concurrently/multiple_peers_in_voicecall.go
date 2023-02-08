package multiple_peers_in_voicecall

import (
	"bufio"
	"fmt"
	"github.com/pion/webrtc/v2"
	"github.com/pion/webrtc/v2/examples/internal/signal"
	"log"
	"net"
	"strconv"
	"sync"
)


var peers = make(map[*webrtc.RTPReceiver]*webrtc.PeerConnection)
var peerLock sync.Mutex

type VoiceCall struct {
	peerConnections []*webrtc.PeerConnection
	audioTracks     map[*webrtc.Track]int
}

func StartVoiceCall() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to accept: %s", err)
		}
		go handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})
  
  if err != nil {
		log.Fatalf("Failed to create PeerConnection: %s", err)
	}
	audioTrack, err := peerConnection.NewTrack(webrtc.DefaultPayloadTypeOpus, 12345, "audio", "pion1")
	if err != nil {
		log.Fatalf("Failed to create audio track: %s", err)
	}

	if _, err = peerConnection.AddTrack(audioTrack); err != nil {
		log.Fatalf("Failed to add audio track: %s", err)
	}

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Fatalf("Failed to create offer: %s", err)
	}

	if err = peerConnection.SetLocalDescription(offer); err != nil {
		log.Fatalf("Failed to set local description: %s", err)
	}
  
  // Wait for the remote description to be set
	signal.Decode(offer, &peerConnection)

	// Start the audio stream
	peerConnection.OnTrack(func(track *webrtc.Track, receiver *webrtc.RTPReceiver) {
		if track.PayloadType() == webrtc.DefaultPayloadTypeOpus {
			log.Printf("Received an Opus track")
		}
	})
  
  for i, pc := range peerConnections {
		if i == len(peerConnections)-1 {
			continue
		}
		if _, err := pc.AddTrack(audioTrack); err != nil {
			return err
		}
	}
}

func (vc *VoiceCall) Join(peerConnection *webrtc.PeerConnection) error {
	// Add the new peer connection to the list of peer connections
	vc.peerConnections = append(vc.peerConnections, peerConnection)

	// Create a new audio track
	audioTrack, err := peerConnection.NewTrack(webrtc.DefaultPayloadTypeOpus, 12345, "audio", "pion1")
	if err != nil {
		return err
	}

	// Add the audio track to the peer connection
	if _, err = peerConnection.AddTrack(audioTrack); err != nil {
		return err
	}

	// Add the audio track to the map of audio tracks
	vc.audioTracks[audioTrack] = len(vc.peerConnections) - 1

	// Create an offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		return err
	}

	// Set the local description
	if err = peerConnection.SetLocalDescription(offer); err != nil {
		return err
	}

	// Wait for the remote description to be set
	signal.Decode(offer, &peerConnection)

	// Add the audio track to all other peer connections
	for i, pc := range vc.peerConnections {
		if i == len(vc.peerConnections)-1 {
			continue
		}
		if _, err := pc.AddTrack(audioTrack); err != nil {
			return err
		}
	}

	return nil
}

func (vc *VoiceCall) Leave(peerConnection *webrtc.PeerConnection) error {
	// Remove the peer connection from the list of peer connections
	for i, pc := range vc.peerConnections {
		if pc == peerConnection {
			vc.peerConnections = append(vc.peerConnections[:i], vc.peerConnections[i+1:]...)
			break
		}
	}

	// Remove the audio track from the map of audio tracks
	for track := range peerConnection.GetSenders() {
		if idx, ok := vc.audioTracks[track]; ok {
			delete(vc.audioTracks, track)
			for _, pc := range vc.peerConnections {
				if err := pc.RemoveTrack(pc.GetReceivers()[idx]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

