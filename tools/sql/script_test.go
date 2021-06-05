package psql

import (
	"context"
	"log"
	"testing"
)

type NameHolder struct {
	Name string `db:"name"`
}

type IdHolder struct {
	Id int `db:"id"`
}

var repo *SqlRepository

func simpleExec(t *testing.T) {
	results, err := repo.NamedExecMany(`INSERT INTO "test" ("name") VALUES (:name);
        INSERT INTO "test" ("name") VALUES (:name);`, NameHolder{"Natasha"}, NameHolder{"Innokenty"})
	if err != nil {
		t.Fatal(err)
	}

	if len(results) < 2 {
		t.Fatal("expected 2 results, actually received ", len(results))
	}

	r1, err := results[0].RowsAffected()
	if err != nil {
		t.Fatal(err)
	}

	r2, err := results[1].RowsAffected()
	if err != nil {
		t.Fatal(err)
	}

	if r1 != 1 || r2 != 1 {
		t.Fatal("should be affected 2 rows")
	}
}

func execWithoutParams(t *testing.T) {
	_, err := repo.NamedExecMany(`INSERT INTO "test" ("name") VALUES ('Anatoly');
        INSERT INTO "test" ("name") VALUES ('Ivan');`)
	if err != nil {
		t.Fatal(err)
	}
}

func execWithEmptyLines(t *testing.T) {
	_, err := repo.NamedExecMany(`
        INSERT INTO "test" ("name") VALUES ('Masha');
        INSERT INTO "test" ("name") VALUES ('Anya');
        ;`)
	if err != nil {
		t.Fatal(err)
	}
}

func execTx(t *testing.T) {
	results, err := repo.NamedExecManyContext(context.Background(), `DELETE FROM "test" WHERE "id" = 1; DELETE FROM "test" WHERE "id" = :id;`, nil, IdHolder{3})
	if err != nil {
		t.Fatal(err)
	}

	if len(results) < 2 {
		t.Fatal("expected 2 results, actually received ", len(results))
	}

	r1, err := results[0].RowsAffected()
	if err != nil {
		t.Fatal(err)
	}

	r2, err := results[1].RowsAffected()
	if err != nil {
		t.Fatal(err)
	}

	if r1 != 1 || r2 != 1 {
		t.Fatal("should be affected 2 rows")
	}
}

func failureReferences(t *testing.T) {
	_, err := repo.NamedExecMany(`
		CREATE TABLE a (id uuid);
		CREATE TABLE b(id uuid, a uuid REFERENCES a(id));`)
	if err == nil {
		t.Fatal("expected failure")
	}
	log.Println(err)
}

func rollbackOnFailure(t *testing.T) {
	tx, err := repo.DB.Beginx()
	if err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec(`CREATE TABLE inTxTable (id serial primary key)`)
	if err != nil {
		t.Fatal(err)
	}

	var success bool
	success, _, err = repo.NamedExecTx(tx, `
		CREATE TABLE a (id uuid);
		CREATE TABLE b(id uuid, a uuid REFERENCES a(id));`)
	if err == nil {
		t.Fatal(err)
	}

	if success {
		t.Fatal("expected failure")
	}

	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}

	row := repo.DB.QueryRowx(`
	SELECT 'inTxTable' IN (
    	SELECT table_name
    	FROM information_schema.tables
    	WHERE table_schema = 'public') AS exists;`)

	exists := true
	err = row.Scan(&exists)
	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("expected rollback")
	}
}

func checkOneInstance(t *testing.T) {
	repo2 := createTestRepository(t)

	if *repo.DB != *repo2.DB {
		t.Fatal("few repos has different sqlx.DB instances")
	}
}

func createTestRepository(t *testing.T) *SqlRepository {
	config := &DatabaseConfig{
		URL:            "postgres://localhost:5432/postgres?user=postgres&password=postgres",
		MaxConnections: 15,
	}

	var err error
	repo, err = NewSqlRepository(*config)
	if err != nil {
		t.Fatal(err)
	}

	return repo
}

func TestSqlRepository(t *testing.T) {
	repo = createTestRepository(t)

	defer func() {
		_, err := repo.DB.Exec(`DROP TABLE "test"`)
		if err != nil {
			panic(err)
		}

		err = repo.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err := repo.DB.Exec(`CREATE TABLE "test"("id" serial primary key, "name" text)`)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Simple NamedExec", simpleExec)
	t.Run("NamedExec without params", execWithoutParams)
	t.Run("NamedExec with empty lines", execWithEmptyLines)
	t.Run("NamedExec with transaction", execTx)
	t.Run("NamedExec failure at one query", failureReferences)
	t.Run("Rollback all TX on failure", rollbackOnFailure)
	t.Run("Check union sqlx.DB instance", checkOneInstance)
}
