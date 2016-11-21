package crawler

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func AddEntryIfDoesntExist(ctx context.Context, entry *Entry) (*Entry, bool, error) {
	var dsEntry = &Entry{}

	key := datastore.NewKey(ctx, "Entry", entry.Id, 0, nil)

	e := datastore.Get(ctx, key, dsEntry)

	if e != nil && e != datastore.ErrNoSuchEntity {
		return nil, false, e
	}

	if e == datastore.ErrNoSuchEntity {
		return dsEntry, false, nil
	}

	_, e = datastore.Put(ctx, key, entry)

	if e != nil {
		return nil, false, e
	} else {
		return dsEntry, true, nil
	}
}
