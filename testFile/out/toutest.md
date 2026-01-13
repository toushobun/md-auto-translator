# 📘 第2章：SQL SELECT文を使用したデータの取得

（SQL 基本查询）

------

## 🧩 一、SELECT语句的基本结构

SQL（Structured Query Language）是一种**声明式语言**。
 我们只需告诉数据库“想要什么”，而不用告诉它“怎么做”。

语法结构如下：

```sql
SELECT 列名1, 列名2, ...
FROM 表名;
```

👉 **关键点**：

- **SELECT**：要取哪些列
- **FROM**：从哪张表取
- **分号 ;**：语句结束（必须有）

------

### 🧠 例1：查看所有员工

```sql
SELECT * FROM EMPLOYEES;
```

> `*` 表示“所有列”。

输出示例（例）：

| EMPLOYEE_ID | FIRST_NAME | LAST_NAME | JOB_TITLE  | DEPARTMENT_ID | MANAGER_ID | HIRE_DATE  | SALARY |
| ----------- | ---------- | --------- | ---------- | ------------- | ---------- | ---------- | ------ |
| 101         | 田中       | 太郎      | エンジニア | 10            | 100        | 2022-04-01 | 420000 |
| 102         | 鈴木       | 花子      | 営業担当   | 20            | 105        | 2023-01-15 | 350000 |

------

## 🧩 二、列别名（AS）

列别名可以让输出的标题更易读。

```sql
SELECT first_name AS 名前,
       last_name AS 苗字,
       salary AS 給与
FROM EMPLOYEES;
```

Oracle里 `AS` 可以省略：

```sql
SELECT first_name 名前 FROM EMPLOYEES;
```

------

## 🧩 三、连接运算符（||）

可以把多个列拼成一个字符串。

```sql
SELECT first_name || ' ' || last_name AS 氏名
FROM EMPLOYEES;
```

输出示例：

| 氏名      |
| --------- |
| 田中 太郎 |
| 鈴木 花子 |

------

## 🧩 四、DISTINCT 去重

如果想查看部门编号有哪些（去掉重复）：

```sql
SELECT DISTINCT department_id
FROM EMPLOYEES;
```

输出：

| DEPARTMENT_ID |
| ------------- |
| 10            |
| 20            |
| 30            |

------

## 🧩 五、文字常量（Literal）

除了表中的列，也可以直接写固定值。

```sql
SELECT 'Oracle SQL学習中' AS コメント,
       first_name, last_name
FROM EMPLOYEES;
```

输出示例：

| コメント         | FIRST_NAME | LAST_NAME |
| ---------------- | ---------- | --------- |
| Oracle SQL学習中 | 田中       | 太郎      |
| Oracle SQL学習中 | 鈴木       | 花子      |

------

## 🧩 六、表别名（Table Alias）

在多表查询或长表名时，用别名让语句更简洁：

```sql
SELECT e.first_name, e.last_name, d.department_name
FROM EMPLOYEES e, DEPARTMENTS d
WHERE e.department_id = d.department_id;
```

这里：

- `EMPLOYEES e` 给 EMPLOYEES 表起名 e
- `DEPARTMENTS d` 给 DEPARTMENTS 表起名 d

------

## 🧩 七、练习题（実習）

> 使用你定义的五张表来练习。

### ✅ 练习1：显示所有部门的名称和编号

```sql
SELECT department_id, department_name
FROM DEPARTMENTS;
```

### ✅ 练习2：显示员工的姓名与所在部门编号

```sql
SELECT first_name || ' ' || last_name AS 氏名,
       department_id
FROM EMPLOYEES;
```

### ✅ 练习3：显示项目表中的项目名称（去重）

```sql
SELECT DISTINCT project_name
FROM PROJECTS;
```

### ✅ 练习4：为薪资历史表的每条记录添加说明文字

```sql
SELECT record_id,
       employee_id,
       '給与履歴データ' AS コメント
FROM SALARY_HISTORY;
```

### ✅ 练习5（进阶）：显示员工姓名与对应上司ID（给表取别名）

```sql
SELECT e.first_name || ' ' || e.last_name AS 社員,
       e.manager_id AS 上司ID
FROM EMPLOYEES e;
```

------

## 🧩 八、扩展理解（Oracle特有）

Oracle 特有语法和特点：

