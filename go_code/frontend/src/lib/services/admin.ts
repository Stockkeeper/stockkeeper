export async function createAdminPassword(password: string): Promise<any> {
  let url = "http://localhost:8000/api/v1/admin/password";
  let options = {
    method: "POST",
    headers: {
      "content-type": "application/json",
      "accept": "application/json",
    },
    body: JSON.stringify({ password }),
    mode: "cors" as RequestMode,
  };
  let response = await fetch(url, options);
  let data = await response.json();
  return data;
};

export async function getAdminPassword(): Promise<any> {
  let url = "http://localhost:8000/api/v1/admin/password";
  let options = {
    method: "GET",
    headers: {
      "accept": "application/json",
    },
    mode: "cors" as RequestMode,
  };
  let response = await fetch(url, options);
  let data = await response.json();
  return data;
};
