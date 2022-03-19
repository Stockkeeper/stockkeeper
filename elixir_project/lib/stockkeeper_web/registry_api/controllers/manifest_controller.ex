defmodule StockkeeperWeb.RegistryAPI.ManifestController do
  use StockkeeperWeb, :controller
  import Plug.Conn

  plug StockkeeperWeb.RegistryAPI.Validator, :repo_name
  plug StockkeeperWeb.RegistryAPI.Validator, :manifest_ref

  def show(conn, %{"repo_name" => repo_name, "manifest_ref" => manifest_ref}) do
    conn
    |> put_status(:not_found)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render(
      "error.json",
      code: "MANIFEST_UNKNOWN",
      message:
        "Unable to find manifest with repository name #{repo_name} and manifest reference #{manifest_ref}"
    )
  end

  def update(conn, %{"repo_name" => _repo_name, "manifest_ref" => _manifest_ref}) do
    conn
    |> put_status(:not_found)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render("unsupported.json")
  end

  def delete(conn, %{"repo_name" => repo_name, "manifest_ref" => manifest_ref}) do
    conn
    |> put_status(:not_found)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render(
      "error.json",
      code: "MANIFEST_UNKNOWN",
      message:
        "Unable to find manifest with repository name #{repo_name} and manifest reference #{manifest_ref}"
    )
  end
end
