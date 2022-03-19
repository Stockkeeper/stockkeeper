defmodule StockkeeperWeb.RegistryAPI.BlobController do
  use StockkeeperWeb, :controller
  import Plug.Conn

  plug StockkeeperWeb.RegistryAPI.Validator, :repo_name
  plug StockkeeperWeb.RegistryAPI.Validator, :blob_digest

  def show(conn, %{"repo_name" => repo_name, "blob_digest" => blob_digest}) do
    conn
    |> put_status(:not_found)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render(
      "error.json",
      code: "BLOB_UNKNOWN",
      message:
        "Unable to find blob with repository name #{repo_name} and blob digest #{blob_digest}"
    )
  end

  def delete(conn, %{"repo_name" => repo_name, "blob_digest" => blob_digest}) do
    conn
    |> put_status(:not_found)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render(
      "error.json",
      code: "BLOB_UNKNOWN",
      message:
        "Unable to find blob with repository name #{repo_name} and blob digest #{blob_digest}"
    )
  end
end
