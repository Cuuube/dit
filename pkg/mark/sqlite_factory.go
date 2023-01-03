package mark

import (
	"os"
	"path"

	"github.com/Cuuube/dit/pkg/cli"
	"github.com/Cuuube/dit/pkg/ctrl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db             *gorm.DB
	sqliteFilePath = path.Join(ctrl.IgnoreErr(os.Executable()), "../mark_sqlite.db")
)

func init() {
}

func NewSQLiteGorm() *gorm.DB {
	checkSqliteFile()
	db, err := gorm.Open(sqlite.Open(sqliteFilePath))
	if err != nil {
		panic(err)
	}
	// 不打印gorm日志
	db.Logger = logger.Default.LogMode(logger.Silent)
	return db
}

func GetSQLiteGorm() *gorm.DB {
	if db == nil {
		db = NewSQLiteGorm()
	}
	return db
}

// 检查本地是否有sqlite文件
func checkSqliteFile() {
	_, err := os.Stat(sqliteFilePath)
	if err != nil {
		cli.Println("找不到db", sqliteFilePath)
		panic("")
	}
}
