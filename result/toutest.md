# 📘 第2章：SQL SELECT文を使用したデータの取得

（SQL 基本クエリ）

------

## 🧩 一、SELECT文の基本構造

SQL（Structured Query Language）は、**宣言型言語**です。  
 私たちはデータベースに「何を取りたいか」だけを伝え、どうやって実行するかを伝える必要はありません。

文法構造は次のとおりです：

```sql
SELECT 列名1, 列名2, ...
FROM 表名;
```

👉 **要点**：

- **SELECT**：取得する列
- **FROM**：どのテーブルから取得するか
- **セミコロン ;**：文の終了（必須）

------

### 🧠 例1：全従業員を表示

```sql
SELECT * FROM EMPLOYEES;
```

> `*` は「すべての列」を表します。

出力例（例）：

| EMPLOYEE_ID | FIRST_NAME | LAST_NAME | JOB_TITLE  | DEPARTMENT_ID | MANAGER_ID | HIRE_DATE  | SALARY |
| ----------- | ---------- | --------- | ---------- | ------------- | ---------- | ---------- | ------ |
| 101         | 田中       | 太郎      | エンジニア | 10            | 100        | 2022-04-01 | 420000 |
| 102         | 鈴木       | 花子      | 営業担当   | 20            | 105        | 2023-01-15 | 350000 |

------

## 🧩 二、列別名（AS）

列の別名は出力の見出しを読みやすくします。

```sql
SELECT first_name AS 名前,
       last_name AS 苗字,
       salary AS 給与
FROM EMPLOYEES;
```

Oracle では `AS` を省略できます：

```sql
SELECT first_name 名前 FROM EMPLOYEES;
```

------

## 🧩 三、連結演算子（||）

複数の列を1つの文字列に結合できます。

```sql
SELECT first_name || ' ' || last_name AS 氏名
FROM EMPLOYEES;
```

出力例：

| 氏名      |
| --------- |
| 田中 太郎 |
| 鈴木 花子 |

------

## 🧩 四、DISTINCT 重複排除

部門IDが何通りあるか（重複を除く）を見たい場合：

```sql
SELECT DISTINCT department_id
FROM EMPLOYEES;
```

出力：

| DEPARTMENT_ID |
| ------------- |
| 10            |
| 20            |
| 30            |

------

## 🧩 五、文字列リテラル

表の列の他に固定値を直接書くこともできます。

```sql
SELECT 'Oracle SQL学習中' AS コメント,
       first_name, last_name
FROM EMPLOYEES;
```

出力例：

| コメント         | FIRST_NAME | LAST_NAME |
| ---------------- | ---------- | --------- |
| Oracle SQL学習中 | 田中       | 太郎      |
| Oracle SQL学習中 | 鈴木       | 花子      |

------

## 🧩 六、表別名（Table Alias）

複数テーブルの結合や長いテーブル名の場合、別名を使うと文が簡潔になります：

```sql
SELECT e.first_name, e.last_name, d.department_name
FROM EMPLOYEES e, DEPARTMENTS d
WHERE e.department_id = d.department_id;
```

ここでは：

- `EMPLOYEES e` は EMPLOYEES テーブルに別名 e を付ける
- `DEPARTMENTS d` は DEPARTMENTS テーブルに別名 d を付ける

------

## 🧩 七、練習題（実習）

> あなたが定義した5つのテーブルを使って練習します。

### ✅ 練習1：全部門の名称と番号を表示

```sql
SELECT department_id, department_name
FROM DEPARTMENTS;
```

### ✅ 練習2：従業員の氏名と所属部門IDを表示

```sql
SELECT first_name || ' ' || last_name AS 氏名,
       department_id
FROM EMPLOYEES;
```

### ✅ 練習3：プロジェクト表のプロジェクト名（重複排除）

```sql
SELECT DISTINCT project_name
FROM PROJECTS;
```

### ✅ 練習4：給与履歴表の各レコードに説明文を追加

```sql
SELECT record_id,
       employee_id,
       '給与履歴データ' AS コメント
FROM SALARY_HISTORY;
```

### ✅ 練習5（上級）：従業員名と対応する上司IDを表示（表に別名を付ける）

```sql
SELECT e.first_name || ' ' || e.last_name AS 社員,
       e.manager_id AS 上司ID
FROM EMPLOYEES e;
```

------

## 🧩 八、拡張理解（Oracle特有）

Oracle 特有の構文と特徴：

