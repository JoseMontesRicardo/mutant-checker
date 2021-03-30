package types

type AdnSequence = []string

type MutantBody struct {
	Dna AdnSequence `json:"dna" validate:"min=4,max=10"`
}

type MutantResponse struct {
	Status int `json:"status"`
}

type Stats struct {
	Count_human_dna  int     `bson:"count_human_dna" json:"count_human_dna"`
	Count_hutant_dna int     `bson:"count_mutant_dna" json:"count_hutant_dna"`
	Ratio            float32 `bson:"ratio,truncate" json:"ratio"`
}
