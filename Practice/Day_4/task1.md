**Задание**  

Понять, почему тесты падают с ошибкой и сделать фикс(исправить её).

###Ошибка первая:

``` golang
var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	// соединение с бд
	var err error
    // Значение присваеивается локальной пременной. Глобальноая остается равной nil
	//testDB, err := pgxpool.New(context.Background(), dbSource) 
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can not connection to db", err)
	}
	//закрываем соединеие
	defer testDB.Close()
	//Вызваем конструктор создания экзампляра типа Queries
	testQueries = New(testDB)
	//Запускаем subtest(тесты) и итоговый код выполнения передаем в Exit()
	os.Exit(m.Run())
}
```
### Вторая ошибка:
_Ошибка 40001 (could not serialize access due to concurrent update) в PostgreSQL возникает, когда две или более транзакции одновременно пытаются изменить одни и те же данные, и из-за этого нарушается согласованность. PostgreSQL не может гарантировать сериализуемость (эквивалентность некоторому последовательному выполнению) и прерывает одну из транзакций, предлагая повторить попытку._
``` golang
// Step 3. Update account's balance
		// Меняем порядок вызова для пар аккаунтов
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = AddMoney(
				ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount,
			)
		} else {
			result.FromAccount, result.ToAccount, err = AddMoney(
				ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount,
			)
		}

```
