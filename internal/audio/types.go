package audio

type InputDevice struct {
	Card        int    `json:"card"`
	Name        string `json:"name"`
	Device      int    `json:"device"`
	Description string `json:"description"`
}