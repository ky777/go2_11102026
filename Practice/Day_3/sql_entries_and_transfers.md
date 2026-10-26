*Задание**  
Создать `CRUD`'ы для типов _Entry_ и _Transfer_ с помощью библиотеки `sqlc`  


Тип **Transfer**

- ListTransfers выдает список transfer для параметоров `from_account_id` или `to_account_id` с учетом `LIMIT` и `OFFSET`.  
  Пример содержимого WHERE:
  ```
    WHERE 
        from_account_id = $1 OR
        to_account_id = $2
  ```  
- UpdateTransfer изменяет поля: `amount` и `to_account_id`