| 特性                                                         | 説明 |
| ------------------------------------------------------------ | ---- |
| SQLは大小文字を区別しないが、キーワードは通常大文字。           |     |
| 文字列は `' '` で囲み、列名は囲まない。                           |     |
| 各文は `;` で終了。                                           |     |
| Oracle はデフォルトで `DUAL` テーブルを使用して定数をテストします：例：`SELECT SYSDATE FROM DUAL;`（システム日付を取得） |     |

------

## 🧭 本章のまとめ

| 知識点 | 关键語法                |
| ------ | ----------------------- |
| 列を取得 | `SELECT 列名 FROM テーブル名` |
| 全件取得 | `SELECT *`              |
| 別名 | `AS` または直接別名を記述       |
| 結合 | 連結を表す                   |
| 重複排除 | `DISTINCT`              |
| テーブル別名 | `FROM テーブル名 別名`        |

------

# 📘 第3章：データの制限とソート

（Restricting and Sorting Data）

------

## 🧩 一、WHERE句とは

**WHERE句**は、クエリ結果として返される行を制限（フィルタ）します。

基本構文：

```sql
SELECT 列名
FROM 表名
WHERE 条件式;
```

### 🧠 例1：開発部に所属する従業員を検索

```sql
SELECT first_name, last_name, department_id
FROM EMPLOYEES
WHERE department_id = 10;
```

出力例：

| FIRST_NAME | LAST_NAME | DEPARTMENT_ID |
| ---------- | --------- | ------------- |
| 田中       | 太郎      | 10            |
| 伊藤       | 健        | 10            |

------

## 🧩 二、条件式（Comparison Operators）

| 演算子                    | 説明     | 例                               |
| ------------------------- | -------- | ---------------------------------- |
| `=`                       | イコール | `department_id = 10`               |
| `>` `<` `>=` `<=`         | 比較     | `salary >= 400000`                 |
| `!=` または `<>`           | 不等号   | `job_title <> '営業担当'`          |
| `BETWEEN ... AND ...`     | 範囲     | `salary BETWEEN 300000 AND 500000` |
| `IN (...)`                | リスト一致 | `department_id IN (10, 20, 30)`    |
| `LIKE`                    | パターン一致 | `last_name LIKE '田%'`             |
| `IS NULL` / `IS NOT NULL` | NULL判定 | `manager_id IS NULL`               |

------

### 🧠 例2：給与が40万を超える従業員

```sql
SELECT first_name, last_name, salary
FROM EMPLOYEES
WHERE salary > 400000;
```

### 🧠 例3：給与が30万〜50万の従業員

```sql
SELECT first_name, salary
FROM EMPLOYEES
WHERE salary BETWEEN 300000 AND 500000;
```

### 🧠 例4：姓が「田」で始まる人

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE last_name LIKE '田%';
```

> `%` はワイルドカードで、任意の文字列を表します。

------

## 🧩 三、論理演算子（AND / OR / NOT）

| 演算子 | 説明             | 例                                        |
| ------ | ---------------- | ----------------------------------------- |
| AND    | 複数条件を同時満たす | `salary > 400000 AND department_id = 10`   |
| OR     | 条件のいずれかを満たす | `department_id = 10 OR department_id = 20` |
| NOT    | 否定             | `NOT job_title = '営業担当'`               |

### 🧠 例5：開発部または営業部の従業員

```sql
SELECT first_name, department_id
FROM EMPLOYEES
WHERE department_id IN (10, 20);
```

### 🧠 例6：給与が40万を超え、職位が営業担当ではない

```sql
SELECT first_name, job_title, salary
FROM EMPLOYEES
WHERE salary > 400000
  AND job_title <> '営業担当';
```

------

## 🧩 四、ORDER BY句（ソート）

ソートの構文：

```sql
SELECT ...
FROM ...
ORDER BY 列名 [ASC|DESC];
```

| キーワード | 説明         |
| ------ | ------------ |
| ASC    | 昇順（デフォルト） |
| DESC   | 降順         |

### 🧠 例7：給与の高い順にソート

```sql
SELECT first_name, salary
FROM EMPLOYEES
ORDER BY salary DESC;
```

### 🧠 例8：部門昇順 + 給与降順

```sql
SELECT first_name, department_id, salary
FROM EMPLOYEES
ORDER BY department_id ASC, salary DESC;
```

------

## 🧩 五、NULLの扱い（空値）

NULL は「値が無い」または「未知」を示します。

⚠️ 注意：

- `salary = NULL` ❌ 無効
- 使用すべきは `IS NULL` または `IS NOT NULL`

例：

```sql
SELECT first_name, manager_id
FROM EMPLOYEES
WHERE manager_id IS NULL;
```

------

## 🧩 六、SQL行制限句（ROWNUM）

Oracle特有の機能で、返される行数を制限します。

例：給与が最も高い上位3名を表示

```sql
SELECT first_name, salary
FROM EMPLOYEES
ORDER BY salary DESC
FETCH FIRST 3 ROWS ONLY;
```

（Oracle 12c以降の標準構文、旧バージョンは `ROWNUM`）

------

## 🧩 七、置換変数（替换变量）※SQL*PlusまたはOracle Live SQL用

SQL実行時に値の入力を促す例：

```sql
SELECT * FROM EMPLOYEES
WHERE department_id = &DEPT_ID;
```

実行時にシステムは `DEPT_ID` の入力を求めます。

------

## 🧩 八、練習題（実習）

あなたの5つのテーブルを基に演習問題：

### ✅ 練習1：開発部（department_id=10）のすべての従業員を表示

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE department_id = 10;
```

