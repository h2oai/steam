#'
#' Login to a steam client
#'
#' Login provides a connection for the user to communicate with a Steam instance
#' @export
steam.login <- function(ip = "localhost", port = 9000, verify_ssl = TRUE, username = NA_character_,
                        password = NA_character_, login_file = NA_character_,
                        login_file_passphrase = NA_character_) {
    conn <- structure(
        list(ip=ip, port=port, verify_ssl=verify_ssl, username=username, password=password, uid=0),
        class="SteamConnection"
    )

    .pingServer(conn, "check")
    conn
}

#'
#' Start an H2O cluster through Steam
#'
#' @export
steam.startCluster <- function(conn, cluster_name, num_nodes = 3, memory_per_node="10g", h2o_version) {
    # Fetch/Start entities through Steam
    e <- .getEngineByVersion(conn, h2o_version)$engine
    id <- .startClusterOnYarn(
        conn         = conn,
        cluster_name = cluster_name,
        engine_id    = engine$id,
        size         = num_nodes,
        memory       = memory_per_node,
        secure       = TRUE,
        keytab       = NA_character_
    )$cluster_id
    h <- .getCluster(conn, id)$cluster

    # Fetch ouput entities
    x <- unlist(strsplit(h$address, ":"), recursive = FALSE)
    c <- .getConfig(conn)$config
    list(
        id = id,
        ip = conn$ip,
        port = substr(c$cluster_proxy_address, 2, 6),
        startH2O = FALSE,
        cookies = structure(list(h$token), names=h$name),
        context_path = h$context_path
    )
}

#'
#' Stop an H2O cluster through steam
#'
#' @export
steam.stopCluster <- function(conn, config) {
    void <- .stopClusterOnYarn(conn, config$id, NA_character_)
}
