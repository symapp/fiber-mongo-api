package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Length      string             `json:"length,omitempty" bson:"length,omitempty"`
	ArtistId    string                `json:"artist_id,omitempty" bson:"artist_id,omitempty"`
	AmtSongs    int                `json:"amt_songs,omitempty" bson:"amt_songs,omitempty"`
	ReleaseYear int                `json:"release_year,omitempty" bson:"release_year,omitempty"`
	Sales       int                `json:"sales,omitempty" bson:"sales,omitempty"`
	Producer    string             `json:"producer,omitempty" bson:"producer,omitempty"`
	Genre       string             `json:"genre,omitempty" bson:"genre,omitempty"`
	Label       string             `json:"label,omitempty" bson:"label,omitempty"`
	Studio      string             `json:"studio,omitempty" bson:"studio,omitempty"`
}
