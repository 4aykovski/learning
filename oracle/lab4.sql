SELECT *
FROM hr.employees e1
         JOIN hr.employees e2 on e1.manager_id = e2.employee_id
WHERE e1.hire_date < e2.hire_date and (TO_CHAR(CURRENT_DATE, 'YYYY') - TO_CHAR(e1.hire_date, 'YYYY') > 17);

SELECT e1.last_name, e1.first_name, e1.phone_number, e2.last_name, e2.first_name, e2.phone_number
FROM hr.employees e1
         LEFT JOIN hr.employees e2 on e1.manager_id = e2.employee_id

SELECT first_name, last_name, salary
FROM hr.employees e1
         JOIN hr.departments d on d.manager_id = e1.employee_id
WHERE (TO_CHAR(CURRENT_DATE, 'YYYY') - TO_CHAR(e1.hire_date, 'YYYY') > 16)

SELECT department_name, first_name, last_name, email, phone_number, hire_date, salary
FROM hr.employees e1
         JOIN hr.departments d on d.manager_id = e1.employee_id

SELECT d.department_id, d.department_name, d.manager_id, d.location_id
FROM hr.departments d
WHERE (SELECT COUNT(*)
       FROM hr.employees e
       WHERE e.department_id = d.department_id
       GROUP BY department_id) > 5

SELECT d.department_id, d.department_name, d.manager_id, d.location_id
FROM hr.departments d
WHERE (SELECT max(salary)
       FROM hr.employees e
       WHERE e.department_id = d.department_id) > 10000

SELECT d.department_id, d.department_name, d.manager_id, d.location_id
FROM hr.departments d
WHERE (SELECT min(salary)
       FROM hr.employees e
       WHERE e.department_id = d.department_id) < 5000

-----
SELECT d.department_id, d.department_name, d.manager_id, d.location_id,
       (SELECT avg(salary)
        FROM hr.employees e
        WHERE e.department_id=d.department_id) average
FROM hr.departments d;

SELECT SUM(*)
FROM (SELECT COUNT(*) com_count
      FROM hr.employees e
      WHERE commission_pct NOT LIKE '%-%'
      GROUP BY commission_pct)
----