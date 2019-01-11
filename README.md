# Send Push messsage through FCM and APN
APN and FCM micro service

## FCM - HOW TO USE
``` go
package main

import (
	"apnfcm"
	"apnfcm/models"
	"fmt"
)

func main() {
	deviceIds := []string{""}
	apnfcm.InitAndroidAPN("")
	responses, err := apnfcm.SendAndroid(models.AndroidAPN{RegistrationIds: deviceIds, Notification:models.AndroidNotification{Title:"MY TITLE", Body:"MY BODY"}, Sound: true, Priority:"High"})
	if err != nil {
		fmt.Println(err)
	}
	for _, response := range responses {
		fmt.Println(response)
	}
}
```
## APN - HOW TO USE
``` go
package main

import (
	"apnfcm"
	"apnfcm/models"
	"fmt"
)

func main() {
	deviceIds := []string{""}
	apnfcm.InitIosAPN("","","","")
	responses, err := apnfcm.SendIOS(deviceIds, models.IOSAPS{Alert:models.Alert{Body:"BODY", Title:"TITLE"}})
	if err != nil {
		fmt.Println(err)
	}
	for _, response := range responses {
		fmt.Println(response)
	}
}
```