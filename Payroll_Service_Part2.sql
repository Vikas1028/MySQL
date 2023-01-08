#UC1
create database payroll_service;
show databases;
use payroll_service;

#UC2
create table employee_payroll(
Id int auto_increment primary key,
Name varchar(50) not null,
Salary int8 not null,
Start_date date not null);

#UC3
insert into employee_payroll (Name, Salary, Start_date) value ('Vikas', 30000, '2022-10-28'), ('Akash', 50000, '2019-9-27'), ('Sanket', 40000, '2020-5-8');

#UC4
select * from employee_payroll;

#UC5
select Salary from employee_payroll where Name = 'Vikas';
select Name from employee_payroll where Start_date between CAST('2018-01-01' AS DATE) AND DATE(NOW());

#UC6
alter table employee_payroll add Gender char(1) not null;
SET SQL_SAFE_UPDATES=0;
update employee_payroll set Gender = 'M' where Name = 'Vikas' or Name = 'Akash' or Name = 'Sanket';

#UC7
select sum(Salary) from employee_payroll where Gender = 'M' group by Gender;
select avg(Salary) from employee_payroll where Gender = 'M' group by Gender;
select min(Salary) from employee_payroll where Gender = 'M' group by Gender;
select max(Salary) from employee_payroll where Gender = 'M' group by Gender;
select count(*) from employee_payroll where Gender = 'M' group by Gender;

#UC8
alter table employee_payroll
add Phone int8 not null,
add Address varchar(100) not null default 'N/A',
add Department varchar(50) not null;

#UC9
alter table employee_payroll
add BasicPay int8 not null,
add Deductions int8 not null,
add TaxablePay int8 not null,
add IncomeTax int8 not null,
add NetPay int8 not null;

#UC10
alter table employee_payroll drop column Salary;
alter table employee_payroll drop column Phone;
alter table employee_payroll drop column Address;
desc employee_payroll;
select * from employee_payroll;
alter table employee_payroll modify start_date int8 not null after NetPay;
insert into employee_payroll
(Name, department, Gender, BasicPay, Deductions, TaxablePay, IncomeTax, NetPay, start_date) VALUES
('Terisa', 'Marketting', 'F', 3000000.00, 1000000.00, 2000000.00, 500000.00, 1500000.00, '2018-3-17' );
insert into employee_payroll
(Name, department, Gender, BasicPay, Deductions, TaxablePay, IncomeTax, NetPay, start_date) VALUES
('Terisa', 'Sales', 'F', 3000000.00, 1000000.00, 2000000.00, 500000.00, 1500000.00, '2018-3-17' );

#UC10
#Generate ER diagram




