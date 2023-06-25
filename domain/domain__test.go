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
			prepare  func()
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
				prepare: func() {
					// delete data before create
					farm := entity.Farm{}
					dbgorm.Where("id = ?", "integtest").Delete(&farm)
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
				prepare: func() {},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
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
			prepare  func()
		}{
			{
				testID:   1,
				testDesc: "Success get farm by id",
				testType: "P",
				farmID:   "integ-test",
				prepare: func() {
					// insert data before get
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)
				},
			},
			{
				testID:   2,
				testDesc: "Success get farm by id, not found",
				testType: "P",
				farmID:   "invalid-id",
				prepare:  func() {},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
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
			prepare  func()
		}{
			{
				testID:   1,
				testDesc: "Success get farm",
				testType: "P",
				prepare: func() {
					// insert data before get
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)
				},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
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

func TestUpdateFarm(t *testing.T) {
	Convey("TestUpdateFarm", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			in       struct {
				farmID  string
				payload entity.UpdateFarmRequest
			}
			prepare func()
		}{
			{
				testID:   1,
				testDesc: "Success update farm",
				testType: "P",
				in: struct {
					farmID  string
					payload entity.UpdateFarmRequest
				}{
					farmID: "integ-test",
					payload: entity.UpdateFarmRequest{
						Name:        "test-update",
						Description: "test-update",
					},
				},
				prepare: func() {
					// insert data before create
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)
				},
			},
			{
				testID:   2,
				testDesc: "Success upsert farm",
				testType: "P",
				in: struct {
					farmID  string
					payload entity.UpdateFarmRequest
				}{
					farmID: "integ-test",
					payload: entity.UpdateFarmRequest{
						Name:        "test-update",
						Description: "test-update",
					},
				},
				prepare: func() {
					// delete data before update
					farm := entity.Farm{}
					dbgorm.Where("id = ?", "integ-test").Delete(&farm)
				},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.UpdateFarm(tc.in.farmID, tc.in.payload)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestDeleteFarmByID(t *testing.T) {
	Convey("TestDeleteFarmByID", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			in       struct {
				farmID string
			}
			prepare func()
		}{
			{
				testID:   1,
				testDesc: "Success delete farm",
				testType: "P",
				in: struct {
					farmID string
				}{
					farmID: "integ-test",
				},
				prepare: func() {
					// insert data before create
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)
				},
			},
			{
				testID:   2,
				testDesc: "failed delete farm",
				testType: "N",
				in: struct {
					farmID string
				}{
					farmID: "invalid",
				},
				prepare: func() {
				},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			err := dom.DeleteFarmByID(tc.in.farmID)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}