| 特性                                                         | 说明 |
| ------------------------------------------------------------ | ---- |
| SQL 不区分大小写，但关键字通常大写。                         |      |
| 字符串用 `' '` 包裹，列名不用。                              |      |
| 每条语句必须以 `;` 结束。                                    |      |
| Oracle 默认使用 `DUAL` 表测试常量： 例：`SELECT SYSDATE FROM DUAL;`（获取系统日期） |      |

------

## 🧭 本章小结

| 知识点 | 关键语法                |
| ------ | ----------------------- |
| 取列   | `SELECT 列名 FROM 表名` |
| 全取   | `SELECT *`              |
| 别名   | `AS` 或直接写别名       |
| 拼接   | `                       |
| 去重   | `DISTINCT`              |
| 表别名 | `FROM 表名 别名`        |

------

# 📘 第3章：データの制限とソート

（Restricting and Sorting Data）

------

## 🧩 一、WHERE句とは

**WHERE子句**用来限制（筛选）查询结果中返回的行。

基本语法：

```sql
SELECT 列名
FROM 表名
WHERE 条件式;
```

### 🧠 例1：查找属于开发部的员工

```sql
SELECT first_name, last_name, department_id
FROM EMPLOYEES
WHERE department_id = 10;
```

输出示例：

| FIRST_NAME | LAST_NAME | DEPARTMENT_ID |
| ---------- | --------- | ------------- |
| 田中       | 太郎      | 10            |
| 伊藤       | 健        | 10            |

------

## 🧩 二、条件表达式（Comparison Operators）

| 运算符                    | 说明     | 例子                               |
| ------------------------- | -------- | ---------------------------------- |
| `=`                       | 等于     | `department_id = 10`               |
| `>` `<` `>=` `<=`         | 大小比较 | `salary >= 400000`                 |
| `!=` 或 `<>`              | 不等于   | `job_title <> '営業担当'`          |
| `BETWEEN ... AND ...`     | 范围     | `salary BETWEEN 300000 AND 500000` |
| `IN (...)`                | 列表匹配 | `department_id IN (10, 20, 30)`    |
| `LIKE`                    | 模糊匹配 | `last_name LIKE '田%'`             |
| `IS NULL` / `IS NOT NULL` | 是否为空 | `manager_id IS NULL`               |

------

### 🧠 例2：工资高于40万的员工

```sql
SELECT first_name, last_name, salary
FROM EMPLOYEES
WHERE salary > 400000;
```

### 🧠 例3：工资在30万～50万之间的员工

```sql
SELECT first_name, salary
FROM EMPLOYEES
WHERE salary BETWEEN 300000 AND 500000;
```

### 🧠 例4：姓氏以“田”开头的人

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE last_name LIKE '田%';
```

> `%` 是通配符，代表任意字符序列。

------

## 🧩 三、逻辑运算符（AND / OR / NOT）

| 运算符 | 说明             | 示例                                       |
| ------ | ---------------- | ------------------------------------------ |
| AND    | 同时满足多个条件 | `salary > 400000 AND department_id = 10`   |
| OR     | 满足其中一个条件 | `department_id = 10 OR department_id = 20` |
| NOT    | 否定             | `NOT job_title = '営業担当'`               |

### 🧠 例5：开发部或营业部的员工

```sql
SELECT first_name, department_id
FROM EMPLOYEES
WHERE department_id IN (10, 20);
```

### 🧠 例6：工资高于40万且职位不是营业担当

```sql
SELECT first_name, job_title, salary
FROM EMPLOYEES
WHERE salary > 400000
  AND job_title <> '営業担当';
```

------

## 🧩 四、ORDER BY句（排序）

排序语法：

```sql
SELECT ...
FROM ...
ORDER BY 列名 [ASC|DESC];
```

| 关键字 | 说明         |
| ------ | ------------ |
| ASC    | 升序（默认） |
| DESC   | 降序         |

### 🧠 例7：按照工资从高到低排序

```sql
SELECT first_name, salary
FROM EMPLOYEES
ORDER BY salary DESC;
```

### 🧠 例8：部门升序 + 工资降序

```sql
SELECT first_name, department_id, salary
FROM EMPLOYEES
ORDER BY department_id ASC, salary DESC;
```

------

## 🧩 五、NULLの扱い（空值）

NULL 表示「没有值」或「未知」。

⚠️ 注意：

- `salary = NULL` ❌ 无效
- 应使用 `IS NULL` 或 `IS NOT NULL`

例：

```sql
SELECT first_name, manager_id
FROM EMPLOYEES
WHERE manager_id IS NULL;
```

------

## 🧩 六、SQL行制限句（ROWNUM）

