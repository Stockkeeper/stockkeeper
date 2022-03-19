defmodule StockkeeperWeb.RegistryAPI.BlobUploadControllerTest do
  use StockkeeperWeb.ConnCase

  setup do
    {:ok, repository} = Stockkeeper.Registry.create_repository("my-repository")
    {:ok, upload_session} = Stockkeeper.Registry.create_upload_session(repository)

    %{
      repository: repository,
      upload_session: upload_session
    }
  end

  describe "POST /v2/:repo_name/blobs/uploads" do
    test "[error] opening a session when the repository name does not exist", %{conn: conn} do
      conn =
        post(
          conn,
          Routes.registry_api_blob_upload_path(
            conn,
            :create,
            "bad-repository-name"
          )
        )

      assert response(conn, :not_found)
    end

    test "[ok] opening a session", %{conn: conn, repository: repository} do
      conn =
        post(
          conn,
          Routes.registry_api_blob_upload_path(
            conn,
            :create,
            repository.name
          )
        )

      assert response(conn, :accepted)

      [location] = Plug.Conn.get_resp_header(conn, "location")
      assert valid_upload_location?(location, repository)
    end

    test "[ok] uploading a blob", %{conn: conn, repository: repository} do
      blob_data = "012346789"

      blob_hash =
        :crypto.hash(:sha256, blob_data)
        |> Base.encode16()
        |> String.downcase()

      conn =
        conn
        |> Plug.Conn.put_req_header("content-type", "application/octet-stream")
        |> Plug.Conn.put_req_header("content-length", "#{String.length(blob_data)}")
        |> post(
          Routes.registry_api_blob_upload_path(
            conn,
            :create,
            repository.name,
            digest: "sha256:#{blob_hash}"
          ),
          blob_data
        )

      assert response(conn, :created)

      [location] = Plug.Conn.get_resp_header(conn, "location")
      assert valid_blob_location?(location, repository)
    end
  end

  describe "PUT /v2/:repo_name/blobs/uploads/:upload_session_id" do
    # test "fails when repository name does not exist", %{conn: conn} do
    #   conn =
    #     put(
    #       conn,
    #       Routes.registry_api_blob_upload_path(
    #         conn,
    #         :update,
    #         "bad-repository-name",
    #         "bad-upload-session-id"
    #       )
    #     )

    #   assert response(conn, :not_found)
    # end

    # test "fails when upload session ID does not exist", %{
    #   conn: conn,
    #   repository: repository
    # } do
    #   conn =
    #     put(
    #       conn,
    #       Routes.registry_api_blob_upload_path(
    #         conn,
    #         :update,
    #         repository.name,
    #         "bad-upload-session-id"
    #       )
    #     )

    #   assert response(conn, :not_found)
    # end

    #   test "succeeds", %{conn: conn, repository: repository, upload_session: upload_session} do
    #     blob_data = "012346789"
    #     blob_hash =
    #       :crypto.hash(:sha256, blob_data)
    #       |> Base.encode16()
    #       |> String.downcase()

    #     conn
    #     |> Plug.Conn.put_req_header("content-type", "application/octet-stream")
    #     |> Plug.Conn.put_req_header("content-length", length(blob_data))
    #     |> post(
    #       Routes.registry_api_blob_upload_path(
    #         conn,
    #         :update,
    #         repository.name,
    #         upload_session.id
    #       ),
    #       blob_data,
    #       digest: "sha256:#{blob_hash}"
    #     )

    #     assert response(conn, :created)

    #     [location] = Plug.Conn.get_resp_header(conn, "location")
    #     assert valid_blob_location?(location, repository)
    #  end
  end

  defp valid_upload_location?(location, repository) do
    location =~
      ~r"\/v2\/#{repository.name}\/blobs\/uploads\/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
  end

  defp valid_blob_location?(location, repository) do
    location =~
      ~r"\/v2\/#{repository.name}\/blobs\/[0-9a-f]{0,127}"
  end
end
