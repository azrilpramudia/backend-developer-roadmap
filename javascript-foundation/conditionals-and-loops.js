// Conditional (if-else)
const httpStatusCode = 200;

if (httpStatusCode === 200) {
  console.log("Request Success");
} else if (httpStatusCode === 404) {
  console.log("Resource Not Found");
} else {
  console.log("Internal Server Error");
}

// Loop (for & while)
// digunakan untuk perulangan dengan index yang jelas
const ports = [3000, 5000, 8000];

for (let i = 0; i < ports.length; i++) {
  console.log("Checking Port: ", ports[i]);
}

// While
// digunakan untuk perulangan bergantung pada kondisi tertentu, bukan jumlah index
let count = 1;
while (count <= 3) {
  console.log("Ping Server: ", count);
  count++;
}