Oracle 特有的功能，用来限制返回行数。

例：显示工资最高的前3人

```sql
SELECT first_name, salary
FROM EMPLOYEES
ORDER BY salary DESC
FETCH FIRST 3 ROWS ONLY;
```

（Oracle 12c以后的标准语法，旧版本用 `ROWNUM`）

------

## 🧩 七、置換変数（替换变量）※SQL*Plus或Oracle Live SQL用

在执行SQL时提示输入值，例如：

```sql
SELECT * FROM EMPLOYEES
WHERE department_id = &DEPT_ID;
```

执行时系统会提示输入 `DEPT_ID`。

------

## 🧩 八、练习题（実習）

基于你的五张表来做例题：

### ✅ 练习1：显示开发部（department_id=10）的所有员工

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE department_id = 10;
```

### ✅ 练习2：查找薪资在30万到50万之间的员工

```sql
SELECT first_name, salary
FROM EMPLOYEES
WHERE salary BETWEEN 300000 AND 500000;
```

### ✅ 练习3：显示没有上司（manager_id为空）的员工

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE manager_id IS NULL;
```

### ✅ 练习4：找出2024年以后入职的员工

```sql
SELECT first_name, hire_date
FROM EMPLOYEES
WHERE hire_date >= TO_DATE('2024-01-01', 'YYYY-MM-DD');
```

### ✅ 练习5：显示各部门工资最高的前三人（Oracle 12c语法）

```sql
SELECT first_name, department_id, salary
FROM EMPLOYEES
ORDER BY department_id, salary DESC
FETCH FIRST 3 ROWS WITH TIES;
```

------

## 🧭 本章小结

| 功能     | 关键语法                | 示例                |
| -------- | ----------------------- | ------------------- |
| 筛选行   | `WHERE`                 | `salary > 400000`   |
| 范围     | `BETWEEN ... AND ...`   | `BETWEEN 10 AND 20` |
| 列表匹配 | `IN`                    | `IN (10,20,30)`     |
| 模糊匹配 | `LIKE '田%'`            |                     |
| 空值判断 | `IS NULL`               |                     |
| 多条件   | `AND / OR / NOT`        |                     |
| 排序     | `ORDER BY ... ASC/DESC` |                     |

------

# 📘 第4章：単一行関数を使用した出力のカスタマイズ

（Customize Output Using Single-Row Functions）

------

## 🧩 一、什么是单行函数（Single Row Function）

**特点：**

- 每行执行一次，返回一行结果
- 可以嵌套使用
- 在 SELECT、WHERE、ORDER BY、GROUP BY 等子句中使用

### 🧠 例：将名字改为大写

```sql
SELECT UPPER(first_name) AS 名前
FROM EMPLOYEES;
```

------

## 🧩 二、文字函数（Character Functions）

用于处理文字（字符串）。

| 函数                         | 作用       | 例子                                     |
| ---------------------------- | ---------- | ---------------------------------------- |
| `UPPER(str)`                 | 全部大写   | `UPPER('tanaka') → TANAKA`               |
| `LOWER(str)`                 | 全部小写   | `LOWER('TANAKA') → tanaka`               |
| `INITCAP(str)`               | 首字母大写 | `INITCAP('tanaka taro') → Tanaka Taro`   |
| `CONCAT(a, b)`               | 拼接       | `CONCAT(first_name, last_name)`          |
| `SUBSTR(str, start, length)` | 取子串     | `SUBSTR('YAMADA', 1, 3) → YAM`           |
| `LENGTH(str)`                | 字符数     | `LENGTH('山田') → 2`                     |
| `REPLACE(str, old, new)`     | 替换       | `REPLACE('営業部', '部', '課') → 営業課` |

------

### 🧠 例1：显示员工名（首字母大写）

```sql
SELECT INITCAP(first_name || ' ' || last_name) AS 氏名
FROM EMPLOYEES;
```

### 🧠 例2：部门名后加“部”

```sql
SELECT REPLACE(department_name, '部', '部門') AS 部署名
FROM DEPARTMENTS;
```

------

## 🧩 三、数值函数（Number Functions）

对数字（如薪资）进行数学运算。

| 函数            | 说明              | 示例                           |
| --------------- | ----------------- | ------------------------------ |
| `ROUND(num, n)` | 四舍五入到第 n 位 | `ROUND(1234.567, 2) → 1234.57` |
| `TRUNC(num, n)` | 截断              | `TRUNC(1234.567, 1) → 1234.5`  |
| `MOD(a, b)`     | 取余数            | `MOD(10, 3) → 1`               |

