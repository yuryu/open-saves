// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package createrecord

/*

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "github.com/googleforgames/open-saves/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)




/*
type WriteBenchmarker struct {
	records []uuid.UUID
	stats   Stats
}

func (b *WriteBenchmarker) Run(ctx context.Context, client pb.OpenSavesClient, store *pb.Store, numIterations, numThreads int) {
	for i := 0; i < numIterations; i++ {
		key := uuid.New()
		if _, err := client.CreateRecord(ctx, &pb.CreateRecordRequest{
			StoreKey: store.GetKey(),
			Record: &pb.Record{
				Key: key.String(),
			},
		}); err != nil {

		}

	}

}

func cleanup(ctx context.Context, conn *grpc.ClientConn, storeKey string, benchmarker Benchmarker, concurrency int) {
	keyChan := make(chan uuid.UUID, concurrency)
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			client := pb.NewOpenSavesClient(conn)
			for {
				key, more := <-keyChan
				if !more {
					return
				}
				client.DeleteRecord(ctx, &pb.DeleteRecordRequest{StoreKey: storeKey, Key: key.String()})
			}
		}()
	}
	for _, v := range benchmarker.Records() {
		keyChan <- v
	}
	close(keyChan)
	wg.Wait()
	client := pb.NewOpenSavesClient(conn)
	client.DeleteStore(ctx, &pb.DeleteStoreRequest{Key: storeKey})
}

*/
