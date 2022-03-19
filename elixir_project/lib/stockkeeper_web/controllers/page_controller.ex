defmodule StockkeeperWeb.PageController do
  use StockkeeperWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
