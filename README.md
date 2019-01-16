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
	deviceIds := []string{"DEVICE_IDS"}
	apnfcm.InitAndroidAPN("FCM_KEY")
	responses, err := apnfcm.SendAndroid(models.AndroidAPN{RegistrationIds: deviceIds, Notification:models.AndroidNotification{Title:"MY TITLE", Body:"MY BODY"}, Sound: true, Priority:"High"})
	if err != nil {
		fmt.Println(err)
	}
	for _, response := range responses {
		fmt.Println(response)
	}
}
```
## APN With P8 - HOW TO USE
``` go
package main

import (
	"apnfcm"
	"apnfcm/models"
	"fmt"
)

func main() {
	deviceIds := []string{"LIST_OF DEVICE_IDS"}
	apnfcm.InitIosAPN("P8_FILE_PATH","TEAM_ID","KEY_ID","TOPIC_NAME")
	responses, err := apnfcm.SendIOS(deviceIds, models.IOSAPS{Alert:models.Alert{Body:"BODY", Title:"TITLE"}})
	if err != nil {
		fmt.Println(err)
	}
	for _, response := range responses {
		fmt.Println(response)
	}
}
```
## APN With PEM - HOW TO USE
``` go
package main

import (
	"apnfcm"
	"apnfcm/models"
	"fmt"
)

func main() {
	deviceIds := []string{""}
	apnfcm.InitIosPEMAPN("PEM_FILE_PATH","TOPIC_NAME")
	responses, err := apnfcm.SendIOS(deviceIds, models.IOSAPS{Alert:models.Alert{Body:"BODY", Title:"TITLE"}})
	if err != nil {
		fmt.Println(err)
	}
	for _, response := range responses {
		fmt.Println(response)
	}
}
```