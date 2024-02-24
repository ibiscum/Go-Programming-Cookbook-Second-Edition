package storage

import (
	"context"
	"reflect"
	"testing"
)

func TestNewMongoStorage(t *testing.T) {
	type args struct {
		ctx        context.Context
		connection string
		db         string
		collection string
	}
	tests := []struct {
		name    string
		args    args
		want    *MongoStorage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoStorage(tt.args.ctx, tt.args.connection, tt.args.db, tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
