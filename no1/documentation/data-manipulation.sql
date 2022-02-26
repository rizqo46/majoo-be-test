-- Untuk Query Omzet Merchant 
SELECT merchant_name FROM Merchants m WHERE m.id = 1 LIMIT 1
CREATE TEMPORARY TABLE IF NOT EXISTS temp_date(date DATE);

INSERT INTO temp_date(date) VALUES ("2021-11-01"),
                ("2021-11-02"),("2021-11-03"),("2021-11-04"),
                ("2021-11-05"),("2021-11-06"),("2021-11-07"),
                ("2021-11-08"),("2021-11-09"),
                ("2021-11-10"),("2021-11-11"),("2021-11-12"),
                ("2021-11-13"),("2021-11-14"),("2021-11-15"),
                ("2021-11-16"),("2021-11-17"),("2021-11-18"),
                ("2021-11-19"),("2021-11-20"),("2021-11-21"),
                ("2021-11-22"),("2021-11-23"),("2021-11-24"),
                ("2021-11-25"),("2021-11-26"),("2021-11-27"),
                ("2021-11-28"),("2021-11-29"),("2021-11-30");

SELECT 'merchant 1' as merchant_name, DATE_FORMAT(date, "%Y-%m-%d") as date, omzet FROM temp_date AS D LEFT JOIN
                (SELECT m.id, m.merchant_name, SUM(t.bill_total) as omzet, DATE(DATE_FORMAT(t.created_at, "%Y-%m-%d")) AS datef
                FROM Merchants m LEFT JOIN Transactions t ON t.merchant_id = m.id GROUP BY datef HAVING m.id=1) 
                as t ON t.datef = D.date LIMIT 3 OFFSET 0;

-- Untuk Query Omzet Outlet
SELECT outlet_name, merchant_id FROM Outlets o WHERE o.id = 1 LIMIT 1

SELECT merchant_name FROM Merchants m WHERE m.id = 1 LIMIT 1

CREATE TEMPORARY TABLE IF NOT EXISTS temp_date(date DATE);
                INSERT INTO temp_date(date) VALUES ("2021-11-01"),
                ("2021-11-02"),("2021-11-03"),("2021-11-04"),
                ("2021-11-05"),("2021-11-06"),("2021-11-07"),
                ("2021-11-08"),("2021-11-09"),
                ("2021-11-10"),("2021-11-11"),("2021-11-12"),
                ("2021-11-13"),("2021-11-14"),("2021-11-15"),
                ("2021-11-16"),("2021-11-17"),("2021-11-18"),
                ("2021-11-19"),("2021-11-20"),("2021-11-21"),
                ("2021-11-22"),("2021-11-23"),("2021-11-24"),
                ("2021-11-25"),("2021-11-26"),("2021-11-27"),
                ("2021-11-28"),("2021-11-29"),("2021-11-30");


SELECT 'merchant 1' as merchant_name, 'Outlet 1' as outlet_name, DATE_FORMAT(date, "%Y-%m-%d") as date, omzet FROM temp_date AS D LEFT JOIN
                (SELECT o.id, SUM(t.bill_total) as omzet, DATE(DATE_FORMAT(t.created_at, "%Y-%m-%d")) AS datef
                FROM Outlets o LEFT JOIN Transactions t ON t.outlet_id = o.id GROUP BY datef HAVING o.id=1) 
                as t ON t.datef = D.date LIMIT 5 OFFSET 5;