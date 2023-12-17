use crate::models::app_state::AppState;
use crate::models::event::Event;
use crate::models::event::NewEvent;
use axum::debug_handler;
use axum::{extract::State, http::StatusCode, Json};

#[debug_handler]
pub async fn get_events(State(state): State<AppState>) -> (StatusCode, Json<Vec<Event>>) {
    let rows = sqlx::query_as::<_, Event>("SELECT * FROM events");
    let events: Vec<Event> = rows.fetch_all(&state.db_pool).await.unwrap();

    (StatusCode::OK, Json(events))
}

#[debug_handler]
pub async fn add_event(
    State(state): State<AppState>,
    Json(new_event): Json<NewEvent>,
) -> (StatusCode, Json<String>) {
    let result = sqlx::query!("INSERT INTO events (event_name, start_time, end_time, location) VALUES ($1, $2, $3, $4)",
        new_event.event_name,
        new_event.start_time,
        new_event.end_time,
        new_event.location
    ).execute(&state.db_pool)
    .await;

    match result {
        Ok(_) => return (StatusCode::OK, Json("Event added successfully".to_string())),
        Err(e) => {
            eprintln!("Failed to execute query: {:?}", e);
            return (
                StatusCode::INTERNAL_SERVER_ERROR,
                Json("Failed to add event".to_string()),
            );
        }
    };
}
