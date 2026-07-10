async function getProfileUser(userId) {
  try {
    const respone = await fetch(
      `https://jsonplaceholder.typicode.com/users/${userId}`,
    );

    if (respone.status === 200) {
      const userData = await respone.json();
      console.log("Nama User:", userData.name);
      console.log("Email User:", userData.email);
    } else if (respone.status === 404) {
      console.log("User tidak ditemukan...");
    }
  } catch (error) {
    console.error("Gagal koneksi ke server:", error);
  }
}

getProfileUser(1);
