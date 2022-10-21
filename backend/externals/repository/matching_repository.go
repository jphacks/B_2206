package repository

import (
	"context"
	"database/sql"
	"github.com/jphacks/B_2206/api/drivers/db"
	"fmt"
)


type matchingRepository struct {
	client db.Client
}

type MatchingRepository interface {
	Create(context.Context, string, string, string) error
	Destroy(context.Context, string) error
	FindMatchingByAccessToken(context.Context, string) *sql.Row
}

func NewMatchingRepository(client db.Client) MatchingRepository {
	return &matchingRepository{client}
}

// 作成
func (r *matchingRepository) Create(c context.Context, authID string, userID string, accessToken string) error {
	query := "insert into matching (auth_id, user_id, access_token) values (" + authID + ", " + userID + ", '" + accessToken + "')"
	_, err := r.client.DB().ExecContext(c, query)
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return nil
}

// 削除
func (r *matchingRepository) Destroy(c context.Context, accessToken string) error {
	// access tokenで該当のmatchingを削除
	query := "delete from matching where access_token = '" + accessToken + "'"
	_, err := r.client.DB().ExecContext(c, query)
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[36m%s\n", query)
	return nil
}

// アクセストークンからセッションを取得
func (r *matchingRepository) FindMatchingByAccessToken(c context.Context, accessToken string) *sql.Row {
	query := "select * from matching where access_token = '" + accessToken + "'"
	row := r.client.DB().QueryRowContext(c, query)
	fmt.Printf("\x1b[36m%s\n", query)
	return row
}
