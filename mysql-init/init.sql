ALTER USER 'codewarrior'@'%' IDENTIFIED WITH mysql_native_password BY 'codewarx';
GRANT ALL PRIVILEGES ON *.* TO 'codewarrior'@'%';
FLUSH PRIVILEGES;