### ✅ 練習2：給与が30万〜50万の従業員を検索

```sql
SELECT first_name, salary
FROM EMPLOYEES
WHERE salary BETWEEN 300000 AND 500000;
```

### ✅ 練習3：上司がいない（manager_id が NULL）の従業員を表示

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE manager_id IS NULL;
```

### ✅ 練習4：2024年以降に入社した従業員を検索

```sql
SELECT first_name, hire_date
FROM EMPLOYEES
WHERE hire_date >= TO_DATE('2024-01-01', 'YYYY-MM-DD');
```

### ✅ 練習5：各部門で給与が最高の上位3名を表示（Oracle 12c文法）

```sql
SELECT first_name, department_id, salary
FROM EMPLOYEES
ORDER BY department_id, salary DESC
FETCH FIRST 3 ROWS WITH TIES;
```

------

## 🧭 本章のまとめ

| 功能     | 关键语法                | 示例                |
| -------- | ----------------------- | ------------------- |
| 行を絞り込み | `WHERE`                 | `salary > 400000`   |
| 範囲     | `BETWEEN ... AND ...`   | `BETWEEN 10 AND 20` |
| リスト一致 | `IN`                    | `IN (10,20,30)`     |
| ワイルドカード | `LIKE '田%'`            |                     |
| NULL判定 | `IS NULL`               |                     |
| 複数条件   | `AND / OR / NOT`        |                     |
| ソート     | `ORDER BY ... ASC/DESC` |                     |

------

# 📘 第4章：単一行関数を使用した出力のカスタマイズ

（Single Row Functionを用いた出力のカスタマイズ）

------

## 🧩 一、単一行関数（Single Row Function）とは

**特徴：**

- 各行ごとに1回実行され、1行の結果を返します
- ネストして使用できます
- SELECT、WHERE、ORDER BY、GROUP BY などの句で使用します

### 🧠 例：名前を大文字に変換

```sql
SELECT UPPER(first_name) AS 名前
FROM EMPLOYEES;
```

------

## 🧩 二、文字関数（Character Functions）

文字列を処理するための関数。

| 関数                         | 作用       | 例                                      |
| ---------------------------- | ---------- | ---------------------------------------- |
| `UPPER(str)`                 | 全部大文字   | `UPPER('tanaka') → TANAKA`               |
| `LOWER(str)`                 | 全部小文字   | `LOWER('TANAKA') → tanaka`               |
| `INITCAP(str)`               | 首字母大文字 | `INITCAP('tanaka taro') → Tanaka Taro`   |
| `CONCAT(a, b)`               | 連結       | `CONCAT(first_name, last_name)`          |
| `SUBSTR(str, start, length)` | 文字列の一部を取得 | `SUBSTR('YAMADA', 1, 3) → YAM`           |
| `LENGTH(str)`                | 文字数     | `LENGTH('山田') → 2`                     |
| `REPLACE(str, old, new)`     | 置換       | `REPLACE('営業部', '部', '課') → 営業課` |

------

### 🧠 例1：従業員名を表示（頭文字を大文字）

```sql
SELECT INITCAP(first_name || ' ' || last_name) AS 氏名
FROM EMPLOYEES;
```

### 🧠 例2：部門名の後ろに「部」を追加

```sql
SELECT REPLACE(department_name, '部', '部門') AS 部署名
FROM DEPARTMENTS;
```

------

## 🧩 三、数値関数（Number Functions）

数値（給与など）に対して数学演算を行います。

| 関数            | 说明              | 示例                           |
| --------------- | ----------------- | ------------------------------ |
| `ROUND(num, n)` | n位へ四捨五入      | `ROUND(1234.567, 2) → 1234.57` |
| `TRUNC(num, n)` | 切り捨て            | `TRUNC(1234.567, 1) → 1234.5`  |
| `MOD(a, b)`     | 余り               | `MOD(10, 3) → 1`               |

### 🧠 例3：給与を千位に四捨五入

```sql
SELECT first_name, ROUND(salary, -3) AS 給与_丸め
FROM EMPLOYEES;
```

------

## 🧩 四、日付関数（Date Functions）

Oracleの日付（DATE）型は非常に強力で、年/月/日/時/分/秒を格納できます。

| 関数                     | 说明           | 示例                                 |
| ------------------------ | -------------- | ------------------------------------ |
| `SYSDATE`                | 現在のシステム日付 | `SELECT SYSDATE FROM DUAL;`          |
| `ADD_MONTHS(date, n)`    | 月の加算/減算   | `ADD_MONTHS(hire_date, 6)`           |
| `MONTHS_BETWEEN(d1, d2)` | 二つの日付の差を月数で | `MONTHS_BETWEEN(SYSDATE, hire_date)` |
| `NEXT_DAY(date, '曜日')` | 指定曜日の次の日   | `NEXT_DAY(SYSDATE, '金曜日')`        |
| `LAST_DAY(date)`         | 該当月の最終日   | `LAST_DAY(SYSDATE)`                  |

------

### 🧠 例4：入社日から今日までの月数を計算

```sql
SELECT first_name,
       ROUND(MONTHS_BETWEEN(SYSDATE, hire_date), 1) AS 入社月数
