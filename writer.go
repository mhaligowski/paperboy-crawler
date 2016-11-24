package crawler

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"github.com/nu7hatch/gouuid"
	"time"
)

type StreamItem struct {
	StreamItemId string `datastore:"id"`;
	UserId string `datastore:"user_id"`;
	TargetId string `dataastore:"target_id"`;
	Title string `datastore:"title"`;
	OrderSequence int64 `datastore:"order_sequence"`;
}

func AddEntryIfDoesntExist(ctx context.Context, entry Entry) (*Entry, bool, error) {
	var dsEntry = &Entry{}

	key := datastore.NewKey(ctx, "Entry", entry.Id, 0, nil)

	e := datastore.Get(ctx, key, dsEntry)

	if e != nil && e != datastore.ErrNoSuchEntity {
		return nil, false, e
	} else if e == datastore.ErrNoSuchEntity {
		_, e = datastore.Put(ctx, key, &entry)

		if e != nil {
			return nil, false, e
		} else {
			return dsEntry, true, nil
		}
	} else {
		return dsEntry, false, nil
	}
}

func PutStreamEntry(ctx context.Context, entry Entry, userId string) (string, error) {
	keyValue, err := uuid.NewV4();
	if err != nil {
		return "", err;
	}

	key := datastore.NewKey(ctx, "StreamItem", keyValue.String(), 0, nil)

	item := StreamItem{
		Title: entry.Title,
		OrderSequence: time.Now().UnixNano(),
		StreamItemId:keyValue.String(),
		TargetId:entry.Id,
		UserId:userId,
	}

	k, err := datastore.Put(ctx, key, &item)
	return k.String(), err
}