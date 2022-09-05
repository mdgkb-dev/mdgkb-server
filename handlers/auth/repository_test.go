package auth

import (
	"context"
	"log"
	"mdgkb/mdgkb-server/handlers/auth/mocks"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/tests"
	"os"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pro-assistance/pro-assister/config"
	"github.com/pro-assistance/pro-assister/helper"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "auth repository")
}

var _ = Describe("AuthPathPermissions", func() {
	var err error
	c, err := config.LoadTestConfig()
	if err != nil {
		log.Fatal(err)
	}
	h := helper.NewHelper(*c)
	repositoty := NewRepository(h)

	BeforeEach(func() {
		fixture := tests.GetFixtures(repositoty.db(), (*models.PathPermission)(nil), (*models.PathPermissionRole)(nil))
		err := fixture.Load(context.Background(), os.DirFS("./mocks"), "fixtures.yaml")
		Expect(err).To(BeNil())
	})

	Describe("Show all PathPermissions", func() {
		var items models.PathPermissions
		BeforeEach(func() {
			items, err = repositoty.getAllPathPermissions()
			Expect(err).To(BeNil())
		})

		It("Returns path permissions", func() {
			Expect(items).To(HaveLen(2))
		})
	})
	//
	Describe("Delete permission", func() {
		var items models.PathPermissions
		BeforeEach(func() {
			err = repositoty.deleteManyPathPermissions([]uuid.UUID{tests.CreateUUID.UUID})
			Expect(err).To(BeNil())
			items, err = repositoty.getAllPathPermissions()
			Expect(err).To(BeNil())
		})

		It("Returns only one permission", func() {
			Expect(items).To(HaveLen(1))
		})
	})

	It("Creates new PathPermission", func() {
		err := repositoty.upsertManyPathPermissions(models.PathPermissions{&mocks.MockFull})
		Expect(err).To(BeNil())
	})
})
