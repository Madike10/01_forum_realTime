
//Fetch les donnees when POST method is called

export async function postData(url = "", data = {}) {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    console.log("response OK");
    return response.json();
  }


export async function getData(url = '') {
  console.log('getting', url)
  const response = await fetch(url, {
      method: 'GET'
  })
  return response.json()
}