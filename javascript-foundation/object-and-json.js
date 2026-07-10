// Create & Access Object

const userProfile = {
  useranme: "meow",
  role: "Backend Developer",
  isStatusActive: true,
  skills: ["Go", "JavaScript", "Linux", "Docker"],
};

// Access Properti with dot notation
console.log("Username: ", userProfile.useranme);
console.log("Skils: ", userProfile.skills[0]); // take index-0

// how to change mark inside object
userProfile.role = "Go Developer";
console.log("Role Baru: ", userProfile.role);

// Manipulation JSON (Proccess send & receive data)
// Skenario A: Mengirim data ke Backend Go (Serialization / Marshalling)
const jsonstringToSend = JSON.stringify(userProfile);
console.log(jsonstringToSend);

// Skenario B: Menerima data dari Backend Go (Deserialization / Unmarshalling)
const jsonResponeFromBackend =
  '{"status":"success","message":"Data Berhasil disimpan"}';
const parsedObject = JSON.parse(jsonResponeFromBackend);

console.log("Status dari Server: ", parsedObject.status);
console.log("Pesan dari Server: ", parsedObject.message);