FROM EMPLOYEES;
```

### 🧠 例5：入社から6か月後の日付を取得

```sql
SELECT first_name,
       ADD_MONTHS(hire_date, 6) AS 試用期間終了
FROM EMPLOYEES;
```

------

## 🧩 五、ネスト関数（Nested Functions）

関数はネストして使用できます。

### 🧠 例6：名前を大文字にして先頭3文字を取得

```sql
SELECT SUBSTR(UPPER(first_name), 1, 3) AS 名前略
FROM EMPLOYEES;
```

------

## 🧩 六、WHEREでの関数使用

WHERE句内で関数を使って条件を絞ることができます。

### 🧠 例7：姓が「田」で始まる（大文字小文字を区別しない）

```sql
SELECT first_name, last_name
FROM EMPLOYEES
WHERE UPPER(last_name) LIKE '田%';
```

------

## 🧩 七、日付形式化関数 TO_CHAR(date, format)

日付の出力形式を整形します。

```sql
TO_CHAR(hire_date, 'YYYY/MM/DD')
TO_CHAR(SYSDATE, 'YYYY年MM月DD日 HH24:MI:SS')
```

### 🧠 例8：入社日を表示（YYYY年MM月DD日）

```sql
SELECT first_name,
       TO_CHAR(hire_date, 'YYYY年MM月DD日') AS 入社日
FROM EMPLOYEES;
```

------

## 🧩 八、練習題（実習）

> あなたの5つのテーブルを使って練習します：

### ✅ 練習1：すべての従業員名を「頭文字大文字 名前」に表示

```sql
SELECT INITCAP(first_name || ' ' || last_name) AS 氏名
FROM EMPLOYEES;
```

### ✅ 練習2：各従業員の給与を表示（千位に四捨五入）

```sql
SELECT first_name, ROUND(salary, -3) AS 給与_千単位
FROM EMPLOYEES;
```

### ✅ 練習3：従業員の入社から現在までの月数

```sql
SELECT first_name,
       ROUND(MONTHS_BETWEEN(SYSDATE, hire_date), 1) AS 入社月数
FROM EMPLOYEES;
```

### ✅ 練習4：プロジェクト表のプロジェクト名の先頭5文字

```sql
SELECT SUBSTR(project_name, 1, 5) AS プロジェクト略称
FROM PROJECTS;
```

### ✅ 練習5：給与履歴の各レコードの更新日を表示（形式化出力）

```sql
SELECT record_id,
       TO_CHAR(update_date, 'YYYY/MM/DD HH24:MI') AS 更新日時
FROM SALARY_HISTORY;
```

------

## 🧭 本章のまとめ

| 関数カテゴリー   | 例                                          | 用途       |
| ---------- | --------------------------------------------- | ---------- |
| 文字関数   | `UPPER()`, `SUBSTR()`, `REPLACE()`            | 文字処理   |
| 数値関数   | `ROUND()`, `TRUNC()`, `MOD()`                 | 数値処理   |
| 日付関数   | `SYSDATE`, `ADD_MONTHS()`, `MONTHS_BETWEEN()` | 日付操作   |
| 形式化関数 | `TO_CHAR(date, format)`                       | 日付の整形 |