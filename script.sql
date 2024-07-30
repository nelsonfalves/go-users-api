-- Cria o banco de dados
CREATE DATABASE IF NOT EXISTS user_db;

-- Seleciona o banco de dados
USE user_db;

-- Cria a tabela 'user'
CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200),
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    telephone VARCHAR(15) NOT NULL
);

-- Insere um usuário de teste
INSERT INTO user (name, username, email, password, telephone)
VALUES ('test', 'test', 'test@test.com', 'test', '00000000000');

