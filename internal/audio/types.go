package audio

type AudioSource struct {
	Driver              string `json:"driver,omitzero"`
	Index               int    `json:"index,omitzero"`
	Name                string `json:"name,omitzero"`
	SampleSpecification string `json:"sample_specification,omitzero"`
	State               string `json:"state,omitzero"`
}
