package main

import (
	"encoding/json"
	"fmt"

	webpush "github.com/SherClockHolmes/webpush-go"
)

// const (
// 	subscription    = "mailto: <nikhil@chai.finance>"
// 	vapidPublicKey  = "BJkwr8YFq0MvT7PxYFrsapQkS82fSHcGDM5t7aGBF4FBlRCM8Pjy1APatWUjrvx-ICk4io54ks7GN78jPqyhZYo"
// 	vapidPrivateKey = "muV0vo0ls3pC08z8NGdp3YxbleGZLDa3lpt3tCFBaE4"
// )

const (
	subscription = `{
		"endpoint": "https://fcm.googleapis.com/fcm/send/e_CA5sCpUR0:APA91bEVPOK-P6ynJ_mWbBV7NtM92TnMitBx2_baBKJE3ePaAIjCeRVKQXmoEy7cbIakPQqY-o8Sn78VVCXvEDkb4GiXeAENUPqICfY54Qy0NRSDfKwrh93sxKhtTPtdomiix0SOCVBA",
		"expirationTime": null,
		"keys": {
		  "p256dh": "BB634peJSxkrSsahe1TQ_Us25azoT1CFb8R4lL_V4_NHvhob6Ilg4wSdbRpQ_fOUP29GW9HYisIv7XijD-0PC6E",
		  "auth": "mquCOUITvqYgt2_pVjFECQ"
		}
	  }`
	vapidPublicKey  = "BJkwr8YFq0MvT7PxYFrsapQkS82fSHcGDM5t7aGBF4FBlRCM8Pjy1APatWUjrvx-ICk4io54ks7GN78jPqyhZYo"
	vapidPrivateKey = "muV0vo0ls3pC08z8NGdp3YxbleGZLDa3lpt3tCFBaE4"
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Payment Received of 2000 THB"), s, &webpush.Options{
		Subscriber:      "nikhil@chai.finance",
		VAPIDPrivateKey: vapidPrivateKey,
		VAPIDPublicKey:  vapidPublicKey,
		TTL:             30,
	})
	if err != nil {
		fmt.Println("err :: ", err)
		// TODO: Handle error
	}
	defer resp.Body.Close()
}
