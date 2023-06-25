package domain_test

import (
	"log"
	"os"
	"testing"

	"github.com/alvinatthariq/farmsvc-go/domain"
	"github.com/alvinatthariq/farmsvc-go/entity"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbgorm      *gorm.DB
	router      *mux.Router
	redisClient *redis.Client
	err         error

	dom domain.DomainItf
)

func TestMain(t *testing.M) {
	dbgorm, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3307)/farm_db?parseTime=true"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})

	dom = domain.Init(
		dbgorm,
		redisClient,
	)

	exitVal := t.Run()

	os.Exit(exitVal)
}

func TestCreateFarm(t *testing.T) {
	Convey("TestCreateFarm", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entity.CreateFarmRequest
		}{
			{
				testID:   1,
				testDesc: "Success create farm",
				testType: "P",
				payload: entity.CreateFarmRequest{
					ID:          "integtest",
					Name:        "name test",
					Description: "test",
				},
			},
			{
				testID:   2,
				testDesc: "fail create farm, already exist",
				testType: "N",
				payload: entity.CreateFarmRequest{
					ID:          "integtest",
					Name:        "name test",
					Description: "test",
				},
			},
		}

		for _, tc := range testCases {
			// delete data before create
			farm := entity.Farm{}
			if tc.testType == "P" {
				dbgorm.Where("id = ?", tc.payload.ID).Delete(&farm)
			}

			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.CreateFarm(tc.payload)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestGetFarmByID(t *testing.T) {
	Convey("TestGetFarmByID", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			farmID   string
		}{
			{
				testID:   1,
				testDesc: "Success get farm by id",
				testType: "P",
				farmID:   "integ-test",
			},
			{
				testID:   2,
				testDesc: "Success get farm by id, not found",
				testType: "P",
				farmID:   "invalid-id",
			},
		}

		for _, tc := range testCases {
			// insert data before create
			farm := entity.Farm{
				ID:          "integ-test",
				Name:        "integ-test",
				Description: "integ-test",
			}
			if tc.testType == "P" {
				dbgorm.Create(&farm)
			}

			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.GetFarmByID(tc.farmID)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestGetFarm(t *testing.T) {
	Convey("TestGetFarm", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
		}{
			{
				testID:   1,
				testDesc: "Success get farm",
				testType: "P",
			},
		}

		for _, tc := range testCases {
			// insert data before create
			farm := entity.Farm{
				ID:          "integ-test",
				Name:        "integ-test",
				Description: "integ-test",
			}
			if tc.testType == "P" {
				dbgorm.Create(&farm)
			}

			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.GetFarm()
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}
