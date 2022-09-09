package database

import (
	"context"
	"time"

	jp "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/application/domain/job_posting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (dba *Adapter) CreateJobPostings(postings []jp.JobPosting) error {
	postingsMongo := make([]interface{}, 0)
	for _, v := range postings {
		v.CreatedAt = time.Now().UTC().Unix()
		postingsMongo = append(postingsMongo, v)
	}
	_, err := dba.collection.InsertMany(context.Background(), postingsMongo)
	if err != nil {
		return err
	}
	return nil
}

func (dba *Adapter) RetrieveJobPostings(params jp.SearchFilterParams, page int, limit int) ([]jp.JobPosting, error) {

	var query bson.M
	queryAnd := []bson.M{}
	dateProvided := false

	if !params.CreatedBefore.IsZero() {
		dateProvided = true
		queryAnd = append(queryAnd, bson.M{"created_at": bson.M{"$lt": params.CreatedBefore.UTC().Unix()}})
	}

	if !params.CreatedAfter.IsZero() {
		dateProvided = true
		queryAnd = append(queryAnd, bson.M{"created_at": bson.M{"$gt": params.CreatedAfter.UTC().Unix()}})
	}

	if !dateProvided {
		query = bson.M{}
	} else {
		query = bson.M{"$and": queryAnd}
	}

	opts := new(options.FindOptions)
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}

	ctx := context.TODO()

	curs, err := dba.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}

	var postings []jp.JobPosting
	if err = curs.All(ctx, &postings); err != nil {
		return nil, err
	}

	return postings, nil

}

func (dba *Adapter) DeleteJobPostings(postings []jp.JobPosting) error {
	query := []bson.M{}
	for _, posting := range postings {
		query = append(query, bson.M{
			"job_id": posting.JobID, "title": posting.Title,
		})
	}

	_, err := dba.collection.DeleteMany(context.Background(), bson.M{"$or": query})
	if err != nil {
		return err
	}
	return nil
}
