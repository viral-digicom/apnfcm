package models

type AndroidAPN struct {
	RegistrationIds []string
	Data            interface{}
	Notification    AndroidNotification
	Priority        string
	Sound           bool
}

type AndroidNotification struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Icon         string `json:"icon"`
	Click_action string `json:"click_action"`
}

func (a AndroidAPN) Map() map[string]interface{} {
	aps := make(map[string]interface{}, 0)
	if len(a.RegistrationIds) > 0 {
		aps["registration_ids"] = a.RegistrationIds
	}
	aps["notification"] = a.Notification
	aps["priority"] = a.Priority
	aps["sound"] = a.Sound
	if a.Data != nil {
		aps["data"] = a.Data
	}
	return aps
}

type IOSAPS struct {
	Alert            Alert
	Badge            int
	sound            string
	contentAvailable int
	category         string
	threadId         string
}

type Alert struct {
	Title         string
	Body          string
	TitleLockKey  string
	TitleLocArgs  string
	ActionLockKey string
	LockKey       string
	LockArgs      string
	LaunchImage   string
}

func (a IOSAPS) Map() map[string]interface{} {
	aps := make(map[string]interface{}, 0)
	aps["alert"] = a.Alert.Map()
	if a.Badge >= -1 {
		aps["badge"] = a.Badge
	}
	if a.sound != "" {
		aps["sound"] = a.sound
	}
	if a.contentAvailable >= -1 {
		aps["content-available"] = a.contentAvailable
	}
	if a.category != "" {
		aps["category"] = a.category
	}
	return map[string]interface{}{"aps": aps}
}

func (a Alert) Map() map[string]interface{} {
	alert := make(map[string]interface{}, 0)
	if a.Title != "" {
		alert["title"] = a.Title
	}
	if a.Body != "" {
		alert["body"] = a.Body
	}
	if a.TitleLockKey != "" {
		alert["title-loc-key"] = a.TitleLockKey
	}
	if a.TitleLocArgs != "" {
		alert["title-loc-args"] = a.TitleLocArgs
	}
	if a.ActionLockKey != "" {
		alert["action-loc-key"] = a.ActionLockKey
	}
	if a.LockKey != "" {
		alert["loc-key"] = a.LockKey
	}
	if a.LockArgs != "" {
		alert["loc-args"] = a.LockArgs
	}
	if a.LaunchImage != "" {
		alert["launch-image"] = a.LaunchImage
	}
	return alert
}
