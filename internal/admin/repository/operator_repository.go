package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	commonModel "goadmin/internal/common/model"
)

// OperatorRepository defines data access methods for operators.
type OperatorRepository interface {
	GetByAccount(ctx context.Context, account string) (*commonModel.Operator, error)
	UpdateLoginMeta(ctx context.Context, id int, lastIP string, when time.Time) error
}

type operatorRepository struct {
	db *gorm.DB
}

// NewOperatorRepository constructs an OperatorRepository backed by GORM.
func NewOperatorRepository(db *gorm.DB) OperatorRepository {
	return &operatorRepository{db: db}
}

func (r *operatorRepository) GetByAccount(ctx context.Context, account string) (*commonModel.Operator, error) {
	var operator commonModel.Operator
	err := r.db.WithContext(ctx).
		Where("is_del = 0 AND (email = ? OR phone = ? OR nickname = ?)", account, account, account).
		First(&operator).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &operator, nil
}

func (r *operatorRepository) UpdateLoginMeta(ctx context.Context, id int, lastIP string, when time.Time) error {
	return r.db.WithContext(ctx).
		Model(&commonModel.Operator{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"last_login_time": when,
			"last_ip":         lastIP,
			"login_num":       gorm.Expr("COALESCE(login_num, 0) + 1"),
		}).Error
}
