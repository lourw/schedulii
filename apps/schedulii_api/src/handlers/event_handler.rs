use crate::models::app_state::AppState;
use crate::models::event::Event;
use axum::debug_handler;
use axum::{extract::State, http::StatusCode};

#[debug_handler]
pub async fn get_events(State(state): State<AppState>) -> (StatusCode, String) {
    let rows = sqlx::query_as::<_, Event>("SELECT * FROM events");
    let events: Vec<Event> = rows.fetch_all(&state.db_pool).await.unwrap();

    println!("{:?}", events);
    return (StatusCode::OK, "Hello, WORLD!".to_string());
}
