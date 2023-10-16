package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/md5"
	"encoding/hex"
	"goBlog/domain"
	"time"
)

type PostDoc struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type PostFirebaseDao struct {
	FirestoreClient   *firestore.Client
	CollectionsPrefix string
	Context           context.Context
}

const collectionName = "Post"

func (r PostFirebaseDao) Create(post domain.Post, user string) (string, error) {
	postDoc := map2Doc(post)
	postId := getPostHash(user, postDoc)
	docRef := r.createDocRef(postId, user, r.CollectionsPrefix, collectionName)
	_, werr := docRef.Create(r.Context, postDoc)
	if werr != nil {
		return "", werr
	}
	return postId, nil
}

func (r PostFirebaseDao) createDocRef(id string, user string, prefix string, collName string) *firestore.DocumentRef {
	return r.FirestoreClient.Doc(prefix + "_" + collName + "/" + id)
}

func getPostHash(user string, doc *PostDoc) string {
	hash := md5.Sum([]byte(user + doc.Title + doc.CreatedAt.String()))
	return hex.EncodeToString(hash[:])
}

func map2Doc(post domain.Post) *PostDoc {
	return &PostDoc{
		Title:     post.Title,
		Content:   post.Content,
		Author:    post.Author,
		CreatedAt: time.Now(),
	}
}
