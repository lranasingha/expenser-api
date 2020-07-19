CREATE TABLE expense_schema.user_expense
(
    user_id     SERIAL      not null
        CONSTRAINT pk_exp_uid PRIMARY KEY,
    category    varchar(10) not null,
    name        varchar(100),
    description varchar(250),
    payload     bytea       not null
);

