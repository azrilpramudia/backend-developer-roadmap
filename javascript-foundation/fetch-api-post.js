async function registerUserBaru() {
  const dataForm = {
    name: "meow",
    username: "meowhx",
    role: "backend",
  };

  try {
    const respone = await fetch("https://jsonplaceholder.typicode.com/posts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataForm),
    });

    if (respone.status === 201) {
      const result = await respone.json();
      console.log("Created 201");
      console.log("Response dari server", result);
    } else {
      console.log("Gagal membuat data baru. Status:", respone.status);
    }
  } catch (error) {
    console.error("Gagal melakukan registrasi:", error);
  }
}

registerUserBaru();
