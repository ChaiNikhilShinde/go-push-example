package main

import (
	"encoding/json"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	subscription    = "mailto: <nikhil@chai.finance>"
	vapidPublicKey  = "BJkwr8YFq0MvT7PxYFrsapQkS82fSHcGDM5t7aGBF4FBlRCM8Pjy1APatWUjrvx-ICk4io54ks7GN78jPqyhZYo"
	vapidPrivateKey = "muV0vo0ls3pC08z8NGdp3YxbleGZLDa3lpt3tCFBaE4"
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      subscription,
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		// TODO: Handle error
	}
	defer resp.Body.Close()
}
