defmodule StockkeeperWeb.RegistryAPI.Validator do
  use StockkeeperWeb, :controller
  import Plug.Conn

  @blob_digest_pattern "^[a-f0-9]{64}$"
  @manifest_ref_pattern "^[a-zA-Z0-9_][a-zA-Z0-9._-]{0,127}$"
  @repo_name_pattern "^[a-z0-9]+([._-][a-z0-9]+)*(/[a-z0-9]+([._-][a-z0-9]+)*)*$"

  def init(type), do: type

  def call(%Plug.Conn{path_params: %{"blob_digest" => blob_digest}} = conn, :blob_digest) do
    if Regex.match?(~r"#{@blob_digest_pattern}", blob_digest) do
      conn
    else
      conn
      |> put_status(:bad_request)
      |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
      |> render(
        "error.json",
        code: "DIGEST_INVALID",
        message: "Invalid blob digest #{blob_digest}"
      )
    end
  end

  def call(
        %Plug.Conn{
          path_params: %{"manifest_ref" => manifest_ref},
          assigns: %{repository: repository}
        } = conn,
        :manifest_ref
      ) do
    if not Regex.match?(~r"#{@manifest_ref_pattern}", manifest_ref) do
      conn
      |> put_status(:not_found)
      |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
      |> render(
        "error.json",
        code: "MANIFEST_INVALID",
        message:
          "The manifest reference #{manifest_ref} is invalid. The manifest reference must match the regex pattern #{@manifest_ref_pattern}."
      )
    else
      case Stockkeeper.Registry.Repository.get_artifact_by_ref(repository, manifest_ref) do
        nil ->
          conn
          |> put_status(:not_found)
          |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
          |> render(
            "error.json",
            code: "MANIFEST_UNKNOWN",
            message:
              "Unable to find manifest with repository name #{repository.name} and manifest reference #{manifest_ref}"
          )

        artifact ->
          assign(conn, :artifact, artifact)
      end
    end
  end

  def call(%Plug.Conn{path_params: %{"repo_name" => repo_name}} = conn, :repo_name) do
    if not Regex.match?(~r"#{@repo_name_pattern}", repo_name) do
      conn
      |> put_status(:not_found)
      |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
      |> render(
        "error.json",
        code: "NAME_INVALID",
        message:
          "The repository name #{repo_name} is invalid. The repository name must match the regex pattern #{@repo_name_pattern}."
      )
      |> halt()
    else
      case Stockkeeper.Registry.get_repository_by_name(repo_name) do
        nil ->
          conn
          |> put_status(:not_found)
          |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
          |> render(
            "error.json",
            code: "NAME_UNKNOWN",
            message: "A repository with the name #{repo_name} could not be found."
          )

        %Stockkeeper.Registry.Repository{} = repository ->
          assign(conn, :repository, repository)
      end
    end
  end
end
