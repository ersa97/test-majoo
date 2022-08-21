1. Get User

   - select id, name , user_name , password , created_at, created_by , updated_at , updated_by from users where user_name = 'admin1'  ORDER BY `users`.`id` ASC LIMIT 1

2. Get Merchant Transactions by User id

create view merchantomzet as
select
    `merchants`.`user_id` as `user_id`,
    `merchants`.`merchant_name` as `merchant_name`,
    `transactions`.`created_at` as `created_at`,
    `transactions`.`bill_total` as `bill_total`
from
    (`merchants`
join `transactions`)
where
    (`merchants`.`id` = `transactions`.`merchant_id`)


   - SELECT count(*) FROM `merchantomzet`  WHERE (user_id =  1)

   - select merchant_name, created_at, bill_total from merchantomzet where user_id =  1  ORDER BY user_id asc LIMIT 10 OFFSET 0

3. Get Merchant Outlet Transactions by User id

create view merchantomzetoutlet as
    select
    `merchants`.`user_id` as `user_id`,
    `merchants`.`merchant_name` as `merchant_name`,
    `outlets`.`outlet_name` as `outlet_name`,
    `transactions`.`created_at` as `created_at`,
    `transactions`.`bill_total` as `bill_total`
from
    ((`merchants`
join `transactions`)
join `outlets`)
where
    ((`merchants`.`id` = `transactions`.`merchant_id`)
    and (`merchants`.`id` = `outlets`.`merchant_id`))
order by
    `outlets`.`outlet_name`,
    `transactions`.`created_at`


   - SELECT count(*) FROM `merchantomzetoutlet`  WHERE (user_id = 1)

   - select merchant_name ,outlet_name ,created_at,bill_total from merchantomzetoutlet where user_id = 1 LIMIT 10 OFFSET 0