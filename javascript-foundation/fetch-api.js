const API_BASE_URL = "https://jsonplaceholder.typicode.com";

const ENPOINTS = {
  GET_USER_PROFILE: `${API_BASE_URL}/users/1`,
};

async function getData() {
  try {
    const respone = await fetch(ENPOINTS.GET_USER_PROFILE);
    if (respone.status === 200) {
      const data = await respone.json();
      console.log("Data Berhasil didapat: ", data);
    }
  } catch (error) {
    console.log("Gagal Mengambil Data: ", error);
  }
}

getData();