### 🧠 例3：工资按千位四舍五入

```sql
SELECT first_name, ROUND(salary, -3) AS 給与_丸め
FROM EMPLOYEES;
```

------

## 🧩 四、日期函数（Date Functions）

Oracle 的日期（`DATE`）类型非常强大，能存储 年/月/日/时/分/秒。

| 函数                     | 说明           | 示例                                 |
| ------------------------ | -------------- | ------------------------------------ |
| `SYSDATE`                | 当前系统日期   | `SELECT SYSDATE FROM DUAL;`          |
| `ADD_MONTHS(date, n)`    | 加/减月        | `ADD_MONTHS(hire_date, 6)`           |
| `MONTHS_BETWEEN(d1, d2)` | 两日期相差月数 | `MONTHS_BETWEEN(SYSDATE, hire_date)` |
| `NEXT_DAY(date, '曜日')` | 下一个指定星期 | `NEXT_DAY(SYSDATE, '金曜日')`        |
| `LAST_DAY(date)`         | 所在月最后一天 | `LAST_DAY(SYSDATE)`                  |

------

### 🧠 例4：计算员工入职到今天的月数

```sql
SELECT first_name,
       ROUND(MONTHS_BETWEEN(SYSDATE, hire_date), 1) AS 入社月数
FROM EMPLOYEES;
```

### 🧠 例5：求入职6个月后的日期

```sql
SELECT first_name,
       ADD_MONTHS(hire_date, 6) AS 試用期間終了
FROM EMPLOYEES;
```

------

## 🧩 五、嵌套函数（Nested Functions）

函数可以嵌套使用。

### 🧠 例6：名字转大写后取前3个字符

```sql
SELECT SUBSTR(UPPER(first_name), 1, 3) AS 名前略
FROM EMPLOYEES;
```

------

## 🧩 六、函数在WHERE中的使用

可以在 WHERE 中用函数筛选条件：

### 🧠 例7：取姓氏以“田”开头（不区分大小写）

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE UPPER(last_name) LIKE '田%';
```

------

## 🧩 七、日期格式化函数 TO_CHAR(date, format)

格式化日期输出。

```sql
TO_CHAR(hire_date, 'YYYY/MM/DD')
TO_CHAR(SYSDATE, 'YYYY年MM月DD日 HH24:MI:SS')
```

### 🧠 例8：显示入职日（YYYY年MM月DD日）

```sql
SELECT first_name,
       TO_CHAR(hire_date, 'YYYY年MM月DD日') AS 入社日
FROM EMPLOYEES;
```

------

## 🧩 八、练习题（実習）

> 使用你的五张表来练习：

### ✅ 练习1：把所有员工名显示为“首字母大写 姓名”

```sql
SELECT INITCAP(first_name || ' ' || last_name) AS 氏名
FROM EMPLOYEES;
```

### ✅ 练习2：显示每位员工的薪资（四舍五入到千位）

```sql
SELECT first_name, ROUND(salary, -3) AS 給与_千単位
FROM EMPLOYEES;
```

### ✅ 练习3：显示员工入职到现在的月数

```sql
SELECT first_name,
       ROUND(MONTHS_BETWEEN(SYSDATE, hire_date), 1) AS 入社月数
FROM EMPLOYEES;
```

### ✅ 练习4：显示项目表中，项目名的前5个字符

```sql
SELECT SUBSTR(project_name, 1, 5) AS プロジェクト略称
FROM PROJECTS;
```

### ✅ 练习5：显示薪资历史中每条记录的登记日（格式化输出）

```sql
SELECT record_id,
       TO_CHAR(update_date, 'YYYY/MM/DD HH24:MI') AS 更新日時
FROM SALARY_HISTORY;
```

------

## 🧭 本章小结

| 函数类别   | 示例                                          | 用途       |
| ---------- | --------------------------------------------- | ---------- |
| 文字函数   | `UPPER()`, `SUBSTR()`, `REPLACE()`            | 操作文字   |
| 数值函数   | `ROUND()`, `TRUNC()`, `MOD()`                 | 操作数值   |
| 日期函数   | `SYSDATE`, `ADD_MONTHS()`, `MONTHS_BETWEEN()` | 操作日期   |
| 格式化函数 | `TO_CHAR(date, format)`                       | 日期格式化 |

------



<!-- TODO: translate from testFile/original/toutest.md -->
