package repositories_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/diazharizky/go-mongodb-with-tests/internal/app"
	"github.com/diazharizky/go-mongodb-with-tests/internal/models"
	"github.com/diazharizky/go-mongodb-with-tests/internal/repositories"
)

var _ = Describe("UserRepository Tests", func() {
	var repo app.IUserRepository

	now = getNow()
	userSeeds := []models.User{
		{
			ID:        primitive.NewObjectID(),
			Username:  "user_one",
			Email:     "user_one@mail.com",
			Age:       20,
			FullName:  "User One",
			CreatedAt: now,
		},
		{
			ID:        primitive.NewObjectID(),
			Username:  "user_two",
			Email:     "user_two@mail.com",
			Age:       25,
			FullName:  "User Two",
			CreatedAt: now,
		},
	}

	BeforeEach(func() {
		repo = repositories.NewUserRepository(db)
	})

	AfterEach(func() {
		db.Drop(context.TODO())
	})

	Context("Create Tests", func() {
		When("Creating new user", func() {
			It("Should create new record in database", func() {
				ctx := context.Background()
				insertedID, gotErr := repo.Create(ctx, userSeeds[0])

				userID := userSeeds[0].ID
				Expect(gotErr).To(BeNil())
				Expect(*insertedID).To(Equal(userID.Hex()))

				gotUser, _ := repo.Get(ctx, userID.Hex())

				Expect(gotUser).To(Equal(&userSeeds[0]))
			})
		})
	})

	Context("List Tests", func() {
		When("Querying list of user from database", func() {
			It("Should return list of user", func() {
				ctx := context.Background()

				_, err = repo.Create(ctx, userSeeds[0])
				_, err = repo.Create(ctx, userSeeds[1])

				gotUsers, gotErr := repo.List(ctx)

				Expect(gotErr).To(BeNil())

				for i, gu := range gotUsers {
					Expect(gu).To(Equal(userSeeds[i]))
				}
			})
		})
	})

	Context("Get Test", func() {
		When("Querying user from database", func() {
			It("Should return user", func() {
				ctx := context.Background()

				repo.Create(ctx, userSeeds[0])

				gotUser, gotErr := repo.Get(ctx, userSeeds[0].ID.Hex())

				Expect(gotErr).To(BeNil())
				Expect(gotUser).To(Equal(&userSeeds[0]))
			})
		})
	})

	Context("Update Tests", func() {
		When("Updating an existing record", func() {
			It("Should update the record", func() {
				ctx := context.Background()

				repo.Create(ctx, userSeeds[0])

				now = getNow()
				update := models.User{
					ID:        userSeeds[0].ID,
					Username:  userSeeds[0].Username,
					Email:     "new_user_one@mail.com",
					Age:       30,
					FullName:  userSeeds[0].FullName,
					CreatedAt: userSeeds[0].CreatedAt,
					UpdatedAt: &now,
				}

				userID := userSeeds[0].ID
				gotErr := repo.Update(ctx, userID.Hex(), update)

				Expect(gotErr).To(BeNil())

				updatedUser, _ := repo.Get(ctx, userID.Hex())

				Expect(updatedUser).To(Equal(&update))
			})
		})
	})
})
