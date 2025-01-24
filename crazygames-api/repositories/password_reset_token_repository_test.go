package repositories

import (
	"testing"
	"time"

	"crazygames.io/entities"
	"github.com/stretchr/testify/assert"
)

func createPasswordResetToken(email, token string, expiresAt time.Time, isUsed bool) (*entities.PasswordResetToken, error) {
	resetToken := &entities.PasswordResetToken{
		Email:     email,
		Token:     token,
		ExpiredAt: expiresAt,
		IsUsed:    isUsed,
	}
	err := passwordResetTokenRepository.Create(resetToken)
	if err != nil {
		return nil, err
	}
	return resetToken, nil
}

func Test_CreatePasswordResetToken(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	t.Run("create password reset token should succeed", func(t *testing.T) {
		token := &entities.PasswordResetToken{
			Email:     "testuser@gmail.com",
			Token:     "testtoken",
			ExpiredAt: time.Now().Add(1 * time.Hour),
			IsUsed:    false,
		}

		err := passwordResetTokenRepository.Create(token)
		assert.NoError(t, err, "failed to create password reset token")
		assert.NotZero(t, token.ID, "expected token ID to be set")
	})
}

func Test_UpdatePasswordResetToken(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	token, err := createPasswordResetToken("testuser@gmail.com", "testtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("update password reset token should succeed", func(t *testing.T) {
		token.Token = "updatedtoken"
		updatedToken, err := passwordResetTokenRepository.Update(token)
		assert.NoError(t, err, "failed to update password reset token")
		assert.Equal(t, "updatedtoken", updatedToken.Token, "expected token to be updated")
	})
}

func Test_GetPasswordResetTokenByEmail(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	token, err := createPasswordResetToken("testuser@gmail.com", "testtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("get password reset token by valid email should succeed", func(t *testing.T) {
		retrievedToken, err := passwordResetTokenRepository.GetByEmail("testuser@gmail.com")
		assert.NoError(t, err, "failed to get password reset token by email")
		assert.Equal(t, token.Token, retrievedToken.Token, "expected token to match")
	})

	t.Run("get password reset token by invalid email should fail", func(t *testing.T) {
		_, err := passwordResetTokenRepository.GetByEmail("nonexistent@gmail.com")
		assert.Error(t, err, "expected error when getting token by invalid email")
	})
}

func Test_GetPasswordResetTokenByToken(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	token, err := createPasswordResetToken("testuser@gmail.com", "testtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("get password reset token by valid token should succeed", func(t *testing.T) {
		retrievedToken, err := passwordResetTokenRepository.GetByToken("testtoken")
		assert.NoError(t, err, "failed to get password reset token by token")
		assert.Equal(t, token.Email, retrievedToken.Email, "expected email to match")
	})

	t.Run("get password reset token by invalid token should fail", func(t *testing.T) {
		_, err := passwordResetTokenRepository.GetByToken("invalidtoken")
		assert.Error(t, err, "expected error when getting token by invalid token")
	})
}

func Test_SetResetToken(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	_, err := createPasswordResetToken("testuser@gmail.com", "oldtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("set reset token should succeed", func(t *testing.T) {
		newToken := "newtoken"
		newExpiry := time.Now().Add(2 * time.Hour)

		err := passwordResetTokenRepository.SetResetToken("testuser@gmail.com", newToken, newExpiry)
		assert.NoError(t, err, "failed to set reset token")

		// Verify the token was updated
		retrievedToken, err := passwordResetTokenRepository.GetByEmail("testuser@gmail.com")
		assert.NoError(t, err, "failed to get updated token")
		assert.Equal(t, newToken, retrievedToken.Token, "expected token to be updated")
		assert.Equal(t, newExpiry.Unix(), retrievedToken.ExpiredAt.Unix(), "expected expiry time to be updated")
	})
}

func Test_GetUserByResetToken(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	token, err := createPasswordResetToken("testuser@gmail.com", "validtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("get user by valid reset token should succeed", func(t *testing.T) {
		retrievedToken, err := passwordResetTokenRepository.GetUserByResetToken("validtoken")
		assert.NoError(t, err, "failed to get user by reset token")
		assert.Equal(t, token.Email, retrievedToken.Email, "expected email to match")
	})

	t.Run("get user by expired reset token should fail", func(t *testing.T) {
		expiredToken, err := createPasswordResetToken("expireduser@gmail.com", "expiredtoken", time.Now().Add(-1*time.Hour), false)
		assert.NoError(t, err, "failed to create expired token for test")

		_, err = passwordResetTokenRepository.GetUserByResetToken(expiredToken.Token)
		assert.Error(t, err, "expected error when getting user by expired token")
	})

	t.Run("get user by used reset token should fail", func(t *testing.T) {
		usedToken, err := createPasswordResetToken("useduser@gmail.com", "usedtoken", time.Now().Add(1*time.Hour), true)
		assert.NoError(t, err, "failed to create used token for test")

		_, err = passwordResetTokenRepository.GetUserByResetToken(usedToken.Token)
		assert.Error(t, err, "expected error when getting user by used token")
	})
}

func Test_MarkTokenAsUsed(t *testing.T) {
	db.Exec("DELETE FROM password_reset_tokens")

	token, err := createPasswordResetToken("testuser@gmail.com", "testtoken", time.Now().Add(1*time.Hour), false)
	assert.NoError(t, err, "failed to create password reset token for test")

	t.Run("mark token as used should succeed", func(t *testing.T) {
		err := passwordResetTokenRepository.MarkTokenAsUsed(token.ID)
		assert.NoError(t, err, "failed to mark token as used")

		// Verify the token is marked as used
		retrievedToken, err := passwordResetTokenRepository.GetByToken("testtoken")
		assert.NoError(t, err, "failed to get token after marking as used")
		assert.True(t, retrievedToken.IsUsed, "expected token to be marked as used")
	})
}
