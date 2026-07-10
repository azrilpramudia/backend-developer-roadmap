const hitServerDatabase = (userId) => {
  return new Promise((resolve, reject) => {
    console.log("[Server] Mencari data user di database...");

    // delay internet selama 2 detik
    setTimeout(() => {
      const dataFound = true; // smulasi sukses

      if (dataFound) {
        // Status Promise Sukses (Fulfilled)
        resolve({ id: userId, name: "Meow", role: "Backend Developer" });
      } else {
        // Status Promise gagal (Rejected)
        reject("Error: user tidak di temukan di database!");
      }
    }, 2000);
  });
};

async function prosesDataUser() {
  try {
    console.log("[Client] Memulai Proses...");
    const user = await hitServerDatabase(101);
    console.log(
      `[Client] Sukses! Data = Nama: ${user.name}, Role: ${user.role}`,
    );
  } catch (error) {
    console.error("Terjadi kesalahan sistem:", error);
  } finally {
    console.log("[Client] Proses Selesai");
  }
}

prosesDataUser();
