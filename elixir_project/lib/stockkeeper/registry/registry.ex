defmodule Stockkeeper.Registry do
  alias Stockkeeper.Registry.{Blob, Repository, UploadSession}
  alias Stockkeeper.Repo

  @spec create_repository(String.t()) ::
          {:ok, Repository.t()} | {:error, :repository_name_already_exists}
  def create_repository(name) do
    %Repository{}
    |> Repository.new_changeset(%{name: name})
    |> Repo.insert()
  end

  @spec open_upload_session(Repository.t(), Keyword.t()) ::
          {:ok, UploadSession.t()} | {:error, :repository_not_found}
  def open_upload_session(repository, opts \\ [is_multichunk?: false]) do
    id = Ecto.UUID.generate()
    upload_location = "#{@upload_base_url}/v2/#{repository.name}/blobs/uploads/#{id}"

    %Blob{
      id: id,
      repository_id: repository.id,
      location: location,
      is_multichunk?: opts[:is_multichunk?]
    }
    |> Repo.insert()
    |> case do
      {:ok, blob} ->
        {:ok, blob}

      {:error, changeset} ->
        IO.inspect(changeset)
        {:error, :repository_not_found}
    end
  end

  @spec upload_chunk(Blob.t())
end
