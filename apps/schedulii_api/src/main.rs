use axum::{routing::get, Router, Server};
use axum_prometheus::PrometheusMetricLayer;
use std::net::SocketAddr;
use sqlx::postgres::PgPoolOptions;
use std::env;

#[derive(Clone)]
struct AppState {
    _db_pool: sqlx::PgPool,
}

#[tokio::main]
async fn main() {
    let database_url = match env::var("DATABASE_URL") {
        Ok(val) => val,
        Err(e) => e.to_string(),
    };
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await
        .unwrap();

    let (prometheus_layer, metric_handler) = PrometheusMetricLayer::pair();
    let state = AppState { _db_pool: pool };

    let app = Router::new()
        .route("/", get(|| async { "Hello, World" }))
        .route("/metrics", get(|| async move { metric_handler.render() }))
        .layer(prometheus_layer)
        .with_state(state);

    let addr = SocketAddr::from(([0, 0, 0, 0], 9000));
    Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

#[cfg(test)]
mod tests {
    use axum::body::Body;
    use axum::http::{Request, StatusCode};
    use axum::{routing::get, Router};
    use axum_prometheus::PrometheusMetricLayer;
    use tower::ServiceExt;

    #[tokio::test]
    async fn main() {
        let (prometheus_layer, metric_handler) = PrometheusMetricLayer::pair();

        let app = Router::new()
            .route("/", get(|| async { "Hello World" }))
            .route("/metrics", get(|| async move { metric_handler.render() }))
            .layer(prometheus_layer);

        let response = app
            .oneshot(Request::builder().uri("/").body(Body::empty()).unwrap())
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::OK);
    }
}
