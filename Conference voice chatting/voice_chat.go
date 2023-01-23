package voice_chat

import (
    "github.com/pion/webrtc/v2"
)

// Create a new WebRTC peer connection

peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
    ICEServers: []webrtc.ICEServer{
        {
            URLs: []string{"stun:stun.l.google.com:19302"},
        },
    },
})
if err != nil {
    log.Fatalf("Failed to create peer connection: %v", err)
}

// Create a new audio track for the local user
audioTrack, err := peerConnection.NewTrack(webrtc.DefaultPayloadTypeOpus, rand.Uint32(), "audio", "pion")
if err != nil {
    log.Fatalf("Failed to create audio track: %v", err)
}

//Add the audio track to the peer connection
if _, err = peerConnection.AddTrack(audioTrack); err != nil {
    log.Fatalf("Failed to add audio track: %v", err)
}

// Create an offer and send it to the other users

offer, err := peerConnection.CreateOffer(nil)
if err != nil {
    log.Fatalf("Failed to create offer: %v", err)
}
if err = peerConnection.SetLocalDescription(offer); err != nil {
    log.Fatalf("Failed to set local description: %v", err)
}

// Handle incoming answers and ICE candidates from the other users

peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
    // Send the candidate to the other users
})
peerConnection.OnTrack(func(track *webrtc.Track, receiver *webrtc.RTPReceiver) {
    // Handle the incoming track
})

// Start the audio track and wait for the other users to join
if err = audioTrack.Start(); err != nil {
    log.Fatalf("Failed to start audio track: %v", err)
}

/*

ToDo:
Implement functions to handle disconnection, reconnection, and mute-unmute options.
Make use of github.com/ flashtalking/go-webrtc or github.com/deepch/voice.
Ensure encryption and authentication.

*/

