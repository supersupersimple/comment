package api

import (
	"context"
	"fmt"
	"time"

	"github.com/supersupersimple/comment/ent"
	"golang.org/x/crypto/bcrypt"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConvertPublishedTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < time.Second {
		return "just now"
	}

	if diff < 2*time.Second {
		return "1 second ago"
	}

	if diff < 60*time.Second {
		return fmt.Sprintf("%d seconds ago", int(diff/time.Second))
	}

	if diff < 2*time.Minute {
		return "1 minute ago"
	}

	if diff < 60*time.Minute {
		return fmt.Sprintf("%d minutes ago", int(diff/time.Minute))
	}

	if diff < 2*time.Hour {
		return "1 hour ago"
	}

	if diff < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(diff/time.Hour))
	}

	if diff < 2*24*time.Hour {
		return "1 day ago"
	}

	if diff < 7*24*time.Hour {
		return fmt.Sprintf("%d days ago", int(diff/24/time.Hour))
	}

	if diff < 14*24*time.Hour {
		return "1 week ago"
	}

	if diff < 365*24*time.Hour {
		return fmt.Sprintf("%d weeks ago", int(diff/7/24/time.Hour))
	}

	if diff < 2*365*24*time.Hour {
		return "1 year ago"
	}

	return fmt.Sprintf("%d years ago", int(diff/365/24/time.Hour))
}
