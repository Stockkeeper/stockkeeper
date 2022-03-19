defmodule StockkeeperWeb.RegistryAPI.SpecSupportController do
  use StockkeeperWeb, :controller
  import Plug.Conn

  def show(conn, _params) do
    send_resp(conn, 200, "")
  end
end
