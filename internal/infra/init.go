package infra

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nilsonmart/we-exchange/internal/models"
	"github.com/nilsonmart/we-exchange/internal/repository"
)

func InitSQLite() {

	conn := ConnectionSQLite()

	repo := repository.NewSQLiteRepository(conn)

	if err := repo.MigrateAccount(); err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	if err := repo.MigrateActivity(); err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	if err := repo.MigrateSchema(); err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	for _, v := range mockActivity() {
		_, err := repo.CreateActivity(v)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func mockActivity() []models.Activity {
	list := []models.Activity{
		{ID: 1, OldDate: time.Now().Add(2), NewDate: time.Now().Add(5), Paid: false, Approved: false, UserID: getUserId(), CreationDate: time.Now(), CreationUserID: getUserId()},
		{ID: 2, OldDate: time.Now().Add(5), NewDate: time.Now().Add(10), Paid: true, Approved: true, UserID: 2, CreationDate: time.Now().Add(-2), CreationUserID: 2},
	}
	return list
}

func getUserId() int64 {
	var c *gin.Context
	userId, err := c.Cookie("userid")
	if err != nil {
		panic(err)
	}

	p, err := strconv.ParseInt(userId, 10, 0)
	if err != nil {
		return 0
	}

	return p

}
