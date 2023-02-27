package database

import (
	"bufio"
	"github.com/valentinRog/chess-puzzles-api/config"
	"github.com/valentinRog/chess-puzzles-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var DB *gorm.DB

var NPuzzles uint

func Connect() {
	dsn := config.Config("DB_USER") + ":" + config.Config("MYSQL_ROOT_PASSWORD") + "@tcp(" + config.Config("DB_HOST") + ")/" + config.Config("MYSQL_DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func InitTables() {
	DB.Migrator().DropTable(&models.Puzzle{})
	DB.AutoMigrate(&models.Puzzle{})
}

func Populate() {
	f, err := os.Open(config.Config("CSV_FILENAME"))
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	buff := []models.Puzzle{}
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		rating, _ := strconv.Atoi(arr[3])
		var nPieces uint8
		for _, c := range strings.Split(arr[1], " ")[0] {
			if unicode.IsLetter(c) {
				nPieces++
			}
		}
		puzzle := models.Puzzle{
			FEN:     arr[1],
			Moves:   arr[2],
			Rating:  rating,
			NPieces: nPieces,
		}
		buff = append(buff, puzzle)
	}
	for i := range buff {
		j := rand.Intn(i + 1)
		buff[i], buff[j] = buff[j], buff[i]
	}
	arr := []models.Puzzle{}
	for _, puzzle := range buff {
		arr = append(arr, puzzle)
		NPuzzles++
		if len(arr) == 1000 {
			DB.Create(&arr)
			arr = []models.Puzzle{}
		}
	}
	DB.Create(&arr)
	DB.Exec("CREATE INDEX rating ON puzzles (rating)")
	DB.Exec("CREATE INDEX n_pieces ON puzzles (n_pieces)")
}
