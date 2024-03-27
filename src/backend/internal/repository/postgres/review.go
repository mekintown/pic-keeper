package postgres

import (
	"context"

	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReviewDB struct {
	*BaseDB[model.Review]
}

func NewReviewDB(db *bun.DB) *ReviewDB {
	type T = model.Review

	return &ReviewDB{
		BaseDB: NewBaseDB[T](db),
	}
}

func (p *ReviewDB) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*model.Review, error) {
	var reviews []*model.Review
	if err := p.db.NewSelect().Model(&reviews).Where("customer_id = ?", userId).Scan(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (p *ReviewDB) FindByGalleryId(ctx context.Context, galleryId uuid.UUID) ([]*model.Review, error) {
	var booking model.Booking
	var reviews []*model.Review

	// not sure
	// get all bookings that belong to this galleryId
	allBookingId := p.db.NewSelect().Model(&booking).Where("gallery_id = ?", galleryId).Column("id")

	// get the review that belongs to each of allBookingId
	if err := p.db.NewSelect().Model(&reviews).Where("booking_id IN (?)", allBookingId).Scan(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (p *ReviewDB) FindByPhotographerId(ctx context.Context, photographerId uuid.UUID) ([]*model.Review, error) {
	var gallery model.Gallery
	var booking model.Booking
	var reviews []*model.Review

	// not sure
	// get all galleries that belong to this photographerId
	allGalleryId := p.db.NewSelect().Model(&gallery).Where("photographer_id = ?", photographerId).Column("id")

	// get all bookings that belong to each of allGalleryId
	allBookingId := p.db.NewSelect().Model(&booking).Where("gallery_id IN (?)", allGalleryId).Column("id")

	// get the review that belongs to each of allBookingId
	if err := p.db.NewSelect().Model(&reviews).Where("booking_id IN (?)", allBookingId).Scan(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewDB) CheckExistenceByGalleryId(ctx context.Context, galleryId uuid.UUID) (bool, error) {
	var review model.Review
	exist, err := r.db.NewSelect().Model(&review).Where("gallery_id = ?", galleryId).Exists(ctx)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *ReviewDB) SumAndCountRatingByGalleryId(ctx context.Context, galleryId uuid.UUID) (int, int, error) {
	var review model.Review
	var sum int
	count, err := r.db.NewSelect().Model(&review).Where("gallery_id = ?", galleryId).ColumnExpr("SUM(rating)").ScanAndCount(ctx, &sum)
	if err != nil {
		return 0, 0, err
	}

	return sum, count, nil
}