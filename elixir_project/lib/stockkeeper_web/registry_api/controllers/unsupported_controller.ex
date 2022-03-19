defmodule StockkeeperWeb.RegistryAPI.UnsupportedController do
  use StockkeeperWeb, :controller
  import Plug.Conn

  def error(conn, _params) do
    conn
    |> put_status(501)
    |> put_view(StockkeeperWeb.RegistryAPI.ErrorView)
    |> render("unsupported.json")
  end
end
