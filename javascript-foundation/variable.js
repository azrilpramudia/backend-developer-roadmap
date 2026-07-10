// gunakan const untuk data yang tetap
const serverPort = 3000;
const databaseName = "users_db";

// gunakan let jika data nya akan berubah
let loginAttempts = 0;
loginAttempts = loginAttempts + 1;

console.log("Server Running on Port: " + serverPort);
console.log("Login Attempts: " + loginAttempts);
console.log("Database: " + databaseName);
