package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	ArtistSince      int                `json:"artist_since,omitempty" bson:"artist_since,omitempty"`
	Genres           []string           `json:"genres,omitempty" bson:"genres,omitempty"`
	MonthlyListeners int                `json:"monthly_listeners,omitempty" bson:"monthly_listeners,omitempty"`
	Website          string             `json:"website,omitempty" bson:"website,omitempty"`
}
