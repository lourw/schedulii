mod handlers;
mod models;

use axum::{
    http::header::CONTENT_TYPE, http::Method, routing::delete, routing::get, routing::post, Router,
    Server,
};
use axum_prometheus::PrometheusMetricLayer;
use dotenvy::dotenv;
use models::app_state::AppState;
use sqlx::postgres::PgPoolOptions;
use std::env;
use std::net::SocketAddr;
use tower_http::cors::{Any, CorsLayer};

#[tokio::main]
async fn main() {
    dotenv().ok();
    let database_url = match env::var("DATABASE_URL") {
        Ok(val) => val,
        Err(e) => e.to_string(),
    };
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await
        .expect("Failed to create connection pool.");

    let (prometheus_layer, metric_handler) = PrometheusMetricLayer::pair();
    let state = AppState { db_pool: pool };

    let cors = CorsLayer::new()
        .allow_methods([Method::GET, Method::POST])
        .allow_origin(Any)
        .allow_headers([CONTENT_TYPE]);

    let app = Router::new()
        .route("/", get(|| async { "Hello, World" }))
        .route("/events", get(handlers::event_handler::get_events))
        .route("/events/add", post(handlers::event_handler::add_event))
        .route(
            "/events/delete",
            delete(handlers::event_handler::delete_event),
        )
        .route("/metrics", get(|| async move { metric_handler.render() }))
        .layer(prometheus_layer)
        .layer(cors)
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
    use tower::ServiceExt;

    #[tokio::test]
    async fn main() {
        let app = Router::new().route("/", get(|| async { "Hello World" }));

        let response = app
            .oneshot(Request::builder().uri("/").body(Body::empty()).unwrap())
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::OK);
    }
}
