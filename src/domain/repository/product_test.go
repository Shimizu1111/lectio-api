package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Shimizu1111/lectio/design/entity"
	"github.com/Shimizu1111/lectio/src/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Mockを作成
func InitDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

// 全件取得が正常に行われることを確認
func TestCalcsrvcFindAllProduct(t *testing.T) {
	//MockDBを初期化する
	mockDB, mock, err := InitDBMock()

	if err != nil {
		t.Fatal(err)
	}

	defer mockDB.Close()

	r := repository.Repository{DB: mockDB}

	expect := []entity.Book{
		{
			ID:               1,
			UserID:           1,
			Name:             "この素晴らしい世界に祝福を!",
			Author:           "暁なつめ1",
			Publisher:        "角川スニーカー文庫",
			Price:            700,
			RegistrationDate: "2021-01-01 10:00:00",
			BookStatus:       "favorite",
			IsDeleted:        false,
			CreatedAt:        "2022-07-19 14:19:01",
			UpdatedAt:        "2022-07-19 14:19:01",
		},
		{
			ID:               2,
			UserID:           2,
			Name:             "この素晴らしい世界に祝福を!",
			Author:           "暁なつめ2",
			Publisher:        "角川スニーカー文庫",
			Price:            700,
			RegistrationDate: "2022-07-19 23:19:28",
			BookStatus:       "favorite",
			IsDeleted:        false,
			CreatedAt:        "2022-07-19 14:19:28",
			UpdatedAt:        "2022-07-19 14:19:28",
		},
		{
			ID:               3,
			UserID:           3,
			Name:             "この素晴らしい世界に祝福を!",
			Author:           "暁なつめ3",
			Publisher:        "角川スニーカー文庫",
			Price:            700,
			RegistrationDate: "2022-07-19 23:42:24",
			BookStatus:       "favorite",
			IsDeleted:        false,
			CreatedAt:        "2022-07-19 14:42:24",
			UpdatedAt:        "2022-07-19 14:42:24",
		},
	}

	// Mockで使用するBooksテーブルのカラム
	columns := []string{
		"id",
		"user_id",
		"name",
		"author",
		"publisher",
		"price",
		"registration_date",
		"book_status",
		"is_deleted",
		"created_at",
		"updated_at",
	}

	// Mock設定
	// TODO:AddRowをループで生成できるようにする
	mock.ExpectQuery(regexp.QuoteMeta(`books`)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "1", "この素晴らしい世界に祝福を!", "暁なつめ1", "角川スニーカー文庫", "700", "2021-01-01 10:00:00", "favorite", "0", "2022-07-19 14:19:01", "2022-07-19 14:19:01").
			AddRow("2", "2", "この素晴らしい世界に祝福を!", "暁なつめ2", "角川スニーカー文庫", "700", "2022-07-19 23:19:28", "favorite", "0", "2022-07-19 14:19:28", "2022-07-19 14:19:28").
			AddRow("3", "3", "この素晴らしい世界に祝福を!", "暁なつめ3", "角川スニーカー文庫", "700", "2022-07-19 23:42:24", "favorite", "0", "2022-07-19 14:42:24", "2022-07-19 14:42:24"),
		)

	// 全件取得実行
	var books []*entity.Book
	res := r.FindAllProduct(books)

	// Mockのデータと期待値が一致していることを確認
	for i, v := range expect {

		// TODO:構造体をループで取得する
		if res[i].ID != v.ID {
			t.Errorf("取得結果不一致  BookID:%+v", res[i].ID)
		}
		if res[i].UserID != v.UserID {
			t.Errorf("取得結果不一致  UserID:%+v", res[i].UserID)
		}
		if res[i].Name != v.Name {
			t.Errorf("取得結果不一致  Name:%+v", res[i].Name)
		}
		if res[i].Author != v.Author {
			t.Errorf("取得結果不一致  Author:%+v", res[i].Author)
		}
		if res[i].Publisher != v.Publisher {
			t.Errorf("取得結果不一致  Publisher:%+v", res[i].Publisher)
		}
		if res[i].Price != v.Price {
			t.Errorf("取得結果不一致  Price:%+v", res[i].Price)
		}
		if res[i].RegistrationDate != v.RegistrationDate {
			t.Errorf("取得結果不一致  RegistrationDate:%+v", res[i].RegistrationDate)
		}
		if res[i].BookStatus != v.BookStatus {
			t.Errorf("取得結果不一致  BookStatus:%+v", res[i].BookStatus)
		}
		if res[i].IsDeleted != v.IsDeleted {
			t.Errorf("取得結果不一致  IsDeleted:%+v", res[i].IsDeleted)
		}
		if res[i].CreatedAt != v.CreatedAt {
			t.Errorf("取得結果不一致  CreatedAt:%+v", res[i].CreatedAt)
		}
		if res[i].UpdatedAt != v.UpdatedAt {
			t.Errorf("取得結果不一致  UpdatedAt:%+v", res[i].UpdatedAt)
		}
	}
}
