package models

import "time"

type Vocabulary struct {
	Word         string        `json:"word" bson:"word"`
	Type         string        `json:"type" bson:"type"`
	Meaning      string        `json:"meaning" bson:"meaning"`
	Examples     []Example     `json:"examples" bson:"examples"`
	RelatedWords []RelatedWord `json:"related_words" bson:"related_words"`
	Notes        []Note        `json:"notes" bson:"notes"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

type Example struct {
	English    string `json:"english"`
	Vietnamese string `json:"vietnamese"`
}

type RelatedWord struct {
	Word    string `json:"word"`
	Type    string `json:"type"`
	Meaning string `json:"meaning"`
	Tag     string `json:"tag"`
}

type Note struct {
	Word string `json:"word"`
	Note string `json:"note"`
}
