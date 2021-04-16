package api

type event struct {
	eventType string `json:"type"`
	game      game   `json:"game"`
}
type game struct {
	id string `json:"id"`
}

func temp() {

}
