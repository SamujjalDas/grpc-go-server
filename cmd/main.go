package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"

	dbmigration "github.com/SamujjalDas/grpc-go-server/db"
	mydb "github.com/SamujjalDas/grpc-go-server/internal/adapter/database"
	mygrpc "github.com/SamujjalDas/grpc-go-server/internal/adapter/grpc"
	app "github.com/SamujjalDas/grpc-go-server/internal/application"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(log.Writer())

	sqlDb, err := sql.Open("pgx", "postgres://postgres:test@localhost:5432/postgres?sslmode=disable")
	//sqlDb, err := sql.Open("pgx", "postgres://snaply_db_user:PAx8ZBorw7cpQN8p6KpijzeJJNe79Soc@dpg-co4iglsf7o1s738sm8tg-a.oregon-postgres.render.com/snaply_db?sslmode=disable")

	if err != nil {
		log.Fatalln("Cant connect to database : ", err)
	}

	dbmigration.Migrate(sqlDb)

	databaseAdapter, err := mydb.NewDatabaseAdapter(sqlDb)

	if err != nil {
		log.Fatalln("Cant create to database adapter : ", err)
	}

	//runDummyOrm(databaseAdapter)

	hs := &app.HelloService{}
	bs := app.NewBankService(databaseAdapter)

	grpcAdapter := mygrpc.NewGrpcAdapter(hs, bs, 9090)
	grpcAdapter.Run()
}

func runDummyOrm(da *mydb.DatabaseAdapter) {
	now := time.Now()

	uuid, _ := da.Save(
		&mydb.DummyOrm{
			UserId:    uuid.New(),
			UserName:  "Samujjal Das New Entry" + time.Now().Format("15:04:05"),
			CreatedAt: now,
			UpdatedAt: now,
		},
	)

	res, _ := da.GetByUuid(&uuid)

	log.Println("res : ", res)
}
