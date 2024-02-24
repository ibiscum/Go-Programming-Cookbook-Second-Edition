package storage

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoStorage_GetByName(t *testing.T) {
	type fields struct {
		Client     *mongo.Client
		DB         string
		Collection string
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoStorage{
				Client:     tt.fields.Client,
				DB:         tt.fields.DB,
				Collection: tt.fields.Collection,
			}
			got, err := m.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoStorage.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoStorage.GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
