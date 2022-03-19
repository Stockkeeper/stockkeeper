defmodule StockkeeperWeb.RegistryAPI.BlobUploadController do
  use StockkeeperWeb, :controller
  import Plug.Conn

  def create(conn, params) do
    conn
    |> try_get_repository(params)
    |> try_open_blob()
    |> try_upload_chunk(params)
  end

  defp try_get_repository(%Plug.Conn{halted: true} = conn, _params), do: conn

  defp try_get_repository(conn, %{"repo_name" => repo_name}) do
    case Stockkeeper.Registry.Repository.get_by_name(repo_name) do
      nil ->
        conn
        |> put_status(:not_found)
        |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
        |> render("error.json",
          code: "NAME_UNKNOWN",
          message: "Unable to find a repository with the name \"#{repo_name}\""
        )
        |> halt()

      %Stockkeeper.Registry.Repository{} = repository ->
        assign(conn, :repository, repository)
    end
  end

  defp try_open_blob(%Plug.Conn{halted: true} = conn), do: conn

  defp try_open_blob(%Plug.Conn{assigns: %{repository: repository}} = conn) do
    case Stockkeeper.Registry.Repository.open_blob(repository) do
      {:error, :repository_not_found} ->
        conn
        |> put_status(:not_found)
        |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
        |> render("error.json",
          code: "NAME_UNKNOWN",
          message: "Unable to find a repository with the name \"#{repository.name}\""
        )
        |> halt()

      {:ok, blob} ->
        assign(conn, :blob, blob)
    end
  end

  defp try_upload_chunk(%Plug.Conn{halted: true} = conn, _params), do: conn

  defp try_upload_chunk(conn, %{"digest" => _digest}) do
    conn
    |> put_resp_header("location", "/blob-location")
    |> send_resp(:created, [])
  end

  defp try_upload_chunk(conn, _params) do
    conn
    |> put_resp_header("location", "/upload-location")
    |> send_resp(:accepted, [])
  end

  def update(conn, _params) do
    conn
    |> put_status(501)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render("unsupported.json")
  end

  def finish(conn, _params) do
    conn
    |> put_status(501)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render("unsupported.json")
  end
end
