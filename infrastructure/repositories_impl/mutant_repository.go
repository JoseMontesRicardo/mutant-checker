package repositories_impl

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test.com/mutant-checker/domain/models"
	"test.com/mutant-checker/domain/types"
)

type DbInstance struct {
	collection *mongo.Collection
}

var MutantRepo DbInstance

// Setup prepares the collection
func Setup(db *mongo.Database) {
	MutantRepo.collection = db.Collection("mutant")
}

func (mr *DbInstance) Insert(doc models.Mutant) (id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, _ := mr.collection.InsertOne(ctx, doc)
	return fmt.Sprintf("%v", res.InsertedID)
}

func (mr *DbInstance) GetStats() types.Stats {
	matchStep := bson.D{{"$match", bson.D{{}}}}
	groupStep := bson.D{
		{
			"$group", bson.M{
				"_id": "mutant_counter",
				"count_human_dna": bson.M{
					"$sum": bson.M{
						"$cond": bson.A{
							bson.M{
								"$eq": bson.A{"$ismutant", false},
							},
							1,
							0,
						},
					},
				},
				"count_mutant_dna": bson.M{
					"$sum": bson.M{
						"$cond": bson.A{
							bson.M{
								"$eq": bson.A{"$ismutant", true},
							},
							1,
							0,
						},
					},
				},
			},
		},
	}
	project1 := bson.D{
		{
			"$project", bson.M{
				"count_human_dna": bson.M{
					"$cond": bson.A{
						bson.M{
							"$eq": bson.A{
								"$count_human_dna",
								0,
							},
						},
						1,
						"$count_human_dna",
					},
				},
				"count_mutant_dna": "$count_mutant_dna",
			},
		},
	}
	project2 := bson.D{
		{
			"$project", bson.M{
				"count_human_dna":  1,
				"count_mutant_dna": 1,
				"ratio": bson.M{
					"$divide": bson.A{
						"$count_mutant_dna",
						"$count_human_dna",
					},
				},
			},
		},
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := mr.collection.Aggregate(ctx, mongo.Pipeline{matchStep, groupStep, project1, project2})
	if err != nil {
		panic(err)
	}
	var showsLoaded []types.Stats
	if err = cursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}
	return showsLoaded[0]
	// return showsLoaded[0]
}
