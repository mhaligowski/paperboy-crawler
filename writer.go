package crawler

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func addEntryIfDoesntExist(ctx context.Context, entry Entry) (*Entry, bool, error) {
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

func writeEntries(ctx context.Context, feed Feed) ([]Entry, error) {
	newEntries := make([]Entry, len(feed.Entries))
	for _, entry := range feed.Entries {
		entry.Summary = entry.Summary[:1500]

		_, created, err := addEntryIfDoesntExist(ctx, entry)
		if err != nil {
			log.Errorf(ctx, "could not put feed entry %s from feed %s: %v", entry.Id, feed.Id, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if created {
			append(newEntries, entry)
		}
	}

	return newEntries, nil
}

