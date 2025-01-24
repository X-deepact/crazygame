package repositories

import (
	"strconv"
	"testing"
	"time"

	"crazygames.io/entities"
	"github.com/stretchr/testify/assert"
)

func createUser(username, email, password, role string) (*entities.User, error) {
	user := &entities.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
	err := userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Test_CreateUser(t *testing.T) {
	db.Exec("DELETE FROM users")

	t.Run("create user without username should fail", func(t *testing.T) {
		user := &entities.User{
			Email:    "test1@gmail.com",
			Password: "password",
			Role:     "player",
		}

		err := userRepository.Create(user)

		assert.Error(t, err, "expected error when creating user without a username")
	})

	t.Run("create user without email address should fail", func(t *testing.T) {
		user := &entities.User{
			Username: "hello",
			Password: "password",
			Role:     "player",
		}

		err := userRepository.Create(user)

		assert.Error(t, err, "expected error when creating user without email address")
	})

	t.Run("create user without password should fail", func(t *testing.T) {
		user := &entities.User{
			Username: "hello",
			Email:    "test1@gmail.com",
			Role:     "player",
		}

		err := userRepository.Create(user)

		assert.Error(t, err, "expected error when creating user without password")
	})

	t.Run("create user without role should set default value as player", func(t *testing.T) {
		user := &entities.User{
			Username: "testuser1",
			Email:    "test1@gmail.com",
			Password: "password",
		}

		err := userRepository.Create(user)

		assert.Equal(t, err, nil)
		assert.Equal(t, user.Role, "player")
	})

	t.Run("create user with all required fields should succeed", func(t *testing.T) {
		user, err := createUser("testuser2", "test2@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		assert.Equal(t, user.Username, "testuser2")
		assert.Equal(t, user.Email, "test2@gmail.com")
	})
}

func Test_GetUserBy(t *testing.T) {
	db.Exec("DELETE FROM users")

	t.Run("get user by valid ID should succeed", func(t *testing.T) {
		user, err := createUser("testuser3", "test3@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		retrievedUser, err := userRepository.GetByID(user.ID)

		assert.NoError(t, err, "failed to fetch user by ID")
		assert.Equal(t, user.Username, retrievedUser.Username)
		assert.Equal(t, user.Email, retrievedUser.Email)
	})

	t.Run("get user by invalid ID should fail", func(t *testing.T) {
		_, err := createUser("testuser4", "test4@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		_, err = userRepository.GetByID(uint(100_000))

		assert.Error(t, err, "failed to fetch user by ID")
	})

	t.Run("get user by valid Username should succeed", func(t *testing.T) {
		user, err := createUser("testuser5", "test5@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		retrievedUser, err := userRepository.GetByUsername(user.Username)

		assert.NoError(t, err, "failed to fetch user by Username")
		assert.Equal(t, user.Username, retrievedUser.Username)
		assert.Equal(t, user.Email, retrievedUser.Email)
	})

	t.Run("get user by invalid Username should fail", func(t *testing.T) {
		_, err := createUser("testuser6", "test6@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")
		randomUsername := "good" + time.Now().String()

		_, err = userRepository.GetByUsername(randomUsername)

		assert.Error(t, err, "failed to fetch user by randomUsername")
	})
}

func Test_GetAllUsers(t *testing.T) {
	db.Exec("DELETE FROM users")

	var users = make([]*entities.User, 0)

	for i := 1; i <= 10; i++ {
		user, err := createUser("testlistuser"+strconv.Itoa(i), "testlistuser"+strconv.Itoa(i)+"@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")
		users = append(users, user)
	}

	assert.Equal(t, len(users), 10)
}

func Test_UpdateUser(t *testing.T) {
	db.Exec("DELETE FROM users")

	user, err := createUser("testupdateuser1", "testupdateuser1@gmail.com", "password", "player")
	assert.NoError(t, err, "failed to create user for test")

	t.Run("update user's username to existing username should fail", func(t *testing.T) {
		existingUser, err := createUser("testupdateuser2", "testupdateuser2@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		existingUser.Username = user.Username
		_, err = userRepository.Update(existingUser)
		assert.Error(t, err, "expected error when updating user with existing username")
	})

	t.Run("update user's email address to existing email should fail", func(t *testing.T) {
		existingUser, err := createUser("testupdateuser3", "testupdateuser3@gmail.com", "password", "player")
		assert.NoError(t, err, "failed to create user for test")

		existingUser.Email = user.Email
		_, err = userRepository.Update(existingUser)
		assert.Error(t, err, "expected error when updating user with existing email")
	})

	t.Run("update user succeed", func(t *testing.T) {
		user.Email = "something@gmail.com"
		user, err = userRepository.Update(user)
		assert.NoError(t, err, "failed to update user for test")
		assert.Equal(t, user.Email, "something@gmail.com")
	})
}

func Test_DeleteUser(t *testing.T) {
	db.Exec("DELETE FROM users")

	user, err := createUser("testdeleteuser1", "testdeleteuser1@gmail.com", "password", "player")
	assert.NoError(t, err, "failed to create user for test")

	t.Run("delete user by id should succeed", func(t *testing.T) {
		err := userRepository.Delete(user.ID)
		assert.NoError(t, err, "failed to delete user")
	})
}

func Test_DeleteNonExistingUser(t *testing.T) {
	db.Exec("DELETE FROM users")

	t.Run("delete non-existing user should fail", func(t *testing.T) {
		err := userRepository.Delete(99999) // Non-existing user ID
		assert.NoError(t, err, "deleting non-existing user should not return an error")
	})
}

func Test_GetByEmail(t *testing.T) {
	db.Exec("DELETE FROM users")

	user, err := createUser("testemailuser", "testemailuser@gmail.com", "password", "player")
	assert.NoError(t, err, "failed to create user for test")

	t.Run("get user by valid email should succeed", func(t *testing.T) {
		retrievedUser, err := userRepository.GetByEmail(user.Email)
		assert.NoError(t, err, "failed to fetch user by email")
		assert.Equal(t, user.Email, retrievedUser.Email)
		assert.Equal(t, user.Username, retrievedUser.Username)
	})

	t.Run("get user by invalid email should fail", func(t *testing.T) {
		_, err := userRepository.GetByEmail("nonexistent@gmail.com")
		assert.Error(t, err, "expected error when fetching user by invalid email")
	})
}

func Test_GetAllUsersEdgeCases(t *testing.T) {
	db.Exec("DELETE FROM users")

	t.Run("get all users when no users exist should return empty list", func(t *testing.T) {
		users, err := userRepository.GetAll(1, 10)
		assert.NoError(t, err, "failed to fetch users")
		assert.Equal(t, len(users), 0, "expected no users in the database")
	})

	// t.Run("get all users with invalid page/limit should return empty list", func(t *testing.T) {
	// 	// Create some users
	// 	for i := 1; i <= 5; i++ {
	// 		_, err := createUser("testuser"+strconv.Itoa(i), "testuser"+strconv.Itoa(i)+"@gmail.com", "password", "player")
	// 		assert.NoError(t, err, "failed to create user for test")
	// 	}

	// 	users, err := userRepository.GetAll(0, -1) // Invalid page/limit
	// 	assert.NoError(t, err, "failed to fetch users with invalid page/limit")
	// 	assert.Equal(t, len(users), 0, "expected no users with invalid page/limit")
	// })
}

func Test_OauthCreate(t *testing.T) {
	db.Exec("DELETE FROM users")

	t.Run("oauth create user should succeed", func(t *testing.T) {
		user := &entities.User{
			Username: "oauthuser",
			Email:    "oauthuser@gmail.com",
			Password: "password",
			Role:     "player",
		}

		err := userRepository.OauthCreate(user)
		assert.NoError(t, err, "failed to create user using OauthCreate")

		// Verify the user was created
		retrievedUser, err := userRepository.GetByEmail(user.Email)
		assert.NoError(t, err, "failed to fetch user after OauthCreate")
		assert.Equal(t, user.Username, retrievedUser.Username)
		assert.Equal(t, user.Email, retrievedUser.Email)
	})
}

func Test_UpdatePassword(t *testing.T) {
	db.Exec("DELETE FROM users")

	user, err := createUser("testpassworduser", "testpassworduser@gmail.com", "password", "player")
	assert.NoError(t, err, "failed to create user for test")

	t.Run("update password for existing user should succeed", func(t *testing.T) {
		newPassword := "newhashedpassword"
		err := userRepository.UpdatePassword(user.Email, newPassword)
		assert.NoError(t, err, "failed to update password")

		// Verify the password was updated
		updatedUser, err := userRepository.GetByEmail(user.Email)
		assert.NoError(t, err, "failed to fetch user after password update")
		assert.Equal(t, updatedUser.Password, newPassword)
	})
}
