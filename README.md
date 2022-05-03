# Balance Api 
### Задача: 
https://github.com/avito-tech/autumn-2021-intern-assignment

---

### Цель:
Написать api для взаимодействия с балансом пользователя

---

### Endpoints: 
<strong>GET</strong> /balance:id - получить текущий баланс пользователя по id
</Br>
<strong>POST</strong> /balance - добавить баланс на счет пользователя, в теле id, amount
</Br>
<strong>PUT</strong> /balance:id&:amount - начислить средства на баланс пользователя по id
</Br>
<strong>DELETE</strong> /balance/:id/:amount - списать часть средств с баланса
</Br>
