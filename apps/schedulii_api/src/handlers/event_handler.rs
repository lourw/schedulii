use crate::models::app_state::AppState;
use crate::models::event::Event;
use axum::debug_handler;
use axum::{extract::Query, extract::State, http::StatusCode, extract::Json};
use serde::Deserialize;
use sqlx::types::chrono::{DateTime, Utc};

#[debug_handler]
pub async fn get_events(State(state): State<AppState>) -> (StatusCode, Json<Vec<Event>>) {
    let rows = sqlx::query_as::<_, Event>("SELECT * FROM events");
    let events: Vec<Event> = rows.fetch_all(&state.db_pool).await.unwrap();

    (StatusCode::OK, Json(events))
}

#[derive(Deserialize)]
pub struct AddEventParams {
  pub event_name: String,
  pub start_time: DateTime<Utc>,
  pub end_time: DateTime<Utc>,
  pub location: String
}

#[debug_handler]
pub async fn add_event(
    State(state): State<AppState>,
    Json(new_event): Json<AddEventParams>,
) -> (StatusCode, Json<String>) {
    let result = sqlx::query!(
        "INSERT INTO events (event_name, start_time, end_time, location) VALUES ($1, $2, $3, $4)",
        new_event.event_name,
        new_event.start_time,
        new_event.end_time,
        new_event.location
    )
    .execute(&state.db_pool)
    .await;

    match result {
        Ok(_) => (StatusCode::OK, Json("Event added successfully".to_string())),
        Err(e) => {
            eprintln!("Failed to execute query: {:?}", e);
            (
                StatusCode::INTERNAL_SERVER_ERROR,
                Json("Failed to add event".to_string()),
            )
        }
    }
}

#[derive(Deserialize)]
pub struct DeleteEventParams {
    pub event_id: i32,
}

#[debug_handler]
pub async fn delete_event(
    State(state): State<AppState>,
    Query(params): Query<DeleteEventParams>,
) -> StatusCode {
    let result = sqlx::query!("DELETE FROM events WHERE event_id = $1", params.event_id)
        .execute(&state.db_pool)
        .await;

    match result {
        Ok(_) => StatusCode::OK,
        Err(e) => {
            eprintln!("Failed to execute query: {:?}", e);
            StatusCode::INTERNAL_SERVER_ERROR
        }
    }
}
