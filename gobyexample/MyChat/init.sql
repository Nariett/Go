CREATE TABLE Users (
    Id SERIAL PRIMARY KEY,
    Name CHARACTER VARYING(30) UNIQUE,
    Password CHARACTER VARYING(30)
);

INSERT INTO Users (Name, Password) VALUES 
('Саша', '123456'),
('Олег', '1111');