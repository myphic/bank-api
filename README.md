# Balance Api 
### Задача: 
https://github.com/avito-tech/autumn-2021-intern-assignment

---

### Цель:
Написать api для взаимодействия с балансом пользователя

---

### Технологии:
<ul>
    <li>Golang</li>
    <li>Fiber</li>
    <li>Docker</li>
    <li>Redis</li>
    <li>Postgres</li>
    <li>Gorm</li>
    <li>Viper</li>
    <li>Logrus</li>
</ul>

---

### Endpoints: 
<table>
<tr>
<td>
<strong>GET</strong>
</td>
<td>
/api/balance/ - получить текущий баланс пользователя по id
</td>
<td>
Параметры запроса: id пользователя, currency валюта опционально
</td>
</tr>
<tr>
<td>
<strong>POST</strong>
</td>
<td>
/api/deposit/ - добавить баланс на счет пользователя
</td>
<td>
Тело запроса: id пользователя, amount сумма начислений
</td>
</tr>
<tr>
<td>
<strong>POST</strong>
</td>
<td>
/api/transfer/ - перевод средств с одного кошелька на другой
</td>
<td>
Тело запроса: amount сумма, fromUserId откуда перевести, toUserId куда перевести
</td>
</tr>
<tr>
<td>
<strong>DELETE</strong>
</td>
<td>
/api/decrease/ - списать часть средств с баланса
</td>
<td>
Параметры запроса: id пользователя, amount сумма списывания
</td>
</tr>
</table>
