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

	dbgorm.AutoMigrate(&entity.Farm{})
	dbgorm.AutoMigrate(&entity.Pond{})

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

func TestCreatePond(t *testing.T) {
	Convey("TestCreatePond", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entity.CreatePondRequest
			prepare  func()
		}{
			{
				testID:   1,
				testDesc: "Success create pond",
				testType: "P",
				payload: entity.CreatePondRequest{
					ID:          "integtest",
					FarmID:      "integ-test",
					Name:        "name test",
					Description: "test",
				},
				prepare: func() {
					// delete data before create
					pond := entity.Pond{}
					dbgorm.Where("id = ?", "integtest").Delete(&pond)

					// insert data farm before create pond
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
				testDesc: "fail create pond, invalid farm id",
				testType: "N",
				payload: entity.CreatePondRequest{
					ID:          "integtest",
					FarmID:      "invalidd",
					Name:        "name test",
					Description: "test",
				},
				prepare: func() {},
			},
			{
				testID:   3,
				testDesc: "fail create pond, pond already exist",
				testType: "N",
				payload: entity.CreatePondRequest{
					ID:          "integtest",
					FarmID:      "integtest",
					Name:        "name test",
					Description: "test",
				},
				prepare: func() {
					// insert data farm before create pond
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
			_, err := dom.CreatePond(tc.payload)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestGetPondByID(t *testing.T) {
	Convey("TestGetPondByID", t, FailureHalts, func() {
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
					// insert data before get
					pond := entity.Pond{
						ID:          "integ-test",
						FarmID:      "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&pond)
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
			_, err := dom.GetPondByID(tc.farmID)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestGetPond(t *testing.T) {
	Convey("TestGetPond", t, FailureHalts, func() {
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
					// insert data before get
					pond := entity.Pond{
						ID:          "integ-test",
						FarmID:      "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&pond)
				},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.GetPond()
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestUpdatePond(t *testing.T) {
	Convey("TestUpdatePond", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			in       struct {
				pondID  string
				payload entity.UpdatePondRequest
			}
			prepare func()
		}{
			{
				testID:   1,
				testDesc: "Success update pond",
				testType: "P",
				in: struct {
					pondID  string
					payload entity.UpdatePondRequest
				}{
					pondID: "integ-test",
					payload: entity.UpdatePondRequest{
						FarmID:      "integ-test",
						Name:        "test-update",
						Description: "test-update",
					},
				},
				prepare: func() {
					// insert data farm
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)
					// insert data pond
					pond := entity.Pond{
						ID:          "integ-test",
						FarmID:      "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&pond)
				},
			},
			{
				testID:   2,
				testDesc: "Success update pond, not exist, create",
				testType: "P",
				in: struct {
					pondID  string
					payload entity.UpdatePondRequest
				}{
					pondID: "integ-test",
					payload: entity.UpdatePondRequest{
						FarmID:      "integ-test",
						Name:        "test-update",
						Description: "test-update",
					},
				},
				prepare: func() {
					// insert farm data
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)

					// delete pond
					pond := entity.Pond{}
					dbgorm.Where("id = ?", "integ-test").Delete(&pond)
				},
			},
			{
				testID:   3,
				testDesc: "Failed update farm, farm not found",
				testType: "N",
				in: struct {
					pondID  string
					payload entity.UpdatePondRequest
				}{
					pondID: "integ-test",
					payload: entity.UpdatePondRequest{
						FarmID:      "invalidd",
						Name:        "test-update",
						Description: "test-update",
					},
				},
				prepare: func() {
					// insert data farm
					farm := entity.Farm{
						ID:          "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&farm)

					// insert data pond
					pond := entity.Pond{
						ID:          "integ-test",
						FarmID:      "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&pond)
				},
			},
		}

		for _, tc := range testCases {
			tc.prepare()
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.UpdatePond(tc.in.pondID, tc.in.payload)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestDeletePondByID(t *testing.T) {
	Convey("TestDeletePondByID", t, FailureHalts, func() {
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

					// insert data before create
					pond := entity.Pond{
						ID:          "integ-test",
						FarmID:      "integ-test",
						Name:        "integ-test",
						Description: "integ-test",
					}
					dbgorm.Create(&pond)
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
			err := dom.DeletePondByID(tc.in.farmID)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestUpsertAPIStatistic(t *testing.T) {
	Convey("TestUpsertAPIStatistic", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			in       struct {
				apiPath   string
				userAgent string
			}
		}{
			{
				testID:   1,
				testDesc: "Success upsert",
				testType: "P",
				in: struct {
					apiPath   string
					userAgent string
				}{
					apiPath:   "POST /a",
					userAgent: "agent-a",
				},
			},
		}

		for _, tc := range testCases {
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			err := dom.UpsertAPIStatistic(tc.in.apiPath, tc.in.userAgent)
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}

func TestGetAPIStatistic(t *testing.T) {
	Convey("TestGetAPIStatistic", t, FailureHalts, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
		}{
			{
				testID:   1,
				testDesc: "Success get",
				testType: "P",
			},
		}

		for _, tc := range testCases {
			t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
			_, err := dom.GetAPIStatistic()
			if tc.testType == "P" {
				So(err, ShouldBeNil)
			} else {
				So(err, ShouldNotBeNil)
			}
		}
	})
}
