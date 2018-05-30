package easycall

type SendOptions struct {
	Receiver    int    `json:"receiver"`
	Message     string `json:"message"`
	Port        string `json:"port,omitempty"`
	Form        string `json:"type,omitempty"`
	Bps         int    `json:"bsp,omitemty"`
	Polarity    string `json:"polarity,omitempty"`
	Tone        string `json:"tone,omitempty"`
	AutoShorten bool   `json:"autoShorten,omitempty"`
}
